package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start The Http Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start The Http Server")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
