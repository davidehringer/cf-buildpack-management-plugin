package buildpacks

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type BuildpackManifestRepo interface {
	ReadManifest(string) ([]Buildpack, error)
}

type FilesystemBuildpackManifestRepo struct {
}

func NewFilesystemBuildpackManifestRepo() (repo FilesystemBuildpackManifestRepo) {
	return FilesystemBuildpackManifestRepo{}
}

func (repo FilesystemBuildpackManifestRepo) ReadManifest(path string) ([]Buildpack, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buildpacks, err := parseManifest(file)
	err = validateFiles(buildpacks)
	if err != nil {
		return nil, err
	}

	return buildpacks, err
}

func parseManifest(file io.Reader) (buildpacks []Buildpack, err error) {
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(contents, &m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for _, buildpack := range m["buildpacks"].([]interface{}) {
		enabled := true
		if buildpack.(map[interface{}]interface{})["enabled"] != nil {
			enabled = buildpack.(map[interface{}]interface{})["enabled"].(bool)
		}
		locked := false
		if buildpack.(map[interface{}]interface{})["locked"] != nil {
			locked = buildpack.(map[interface{}]interface{})["locked"].(bool)
		}

		buildpacks = append(buildpacks, Buildpack{
			Name: buildpack.(map[interface{}]interface{})["name"].(string),
			Position: buildpack.(map[interface{}]interface{})["position"].(int),
			Filename: buildpack.(map[interface{}]interface{})["filename"].(string),
			Enabled: enabled,
			Locked: locked,
		})
	}

	return
}

func validateFiles(buildpacks []Buildpack) error {
	var invalid bool
	for _, buildpack := range buildpacks {
		_, err := os.Stat(buildpack.Filename)
		if err != nil {
			invalid = true
			dir, err := filepath.Abs(buildpack.Filename)
			if err != nil {
				return err
			}
			fmt.Printf("Invalid filename '%v' for buildpack '%v'. File does not exist.\n", dir, buildpack.Name)
		}
	}
	if invalid {
		return errors.New("Invalid files referenced in configuration")
	}
	return nil
}