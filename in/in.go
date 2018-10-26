package in

import (
	"context"
	"os"
	"github.com/miclip/dotnet-resource/nuget"
	"github.com/miclip/dotnet-resource"
)

//Execute - provides in capability
func Execute(request Request, targetDir string) (Response, []byte, error) {
	out := []byte{}

	if request.Version.PackageID == "" {
		return Response{}, out, nil
	}

	err := os.MkdirAll(targetDir, 0755)
	if err != nil {
		return Response{}, out, err
	}
	
	nugetclient := nuget.NewNugetClient(request.Source.NugetSource)
	err = nugetclient.DownloadPackage(context.Background(), request.Version.PackageID, request.Version.Version, targetDir)
	if err != nil {
		return Response{}, out, err
	}
	dotnetresource.Sayf("downloaded package %s %s \n",request.Version.PackageID, request.Version.Version)

	client := dotnetresource.NewDotnetClient("", request.Source.Framework, request.Source.Runtime, targetDir)
	unpackOut, err := client.ManualUnpack(request.Version.PackageID, request.Version.Version)
	out = append(out, unpackOut...)
	if err != nil {
		return Response{}, out, err
	}

	return Response{}, out, nil
}