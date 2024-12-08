package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate cloud resource tags against defined policies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Validating AWS resources tags...")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
