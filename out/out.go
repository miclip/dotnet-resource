package out

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/miclip/dotnet-resource"
	"github.com/miclip/dotnet-resource/nuget"
)

//Execute - provides out capability
func Execute(request Request, sourceDir string) ([]byte, error) {
	out := []byte{}
	
	client := dotnetresource.NewDotnetClient(request.Params.Project, request.Source.Framework, request.Source.Runtime, sourceDir)

	dotnetresource.Sayf("dotnet build...\n")
	out, err := client.Build()
	if err != nil {
		return out, err
	}

	dotnetresource.Sayf("dotnet test...\n")
	testOut, err := client.Test(request.Params.TestFilter)
	out = append(out, testOut...)
	if err != nil {
		return out, err
	}

	if strings.Contains(request.Params.Version, "*") {
		dotnetresource.Sayf("calculating version...\n")
		err = generateNextVersion(&request)
		if err != nil {
			return out, err
		}
	}

	dotnetresource.Sayf("dotnet pack...\n")
	packOut, err := client.Pack(request.Params.Version)
	out = append(out, packOut...)
	if err != nil {
		return out, err
	}

	dotnetresource.Sayf("dotnet nuget push...\n")
	pushOut, err := client.Push(request.Source.NugetSource, request.Source.NugetAPIKey)
	out = append(out, pushOut...)
	if err != nil {
		return out, err
	}

	return out, nil
}

func generateNextVersion(request *Request) error {

	nugetclient := nuget.NewNugetClient(request.Source.NugetSource)
	pv, err := nugetclient.GetPackageVersion(context.Background(), request.Params.PackageName, true)
	if err != nil {
		pv = &nuget.PackageVersion{
			ID:      request.Params.PackageName,
			Version: "1.0.0",
		}
	}
	latestVersion := strings.Split(pv.Version, ".")
	specVersion := strings.Split(request.Params.Version, ".")
	if len(latestVersion) != len(specVersion) {
		return fmt.Errorf("Version semantics don't match \n %v", err)
	}
	for index := 0; index < len(specVersion); index++ {
		if specVersion[index] == "*" {
			i, _ := strconv.Atoi(latestVersion[index])
			latestVersion[index] = strconv.Itoa(i + 1)
		}
	}
	request.Params.Version = strings.Join(latestVersion, ".")

	return nil
}
