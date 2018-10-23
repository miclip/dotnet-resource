package main

import (
	"time"
	"os"
	"log"
	"encoding/json"
	"github.com/miclip/dotnet-resource"
	"github.com/miclip/dotnet-resource/out"
)

func main() {

	if len(os.Args) < 2 {
		dotnetresource.Sayf("usage: %s <sources directory>\n", os.Args[0])
		os.Exit(1)
	}

	var request out.Request
	inputRequest(&request)

	output, err := out.Execute(request, os.Args[1])
	dotnetresource.Sayf(string(output))
	if err != nil {
		log.Fatal(err)
	}	
	response := out.Response{
		Version: dotnetresource.Version{
			Timestamp: time.Now(),
		},
		Metadata: []dotnetresource.MetadataPair{
			{
				Name:  "project",
				Value: request.Params.Project,
			},
			{
				Name:  "framework",
				Value: request.Source.Framework,
			},
			{
				Name:  "version",
				Value: request.Params.Version,
			},
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