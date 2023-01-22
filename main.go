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

// Paths holds the relative paths to the individual endpoints and their operations. The path is appended to the URL from the Server Object in order to construct the full URL. The Paths MAY be empty, due to Access Control List (ACL) constraints.
type Paths struct {}

// ExternalDocumentation allows referencing an external resource for extended documentation.
type ExternalDocumentation struct {
	// A description of the target documentation. CommonMark syntax MAY be used for rich text representation.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`
	// REQUIRED. The URL for the target documentation. This MUST be in the form of a URL.
	URL	string	`yaml:"url,omitempty" json:"url,omitempty"`
}

// ParameterOrReference ...
type ParameterOrReference interface {
	isParameterOrReference()
}

// RequestBodyOrReference ...
type RequestBodyOrReference interface {
	isRequestBodyOrReference()
}

// ResponseOrReference ...
type ResponseOrReference interface {
	isResponseOrReference()
}

// Responses is a container for the expected responses of an operation. The container maps a HTTP response code to the expected response.
type Responses struct {
	// The documentation of responses other than the ones declared for specific HTTP response codes. Use this field to cover undeclared responses.
	Default	ResponseOrReference	`yaml:"default,omitempty" json:"default,omitempty"`
}

// CallbackOrReference ...
type CallbackOrReference interface {
	isCallbackOrReference()
}

// Operation describes a single API operation on a path.
type Operation struct {
	// A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.
	Tags	[]string	`yaml:"tags,omitempty" json:"tags,omitempty"`
	// A short summary of what the operation does.
	Summary	string	`yaml:"summary,omitempty" json:"summary,omitempty"`
	// A verbose explanation of the operation behavior. CommonMark syntax MAY be used for rich text representation.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`
	// Additional external documentation for this operation.
	ExternalDocs	ExternalDocumentation	`yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
	// Unique string used to identify the operation. The id MUST be unique among all operations described in the API. The operationId value is case-sensitive. Tools and libraries MAY use the operationId to uniquely identify an operation, therefore, it is RECOMMENDED to follow common programming naming conventions.
	OperationID	string	`yaml:"operationId,omitempty" json:"operationId,omitempty"`
	// A list of parameters that are applicable for this operation. If a parameter is already defined at the Path Item, the new definition will override it but can never remove it. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object’s components/parameters.
	Parameters	[]ParameterOrReference	`yaml:"parameters,omitempty" json:"parameters,omitempty"`
	// The request body applicable for this operation. The requestBody is fully supported in HTTP methods where the HTTP 1.1 specification [RFC7231] has explicitly defined semantics for request bodies. In other cases where the HTTP spec is vague (such as GET, HEAD and DELETE), requestBody is permitted but does not have well-defined semantics and SHOULD be avoided if possible.
	RequestBody	RequestBodyOrReference	`yaml:"requestBody,omitempty" json:"requestBody,omitempty"`
	// The list of possible responses as they are returned from executing this operation.
	Responses	Responses	`yaml:"responses,omitempty" json:"responses,omitempty"`
	// A map of possible out-of band callbacks related to the parent operation. The key is a unique identifier for the Callback Object. Each value in the map is a Callback Object that describes a request that may be initiated by the API provider and the expected responses.
	Callbacks	map[string]CallbackOrReference	`yaml:"callbacks,omitempty" json:"callbacks,omitempty"`
	// Declares this operation to be deprecated. Consumers SHOULD refrain from usage of the declared operation. Default value is false.
	Deprecated	bool	`yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
	// A declaration of which security mechanisms can be used for this operation. The list of values includes alternative security requirement objects that can be used. Only one of the security requirement objects need to be satisfied to authorize a request. To make security optional, an empty security requirement ({}) can be included in the array. This definition overrides any declared top-level security. To remove a top-level security declaration, an empty array can be used.
	Security	[]SecurityRequirement	`yaml:"security,omitempty" json:"security,omitempty"`
	// An alternative server array to service this operation. If an alternative server object is specified at the Path Item Object or Root level, it will be overridden by this value.
	Servers	[]Server	`yaml:"servers,omitempty" json:"servers,omitempty"`
}

// PathItemOrReference ...
type PathItemOrReference interface {
	isPathItemOrReference()
}

