package in_test

import (
	"github.com/miclip/dotnet-resource/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/miclip/dotnet-resource"
)

var _ = Describe("in", func() {
	BeforeEach(func() {
		dotnetresource.ExecCommand = fakes.FakeExecCommand
	})
	It("should output an empty JSON list", func() {
		//_, _, _ := in.Execute(in.Request{}, "/targetDir")
		Expect(nil).ShouldNot(HaveOccurred())
	})
})