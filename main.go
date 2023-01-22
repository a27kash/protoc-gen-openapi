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

// ServerVariable is an object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
	Enum []string `yaml:"enum,omitempty" json:"enum,omitempty"`
	// REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. Note this behavior is different than the Schema Object’s treatment of default values, because in those cases parameter values are optional. If the enum is defined, the value MUST exist in the enum’s values.
	Default string `yaml:"default,omitempty" json:"default,omitempty"`
	// An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

// Server is an object representing a Server.
type Server struct {
	// REQUIRED. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenAPI document is being served. Variable substitutions will be made when a variable is named in {brackets}.
	URL string `yaml:"url,omitempty" json:"url,omitempty"`
	// An optional string describing the host designated by the URL. CommonMark syntax MAY be used for rich text representation.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	// A map between a variable name and its value. The value is used for substitution in the server’s URL template.
	Variables map[string]ServerVariable `yaml:"variables,omitempty" json:"variables,omitempty"`
}

// SecurityRequirement lists the required security schemes to execute this operation. The name used for each property MUST correspond to a security scheme declared in the Security Schemes under the Components Object.
type SecurityRequirement struct {}

// ExternalDocumentation allows referencing an external resource for extended documentation.
type ExternalDocumentation struct {
	// A description of the target documentation. CommonMark syntax MAY be used for rich text representation.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`
	// REQUIRED. The URL for the target documentation. This MUST be in the form of a URL.
	URL	string	`yaml:"url,omitempty" json:"url,omitempty"`
}

// Tag adds metadata to a single tag that is used by the Operation Object. It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.
type Tag struct {
	// REQUIRED. The name of the tag.
	Name	string	`yaml:"name,omitempty" json:"name,omitempty"`
	// A description for the tag. CommonMark syntax MAY be used for rich text representation.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`
	// Additional external documentation for this tag.
	ExternalDocs	ExternalDocumentation	`yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
}

// OpenAPI is the root object of the OpenAPI document.
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI document. This is not related to the API info.version string.
	OpenAPI string `yaml:"openapi,omitempty" json:"openapi,omitempty"`
	// REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
	Info Info `yaml:"info,omitempty" json:"info,omitempty"`
	// The default value for the $schema keyword within Schema Objects contained within this OAS document. This MUST be in the form of a URI.
	JSONSchemaDialect string `yaml:"jsonSchemaDialect,omitempty" json:"jsonSchemaDialect,omitempty"`
	// An array of Server Objects, which provide connectivity information to a target server. If the servers property is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
	Servers []Server `yaml:"servers,omitempty" json:"servers,omitempty"`

	// A declaration of which security mechanisms can be used across the API. The list of values includes alternative security requirement objects that can be used. Only one of the security requirement objects need to be satisfied to authorize a request. Individual operations can override this definition. To make security optional, an empty security requirement ({}) can be included in the array.
	Security	[]SecurityRequirement	`yaml:"security,omitempty" json:"security,omitempty"`
	// A list of tags used by the document with additional metadata. The order of the tags can be used to reflect on their order by the parsing tools. Not all tags that are used by the Operation Object must be declared. The tags that are not declared MAY be organized randomly or based on the tools’ logic. Each tag name in the list MUST be unique.
	Tags	[]Tag	`yaml:"tags,omitempty" json:"tags,omitempty"`
	// Additional external documentation.
	ExternalDocs	ExternalDocumentation	`yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
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