// PathItem describes the operations available on a single path. A Path Item MAY be empty, due to ACL constraints. The path itself is still exposed to the documentation viewer but they will not know which operations and parameters are available.
type PathItem struct {
	// Allows for a referenced definition of this path item. The referenced structure MUST be in the form of a Path Item Object. In case a Path Item Object field appears both in the defined object and the referenced object, the behavior is undefined. See the rules for resolving Relative References.
	Ref	string	`yaml:"$ref,omitempty" json:"$ref,omitempty"`
	// An optional, string summary, intended to apply to all operations in this path.
	Summary	string	`yaml:"summary,omitempty" json:"summary,omitempty"`
	// An optional, string description, intended to apply to all operations in this path. CommonMark syntax MAY be used for rich text representation.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`


	Get	Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Put	Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Post Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Delete Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Options	Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Head	Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Patch	Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Trace	Operation	`yaml:"url,omitempty" json:"url,omitempty"`
	Servers	[]Server	`yaml:"url,omitempty" json:"url,omitempty"`
	Parameters	[]ParameterOrReference	`yaml:"url,omitempty" json:"url,omitempty"`
}

func (p PathItem) isPathItemOrReference() {}

// Discriminator is a specific object in a schema which is used to inform the consumer of the document of an alternative schema based on the value associated with it.
type Discriminator struct {
	// REQUIRED. The name of the property in the payload that will hold the discriminator value.
	PropertyName	string	`yaml:"propertyName,omitempty" json:"propertyName,omitempty"`
	// An object to hold mappings between payload values and schema names or references.
	Mapping	map[string]string	`yaml:"mapping,omitempty" json:"mapping,omitempty"`
}

// XML is a metadata object that allows for more fine-tuned XML model definitions.
type XML struct {
	// Replaces the name of the element/attribute used for the described schema property. When defined within items, it will affect the name of the individual XML elements within the list. When defined alongside type being array (outside the items), it will affect the wrapping element and only if wrapped is true. If wrapped is false, it will be ignored.
	Name	string	`yaml:"name,omitempty" json:"name,omitempty"`
	// The URI of the namespace definition. This MUST be in the form of an absolute URI.
	Namespace	string	`yaml:"namespace,omitempty" json:"namespace,omitempty"`
	// The prefix to be used for the name.
	Prefix	string	`yaml:"prefix,omitempty" json:"prefix,omitempty"`
	// Declares whether the property definition translates to an attribute instead of an element. Default value is false.
	Attribute	bool	`yaml:"attribute,omitempty" json:"attribute,omitempty"`
	// MAY be used only for an array definition. Signifies whether the array is wrapped (for example, <books><book/><book/></books>) or unwrapped (<book/><book/>). Default value is false. The definition takes effect only when defined alongside type being array (outside the items).
	Wrapped	bool	`yaml:"wrapped,omitempty" json:"wrapped,omitempty"`
}

// Schema allows the definition of input and output data types. These types can be objects, but also primitives and arrays. This object is a superset of the JSON Schema Specification Draft 2020-12.
type Schema struct {
	// Adds support for polymorphism. The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description. See Composition and Inheritance for more details.
	Discriminator	Discriminator	`yaml:"discriminator,omitempty" json:"discriminator,omitempty"`
	// This MAY be used only on properties schemas. It has no effect on root schemas. Adds additional metadata to describe the XML representation of this property.
	XML	XML	`yaml:"xml,omitempty" json:"xml,omitempty"`
	// Additional external documentation for this schema.
	ExternalDocs	ExternalDocumentation	`yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`

	// A free-form property to include an example of an instance for this schema. To represent examples that cannot be naturally represented in JSON or YAML, a string value can be used to contain the example with escaping where necessary.
	// Example	Any	`yaml:"example,omitempty" json:"example,omitempty"`
}

// HeaderOrReference ...
type HeaderOrReference interface {
	isHeaderOrReference()
}

