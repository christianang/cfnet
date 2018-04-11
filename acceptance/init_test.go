package acceptance_test

import (
	"os/exec"
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance")
}

var pathToCFNet string

var _ = BeforeSuite(func() {
	var err error

	pathToCFNet, err = gexec.Build("github.com/christianang/cfnet/cmd/cfnet")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func cfnet(args ...string) *gexec.Session {
	cmd := exec.Command(pathToCFNet, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
