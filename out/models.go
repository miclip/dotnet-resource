package out

import (
	"github.com/miclip/dotnet-resource"
)

type Request struct {
	Source dotnetresource.Source `json:"source"`
	Params Params                `json:"params"`
}

type Params struct {
	Project     string `json:"project"`
	TestFilter  string `json:"test_filter"`
	Version     string `json:"version"`
	PackageType string `json:"package_type"`
}

type Response struct {
	Version  dotnetresource.VersionTime        `json:"version"`
	Metadata []dotnetresource.MetadataPair `json:"metadata"`
}

type ProjectMetaData struct {
	PackageID   string
	Path        string
	Author      string
	Owner       string
	Description string
}
