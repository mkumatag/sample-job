package job

import (
	"flag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/test/e2e/framework"
	frameworkconfig "k8s.io/kubernetes/test/e2e/framework/config"
	"testing"
)

// handleFlags sets up all flags and parses the command line.
func handleFlags() {
	frameworkconfig.CopyFlags(frameworkconfig.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	flag.Parse()
	framework.AfterReadingAllFlags(&framework.TestContext)
}
func init() {
	testing.Init()
	handleFlags()
	if framework.TestContext.KubeConfig == "" {
		klog.Infof("Couldnt get kubeconfig. set environment variable KUBECONFIG,Path to kubeconfig containing embedded authinfo")
		klog.Infof("Will try to use in-cluster config if KUBECONFIG is not available")
	}
}
func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}
