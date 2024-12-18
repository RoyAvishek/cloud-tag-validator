package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"cloud-tag-validator/pkg/aws"
)

type Policy struct {
	MandatoryTags []string            `yaml:"mandatory_tags"`
	AllowedValues map[string][]string `yaml:"allowed_values"`
}

var policy Policy

func GetPolicyDetails(pathValue string) {

	// adding debug statement to test execution is in this block
	fmt.Println("Fetching policy details")

	yamlData, err := os.ReadFile(pathValue)
	if err != nil {
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlData, &policy); err != nil {
		os.Exit(1)
	}

	// Print the unmarshaled data
	// fmt.Printf("%+v\n", policy)

}

// Function to validate tags against the policy
func ValidateTags() {
	var resourceTags = aws.GetMockTags()

	for _, tags := range resourceTags {
		fmt.Println()
		for key, value := range tags {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

}
