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
func Execute(request *Request, sourceDir string) ([]byte, error) {
	out := []byte{}
	client := dotnetresource.NewDotnetClient(request.Params.Project, request.Source.Framework, request.Source.Runtime, sourceDir)
	out, err := client.Build()
	if err != nil {
		return out, err
	}
	testOut, err := client.Test(request.Params.TestFilter)
	out = append(out, testOut...)
	if err != nil {
		return out, err
	}

	if strings.Contains(request.Params.Version, "*") {
		nugetclient := nuget.NewNugetClient(request.Source.NugetSource)
		pv, err := nugetclient.GetPackageVersion(context.Background(), request.Params.PackageName, true)
		if err != nil {
			pv = &nuget.PackageVersion{
				ID: request.Params.PackageName,
				Version: "1.0.0",
			}
		}
		latestVersion := strings.Split(pv.Version, ".")
		specVersion := strings.Split(request.Params.Version, ".")
		if len(latestVersion) != len(specVersion) {
			return nil, fmt.Errorf("Version semantics don't match \n %v", err)
		}
		for index := 0; index < len(specVersion); index++ {
			if specVersion[index] == "*" {
				i, _ := strconv.Atoi(specVersion[index])
				latestVersion[index] = strconv.Itoa(i + 1)
			}
		}
		request.Params.Version = strings.Join(latestVersion, ".")
	}

	return out, nil
}
