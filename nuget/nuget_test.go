package nuget_test

import (
	"context"
	"encoding/xml"

	"github.com/miclip/dotnet-resource/nuget"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("SearchQueryService", func() {
	var server *ghttp.Server
	var returnedSearchResults nuget.SearchResults
	var statusCode int
	var client nuget.NugetClient

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = nuget.NewNugetClient(server.URL() + "/somefeed/api/v3/index.json")
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/somefeed/api/v3/query"),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, &returnedSearchResults),
			),
		)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When the server returns a Search Result", func() {
		BeforeEach(func() {

			returnedSearchResults = nuget.SearchResults{
				TotalHits:  1000,
				Index:      "index",
				LastReopen: "2018-10-22T22:45:00.425508Z",
				Data: []nuget.SearchResult{
					nuget.SearchResult{
						ID:          "Some.Package.Name",
						Version:     "2.0.1",
						Description: "A test package description",
					},
				},
			}
			statusCode = 200
		})

		It("returns Service Index of a nuget feed", func() {
			r, err := client.SearchQueryService(context.Background(), server.URL()+"/somefeed/api/v3/query", "Some.Package.Name", true)
			Ω(err).Should(Succeed())
			Ω(server.ReceivedRequests()).Should(HaveLen(1))
			Ω(r).ShouldNot(BeNil())

		})
	})

	Context("when the server returns 500", func() {
		BeforeEach(func() {
			statusCode = 500
		})
		It("errors", func() {
			r, err := client.SearchQueryService(context.Background(), server.URL()+"/somefeed/api/v3/query", "Some.Package.Name", true)
			Ω(err).To(HaveOccurred())
			Ω(r).To(BeNil())

		})
	})

	Context("when the server returns 503", func() {
		BeforeEach(func() {
			statusCode = 503
		})

		It("errors", func() {
			r, err := client.SearchQueryService(context.Background(), server.URL()+"/somefeed/api/v3/query", "Some.Package.Name", true)
			Ω(err).To(HaveOccurred())
			Ω(r).To(BeNil())
		})
	})
})

var _ = Describe("ServiceIndex", func() {
	var server *ghttp.Server
	var returnedServiceIndex nuget.ServiceIndex
	var statusCode int
	var client nuget.NugetClient

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = nuget.NewNugetClient(server.URL() + "/somefeed/api/v3/index.json")
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/somefeed/api/v3/index.json"),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, &returnedServiceIndex),
			),
		)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When when the server returns a Service Index", func() {
		BeforeEach(func() {
			returnedServiceIndex = nuget.ServiceIndex{
				Version: "3.0.0",
				Resources: []nuget.Resource{
					nuget.Resource{
						ID:      "https://www.nuget.org/somefeed/api/v3/query",
						Type:    "SearchQueryService",
						Comment: "Query endpoint of NuGet Search service.",
					},
				},
			}
			statusCode = 200
		})

		It("returns Service Index of a nuget feed", func() {
			r, err := client.GetServiceIndex(context.Background())
			Ω(err).Should(Succeed())
			Ω(server.ReceivedRequests()).Should(HaveLen(1))
			Ω(r).ShouldNot(BeNil())
			Ω(r.Resources).Should(HaveLen(1))
			Ω(r.Resources[0].Type).To(Equal("SearchQueryService"))
			Ω(r.Version).To(Equal("3.0.0"))
		})
	})

	Context("when the server returns 500", func() {
		BeforeEach(func() {
			statusCode = 500
		})

		It("errors", func() {
			r, err := client.GetServiceIndex(context.Background())
			Ω(err).To(HaveOccurred())
			Ω(r).To(BeNil())

		})
	})

	Context("when the server returns 503", func() {
		BeforeEach(func() {
			statusCode = 503
		})

		It("errors", func() {
			r, err := client.GetServiceIndex(context.Background())
			Ω(err).To(HaveOccurred())
			Ω(r).To(BeNil())
		})
	})
})

var _ = Describe("GetPackageVersion", func() {
	var server *ghttp.Server
	var returnedServiceIndex nuget.ServiceIndex
	var returnedSearchResults nuget.SearchResults
	var statusCode int
	var client nuget.NugetClient

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = nuget.NewNugetClient(server.URL() + "/somefeed/api/v3/index.json")
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/somefeed/api/v3/index.json"),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, &returnedServiceIndex),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/somefeed/api/v3/query"),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, &returnedSearchResults),
			),
		)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When the client returns a PackageVersion", func() {
		BeforeEach(func() {
			returnedServiceIndex = nuget.ServiceIndex{
				Version: "3.0.0",
				Resources: []nuget.Resource{
					nuget.Resource{
						ID:      server.URL() + "/somefeed/api/v3/query",
						Type:    "SearchQueryService",
						Comment: "Query endpoint of NuGet Search service.",
					},
					nuget.Resource{
						ID:      server.URL() + "/somefeed/api/v3/query",
						Type:    "SearchQueryService/3.0.0-beta",
						Comment: "Query endpoint of NuGet Search service.",
					},
				},
			}
			returnedSearchResults = nuget.SearchResults{
				TotalHits:  1000,
				Index:      "index",
				LastReopen: "2018-10-22T22:45:00.425508Z",
				Data: []nuget.SearchResult{
					nuget.SearchResult{
						ID:          "Some.Package.Name",
						Version:     "2.0.1",
						Description: "A test package description",
					},
					nuget.SearchResult{
						ID:          "Some.Other.Package.Name",
						Version:     "2.0.10",
						Description: "A test package description",
					},
				},
			}
			statusCode = 200
		})

		It("returns PackageVersion for a particular package", func() {
			r, err := client.GetPackageVersion(context.Background(), "Some.Package.Name", false)
			Ω(err).Should(Succeed())
			Ω(server.ReceivedRequests()).Should(HaveLen(2))
			Ω(r).ShouldNot(BeNil())
			Ω(r.ID).To(Equal("Some.Package.Name"))
			Ω(r.Version).To(Equal("2.0.1"))
		})
	})

})

var _ = Describe("CreateNuspec", func() {
	Context("nuspec encoding", func() {
		It("returns a valid nuspec ", func() {
			client := nuget.NewNugetClient("http://nuget.org/somefeed/api/v3/index.json")
			nuspec := client.CreateNuspec("packageID", "1.0.0", "Michael Lipscombe", "description", "owner")
			output, _ := xml.Marshal(nuspec)
			Ω(string(output)).To(Equal("<package xmlns=\"http://schemas.microsoft.com/packaging/2013/05/nuspec.xsd\"><metadata><id>packageID</id><version>1.0.0</version><authors>Michael Lipscombe</authors><owners>owner</owners><requireLicenseAcceptance>false</requireLicenseAcceptance><description>description</description></metadata></package>"))
		})
	})
})
