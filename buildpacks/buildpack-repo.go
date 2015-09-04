package buildpacks

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"strings"
)

type BuildpackRepository interface {
	ListBuildpacks() ([]Buildpack, error)
}

type CliBuildpackRepository struct {
	cliConnection plugin.CliConnection
}

func NewCliBuildpackRepository(cliConnection plugin.CliConnection) (repo CliBuildpackRepository) {
	repo.cliConnection = cliConnection
	return
}

func (repo CliBuildpackRepository) ListBuildpacks() ([]Buildpack, error) {

	output, err := repo.cliConnection.CliCommandWithoutTerminalOutput("curl", "/v2/buildpacks")

	if err != nil {
		fmt.Println("PLUGIN ERROR: Error from CliCommand: ", err)
		return nil, err
	}

	data := strings.Join(output, "\n")
	var r result
	json.Unmarshal([]byte(data), &r)

	buildpacks := make([]Buildpack, 0)
	for _, b := range r.Resources {
		var bp Buildpack
		bp.Name = b.Entity.Name
		bp.Position = b.Entity.Position
		bp.Enabled = b.Entity.Enabled
		bp.Locked = b.Entity.Locked
		bp.Filename = b.Entity.Filename
		buildpacks = append(buildpacks, bp)
	}
	return buildpacks, nil
}

type result struct {
	Resources []resource
}

type resource struct {
	Entity *buildpackResource
}

type buildpackResource struct {
	Name     string
	Position int
	Enabled  bool
	Locked   bool
	Filename string
}
