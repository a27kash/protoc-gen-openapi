package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		outputFile := gen.NewGeneratedFile("openapi.yaml", "")
		outputFile.Write(nil)
		return nil
	})
}

