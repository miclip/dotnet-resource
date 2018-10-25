package nuget

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DotnetClient ...
type NugetClient interface {
	GetServiceIndex(ctx context.Context) (*ServiceIndex, error)
	SearchQueryService(ctx context.Context, searchQueryURL string, query string, preRelease bool) (*SearchResults, error)
	GetPackageVersion(ctx context.Context, name string, preRelease bool) (*PackageVersion, error)
	CreateNuspec(packageID string, version string, author string, description string, owner string) Nuspec
}

type nugetclient struct {
	FeedURL      string
	ServiceIndex ServiceIndex
}

func NewNugetClient(
	feedurl string,
) NugetClient {

	return &nugetclient{
		FeedURL: feedurl,
	}
}

func (client *nugetclient) GetServiceIndex(ctx context.Context) (*ServiceIndex, error) {

	req, err := http.NewRequest(http.MethodGet, client.FeedURL, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Add("accept", "application/json")
	var netClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := netClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error getting Service Index %d", res.StatusCode)
	}
	defer res.Body.Close()

	var r ServiceIndex
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (client *nugetclient) SearchQueryService(ctx context.Context, searchQueryURL string, query string, preRelease bool) (*SearchResults, error) {
	queryParams := fmt.Sprintf("?q=%s&prerelease=%t", query, preRelease)
	req, err := http.NewRequest(http.MethodGet, searchQueryURL+queryParams, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Add("accept", "application/json")
	var netClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := netClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error querying Service %d", res.StatusCode)
	}
	defer res.Body.Close()

	var r SearchResults
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (client *nugetclient) CreateNuspec(packageID string, version string, author string, description string, owner string) Nuspec {
	return Nuspec{
		Xmlns:                    "http://schemas.microsoft.com/packaging/2013/05/nuspec.xsd",
		ID:                       packageID,
		Version:                  version,
		Authors:                  author,
		Owners:                   owner,
		RequireLicenseAcceptance: false,
		Description:              description,
	}
}

func (client *nugetclient) GetPackageVersion(ctx context.Context, name string, preRelease bool) (*PackageVersion, error) {
	serviceIndex, err := client.GetServiceIndex(ctx)
	if err != nil {
		return nil, err
	}
	var searchQueryService string
	for _, resource := range serviceIndex.Resources {
		if resource.Type == "SearchQueryService" {
			searchQueryService = resource.ID
		}
	}

	if searchQueryService == "" {
		return nil, fmt.Errorf("Could not find SearchQueryService URL")
	}

	searchResults, err := client.SearchQueryService(ctx, searchQueryService, name, preRelease)
	if err != nil {
		return nil, err
	}

	if searchResults == nil {
		return nil, nil
	}

	for _, result := range searchResults.Data {
		if result.ID == name {
			return &PackageVersion{
				ID:          result.ID,
				Version:     result.Version,
				Description: result.Description,
			}, nil
		}
	}

	return nil, fmt.Errorf("Package not found name: %s prerelease: %t", name, preRelease)

}
