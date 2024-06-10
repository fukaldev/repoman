/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rpmCmd represents the rpm command
var rpmCmd = &cobra.Command{
	Use:   "rpm",
	Short: "Subset of commands for RPM repositories",
	Long: `Subset of commands for RPM repositories. This will allow you to create, delete, and
	list RPM repositories.
	
	Example:
	$ repoman rpm create
	$ Repo Name: my-rpm-repo
	$ RPM repo created at s3://my-rpm-repo
	$ repoman rpm delete
	$ Repo Name: my-rpm-repo
	$ RPM repo deleted at s3://my-rpm-repo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rpm called")
	},
}

func init() {
	rootCmd.AddCommand(rpmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rpmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rpmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
