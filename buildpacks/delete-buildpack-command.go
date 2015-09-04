package buildpacks

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
)

type CliDeleteBuildpackCommand struct {
	cliConnection plugin.CliConnection
	buildpackName string
}

func NewCliDeleteBuildpackCommand(cliConnection plugin.CliConnection, buildpackName string) (cmd CliDeleteBuildpackCommand) {
	cmd.cliConnection = cliConnection
	cmd.buildpackName = buildpackName
	return
}

func (cmd CliDeleteBuildpackCommand) Action() string {
	return DELETE
}

func (cmd CliDeleteBuildpackCommand) BuildpackName() string {
	return cmd.buildpackName
}

func (cmd CliDeleteBuildpackCommand) Execute() {
	args := make([]string, 3)
	args[0] = "delete-buildpack"
	args[1] = cmd.buildpackName
	args[2] = "-f"

	fmt.Println("Executing: cf ", args)
	_, err := cmd.cliConnection.CliCommand(args...)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
