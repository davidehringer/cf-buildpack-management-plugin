package buildpacks

import (
	"github.com/cloudfoundry/cli/plugin"
	"sort"
)

type CommandGenerator struct {
	cliConnection plugin.CliConnection
}

func NewCommandGenerator(cliConnection plugin.CliConnection) (generator CommandGenerator) {
	generator.cliConnection = cliConnection
	return
}

func (generator CommandGenerator) GenerateCommands(actualState []Buildpack, desiredState []Buildpack) []BuildpackCommand {
	cmds := make([]BuildpackCommand, 0)
	cmds = append(cmds, generator.determineDeletesAndRenames(actualState, desiredState)...)
	cmds = append(cmds, generator.determineAdds(actualState, desiredState)...)
	cmds = append(cmds, generator.updateAll(desiredState)...)
	sort.Sort(byAction(cmds))
	return cmds
}

func (generator CommandGenerator) determineDeletesAndRenames(actualState []Buildpack, desiredState []Buildpack) (cmds []BuildpackCommand) {
	for _, actual := range actualState {
		delete := true
		for _, desired := range desiredState {
			if actual.Filename == desired.Filename {
				delete = false
				if actual.Name != desired.Name {
					cmds = append(cmds, NewCliRenameBuildpackCommand(generator.cliConnection, actual.Name, desired.Name))
				}
			}
		}
		if delete {
			cmds = append(cmds, NewCliDeleteBuildpackCommand(generator.cliConnection, actual.Name))
		}
	}
	return cmds
}

func (generator CommandGenerator) determineAdds(actualState []Buildpack, desiredState []Buildpack) (cmds []BuildpackCommand) {
	for _, desired := range desiredState {
		add := true
		for _, actual := range actualState {
			if actual.Filename == desired.Filename {
				add = false
			}
		}
		if add {
			cmds = append(cmds, NewCliAddBuildpackCommand(generator.cliConnection, desired))
		}
	}
	return cmds
}

func (generator CommandGenerator) updateAll(desiredState []Buildpack) (cmds []BuildpackCommand) {
	for _, desired := range desiredState {
		cmds = append(cmds, NewCliUpdateBuildpackCommand(generator.cliConnection, desired))
	}
	return cmds
}

type byAction []BuildpackCommand

func (a byAction) Len() int {
	return len(a)
}

func (a byAction) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAction) Less(i, j int) bool {
	return actionPriority(a[i].Action()) < actionPriority(a[j].Action())
}

func actionPriority(action string) int {
	switch action {
	case DELETE:
		return 1
	case RENAME:
		return 2
	case ADD:
		return 3
	case UPDATE:
		return 4
	}
	return 5
}
