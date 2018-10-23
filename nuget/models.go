package nuget

// ServiceIndex response type
type ServiceIndex struct {
	Version   string     `json:"version"`
	Resources []Resource `json:"resources"`
}

// Resource nuget Resource type
type Resource struct {
	ID      string `json:"@id"`
	Type    string `json:"@type"`
	Comment string `json:"comment"`
}

// SearchResults from the nuget api
type SearchResults struct {
	TotalHits  int            `json:"totalHits"`
	Index      string         `json:"index"`
	LastReopen string         `json:"lastReopen"`
	Data    []SearchResult `json:"data"`
}

// SearchResult from the nuget api
type SearchResult struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type PackageVersion struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	Description string `json:"description"`
}
