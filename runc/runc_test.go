package runc_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/christianang/cfnet/runc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("runc", func() {
	var (
		runcRoot    string
		containerID string

		runC runc.RunC
	)

	BeforeEach(func() {
		var err error
		runcRoot, err = ioutil.TempDir("", "")
		containerID = "some-container-id"

		err = os.Mkdir(filepath.Join(runcRoot, containerID), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())

		containerStateJson := []byte(`{ "namespace_paths": { "NEWNET": "some-ns-namespace-path" } }`)
		err = ioutil.WriteFile(filepath.Join(runcRoot, containerID, "state.json"), containerStateJson, os.ModePerm)
		Expect(err).NotTo(HaveOccurred())

		runC = runc.NewRunC(runcRoot)
	})

	Describe("GetNetNSPath", func() {
		It("returns the ns path provided a container id", func() {
			actualNamespacePath, err := runC.GetNetNSPath(containerID)
			Expect(err).NotTo(HaveOccurred())

			Expect(actualNamespacePath).To(Equal("some-ns-namespace-path"))
		})

		Context("when an error occurs", func() {
			Context("when ioutil.ReadFile fails", func() {
				It("returns an error", func() {
					_, err := runC.GetNetNSPath("non-existant-container")
					Expect(err).To(MatchError(fmt.Sprintf("failed to open runc state.json: open %s: no such file or directory", filepath.Join(runcRoot, "non-existant-container", "state.json"))))
				})
			})

			Context("when json.Unmarshal fails", func() {
				BeforeEach(func() {
					err := os.Mkdir(filepath.Join(runcRoot, "malformed"), os.ModePerm)
					Expect(err).NotTo(HaveOccurred())

					containerStateJson := []byte(`malformed-json`)
					err = ioutil.WriteFile(filepath.Join(runcRoot, "malformed", "state.json"), containerStateJson, os.ModePerm)
					Expect(err).NotTo(HaveOccurred())
				})

				It("returns an error", func() {
					_, err := runC.GetNetNSPath("malformed")
					Expect(err).To(MatchError("failed to unmarshal json: invalid character 'm' looking for beginning of value"))
				})
			})
		})
	})
})
