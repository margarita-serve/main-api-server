package config

// RootOpenAPIDefinition represent RootOpenAPIDefinition
type RootOpenAPIDefinition struct {
	Info OpenAPIInfo `json:"info" yaml:"info"`
}

// OpenAPIInfo represent OpenAPIInfo
type OpenAPIInfo struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Version     string `json:"version" yaml:"version"`
	BasePath    string `json:"basePath" yaml:"basePath"`
}

// OpenAPIContact represent OpenAPIContact
type OpenAPIContact struct {
	Name  string `json:"name" yaml:"name"`
	URL   string `json:"url" yaml:"url"`
	Email string `json:"email" yaml:"email"`
}

// OpenAPILicense represent OpenAPILicense
type OpenAPILicense struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}
