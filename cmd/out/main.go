package main

import (
	"fmt"

	"github.com/miclip/dotnet-resource/out"
)

func main() {
	output, err := out.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}