package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KubevirtHookVirtXml", func() {
	testDomainXML := `<domain type='kvm'>
  <name>default_vmi-lun</name>
  <uuid>db81dffe-76c9-4b13-8903-76b7fce0dae7</uuid>
  <metadata>
    <kubevirt xmlns="http://kubevirt.io">
      <uid>6b78f958-efc7-468c-badd-c7fab7f7f0f9</uid>
      <graceperiod>
        <deletionGracePeriodSeconds>0</deletionGracePeriodSeconds>
      </graceperiod>
    </kubevirt>
  </metadata>
  <memory unit='KiB'>131072</memory>
  <currentMemory unit='KiB'>131072</currentMemory>
  <vcpu placement='static'>1</vcpu>
  <iothreads>1</iothreads>
  <sysinfo type='smbios'>
    <system>
      <entry name='manufacturer'>KubeVirt</entry>
      <entry name='product'>None</entry>
      <entry name='uuid'>db81dffe-76c9-4b13-8903-76b7fce0dae7</entry>
      <entry name='family'>KubeVirt</entry>
    </system>
  </sysinfo>
  <os>
    <type arch='x86_64' machine='pc-q35-rhel8.4.0'>hvm</type>
    <boot dev='hd'/>
    <smbios mode='sysinfo'/>
  </os>
  <features>
    <acpi/>
  </features>
  <cpu mode='host-model' check='partial'>
    <topology sockets='1' dies='1' cores='1' threads='1'/>
  </cpu>
  <clock offset='utc'/>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <devices>
    <emulator>/usr/libexec/qemu-kvm</emulator>
    <disk type='file' device='disk' model='virtio-non-transitional'>
      <driver name='qemu' type='qcow2' cache='none' error_policy='stop' discard='unmap'/>
      <source file='/var/run/kubevirt-ephemeral-disks/disk-data/containerdisk/disk.qcow2'/>
      <backingStore type='file'>
        <format type='qcow2'/>
        <source file='/var/run/kubevirt/container-disks/disk_0.img'/>
      </backingStore>
      <target dev='vda' bus='virtio'/>
      <alias name='ua-containerdisk'/>
      <address type='pci' domain='0x0000' bus='0x04' slot='0x00' function='0x0'/>
      <iotune>
        <total_bytes_sec>52428800</total_bytes_sec>
        <total_iops_sec>1000</total_iops_sec>
      </iotune>
    </disk>
    <disk type='file' device='disk' model='virtio-non-transitional'>
      <driver name='qemu' type='raw' cache='none' error_policy='stop' discard='unmap'/>
      <source file='/var/run/kubevirt-ephemeral-disks/cloud-init-data/default/vmi-lun/noCloud.iso'/>
      <target dev='vdb' bus='virtio'/>
      <alias name='ua-cloudinitdisk'/>
      <address type='pci' domain='0x0000' bus='0x05' slot='0x00' function='0x0'/>
    </disk>
    <disk type='block' device='lun'>
      <driver name='qemu' type='raw' cache='none' error_policy='stop' io='native' discard='unmap'/>
      <source dev='/dev/blockpvcdisk'/>
      <target dev='sda' bus='scsi'/>
      <alias name='ua-blockpvcdisk'/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='none'/>
    <controller type='scsi' index='0' model='virtio-non-transitional'>
      <address type='pci' domain='0x0000' bus='0x02' slot='0x00' function='0x0'/>
    </controller>
    <controller type='virtio-serial' index='0' model='virtio-non-transitional'>
      <address type='pci' domain='0x0000' bus='0x03' slot='0x00' function='0x0'/>
    </controller>
    <controller type='sata' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x1f' function='0x2'/>
    </controller>
    <controller type='pci' index='0' model='pcie-root'/>
    <controller type='pci' index='1' model='pcie-root-port'>
      <model name='pcie-root-port'/>
      <target chassis='1' port='0x10'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0' multifunction='on'/>
    </controller>
    <controller type='pci' index='2' model='pcie-root-port'>
      <model name='pcie-root-port'/>
      <target chassis='2' port='0x11'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x1'/>
    </controller>
    <controller type='pci' index='3' model='pcie-root-port'>
      <model name='pcie-root-port'/>
      <target chassis='3' port='0x12'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x2'/>
    </controller>
    <controller type='pci' index='4' model='pcie-root-port'>
      <model name='pcie-root-port'/>
      <target chassis='4' port='0x13'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x3'/>
    </controller>
    <controller type='pci' index='5' model='pcie-root-port'>
      <model name='pcie-root-port'/>
      <target chassis='5' port='0x14'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x4'/>
    </controller>
    <controller type='pci' index='6' model='pcie-root-port'>
      <model name='pcie-root-port'/>
      <target chassis='6' port='0x15'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x5'/>
    </controller>
    <controller type='pci' index='7' model='pcie-root-port'>
      <model name='pcie-root-port'/>
      <target chassis='7' port='0x16'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x6'/>
    </controller>
    <interface type='ethernet'>
      <mac address='02:76:4a:44:00:a6'/>
      <target dev='tap0' managed='no'/>
      <model type='virtio-non-transitional'/>
      <mtu size='1480'/>
      <alias name='ua-default'/>
      <rom enabled='no'/>
      <address type='pci' domain='0x0000' bus='0x01' slot='0x00' function='0x0'/>
    </interface>
    <serial type='unix'>
      <source mode='bind' path='/var/run/kubevirt-private/6b78f958-efc7-468c-badd-c7fab7f7f0f9/virt-serial0'/>
      <target type='isa-serial' port='0'>
        <model name='isa-serial'/>
      </target>
    </serial>
    <console type='unix'>
      <source mode='bind' path='/var/run/kubevirt-private/6b78f958-efc7-468c-badd-c7fab7f7f0f9/virt-serial0'/>
      <target type='serial' port='0'/>
    </console>
    <channel type='unix'>
      <target type='virtio' name='org.qemu.guest_agent.0'/>
      <address type='virtio-serial' controller='0' bus='0' port='1'/>
    </channel>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='vnc' socket='/var/run/kubevirt-private/6b78f958-efc7-468c-badd-c7fab7f7f0f9/virt-vnc'>
      <listen type='socket' socket='/var/run/kubevirt-private/6b78f958-efc7-468c-badd-c7fab7f7f0f9/virt-vnc'/>
    </graphics>
    <audio id='1' type='none'/>
    <video>
      <model type='vga' vram='16384' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x0'/>
    </video>
    <memballoon model='virtio-non-transitional'>
      <stats period='10'/>
      <address type='pci' domain='0x0000' bus='0x06' slot='0x00' function='0x0'/>
    </memballoon>
  </devices>
</domain>
`
	Context("on conversion attempt", func() {
		// It("should set vcpus value", func() {

		// 	xml, err := MergeKubeVirtXMLWithProvidedXML([]byte(testDomainXML), []string{"--vcpus=10"})
		// 	Expect(err).To(BeNil())
		// 	Expect(string(xml)).Should(ContainSubstring(`<vcpu placement="static">10</vcpu>`))
		// })
		// It("should set smbios", func() {

		// 	xml, err := MergeKubeVirtXMLWithProvidedXML([]byte(testDomainXML), []string{"--sysinfo=bios.vendor=MyVendor,bios.version=1.2.3"})
		// 	Expect(err).To(BeNil())
		// 	Expect(string(xml)).Should(ContainSubstring(`<entry name="vendor">MyVendor</entry>`))
		// 	Expect(string(xml)).Should(ContainSubstring(`<entry name="version">1.2.3</entry>`))
		// })
		It("should set iotune values", func() {

			xml, err := AddBlkiotune([]byte(testDomainXML), []string{"--disk=iotune.total_bytes_sec=52428800", "--edit=all"})
			Expect(err).To(BeNil())
			Expect(string(xml)).Should(ContainSubstring(`<total_bytes_sec>52428800</total_bytes_sec>`))
		})
		// It("should fail with multiple changing options", func() {

		// 	_, err := MergeKubeVirtXMLWithProvidedXML([]byte(testDomainXML), []string{"--sysinfo=bios.vendor=MyVendor,bios.version=1.2.3", "--vcpus=10"})
		// 	Expect(err).Should(HaveOccurred())
		// })
	})
})
