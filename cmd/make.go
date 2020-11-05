package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "This command helps you make components for your web application.",
	Long:  `This command helps you make components for your web application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("make called")
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
