package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"gopkg.in/yaml.v3"
)

// Contact information for the exposed API.
type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// The URL pointing to the contact information. This MUST be in the form of a URL.
	URL string `yaml:"url,omitempty" json:"url,omitempty"`
	// The email address of the contact person/organization. This MUST be in the form of an email address.
	Email string `yaml:"email,omitempty" json:"email,omitempty"`
}

// License information for the exposed API.
type License struct {
	// REQUIRED. The license name used for the API.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// An SPDX license expression for the API. The identifier field is mutually exclusive of the url field.
	Identifier string `yaml:"identifier,omitempty" json:"identifier,omitempty"`
	// A URL to the license used for the API. This MUST be in the form of a URL. The url field is mutually exclusive of the identifier field.
	URL string `yaml:"url,omitempty" json:"url,omitempty"`
}

// Info provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
type Info struct {
	// REQUIRED. The title of the API.
	Title string `yaml:"title,omitempty" json:"title,omitempty"`
	// A short summary of the API.
	Summary string `yaml:"summary,omitempty" json:"summary,omitempty"`
	// A description of the API. CommonMark syntax MAY be used for rich text representation.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	// A URL to the Terms of Service for the API. This MUST be in the form of a URL.
	TermsOfService string `yaml:"termsOfService,omitempty" json:"termsOfService,omitempty"`
	// The contact information for the exposed API.
	Contact Contact `yaml:"contact,omitempty" json:"contact,omitempty"`
	// The license information for the exposed API.
	License License `yaml:"license,omitempty" json:"license,omitempty"`
	// REQUIRED. The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

// OpenAPI is the root object of the OpenAPI document.
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI document. This is not related to the API info.version string.
	OpenAPI string `yaml:"openapi,omitempty" json:"openapi,omitempty"`
	// REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
	Info Info `yaml:"info,omitempty" json:"info,omitempty"`
}

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		d := OpenAPI{}
		d.OpenAPI = "3.1.0"
		bytes, err := yaml.Marshal(d)
		if err != nil {
			return fmt.Errorf("failed to marshal yaml: %s", err.Error())
		}
		outputFile := gen.NewGeneratedFile("openapi.yaml", "")
		outputFile.Write(bytes)
		return nil
	})
}
