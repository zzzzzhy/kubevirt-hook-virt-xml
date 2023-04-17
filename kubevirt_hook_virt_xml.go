package main

import (
	"context"
	"net"
	"os"
	"strings"

	"github.com/beevik/etree"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	"kubevirt.io/client-go/log"
	"kubevirt.io/kubevirt/pkg/hooks"
	hooksInfo "kubevirt.io/kubevirt/pkg/hooks/info"
	hooksV1alpha2 "kubevirt.io/kubevirt/pkg/hooks/v1alpha2"
)

const (
	onDefineDomainLoggingMessage = "Hook's OnDefineDomain callback method has been called"
	// usage                           = `updater --args '--vcpus 10'`
	readBytesSec  = "virt.xml.hook/readBytesSec"
	writeBytesSec = "virt.xml.hook/writeBytesSec"
	readIopsSec   = "virt.xml.hook/readIopsSec"
	writeIopsSec  = "virt.xml.hook/writeIopsSec"
)

type infoServer struct{}

func (s infoServer) Info(ctx context.Context, params *hooksInfo.InfoParams) (*hooksInfo.InfoResult, error) {
	log.Log.Info("Hook's Info method has been called")

	return &hooksInfo.InfoResult{
		Name: "update-xml",
		Versions: []string{
			hooksV1alpha2.Version,
		},
		HookPoints: []*hooksInfo.HookPoint{
			{
				Name:     hooksInfo.OnDefineDomainHookPointName,
				Priority: 0,
			},
		},
	}, nil
}

// const virtXML = "virt-xml"

type v1alpha2Server struct {
	args []string
}

func AddBlkiotune(domainXML []byte, args []string) ([]byte, error) {
	// vmiSpec := vmSchema.VirtualMachineInstance{}
	// err := json.Unmarshal(vmiJSON, &vmiSpec)
	// if err != nil {
	// 	log.Log.Reason(err).Errorf("Failed to unmarshal given VMI spec: %s", vmiJSON)
	// 	panic(err)
	// }

	// annotations := vmiSpec.GetAnnotations()

	// if _, found := annotations[baseBoardManufacturerAnnotation]; !found {
	// 	log.Log.Info("SM BIOS hook sidecar was requested, but no attributes provided. Returning original domain spec")
	// 	return domainXML, nil
	// }

	// domainSpec := domainSchema.DomainSpec{}
	// err := xml.Unmarshal(domainXML, &domainSpec)
	// if err != nil {
	// 	log.Log.Reason(err).Errorf("Failed to unmarshal given domain spec: %s", domainXML)
	// 	panic(err)
	// }

	// // domainSpec.OS.SMBios = &domainSchema.SMBios{Mode: "sysinfo"}

	// if domainSpec.Devices.Disks == nil {
	// 	log.Log.Reason(err).Errorf("Not found disks")
	// 	return nil, nil
	// }
	// if args[1] == "all" {
	// 	// _args := strings.Split(args[0], ",")
	// 	for _, v := range domainSpec.Devices.Disks {
	// 		print(v)
	// 	}
	// } else {

	// }
	// domainSpec.SysInfo.Type = "smbios"
	// if baseBoardManufacturer, found := annotations[baseBoardManufacturerAnnotation]; found {
	// domainSpec.SysInfo.BaseBoard = append(domainSpec.SysInfo.BaseBoard, domainSchema.Entry{
	// 	Name:  "manufacturer",
	// 	Value: "baseBoardManufacturer",
	// })
	// }

	domainSpec := etree.NewDocument()
	if err := domainSpec.ReadFromBytes(domainXML); err != nil {
		panic(err)
	}
	devices := domainSpec.SelectElement("domain").SelectElement("devices")
	for _, v := range devices.SelectElements("disk") {
		iotune := v.CreateElement("iotune")
		total_bytes_sec := iotune.CreateElement("total_bytes_sec")
		total_bytes_sec.CreateText("5120000")
	}
	// if args[1] == "all" {
	// 	// _args := strings.Split(args[0], ",")
	// 	for _, v := range devices.SelectElements("disk") {
	// 		print(v)
	// 	}
	// }else{}
	newDomainXML, err := domainSpec.WriteToBytes()
	// newDomainXML, err := xml.Marshal(domainSpec)
	if err != nil {
		log.Log.Reason(err).Errorf("Failed to marshal updated domain spec: %+v", domainSpec)
		panic(err)
	}

	log.Log.Info("Successfully updated original domain spec with requested SMBIOS attributes")
	// domainSpec.Indent(2)
	// domainSpec.WriteTo(os.Stdout)
	return newDomainXML, nil
}

func (s v1alpha2Server) OnDefineDomain(ctx context.Context, params *hooksV1alpha2.OnDefineDomainParams) (*hooksV1alpha2.OnDefineDomainResult, error) {
	log.Log.Info(onDefineDomainLoggingMessage)
	newDomainXML, err := AddBlkiotune(params.GetDomainXML(), s.args)
	if err != nil {
		return nil, err
	}
	return &hooksV1alpha2.OnDefineDomainResult{
		DomainXML: newDomainXML,
	}, nil
}

func (s v1alpha2Server) PreCloudInitIso(_ context.Context, params *hooksV1alpha2.PreCloudInitIsoParams) (*hooksV1alpha2.PreCloudInitIsoResult, error) {
	return &hooksV1alpha2.PreCloudInitIsoResult{
		CloudInitData: params.GetCloudInitData(),
	}, nil
}

func main() {
	log.InitializeLogging("xml update")

	var version string
	pflag.StringVar(&version, "version", "", "hook version to use")

	var options string
	pflag.StringVar(&options, "args", "", "params to pass to virt-xml")
	pflag.Parse()

	args := strings.Split(options, "|")

	socketPath := hooks.HookSocketsSharedDirectory + "/update.sock"
	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Log.Reason(err).Errorf("Failed to initialized socket on path: %s", socket)
		log.Log.Error("Check whether given directory exists and socket name is not already taken by other file")
		panic(err)
	}
	defer os.Remove(socketPath)
	// if options == "" {
	// 	panic(fmt.Errorf(usage))
	// }

	server := grpc.NewServer([]grpc.ServerOption{}...)
	hooksInfo.RegisterInfoServer(server, infoServer{})
	hooksV1alpha2.RegisterCallbacksServer(server, v1alpha2Server{args: args})
	log.Log.Infof("Starting hook server exposing 'info' and 'v1alpha2' services on socket %s", socketPath)
	server.Serve(socket)
}
