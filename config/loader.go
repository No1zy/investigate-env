package config

type Loader interface{
	Load(string) error
}

// configuration loader
func Load(path string) error {
	var loader Loader
	loader = YamlLoader{}
	return loader.Load(path)
}
