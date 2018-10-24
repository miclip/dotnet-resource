package out

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/miclip/dotnet-resource"
	"github.com/miclip/dotnet-resource/nuget"
	"gopkg.in/xmlpath.v1"
)

var ExecCommand = exec.Command

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

	packages := FindAllIsPackable(sourceDir)
	for _, p := range packages {
		var version = request.Params.Version
		if strings.Contains(version, "*") {
			dotnetresource.Sayf("calculating version for %s...\n", p.PackageID)
			version, err = GenerateNextVersion(request, p.PackageID)
			if err != nil {
				return out, err
			}
		}
		dotnetresource.Sayf("dotnet pack for %s...\n", p.PackageID)
		packOut, err := client.Pack(p.Path, version)
		out = append(out, packOut...)
		if err != nil {
			return out, err
		}
	}

	dotnetresource.Sayf("dotnet nuget push...\n")
	pushOut, err := client.Push(request.Source.NugetSource, request.Source.NugetAPIKey)
	out = append(out, pushOut...)
	if err != nil {
		return out, err
	}

	return out, nil
}

// GenerateNextVersion ....
func GenerateNextVersion(request Request, packageName string) (string, error) {
	nugetclient := nuget.NewNugetClient(request.Source.NugetSource)
	pv, err := nugetclient.GetPackageVersion(context.Background(), packageName, true)
	if err != nil {
		pv = &nuget.PackageVersion{
			ID:      packageName,
			Version: "1.0.0",
		}
	}
	latestVersion := strings.Split(pv.Version, ".")
	specVersion := strings.Split(request.Params.Version, ".")
	if len(latestVersion) != len(specVersion) {
		return "", fmt.Errorf("Version semantics don't match \n %v", err)
	}
	for index := 0; index < len(specVersion); index++ {
		if specVersion[index] == "*" {
			i, _ := strconv.Atoi(latestVersion[index])
			latestVersion[index] = strconv.Itoa(i + 1)
		}
	}
	version := strings.Join(latestVersion, ".")
	return version, nil
}

// FindAllIsPackable ...
func FindAllIsPackable(sourceDir string) []ProjectMetaData {
	var packages []ProjectMetaData
	isPackablePath := xmlpath.MustCompile("/Project/PropertyGroup/IsPackable")
	packageID := xmlpath.MustCompile("/Project/PropertyGroup/PackageId")

	cmd := ExecCommand("find", sourceDir, "-type", "f", "-name", "*.csproj")
	out, err := cmd.CombinedOutput()
	if err != nil {
		dotnetresource.Fatal("error searching for test projects: \n"+string(out), err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		filePath := scanner.Text()
		file, err := os.Open(filePath)
		root, err := xmlpath.Parse(file)
		if err != nil {
			dotnetresource.Fatal("", err)
		}
		if isPackable, ok := isPackablePath.String(root); ok {
			if isPackable == "true" {
				if value, ok := packageID.String(root); ok {
					packageMetadata := ProjectMetaData{
						PackageID: value,
						Path:      filePath,
					}
					packages = append(packages, packageMetadata)
				}
			}
		}
	}
	return packages
}
