package api_test

import (
	"errors"

	"github.com/christianang/cfnet/api"
	"github.com/christianang/cfnet/api/fakes"
	"github.com/containernetworking/plugins/pkg/ns"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate counterfeiter -o ./fakes/net_ns.go --fake-name NetNS . netNS
type netNS interface {
	ns.NetNS
}

var _ = Describe("ContainerNetNSExec", func() {
	var (
		netns         *fakes.NetNS
		nsadapter     *fakes.NSAdapter
		runc          *fakes.RunC
		commandRunner *fakes.CommandRunner

		containerNetNSExec api.ContainerNetNSExec

		containerID string
		command     string
		commandArgs []string
	)

	Describe("Execute", func() {
		BeforeEach(func() {
			netns = &fakes.NetNS{}
			nsadapter = &fakes.NSAdapter{}
			commandRunner = &fakes.CommandRunner{}
			runc = &fakes.RunC{}

			containerNetNSExec = api.NewContainerNetNSExec(runc, nsadapter, commandRunner)

			netns.DoStub = func(f func(ns.NetNS) error) error {
				return f(netns)
			}
			runc.GetNetNSPathReturns("some/ns/path", nil)
			nsadapter.GetNSReturns(netns, nil)
			containerID = "some-container-id"
			command = "some-command"
			commandArgs = []string{"some-arg", "some-arg-2"}
		})

		It("executes a command within the network namespace of the container provided", func() {
			err := containerNetNSExec.Execute(containerID, command, commandArgs)
			Expect(err).NotTo(HaveOccurred())

			Expect(runc.GetNetNSPathCallCount()).To(Equal(1))
			Expect(runc.GetNetNSPathArgsForCall(0)).To(Equal(containerID))

			Expect(commandRunner.RunCallCount()).To(Equal(1))
			actualCommand, actualCommandArgs := commandRunner.RunArgsForCall(0)
			Expect(actualCommand).To(Equal(command))
			Expect(actualCommandArgs).To(Equal(commandArgs))

			Expect(netns.DoCallCount()).To(Equal(1))
		})

		Context("when an error occurs", func() {
			Context("when runc.GetNetNSPath fails", func() {
				BeforeEach(func() {
					runc.GetNetNSPathReturns("", errors.New("something bad happened"))
				})

				It("returns an error", func() {
					err := containerNetNSExec.Execute("", "", []string{})
					Expect(err).To(MatchError("failed to get net ns path: something bad happened"))
				})
			})

			Context("when ns.GetNS fails", func() {
				BeforeEach(func() {
					nsadapter.GetNSReturns(nil, errors.New("something bad happened"))
				})

				It("returns an error", func() {
					err := containerNetNSExec.Execute("", "", []string{})
					Expect(err).To(MatchError("failed to get net ns: something bad happened"))
				})
			})

			Context("when commandRunner.Run fails", func() {
				BeforeEach(func() {
					commandRunner.RunReturns(errors.New("something bad happened"))
				})

				It("returns an error", func() {
					err := containerNetNSExec.Execute("", "", []string{})
					Expect(err).To(MatchError("failed to run command in net ns: something bad happened"))
				})
			})
		})
	})
})
