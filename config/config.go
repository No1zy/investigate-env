package config

// Template variable
type Variable struct {
	//URL string
	VARIABLE string
}

var Conf DockerCopomseConfig

type DockerCopomseConfig struct {
	Services []Service
}

type Service struct {
	Name string
}
