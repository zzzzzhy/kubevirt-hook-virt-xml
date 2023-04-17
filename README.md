参数传递方法
使用args传递:
```
hooks.kubevirt.io/hookSidecars: '[{"args": ["--args","all"], "image": "docker.io/rubyroes/virt-xml-hook:latest"}]'
```
使用annotations传递:
virt.xml.hook/readBytesSec: '5120000'

接收参数示例:
args传递:
```
type v1alpha2Server struct {
	args []string
}
func main() {
	log.InitializeLogging("xml update")

	var version string
	pflag.StringVar(&version, "version", "", "hook version to use")

	var options string
	pflag.StringVar(&options, "args", "", "params to pass to virt-xml")
	pflag.Parse()

	args := strings.Split(options, "|")
    ...
}
```
annotations传递:
```
import {
    vmSchema "kubevirt.io/api/core/v1"
}
const (
	onDefineDomainLoggingMessage = "Hook's OnDefineDomain callback method has been called"
	// usage                           = `updater --args '--vcpus 10'`
	readBytesSec  = "virt.xml.hook/readBytesSec"
	writeBytesSec = "virt.xml.hook/writeBytesSec"
	readIopsSec   = "virt.xml.hook/readIopsSec"
	writeIopsSec  = "virt.xml.hook/writeIopsSec"
    totalBytesSec = "virt.xml.hook/totalBytesSec"
    totalIopsSec = "virt.xml.hook/totalIopsSec"
)
func AddBlkiotune(vmiJSON []byte, domainXML []byte, args []string) ([]byte, error) {
    ...
	vmiSpec := vmSchema.VirtualMachineInstance{}
	err := json.Unmarshal(vmiJSON, &vmiSpec)
	if err != nil {
		log.Log.Reason(err).Errorf("Failed to unmarshal given VMI spec: %s", vmiJSON)
		panic(err)
	}

	annotations := vmiSpec.GetAnnotations()

	if _, found := annotations[readBytesSec]; !found {
		log.Log.Info("SM BIOS hook sidecar was requested, but no attributes provided. Returning original domain spec")
		return domainXML, nil
	}
    ...
}
func (s v1alpha2Server) OnDefineDomain(ctx context.Context, params *hooksV1alpha2.OnDefineDomainParams) (*hooksV1alpha2.OnDefineDomainResult, error) {
	log.Log.Info(onDefineDomainLoggingMessage)
	newDomainXML, err := AddBlkiotune(params.GetVmi(), params.GetDomainXML(), s.args)
    ...
}
```
<details> <summary>最终效果</summary>

```
<domain type='kvm' id='1'>
  <name>default_vmi-with-sidecar-hook-2</name>
  <uuid>42582d40-f317-4c8d-bafe-d5dadefc6623</uuid>
[...]
  <devices>
    <emulator>/usr/libexec/qemu-kvm</emulator>
    <disk type='file' device='disk' model='virtio-non-transitional'>
      <driver name='qemu' type='qcow2' cache='none' error_policy='stop' discard='unmap'/>
      <source file='/var/run/kubevirt-ephemeral-disks/disk-data/containerdisk/disk.qcow2' index='2'/>
      <backingStore type='file' index='3'>
        <format type='qcow2'/>
        <source file='/var/run/kubevirt/container-disks/disk_0.img'/>
      </backingStore>
      <target dev='vda' bus='virtio'/>
      <iotune>
        <total_bytes_sec>52428800</total_bytes_sec>
        <total_iops_sec>1000</total_iops_sec>
      </iotune>
      <alias name='ua-containerdisk'/>
      <address type='pci' domain='0x0000' bus='0x04' slot='0x00' function='0x0'/>
    </disk>
    <disk type='file' device='disk' model='virtio-non-transitional'>
      <driver name='qemu' type='raw' cache='none' error_policy='stop' discard='unmap'/>
      <source file='/var/run/kubevirt-ephemeral-disks/cloud-init-data/default/vmi-with-sidecar-hook-2/noCloud.iso' index='1'/>
      <backingStore/>
      <target dev='vdb' bus='virtio'/>
      <alias name='ua-cloudinitdisk'/>
      <address type='pci' domain='0x0000' bus='0x05' slot='0x00' function='0x0'/>
    </disk>
   [...]
</domain>
```

</details>