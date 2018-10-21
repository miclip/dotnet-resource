package dotnetresource_test


import (
	"github.com/miclip/dotnet-resource/fakes"
	
	"os/exec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/miclip/dotnet-resource"
)



var _ = Describe("dotnetclient", func() {
	BeforeEach(func(){
		dotnetresource.ExecCommand = fakes.FakeExecCommand
	})

	It("should execute dotnet build command", func() {
		defer func() { dotnetresource.ExecCommand = exec.Command }()
		fakes.MockedExitStatus = 0
		fakes.MockedStdout = ""
		expectedCommand := "dotnet build tmp/source/path/project.csproj -f netcoreapp2.1 -r ubuntu.14.04-x64"		
		
		client := dotnetresource.NewDotnetClient("/path/project.csproj","netcoreapp2.1","ubuntu.14.04-x64", "tmp/source")
		_, err := client.Build()
		Ω(fakes.CommandString).Should(Equal(expectedCommand))
		Ω(err).ShouldNot(HaveOccurred())		
	})

	It("should execute dotnet test command", func() {
		defer func() { dotnetresource.ExecCommand = exec.Command }()
		fakes.MockedExitStatus = 0
		fakes.MockedStdout = ""
		expectedCommand := "dotnet test -f netcoreapp2.1 --no-build --no-restore PASS -p:RuntimeIdentifier=ubuntu.14.04-x64"		
		
		client := dotnetresource.NewDotnetClient("/path/project.csproj","netcoreapp2.1","ubuntu.14.04-x64", "/tmp/source")
		_, err := client.Test("A_Filter")
		Ω(fakes.CommandString).Should(Equal(expectedCommand))
		Ω(err).ShouldNot(HaveOccurred())		
	})

})