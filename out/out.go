package out

import (
	"fmt"
	"log"
	
	"github.com/miclip/dotnet-resource"
)

//Execute - provides out capability
func Execute() (string, error) {
	client := dotnetresource.NewDotnetClient("/path/project.csproj","netcoreapp2.1","ubuntu.14.04-x64")
	out, err := client.Build()
	if(err != nil) {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	return "[]", nil
}