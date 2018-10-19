package dotnetresource

type Source struct {
	Framework          string `json:"framework"`
	Runtime      string `json:"runtime"`
}

func (source Source) IsValid() (bool, string) {
	
	return true, ""
}

type Version struct {
	Path      string `json:"path,omitempty"`
	VersionID string `json:"version_id,omitempty"`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}