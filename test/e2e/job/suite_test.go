package job

import (
	"testing"

	"k8s.io/kubernetes/test/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	framework.TestContext.KubeConfig = "/Users/manjunath/.kube/config"
}
func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}
