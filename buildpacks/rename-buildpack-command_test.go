package buildpacks_test

import (
	. "github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/cli/plugin/fakes"
)

var _ = Describe("RenameBuildpackCommand", func() {

	var cliConnection *fakes.FakeCliConnection

	BeforeEach(func() {
		cliConnection = &fakes.FakeCliConnection{}
	})

	It("it calls the rename-buildpack comand", func() {
		command := NewCliRenameBuildpackCommand(cliConnection, "example-bp", "new-example-bp")
		command.Execute()

		Expect(cliConnection.CliCommandCallCount()).To(Equal(1))

		args := cliConnection.CliCommandArgsForCall(0)
		Expect(args).To(Equal([]string{"rename-buildpack", "example-bp", "new-example-bp"}))
	})

})
