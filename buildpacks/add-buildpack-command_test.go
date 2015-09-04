package buildpacks_test

import (
	//. "github.com/davidehringer/cf-buildpack-management-plugin"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	"github.com/cloudfoundry/cli/plugin/fakes"
)

var _ = Describe("AddBuildpackCommand", func() {

	var cliConnection *fakes.FakeCliConnection

	BeforeEach(func() {
		cliConnection = &fakes.FakeCliConnection{}
	})

	It("it calls the add-buildpack comand", func() {

		//cf create-buildpack BUILDPACK PATH POSITION [--enable|--disable]
		// cliConnection.CliCommandReturns([]string{ "" }, nil)

		// repo := NewCliBuildpackRepository(cliConnection)
		// buildpacks := repo.ListBuildpacks()

		// Expect(buildpacks[0].Name).To(Equal("java"))
		// Expect(buildpacks[0].Position).To(Equal(1))
		// Expect(buildpacks[0].Enabled).To(Equal(true))
		// Expect(buildpacks[0].Locked).To(Equal(false))
		// Expect(buildpacks[0].Filename).To(Equal("java-buildpack-offline-v3.0.zip"))

		// Expect(len(buildpacks)).To(Equal(10))
	})

})
