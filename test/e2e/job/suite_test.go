package job 
import (
	"flag"
	"testing"
	 "k8s.io/klog/v2"
	"k8s.io/kubernetes/test/e2e/framework"
	frameworkconfig "k8s.io/kubernetes/test/e2e/framework/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
// handleFlags sets up all flags and parses the command line.
func handleFlags() {
        frameworkconfig.CopyFlags(frameworkconfig.Flags, flag.CommandLine)
        framework.RegisterCommonFlags(flag.CommandLine)
        framework.RegisterClusterFlags(flag.CommandLine)
        flag.Parse()
	 if framework.TestContext.KubeConfig == ""{
                klog.Infof("Couldnt get kubeconfig. set environment variable KUBECONFIG,Path to kubeconfig containing embedded authinfo")
                klog.Infof("Will try to use in-cluster config if KUBECONFIG is not available")
        }

        framework.AfterReadingAllFlags(&framework.TestContext)
}
func init() {
	testing.Init()
	handleFlags()
}
func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}
