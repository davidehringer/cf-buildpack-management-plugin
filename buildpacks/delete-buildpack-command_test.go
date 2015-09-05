package buildpacks_test

import (
	. "github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/cli/plugin/fakes"
)

var _ = Describe("DeleteBuildpackCommand", func() {

	var cliConnection *fakes.FakeCliConnection

	BeforeEach(func() {
		cliConnection = &fakes.FakeCliConnection{}
	})

	It("it calls the delete-buildpack comand with the 'force' flag", func() {
		command := NewCliDeleteBuildpackCommand(cliConnection, "example-bp")
		command.Execute()

		Expect(cliConnection.CliCommandCallCount()).To(Equal(1))

		args := cliConnection.CliCommandArgsForCall(0)
		Expect(args).To(Equal([]string{"delete-buildpack", "example-bp", "-f"}))
	})

})
