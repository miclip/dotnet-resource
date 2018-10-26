package main

import (
	"log"
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

	var request in.Request
	inputRequest(&request)

	response, output, err := in.Execute(request, os.Args[1])
	dotnetresource.Sayf(string(output))

	if err != nil {
		log.Fatal(err)
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