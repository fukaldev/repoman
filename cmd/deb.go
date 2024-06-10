/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// debCmd represents the deb command
var debCmd = &cobra.Command{
	Use:   "deb",
	Short: "Subset of commands for Debian repositories",
	Long: `Subset of commands for Debian repositories. This will allow you to create, delete, and
	list Debian repositories.
	
	Example:
	$ repoman deb create
	$ Repo Name: my-debian-repo
	$ Debian repo created at s3://my-debian-repo
	$ repoman deb delete
	$ Repo Name: my-debian-repo
	$ Debian repo deleted at s3://my-debian-repo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deb called")
	},
}

func init() {
	rootCmd.AddCommand(debCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
