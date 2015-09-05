package buildpacks_test

import (
	. "github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/cli/plugin/fakes"
)

var _ = Describe("UPdateBuildpackCommand", func() {

	var cliConnection *fakes.FakeCliConnection

	BeforeEach(func() {
		cliConnection = &fakes.FakeCliConnection{}
	})

	var _ = Describe("it calls the update-buildpack comand", func() {
		It("to enable and unlock a buildpack", func() {

			buildpack := Buildpack{
				Name:     "example-bp",
				Position: 2,
				Enabled:  true,
				Locked:   false,
				Filename: "example.zip",
			}
			command := NewCliUpdateBuildpackCommand(cliConnection, buildpack)
			command.Execute()

			Expect(cliConnection.CliCommandCallCount()).To(Equal(1))

			args := cliConnection.CliCommandArgsForCall(0)
			Expect(args).To(Equal([]string{"update-buildpack", buildpack.Name, "-i", "2", "--enable", "--unlock"}))
		})
		It("to disable and lock a buildpack", func() {

			buildpack := Buildpack{
				Name:     "example-bp",
				Position: 3,
				Enabled:  false,
				Locked:   true,
				Filename: "example.zip",
			}
			command := NewCliUpdateBuildpackCommand(cliConnection, buildpack)
			command.Execute()

			Expect(cliConnection.CliCommandCallCount()).To(Equal(1))

			args := cliConnection.CliCommandArgsForCall(0)
			Expect(args).To(Equal([]string{"update-buildpack", buildpack.Name, "-i", "3", "--disable", "--lock"}))
		})
	})
})
