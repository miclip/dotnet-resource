package in_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/miclip/dotnet-resource/in"
)

var _ = Describe("in", func() {
	It("should output an empty JSON list", func() {
		_, err := in.Execute()
		Expect(err).ShouldNot(HaveOccurred())
	})
})