// Parameter describes a single operation parameter.
type Parameter struct {
	// REQUIRED. The name of the parameter. Parameter names are case sensitive.
	Name	string	`yaml:"name,omitempty" json:"name,omitempty"`
	// REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
	In	string	`yaml:"in,omitempty" json:"in,omitempty"`
	// A brief description of the parameter. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`
	// Determines whether this parameter is mandatory. If the parameter location is "path", this property is REQUIRED and its value MUST be true. Otherwise, the property MAY be included and its default value is false.
	Required	bool	`yaml:"required,omitempty" json:"required,omitempty"`
	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage. Default value is false.
	Deprecated	bool	`yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
	// Sets the ability to pass empty-valued parameters. This is valid only for query parameters and allows sending a parameter with an empty value. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED, as it is likely to be removed in a later revision.
	AllowEmptyValue	bool	`yaml:"allowEmptyValue,omitempty" json:"allowEmptyValue,omitempty"`
}

func (p Parameter) isParameterOrReference() {}

// Header follows the structure of the Parameter Object.
type Header Parameter

func (h Header) isHeaderOrReference() {}

// ExampleOrReference ...
type ExampleOrReference interface {
	isExampleOrReference()
}

// LinkOrReference ...
type LinkOrReference interface {
	isLinkOrReference()
}

// Reference is a simple object to allow referencing other components in the OpenAPI document, internally and externally.
type Reference struct {
	// REQUIRED. The reference identifier. This MUST be in the form of a URI.
	Ref	string	`yaml:"$ref,omitempty" json:"$ref,omitempty"`
	// A short summary which by default SHOULD override that of the referenced component. If the referenced object-type does not allow a summary field, then this field has no effect.
	Summary string	`yaml:"summary,omitempty" json:"summary,omitempty"`
	// A description which by default SHOULD override that of the referenced component. CommonMark syntax MAY be used for rich text representation. If the referenced object-type does not allow a description field, then this field has no effect.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`
}

func (r Reference) isHeaderOrReference() {}

func (r Reference) isExampleOrReference() {}

func (r Reference) isLinkOrReference() {}

func (r Reference) isResponseOrReference() {}

// Example ...
type Example struct {
	// Short description for the example.
	Summary	string	`yaml:"summary,omitempty" json:"summary,omitempty"`
	// Long description for the example. CommonMark syntax MAY be used for rich text representation.
	Description	string	`yaml:"description,omitempty" json:"description,omitempty"`

	// Embedded literal example. The value field and externalValue field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON or YAML, use a string value to contain the example, escaping where necessary.
	// Value	Any	`yaml:"value,omitempty" json:"value,omitempty"`

	// A URI that points to the literal example. This provides the capability to reference examples that cannot easily be included in JSON or YAML documents. The value field and externalValue field are mutually exclusive. See the rules for resolving Relative References.
	ExternalValue	string	`yaml:"externalValue,omitempty" json:"externalValue,omitempty"`
}

func (e Example) isExampleOrReference() {}

// Encoding is a single encoding definition applied to a single schema property.
type Encoding struct {
	// The Content-Type for encoding a specific property. Default value depends on the property type: for object - application/json; for array – the default is defined based on the inner type; for all other cases the default is application/octet-stream. The value can be a specific media type (e.g. application/json), a wildcard media type (e.g. image/*), or a comma-separated list of the two types.
	ContentType	string	`yaml:"contentType,omitempty" json:"contentType,omitempty"`
	// A map allowing additional information to be provided as headers, for example Content-Disposition. Content-Type is described separately and SHALL be ignored in this section. This property SHALL be ignored if the request body media type is not a multipart.
	Headers	map[string]HeaderOrReference	`yaml:"headers,omitempty" json:"headers,omitempty"`
	// Describes how a specific property value will be serialized depending on its type. See Parameter Object for details on the style property. The behavior follows the same values as query parameters, including default values. This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Style	string	`yaml:"style,omitempty" json:"style,omitempty"`
	// When this is true, property values of type array or object generate separate parameters for each value of the array, or key-value-pair of the map. For other types of properties this property has no effect. When style is form, the default value is true. For all other styles, the default value is false. This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Explode	bool	`yaml:"explode,omitempty" json:"explode,omitempty"`
	// Determines whether the parameter value SHOULD allow reserved characters, as defined by [RFC3986] :/?#[]@!$&'()*+,;= to be included without percent-encoding. The default value is false. This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	AllowReserved	bool	`yaml:"allowReserved,omitempty" json:"allowReserved,omitempty"`
}

