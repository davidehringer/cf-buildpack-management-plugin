package main

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"
	"strings"
	"fmt"
	"os"
)

type BuildpackManager struct{}

func (c *BuildpackManager) GetMetadata() plugin.PluginMetadata {
	primaryUsage := "cf configure-buildpacks PATH_TO_YAML_CONFIG_FILE [-dryRun]"
	secondaryUsage := `   The provided path for the configuration file can be an absolute or relative path to 
   a file.  The file should have a map named "buildpacks" containing an array of buildpacks. The "filename" 
   values for each buildpack can be an absolute or relative path to a file.

   Valid YAML file example:
   ---
   buildpacks:
   - name: java
     position: 1
     enabled: true
     locked: false
     filename: java-buildpack-offline-v3.0.zip
   - name: ruby_buildpack
     position: 2
     enabled: true
     locked: false
     filename: ruby_buildpack-cached-v1.3.0.zip
    `
    flags := make(map[string]string)
    flags["dryRun"] = "stop before making any changes"

	return plugin.PluginMetadata{
		Name: "buildpack-management",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 9,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "configure-buildpacks",
				HelpText: "Configures system buildpacks using a declarative configuration file.",
				UsageDetails: plugin.Usage{
					Usage: strings.Join([]string{primaryUsage, secondaryUsage}, "\n\n"),
					Options: flags,
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BuildpackManager))
}

func (c *BuildpackManager) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "configure-buildpacks" {
		if len(args) < 2 {
			// TODO this is useless validation. Make useful
			fmt.Println("Incorrect Usage. \n\nPATH_TO_YAML_CONFIG_FILE is a required argument")
			os.Exit(1)
		}

		manifestRepo := buildpacks.NewFilesystemBuildpackManifestRepo()
		buildpackRepo := buildpacks.NewCliBuildpackRepository(cliConnection)

		config, err := manifestRepo.ReadManifest(args[1])
		if(err != nil){
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		service := buildpacks.NewCliBuildpackService(cliConnection, buildpackRepo)
		service.ConfigureBuildpacks(config, dryRun(args))
	}
}

func dryRun(args []string) (dryRun bool) {
	for _, arg := range args {
		if(arg == "--dryRun") {
			dryRun = true
		}
	}
	return 
}

type Result struct {
	Resources []Resource
}

type Resource struct {
	Entity *Buildpack
}

type Buildpack struct {
	Name     string
	Position int
	Enabled  bool
	Filename string
}
