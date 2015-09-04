package buildpacks_test

import (
	. "github.com/davidehringer/cf-buildpack-management-plugin/buildpacks"

	"fmt"
	"github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("CommandGenerator", func() {

	var cliConnection *fakes.FakeCliConnection
	var generator CommandGenerator

	desired := []Buildpack{
		Buildpack{
			Name:     "java",
			Position: 1,
			Enabled:  true,
			Locked:   false,
			Filename: "java-buildpack-offline-v3.2.zip",
		},
		Buildpack{
			Name:     "ruby",
			Position: 2,
			Enabled:  true,
			Locked:   false,
			Filename: "ruby_buildpack-cached-v1.3.1.zip",
		},
		Buildpack{
			Name:     "java_3_0",
			Position: 1,
			Enabled:  true,
			Locked:   false,
			Filename: "java-buildpack-offline-v3.0.zip",
		},
	}

	actual := []Buildpack{
		Buildpack{
			Name:     "java",
			Position: 1,
			Enabled:  true,
			Locked:   false,
			Filename: "java-buildpack-offline-v3.0.zip",
		},
		Buildpack{
			Name:     "ruby",
			Position: 2,
			Enabled:  true,
			Locked:   false,
			Filename: "ruby_buildpack-cached-v1.3.0.zip",
		},
	}

	BeforeEach(func() {
		cliConnection = &fakes.FakeCliConnection{}
	})

	It("generates a lists of ordered buildpack commands that will put the systems in the desired state", func() {
		commands := generator.GenerateCommands(actual, desired)

		Expect(commands).To(ContainElement(MatchingBuildpackCommand(DELETE, "ruby")))
		Expect(commands).To(ContainElement(MatchingBuildpackCommand(RENAME, "java_3_0")))
		Expect(commands).To(ContainElement(MatchingBuildpackCommand(ADD, "java")))
		Expect(commands).To(ContainElement(MatchingBuildpackCommand(ADD, "ruby")))
		Expect(commands).To(ContainElement(MatchingBuildpackCommand(UPDATE, "java")))
		Expect(commands).To(ContainElement(MatchingBuildpackCommand(UPDATE, "ruby")))
		Expect(commands).To(ContainElement(MatchingBuildpackCommand(UPDATE, "java_3_0")))

		// asserting order of actions
		Expect(commands[0].Action()).To(Equal(DELETE))
		Expect(commands[1].Action()).To(Equal(RENAME))
		Expect(commands[2].Action()).To(Equal(ADD))
		Expect(commands[3].Action()).To(Equal(ADD))
		Expect(commands[4].Action()).To(Equal(UPDATE))
		Expect(commands[5].Action()).To(Equal(UPDATE))
		Expect(commands[6].Action()).To(Equal(UPDATE))
	})
})

// Custom Matchers

func MatchingBuildpackCommand(action string, buildpackName string) types.GomegaMatcher {
	return &buildpackCommandMatcher{
		buildpackName: buildpackName,
		action:        action,
	}
}

type buildpackCommandMatcher struct {
	buildpackName string
	action        string
}

func (matcher *buildpackCommandMatcher) Match(actual interface{}) (success bool, err error) {
	command := actual.(BuildpackCommand)
	return command.BuildpackName() == matcher.buildpackName && command.Action() == matcher.action, nil
}

func (matcher *buildpackCommandMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto contain a command of type \t%#v", actual, matcher.action)
}

func (matcher *buildpackCommandMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to contain command of type \t%#v", actual, matcher.action)
}
