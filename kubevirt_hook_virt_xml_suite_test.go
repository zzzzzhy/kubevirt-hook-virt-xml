package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKubevirtHookVirtXml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KubevirtHookVirtXml Suite")
}