// MediaType Object provides schema and examples for the media type identified by its key.
type MediaType struct {
	// The schema defining the content of the request, response, or parameter.
	Schema	Schema	`yaml:"schema,omitempty" json:"schema,omitempty"`

	// Example of the media type. The example object SHOULD be in the correct format as specified by the media type. The example field is mutually exclusive of the examples field. Furthermore, if referencing a schema which contains an example, the example value SHALL override the example provided by the schema.
	// Example	Any	`yaml:"example,omitempty" json:"example,omitempty"`

	// Examples of the media type. Each example object SHOULD match the media type and specified schema if present. The examples field is mutually exclusive of the example field. Furthermore, if referencing a schema which contains an example, the examples value SHALL override the example provided by the schema.
	Examples	map[string]ExampleOrReference	`yaml:"examples,omitempty" json:"examples,omitempty"`
	// A map between a property name and its encoding information. The key, being the property name, MUST exist in the schema as a property. The encoding object SHALL only apply to requestBody objects when the media type is multipart or application/x-www-form-urlencoded.
	Encoding	map[string]Encoding	`yaml:"encoding,omitempty" json:"encoding,omitempty"`
}

// Link object represents a possible design-time link for a response. The presence of a link does not guarantee the caller’s ability to successfully invoke it, rather it provides a known relationship and traversal mechanism between responses and other operations.
type Link struct {
	// A relative or absolute URI reference to an OAS operation. This field is mutually exclusive of the operationId field, and MUST point to an Operation Object. Relative operationRef values MAY be used to locate an existing Operation Object in the OpenAPI definition. See the rules for resolving Relative References.
	OperationRef	string	`yaml:"operationRef,omitempty" json:"operationRef,omitempty"`
	// The name of an existing, resolvable OAS operation, as defined with a unique operationId. This field is mutually exclusive of the operationRef field.
	OperationID	string	`yaml:"operationId,omitempty" json:"operationId,omitempty"`
	// parameters	map[]	`yaml:"parameters,omitempty" json:"parameters,omitempty"`
	// requestBody		`yaml:"requestBody,omitempty" json:"requestBody,omitempty"`
	// description	`yaml:"description,omitempty" json:"description,omitempty"`
	// server	`yaml:"server,omitempty" json:"server,omitempty"`
}

func (l Link) isLinkOrReference() {}

// Response describes a single response from an API Operation, including design-time, static links to operations based on the response.
type Response struct {
	// REQUIRED. A description of the response. CommonMark syntax MAY be used for rich text representation.
	Description	string		`yaml:"description,omitempty" json:"description,omitempty"`
	// Maps a header name to its definition. [RFC7230] states header names are case insensitive. If a response header is defined with the name "Content-Type", it SHALL be ignored.
	Headers	map[string]HeaderOrReference	`yaml:"headers,omitempty" json:"headers,omitempty"`
	// A map containing descriptions of potential response payloads. The key is a media type or media type range and the value describes it. For responses that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
	Content	map[string]MediaType	`yaml:"content,omitempty" json:"content,omitempty"`
	// A map of operations links that can be followed from the response. The key of the map is a short name for the link, following the naming constraints of the names for Component Objects.
	Link	map[string]LinkOrReference	`yaml:"link,omitempty" json:"link,omitempty"`
}

// Components holds a set of reusable objects for different aspects of the OAS. All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
type Components struct {
	// An object to hold reusable Schema Objects.
	Schemas	map[string]Schema	`yaml:"schemas,omitempty" json:"schemas,omitempty"`
	// An object to hold reusable Response Objects.
	responses	map[string]ResponseOrReference	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// parameters	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// examples	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// requestBodies	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// headers	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// securitySchemes	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// links	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// callbacks	`yaml:"xml,omitempty" json:"xml,omitempty"`


	// pathItems	`yaml:"xml,omitempty" json:"xml,omitempty"`
}

// SecurityRequirement lists the required security schemes to execute this operation. The name used for each property MUST correspond to a security scheme declared in the Security Schemes under the Components Object.
type SecurityRequirement struct {}

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
	// The available paths and operations for the API.
	Paths	Paths	`yaml:"paths,omitempty" json:"paths,omitempty"`
	// The incoming webhooks that MAY be received as part of this API and that the API consumer MAY choose to implement. Closely related to the callbacks feature, this section describes requests initiated other than by an API call, for example by an out of band registration. The key name is a unique string to refer to each webhook, while the (optionally referenced) Path Item Object describes a request that may be initiated by the API provider and the expected responses. An example is available.
	Webhooks	map[string]PathItemOrReference	`yaml:"webhooks,omitempty" json:"webhooks,omitempty"`
	// An element to hold various schemas for the document.
	Components	Components	`yaml:"components,omitempty" json:"components,omitempty"`
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
