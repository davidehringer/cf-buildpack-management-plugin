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

			Expect(buildpacks[1].Name).To(Equal("ruby_buildpack"))
			Expect(buildpacks[1].Position).To(Equal(2))
			Expect(buildpacks[1].Enabled).To(Equal(true))
			Expect(buildpacks[1].Locked).To(Equal(false))
			Expect(buildpacks[1].Filename).To(Equal("../fixtures/buildpacks/example-2.zip"))
		})

		It("returns an error if the file doens't exist", func() {

			buildpacks, err := repo.ReadManifest("../fixtures/buildpacks/does-not-exist.yml")

			Expect(err).To(HaveOccurred())
			Expect(buildpacks).To(BeNil())
		})

	})
})
