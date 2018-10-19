package main

import (
	"os"
	"log"
	"encoding/json"
	"github.com/miclip/dotnet-resource"
	"github.com/miclip/dotnet-resource/out"
)

func main() {
	dotnetresource.Sayf("arg 0 %v \n",os.Args[1])

	if len(os.Args) < 2 {
		dotnetresource.Sayf("usage: %s <sources directory>\n", os.Args[0])
		os.Exit(1)
	}

	var request out.Request
	inputRequest(&request)

	output, err := out.Execute(os.Args[1],request.Params.Project, request.Source.Framework, request.Source.Runtime)
	dotnetresource.Sayf(string(output))
	if err != nil {
		log.Fatal(err)
	}	
	response := out.Response{
		Version : dotnetresource.Version{Path:"path", VersionID: "1.0"},
		Metadata : []dotnetresource.MetadataPair{ dotnetresource.MetadataPair{ Name:"name", Value:"value"},
		},
	}

	outputResponse(response)
}

func inputRequest(request *out.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		dotnetresource.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response out.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		dotnetresource.Fatal("writing response to stdout", err)
	}
}