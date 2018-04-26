package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/onsi/gomega/gexec"

	"github.com/containernetworking/plugins/pkg/ns"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("exec", func() {
	var (
		runcRoot         string
		containerID      string
		networkNamespace ns.NetNS
	)

	BeforeEach(func() {
		var err error
		runcRoot, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		containerID = "some-container"
		err = os.Mkdir(filepath.Join(runcRoot, containerID), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())

		networkNamespace, err = ns.NewNS()
		Expect(err).NotTo(HaveOccurred())

		containerStateJson := []byte(fmt.Sprintf(`{ "namespace_paths": { "NEWNET": %q } }`, networkNamespace.Path()))
		err = ioutil.WriteFile(filepath.Join(runcRoot, containerID, "state.json"), containerStateJson, os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(networkNamespace.Close()).To(Succeed())
	})

	It("executes a command in the network namespace and returns the results", func() {
		session := cfnet("--runc-root", runcRoot, "exec", containerID, "ip", "addr", "show")
		Eventually(session).Should(gexec.Exit(0))
		Expect(string(session.Out.Contents())).To(ContainSubstring("1: lo: <LOOPBACK>"))
	})
})
