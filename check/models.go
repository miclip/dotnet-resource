package check

import "github.com/miclip/dotnet-resource"

type Request struct {
	Source  dotnetresource.Source  `json:"source"`
	Version dotnetresource.Version `json:"version"`
}

type Response []dotnetresource.Version