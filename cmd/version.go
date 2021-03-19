package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version for rseason",
	Long:  `version for rseason`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1-20210319")
	},
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
