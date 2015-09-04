package buildpacks

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"strconv"
)

type CliAddBuildpackCommand struct {
	cliConnection plugin.CliConnection
	buildpack     Buildpack
}

func NewCliAddBuildpackCommand(cliConnection plugin.CliConnection, buildpack Buildpack) (cmd CliAddBuildpackCommand) {
	cmd.cliConnection = cliConnection
	cmd.buildpack = buildpack
	return
}

func (cmd CliAddBuildpackCommand) Action() string {
	return ADD
}

func (cmd CliAddBuildpackCommand) BuildpackName() string {
	return cmd.buildpack.Name
}

func (cmd CliAddBuildpackCommand) Execute() {
	args := make([]string, 5)
	args[0] = "create-buildpack"
	args[1] = cmd.buildpack.Name
	args[2] = cmd.buildpack.Filename
	args[3] = strconv.Itoa(cmd.buildpack.Position)
	if cmd.buildpack.Enabled {
		args[4] = "--enable"
	} else {
		args[4] = "--disable"
	}

	fmt.Println("Executing: cf ", args)
	_, err := cmd.cliConnection.CliCommand(args...)
	if err != nil {
		fmt.Println("Unable to add buildpack: ", err)
	}
}
