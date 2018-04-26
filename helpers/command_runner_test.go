package helpers_test

import (
	"bytes"

	"github.com/christianang/cfnet/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CommandRunner", func() {
	var (
		commandRunner helpers.CommandRunner

		stdout *bytes.Buffer
		stderr *bytes.Buffer
	)

	Describe("Run", func() {
		BeforeEach(func() {
			stdout = bytes.NewBuffer([]byte{})
			stderr = bytes.NewBuffer([]byte{})

			commandRunner = helpers.NewCommandRunner(stdout, stderr)
		})

		It("executes a process", func() {
			err := commandRunner.Run("echo", []string{"hello"})
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("hello"))
		})

		Context("when an error occurs", func() {
			Context("when cmd.Run fails", func() {
				It("returns an error", func() {
					err := commandRunner.Run("false", []string{})
					Expect(err).To(MatchError("exit status 1"))
				})
			})
		})
	})
})
