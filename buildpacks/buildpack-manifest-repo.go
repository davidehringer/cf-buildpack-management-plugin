package buildpacks

import (
	"fmt"
	"os"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"
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

	manifest, err := parseManifest(file)
	err = validateFiles(manifest)
	if err != nil {
		return nil, err
	}

	return manifest.Buildpacks, err
}

func parseManifest(file io.Reader) (manifest buildpackManifest, err error) {
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	err = yaml.Unmarshal(contents, &manifest)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return
}

type buildpackManifest struct {
	Buildpacks []Buildpack
}

func validateFiles(manifest buildpackManifest) error {
	var invalid bool
	for _, buildpack := range manifest.Buildpacks {
		_, err := os.Stat(buildpack.Filename)
		if(err != nil){
			invalid = true
			dir, err := filepath.Abs(buildpack.Filename)
			if err != nil {
				return err
			}
			fmt.Printf("Invalid filename '%v' for buildpack '%v'. File does not exist.\n", dir, buildpack.Name)
		}
	}
	if(invalid){
		return errors.New("Invalid files referenced in configuration")
	}
	return nil
}