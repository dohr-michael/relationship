package schema

type Schema struct {
	// Correspond to the version.
	Id string `json:"id"`
	// Grouping of schema
	Type string `json:"type"`
	// Name of the specifics version of the schema
	Name string `json:"name"`
	// Information about this version.
	VersionNote string `json:"versionNote"`
	// Schema configuration
	Configuration SchemaConfiguration `json:"configuration"`
	// Schema definition
	Definition SchemaDefinition `json:"definition"`
}

type SchemaConfiguration map[string]interface{}
type SchemaDefinition interface{}

func NewSchema(id string, name string, t string, note string, configuration SchemaConfiguration, definition SchemaDefinition) Schema {
	return Schema{
		Id:            id,
		Type:          t,
		Name:          name,
		VersionNote:   note,
		Configuration: configuration,
		Definition:    definition,
	}
}
