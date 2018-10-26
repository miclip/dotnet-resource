package check

import (
	"context"

	"github.com/miclip/dotnet-resource"
	"github.com/miclip/dotnet-resource/nuget"
)

//Execute - provides check capability
func Execute(request Request) (Response, error) {

	nugetclient := nuget.NewNugetClient(request.Source.NugetSource)
	packageVersion, err := nugetclient.GetPackageVersion(context.Background(), request.Source.PackageID, request.Source.PreRelease)
	if err != nil {
		dotnetresource.Fatal("error querying for latest version from nuget.", err)
	}

	if packageVersion == nil {
		dotnetresource.Sayf("package %s not found at %s ", request.Source.PackageID, request.Source.NugetSource)
		return Response{}, nil
	}

	response := []dotnetresource.Version{
		dotnetresource.Version{
			PackageID: request.Source.PackageID,
			Version:   packageVersion.Version,
		},
	}

	return response, nil
}
