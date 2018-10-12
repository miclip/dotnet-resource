package main

import (
	"fmt"

	"github.com/miclip/dotnet-resource/in"
)

func main() {
	output, err := in.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}