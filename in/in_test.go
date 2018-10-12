package in_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"github.com/miclip/dotnet-resource/in"
)

var _ = Describe("in", func() {
	AfterEach(func() {
		CleanupBuildArtifacts()
	})

	It("should compile", func() {
		_, err := Build("github.com/miclip/dotnet-resource/in/cmd")
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("should output an empty JSON list", func() {
		output, err := in.Execute()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(output).Should(MatchJSON("[]"))
	})
})