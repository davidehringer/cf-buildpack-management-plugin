package buildpacks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestManagement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Buildpacks Suite")
}
