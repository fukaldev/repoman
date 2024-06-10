/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pyCmd represents the py command
var pyCmd = &cobra.Command{
	Use:   "py",
	Short: "Subset of commands for Python repositories",
	Long: `Subset of commands for Python repositories. This will allow you to create, delete, and
	list Python repositories.
	
	Example:
	$ repoman py create
	$ Repo Name: my-python-repo
	$ Python repo created at s3://my-python-repo
	$ repoman py delete
	$ Repo Name: my-python-repo
	$ Python repo deleted at s3://my-python-repo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("py called")
	},
}

func init() {
	rootCmd.AddCommand(pyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
