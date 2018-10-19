package out

import (
	"path"
	"github.com/miclip/dotnet-resource"
	
)

//Execute - provides out capability
func Execute(request Request, sourceDir string) ([]byte, error) {	
	out := []byte{}
	client := dotnetresource.NewDotnetClient(path.Join(sourceDir,request.Params.Project),request.Source.Framework,request.Source.Runtime)
	out, err := client.Build()
	if(err!=nil){
		return out, err
	}
	testOut, err := client.Test(request.Params.TestFilter)
	out = append(out, testOut...)
	if(err!=nil){
		return out, err
	}
	return out, nil
}

