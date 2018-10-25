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
	TestFilter  string `json:"testfilter"`
	Version     string `json:"version"`
	PackageType string `json:"packagetype"`
}

type Response struct {
	Version  dotnetresource.Version        `json:"version"`
	Metadata []dotnetresource.MetadataPair `json:"metadata"`
}

type ProjectMetaData struct {
	PackageID   string
	Path        string
	Author      string
	Owner       string
	Description string
}
