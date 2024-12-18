package aws

var mockTags = []map[string]interface{}{
	{
		"ResourceARN": "arn:aws:ec2:us-east-1:123456789012:instance/i-12345678",
		"Tags": map[string]string{
			"Environment": "Production",
			"Owner":       "TeamA",
			"AppName":     "AppA",
			"Product":     "ProductA",
		},
	},
	{
		"ResourceARN": "arn:aws:ec2:us-east-1:0987654321:instance/i-12345679",
		"Tags": map[string]string{
			"Environment": "Development",
			"Owner":       "TeamB",
			"AppName":     "AppB",
			"Product":     "ProductB",
		},
	},
}

func GetMockTags() []map[string]string {
	var tagsList []map[string]string
	for _, tag := range mockTags {
		if t, ok := tag["Tags"].(map[string]string); ok {
			tagsList = append(tagsList, t)
		}
	}
	return tagsList
}
