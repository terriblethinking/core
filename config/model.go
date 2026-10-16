package config

type Model struct {
	name string
	identfier string
	provider string
}

func GetDefaultModel(profile string) Model {
	return Model{
		name:	"Ornith Small",
		identfier: "ornith:9b",
		provider: "ollama",
	}

}
