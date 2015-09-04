package buildpacks_test

import (
	. "github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/cli/plugin/fakes"
)

var getBuildpacksJson = `
 {
   "total_results": 10,
   "total_pages": 1,
   "prev_url": null,
   "next_url": null,
   "resources": [
      {
         "metadata": {
            "guid": "85fbbaa3-4d40-413f-85e0-982a74529a9d",
            "url": "/v2/buildpacks/85fbbaa3-4d40-413f-85e0-982a74529a9d",
            "created_at": "2015-04-14T18:57:46Z",
            "updated_at": "2015-04-14T18:58:26Z"
         },
         "entity": {
            "name": "java",
            "position": 1,
            "enabled": true,
            "locked": false,
            "filename": "java-buildpack-offline-v3.0.zip"
         }
      },
      {
         "metadata": {
            "guid": "3aeabd21-f05c-4438-92a9-d5fc8d810e76",
            "url": "/v2/buildpacks/3aeabd21-f05c-4438-92a9-d5fc8d810e76",
            "created_at": "2015-03-12T19:06:51Z",
            "updated_at": "2015-04-06T02:30:47Z"
         },
         "entity": {
            "name": "java_7",
            "position": 2,
            "enabled": true,
            "locked": false,
            "filename": "java-buildpack-offline-v2.4.zip"
         }
      },
      {
         "metadata": {
            "guid": "116360af-4e4a-42e3-99f8-7ee0cee3ee59",
            "url": "/v2/buildpacks/116360af-4e4a-42e3-99f8-7ee0cee3ee59",
            "created_at": "2015-03-12T19:06:56Z",
            "updated_at": "2015-05-03T19:07:03Z"
         },
         "entity": {
            "name": "ruby_buildpack",
            "position": 3,
            "enabled": true,
            "locked": false,
            "filename": "ruby_buildpack-cached-v1.3.0.zip"
         }
      },
      {
         "metadata": {
            "guid": "bb3ecf11-2461-4c84-a6cb-8f78d5a31ce5",
            "url": "/v2/buildpacks/bb3ecf11-2461-4c84-a6cb-8f78d5a31ce5",
            "created_at": "2015-03-12T19:06:56Z",
            "updated_at": "2015-05-03T19:06:56Z"
         },
         "entity": {
            "name": "nodejs_buildpack",
            "position": 4,
            "enabled": true,
            "locked": false,
            "filename": "nodejs_buildpack-cached-v1.2.0.zip"
         }
      },
      {
         "metadata": {
            "guid": "13020a7e-27fe-4008-a37e-5f5044b2516f",
            "url": "/v2/buildpacks/13020a7e-27fe-4008-a37e-5f5044b2516f",
            "created_at": "2015-03-12T19:07:20Z",
            "updated_at": "2015-05-03T19:07:12Z"
         },
         "entity": {
            "name": "go_buildpack",
            "position": 5,
            "enabled": true,
            "locked": false,
            "filename": "go_buildpack-cached-v1.2.0.zip"
         }
      },
      {
         "metadata": {
            "guid": "caa3de4c-d03f-4b5d-8a9e-ddeab221ba24",
            "url": "/v2/buildpacks/caa3de4c-d03f-4b5d-8a9e-ddeab221ba24",
            "created_at": "2015-03-12T19:07:32Z",
            "updated_at": "2015-05-03T19:07:24Z"
         },
         "entity": {
            "name": "python_buildpack",
            "position": 6,
            "enabled": true,
            "locked": false,
            "filename": "python_buildpack-cached-v1.2.0.zip"
         }
      },
      {
         "metadata": {
            "guid": "87bb96c0-e3fc-42e5-bfc4-654d8b3da9f8",
            "url": "/v2/buildpacks/87bb96c0-e3fc-42e5-bfc4-654d8b3da9f8",
            "created_at": "2015-03-12T19:07:39Z",
            "updated_at": "2015-05-03T19:07:47Z"
         },
         "entity": {
            "name": "php_buildpack",
            "position": 7,
            "enabled": true,
            "locked": false,
            "filename": "php_buildpack-offline-v3.1.0.zip"
         }
      },
      {
         "metadata": {
            "guid": "6c383df8-e36f-4c31-9be3-5ba3406ce717",
            "url": "/v2/buildpacks/6c383df8-e36f-4c31-9be3-5ba3406ce717",
            "created_at": "2015-05-29T11:56:12Z",
            "updated_at": "2015-05-29T12:05:55Z"
         },
         "entity": {
            "name": "java_unlimited_crypto",
            "position": 8,
            "enabled": true,
            "locked": false,
            "filename": "java-buildpack-offline-unlimited-crypto-3.0.zip"
         }
      },
      {
         "metadata": {
            "guid": "e001c591-55ed-4300-a2a0-b9d37b4640e3",
            "url": "/v2/buildpacks/e001c591-55ed-4300-a2a0-b9d37b4640e3",
            "created_at": "2015-06-30T15:09:49Z",
            "updated_at": "2015-06-30T15:09:56Z"
         },
         "entity": {
            "name": "java_buildpack_offline",
            "position": 9,
            "enabled": true,
            "locked": false,
            "filename": "java-buildpack-offline-v2.7.1.zip"
         }
      },
      {
         "metadata": {
            "guid": "9d719acf-961c-4d05-b569-750c01051e2b",
            "url": "/v2/buildpacks/9d719acf-961c-4d05-b569-750c01051e2b",
            "created_at": "2015-07-09T17:00:40Z",
            "updated_at": "2015-07-09T17:00:44Z"
         },
         "entity": {
            "name": "staticfile-buildpack",
            "position": 10,
            "enabled": true,
            "locked": false,
            "filename": "staticfile_buildpack-cached-v1.2.0.zip"
         }
      }
   ]
}`

var _ = Describe("BuildpackRepo", func() {
	Describe("ListAll", func() {
		var cliConnection *fakes.FakeCliConnection

		BeforeEach(func() {
			cliConnection = &fakes.FakeCliConnection{}
		})

		It("outputs a slice of all admin buildpacks", func() {
			cliConnection.CliCommandWithoutTerminalOutputReturns([]string{getBuildpacksJson}, nil)

			repo := NewCliBuildpackRepository(cliConnection)
			buildpacks, _ := repo.ListBuildpacks()

			Expect(buildpacks[0].Name).To(Equal("java"))
			Expect(buildpacks[0].Position).To(Equal(1))
			Expect(buildpacks[0].Enabled).To(Equal(true))
			Expect(buildpacks[0].Locked).To(Equal(false))
			Expect(buildpacks[0].Filename).To(Equal("java-buildpack-offline-v3.0.zip"))

			Expect(len(buildpacks)).To(Equal(10))
		})
	})
})
