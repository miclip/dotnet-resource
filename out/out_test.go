package out_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"github.com/miclip/dotnet-resource/out"
)

var _ = Describe("out", func() {
	AfterEach(func() {
		CleanupBuildArtifacts()
	})

	It("should compile", func() {
		_, err := Build("github.com/miclip/dotnet-resource/out/cmd")
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("should output an empty JSON list", func() {
		output, err := out.Execute()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(output).Should(MatchJSON("[]"))
	})
})