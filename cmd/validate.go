package cmd

import (
	"cloud-tag-validator/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var pathValue string

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate cloud resource tags against defined policies",
	Run: func(cmd *cobra.Command, args []string) {
		pathValue, _ := cmd.Flags().GetString("path")
		// print path details for debugging
		if pathValue != "" {
			utils.ValidateTags()
		} else {
			fmt.Println("Not able to fetch the tags")
		}
	},
}

// init adds the validate command to the root command.
func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.PersistentFlags().StringVarP(&pathValue, "path", "f", "configs/default-policy.yaml", "Path to the policy YAML file")
}
