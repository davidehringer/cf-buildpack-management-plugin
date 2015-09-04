package buildpacks_test

import (
	. "github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/cli/plugin/fakes"
)

var _ = Describe("AddBuildpackCommand", func() {

	var cliConnection *fakes.FakeCliConnection

	BeforeEach(func() {
		cliConnection = &fakes.FakeCliConnection{}
	})

	It("it calls the add-buildpack comand", func() {

		buildpack := Buildpack{
			Name: "example-bp",
			Position: 2,
			Enabled: true,
			Locked: false,
			Filename: "example.zip",
		}
		command := NewCliAddBuildpackCommand(cliConnection, buildpack)
		command.Execute()

		Expect(cliConnection.CliCommandCallCount()).To(Equal(1))

		args := cliConnection.CliCommandArgsForCall(0)
		Expect(args).To(Equal([]string{"create-buildpack", "example-bp", "example.zip", "2", "--enable"}))
	})

})
