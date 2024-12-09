package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Policy struct {
	MandatoryTags []string            `yaml:"mandatory_tags"`
	AllowedValues map[string][]string `yaml:"allowed_values"`
}

var policy Policy

func init() {

	yamlData, err := os.ReadFile("configs/default-policy.yaml")
	if err != nil {
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlData, &policy); err != nil {
		os.Exit(1)
	}

	// Print the unmarshaled data
	fmt.Printf("%+v\n", policy)

}
