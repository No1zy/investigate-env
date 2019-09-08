package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"reflect"
)

type YamlLoader struct {}

// yaml loader
func (c YamlLoader) Load(pathToYAML string) error {
	
	yamlInput, err := ioutil.ReadFile(pathToYAML)

	if err != nil {
		return err
	}

	var config interface{}
	if err := yaml.Unmarshal(yamlInput, &config); err != nil {
		return err
	}
	
	setConfig(config)

	return nil
}

// set config
func setConfig(config interface {}) {
	services := []Service{}

	v := reflect.ValueOf(config)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			if key.Interface().(string) == "services" {
				strct := v.MapIndex(key)
				services = getServices(strct.Interface())
				break
			}
		}
	}

	Conf.Services = services

}

//get services from docker-compose.yml
func getServices(in interface {}) []Service {
	v := reflect.ValueOf(in)

	services := []Service{}

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			services = append(services, Service{Name: key.Interface().(string)})
		}
	}
	return services
}
