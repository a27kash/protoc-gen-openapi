package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"gopkg.in/yaml.v3"
)

type Document struct {
	Openapi	string	`yaml:"openapi"`
}

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		d := Document{}
		d.Openapi = "3.1.0"
		bytes, err := yaml.Marshal(d)
		if err != nil {
			return fmt.Errorf("failed to marshal yaml: %s", err.Error())
		}
		outputFile := gen.NewGeneratedFile("openapi.yaml", "")
		outputFile.Write(bytes)
		return nil
	})
}

