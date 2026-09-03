package config

type ToolParam struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Required bool   `yaml:"required"`
}

type ToolConfig struct {
	Version     string      `yaml:"version"`
	Description string      `yaml:"desc"`
	Name        string      `yaml:"name"`
	Params      []ToolParam `yaml:"params"`
}
