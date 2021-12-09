package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KubevirtHookVirtXml", func() {
	Context("on conversion attempt", func() {
		It("should convert valid input", func() {
			testDomainXML := `<domain type='kvm'>
  <vcpu placement='static'>1</vcpu>
</domain>
`
			xml, err := MergeKubeVirtXMLWithProvidedXML([]byte(testDomainXML), []string{"--vcpus 10"})
			Expect(err).To(BeNil())
			Expect(string(xml)).Should(ContainSubstring("<vcpu placement='static'>10</vcpu>"))
		})
	})
})
