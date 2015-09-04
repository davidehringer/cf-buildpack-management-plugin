package buildpacks

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"strconv"
)

type CliUpdateBuildpackCommand struct {
	cliConnection plugin.CliConnection
	buildpack     Buildpack
}

func NewCliUpdateBuildpackCommand(cliConnection plugin.CliConnection, buildpack Buildpack) (cmd CliUpdateBuildpackCommand) {
	cmd.cliConnection = cliConnection
	cmd.buildpack = buildpack
	return
}

func (cmd CliUpdateBuildpackCommand) Action() string {
	return UPDATE
}

func (cmd CliUpdateBuildpackCommand) BuildpackName() string {
	return cmd.buildpack.Name
}

func (cmd CliUpdateBuildpackCommand) Execute() {
	args := make([]string, 6)
	args[0] = "update-buildpack"
	args[1] = cmd.buildpack.Name
	args[2] = "-i"
	args[3] = strconv.Itoa(cmd.buildpack.Position)
	if cmd.buildpack.Enabled {
		args[4] = "--enable"
	} else {
		args[4] = "--disable"
	}
	if cmd.buildpack.Locked {
		args[5] = "--lock"
	} else {
		args[5] = "--unlock"
	}

	fmt.Println("Executing: cf ", args)
	_, err := cmd.cliConnection.CliCommand(args...)
	if err != nil {
		fmt.Println("Unable to update buildpack: ", err)
	}
}
