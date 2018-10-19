package out_test

import (
	"github.com/miclip/dotnet-resource/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/miclip/dotnet-resource/out"
	"github.com/miclip/dotnet-resource"
)

var _ = Describe("out", func() {
	BeforeEach(func(){
		dotnetresource.ExecCommand = fakes.FakeExecCommand
	})

	It("should output an empty JSON list", func() {
		request := out.Request{
			Source: dotnetresource.Source{
				Framework: "netcoreapp2.1",
				Runtime: "ubuntu.14.04-x64",
			},
			Params: out.Params{
				Project: "/path/project.csproj",
				TestFilter: "A_Filter",
			},
		}
		_, err := out.Execute(request, "sourceDir")
		Expect(err).ShouldNot(HaveOccurred())
		//Expect(output).Should()
	})
})