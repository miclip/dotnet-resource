package main

import (
	"log"
	"os"
	"encoding/json"

	"github.com/miclip/dotnet-resource/check"
	"github.com/miclip/dotnet-resource"
)

func main() {
	var request check.Request
	inputRequest(&request)

	response, err := check.Execute(request)
	if err != nil {
		log.Fatal(err)
	}		
	outputResponse(response)
}

func inputRequest(request *check.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		dotnetresource.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response check.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		dotnetresource.Fatal("writing response to stdout", err)
	}
}