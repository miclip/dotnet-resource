package in

import "github.com/miclip/dotnet-resource"

type Request struct {
	Source  dotnetresource.Source  `json:"source"`
	Version []dotnetresource.Version `json:"version"`
}

type Response struct {
	Version  dotnetresource.Version        `json:"version"`
	Metadata []dotnetresource.MetadataPair `json:"metadata"`
}