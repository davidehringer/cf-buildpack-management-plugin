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

		It("defaults 'enabled' to true and 'locked' to false if not provided", func() {

			buildpacks, _ := repo.ReadManifest("../fixtures/buildpacks/buildpack-config-defaults.yml")

			Expect(len(buildpacks)).To(Equal(3))

			Expect(buildpacks[0].Name).To(Equal("java"))
			Expect(buildpacks[0].Position).To(Equal(1))
			Expect(buildpacks[0].Enabled).To(Equal(false))
			Expect(buildpacks[0].Locked).To(Equal(true))
			Expect(buildpacks[0].Filename).To(Equal("../fixtures/buildpacks/example-1.zip"))

			Expect(buildpacks[1].Name).To(Equal("ruby_buildpack"))
			Expect(buildpacks[1].Position).To(Equal(2))
			Expect(buildpacks[1].Enabled).To(Equal(true))
			Expect(buildpacks[1].Locked).To(Equal(false))
			Expect(buildpacks[1].Filename).To(Equal("../fixtures/buildpacks/example-2.zip"))

			Expect(buildpacks[2].Name).To(Equal("nodejs_buildpack"))
			Expect(buildpacks[2].Position).To(Equal(3))
			Expect(buildpacks[2].Enabled).To(Equal(true))
			Expect(buildpacks[2].Locked).To(Equal(false))
			Expect(buildpacks[2].Filename).To(Equal("../fixtures/buildpacks/example-3.zip"))
		})

	})
})
