package main

import (
	"os"
	"encoding/json"
	"github.com/miclip/dotnet-resource"
	"github.com/miclip/dotnet-resource/in"
)

func main() {

	if len(os.Args) < 2 {
		dotnetresource.Sayf("usage: %s <sources directory>\n", os.Args[0])
		os.Exit(1)
	}

	dotnetresource.Sayf("os.Args[0] %s",os.Args[0])
	dotnetresource.Sayf("os.Args[1] %s",os.Args[1])

	var request in.Request
	inputRequest(&request)

	response := in.Response{
		Version: dotnetresource.Version{
			PackageID : request.Version.PackageID,
			Version: request.Version.Version,
		},
	}

	outputResponse(response)
}

func inputRequest(request *in.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		dotnetresource.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response in.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		dotnetresource.Fatal("writing response to stdout", err)
	}
}