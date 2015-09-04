package buildpacks_test

import (
	. "github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuildpackManifestRepo", func() {
	Describe("ReadManifest", func() {

		var repo BuildpackManifestRepo

		BeforeEach(func() {
			repo = NewFilesystemBuildpackManifestRepo()
		})

		It("reads the manifest from a file", func() {

			buildpacks, _ := repo.ReadManifest("../fixtures/buildpacks/buildpack-config.yml")

			Expect(len(buildpacks)).To(Equal(3))

			Expect(buildpacks[1].Filename).To(Equal("../fixtures/buildpacks/example-2.zip"))
			// TODO additional assertions
		})

		// TODO test invalid file or directory
	})
})
