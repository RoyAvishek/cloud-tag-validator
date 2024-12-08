package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "1.0"

var rootCmd = &cobra.Command{
	Use:     "cloud-resource-tagging-validator",
	Short:   "Cloud Resource Tagging Validator",
	Long:    "Cloud Resource Tagging Validator",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a Cobra based String Utility.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
