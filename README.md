# Dotnet Resource

Builds, tests, versions and pushes dotnet core libraries and applications to a nuget repository.

## Source Configuration

* `framework`: *Required for Out* dotnet core [framework](https://docs.microsoft.com/en-us/dotnet/standard/frameworks)  

* `runtime`: *Required for Out.* the dotnet [runtime identifier](https://docs.microsoft.com/en-us/dotnet/core/rid-catalog)

* `nuget_source`: *Required.* URL for nuget feed, only v3 API's are currently supported.

* `nuget_apikey`: *Required.* Nuget server API Key.

* `nuget_timeout`: *Optional.* The timeout for pushing packages to nuget, defaults to 300 seconds.

* `package_Id`: *Required for Check.* Package Name or PackageID as defined in the nuspec or csproj.

* `prerelease`: *Optional.* Whether the package is prerelease or not

### Versioning

The `Version` can be set via the Params on the `Out` step or if a mask is provided e.g. `1.0.*` the resource will query nuget and increment each section that contains an asterisk. If no version is found in the provided nuget feed the asterisks will be replaced with zeros.

## Behavior

### `check`: Detect a new version of a package

Requires the `package_id` and `prerelease` and nuget details in the source. 

Resource:
```yml
- name: app-nuget
  type: dotnet
  source:
    nuget_source: https://www.nuget.org/myfeed/api/v3/index.json
    nuget_apikey: {{nuget_apikey}}
    package_id: DotnetResource.TestApplication
    prerelease: true
```
Plan:
```yml
  - get: app-nuget
    trigger: true
  - put: cloudfoundry-development
    params:
      manifest: app-nuget/DotnetResource.TestApplication/manifest.yml
      path: app-nuget/DotnetResource.TestApplication
```

### `in`: Fetch a package from nuget.

Downloads and unpacks the package into the destination within a folder named as the PackageID.

### `out`: Build, test, version and push a package.

Given a solution or project file specified by `project`, restore, build, test, version and pack the into nuget packages. Expects the nuspec fields to be populated within the `csproj` for both applications and libraries being packaged.

```xml
    <IsPackable>true</IsPackable>
    <PackageId>DotnetResource.TestApplication</PackageId>
    <Version>1.0.0</Version>
    <Authors>Michael Lipscombe</Authors>
    <Company>Pivotal</Company>
    <Description>A test application for dotnet-resource</Description>
```

#### Parameters

* `project`: *Required.* Path to solution file or project file within the source code repository

* `test_filter`: *Required.* A mask to identify test projects to be tested. Currently the dotnet cli returns an exit code of 1 if a project is tested that doe not contain any unit tests. Providing this mask enables the resource to limit the projects test by `dotnet test` to a specific subset. 

* `version`: *Required.* Either the version to use or provide a mask and the resource will auto increment based on latest package it finds in nuget. 

* `package_type`: *Required.* Whether this repo contains libraries or applications, currently it doesn't now support both nuget library packages and application packages in one solution.

## Example Configuration

### Resource Type

``` yaml
- name: dotnet
  type: docker-image
  source:
    repository: miclip/dotnet-resource
    tag: "latest"
```

### Resource

For build, test and packaging

``` yaml
- name: dotnet-build-push-nuget
  type: dotnet 
  source:
    framework: netcoreapp2.1
    runtime: ubuntu.14.04-x64
    nuget_source: https://www.nuget.org/myfeed/api/v3/index.json
    nuget_apikey: {{nuget_apikey}}
    nuget_timeout: 600
```

or 

for continuous delivery `check` and `in` 

``` yaml
- name: app-nuget
  type: dotnet
  source:
    nuget_source: https://www.nuget.org/myfeed/api/v3/index.json
    nuget_apikey: {{nuget_apikey}}
    nuget_timeout: 600
    package_id: DotnetResource.TestApplication
    prerelease: true
```

### Plan

``` yaml
- put: dotnet-build-push-nuget
  params:
    project: app-repo/DotnetResource.sln
    test_filter: "*.Tests.csproj"
    version: "1.0.*"
    package_type: Application
```

``` yaml
- get: app-nuget
  trigger: true
```

## Development

### Prerequisites

* golang is *required* - version 1.11.x is tested; earlier versions may also
  work.
* docker is *required* - version 18.06.x is tested; earlier versions may also
  work.
* dep is used for dependency management of the golang packages.

### Running the tests

The tests have been embedded with the `Dockerfile`; ensuring that the testing
environment is consistent across any `docker` enabled platform. When the docker
image builds, the test are run inside the docker container, on failure they
will stop the build.

Run the tests with the following command:

```sh
docker build -t dotnet-resource .
```

### Examples 

Dotnet core MVC Application with tests:

https://github.com/miclip/dotnet-resource-test-application

Dotnet core Libraries with tests: 

https://github.com/miclip/dotnet-resource-test-library