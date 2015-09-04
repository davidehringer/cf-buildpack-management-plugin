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

		// TODO
		
	})

})
