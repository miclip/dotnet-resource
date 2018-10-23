package dotnetresource

import (
	"time"
)

type Source struct {
	Framework          string `json:"framework"`
	Runtime      string `json:"runtime"`
	NugetSource string `json:"nuget_source"`
	NugetAPIKey string `json:"nuget_apikey"`
}

func (source Source) IsValid() (bool, string) {
	
	return true, ""
}

type Version struct {
	Timestamp time.Time `json:"timestamp"`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}