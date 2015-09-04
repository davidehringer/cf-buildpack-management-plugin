package buildpacks

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
)

type Buildpack struct {
	Name     string
	Position int
	Enabled  bool
	Locked   bool
	Filename string
}

const (
	ADD    = "add"
	DELETE = "delete"
	RENAME = "rename"
	UPDATE = "update"
)

type BuildpackCommand interface {
	Action() string
	BuildpackName() string
	Execute()
}

type BuildpackService interface {
	ConfigureBuildpacks(desiredState []Buildpack, dryRun bool) error
}

type CliBuildpackService struct {
	cliConnection plugin.CliConnection
	repo          BuildpackRepository
}

func NewCliBuildpackService(cliConnection plugin.CliConnection, repo BuildpackRepository) (service CliBuildpackService) {
	service.cliConnection = cliConnection
	service.repo = repo
	return
}

func (service CliBuildpackService) ConfigureBuildpacks(desiredState []Buildpack, dryRun bool) error {
	actualState, err := service.repo.ListBuildpacks()
	if err != nil {
		return err
	}
	commandGenerator := NewCommandGenerator(service.cliConnection)
	commands := commandGenerator.GenerateCommands(actualState, desiredState)
	fmt.Println("The following actions are required to configure buildpacks:")
	for _, command := range commands {
		fmt.Printf("\t- %v buildpack named %v\n", command.Action(), command.BuildpackName())
	}

	if dryRun {
		fmt.Println("Dry run mode. No actions will be executed.")
		return nil
	} else {
		for _, command := range commands {
			command.Execute()
		}
	}

	return nil
}
