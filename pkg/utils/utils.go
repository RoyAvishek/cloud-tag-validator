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

func GetPolicyDetails(pathValue string) Policy {

	// adding debug statement to test execution is in this block
	fmt.Println("Fetching policy details.....")

	yamlData, err := os.ReadFile(pathValue)
	if err != nil {
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlData, &policy); err != nil {
		os.Exit(1)
	}

	return policy

}

// Function to validate tags against the policy
func ValidateTags() {
	var resourceTags = aws.GetMockTags()

	// Fetch policy details
	policy := GetPolicyDetails("configs/default-policy.yaml")

	for _, tags := range resourceTags {
		resourceARN := tags["ResourceARN"]
		tagMap := tags["Tags"].(map[string]string)

		// Check for missing mandatory tags
		missingTags := []string{}
		for _, mandatoryTag := range policy.MandatoryTags {
			if _, exists := tagMap[mandatoryTag]; !exists {
				missingTags = append(missingTags, mandatoryTag)
			}

		}

		if len(missingTags) > 0 {
			fmt.Printf("Resource %s is missing mandatory tags: %v\n", resourceARN, missingTags)
			continue
		}

		// Check for invalid tag values
		invalidTags := []string{}
		for key, value := range tagMap {
			if allowedValues, exists := policy.AllowedValues[key]; exists {
				if !contains(allowedValues, value) {
					invalidTags = append(invalidTags, fmt.Sprintf("%s: %s", key, value))
				}
			}
		}

		if len(invalidTags) > 0 {
			fmt.Printf("Resource %s has invalid tag values: %v\n", resourceARN, invalidTags)
		} else {
			fmt.Printf("Resource %s has valid tags.\n", resourceARN)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
