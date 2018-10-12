package dotnet_resource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDotnetResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DotnetResource Suite")
}
