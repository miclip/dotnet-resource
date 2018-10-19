package out

import (
	"path"
	"github.com/miclip/dotnet-resource"
	
)

//Execute - provides out capability
func Execute(sourceDir string, project string, framework string, runtime string) ([]byte, error) {	
	client := dotnetresource.NewDotnetClient(path.Join(sourceDir,project),framework,runtime)
	out, err := client.Build()
	return out, err
}

