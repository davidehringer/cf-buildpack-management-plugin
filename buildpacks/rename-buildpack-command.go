package buildpacks

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
)

type CliRenameBuildpackCommand struct {
	cliConnection plugin.CliConnection
	oldName       string
	newName       string
}

func NewCliRenameBuildpackCommand(cliConnection plugin.CliConnection, oldName string, newName string) (cmd CliRenameBuildpackCommand) {
	cmd.cliConnection = cliConnection
	cmd.oldName = oldName
	cmd.newName = newName
	return
}

func (cmd CliRenameBuildpackCommand) Action() string {
	return RENAME
}

func (cmd CliRenameBuildpackCommand) BuildpackName() string {
	return cmd.newName
}

func (cmd CliRenameBuildpackCommand) Execute() {
	args := make([]string, 3)
	args[0] = "rename-buildpack"
	args[1] = cmd.oldName
	args[2] = cmd.newName

	fmt.Println("Executing: cf ", args)
	_, err := cmd.cliConnection.CliCommand(args...)
	if err != nil {
		fmt.Println("Unable to rename buildpack: ", err)
	}
}
