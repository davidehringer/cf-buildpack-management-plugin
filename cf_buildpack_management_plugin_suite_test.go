package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfBuildpackManagementPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfBuildpackManagementPlugin Suite")
}
