/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/repoman/internal/awshelper"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createDebCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Debian repository",
	Long: `Create a new Debian repository. This will create a new S3 bucket with the
	appropriate permissions to be used as a Debian repository.
	
	Example:
	$ repoman create deb
	$ Repo Name: my-debian-repo
	$ Debian repo created at s3://my-debian-repo`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Repo Name: ")
		reponame, _ := reader.ReadString('\n')

		repoPath, err := awshelper.CreateBucketWithPublicAccess(reponame, "deb")

		if err != nil {
			fmt.Println("Error creating Debian repo,", err)
			return
		}

		fmt.Println("Debian repo created at", *repoPath)
	},
}

var createRpmCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new RPM repository",
	Long: `Create a new RPM repository. This will create a new S3 bucket with the
	appropriate permissions to be used as a RPM repository.
	
	Example:
	$ repoman create rpm
	$ Repo Name: my-rpm-repo
	$ RPM repo created at s3://my-rpm-repo`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Repo Name: ")
		reponame, _ := reader.ReadString('\n')

		repoPath, err := awshelper.CreateBucketWithPublicAccess(reponame, "rpm")

		if err != nil {
			fmt.Println("Error creating RPM repo,", err)
			return
		}

		fmt.Println("RPM repo created at", *repoPath)
	},
}

var createPyCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Python repository",
	Long: `Create a new Python repository. This will create a new S3 bucket with the
	appropriate permissions to be used as a Python repository.
	
	Example:
	$ repoman create py
	$ Repo Name: my-python-repo
	$ Python repo created at s3://my-python-repo`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Repo Name: ")
		reponame, _ := reader.ReadString('\n')

		repoPath, err := awshelper.CreateBucketWithPublicAccess(reponame, "py")

		if err != nil {
			fmt.Println("Error creating Python repo,", err)
			return
		}

		fmt.Println("Python repo created at", *repoPath)
	},
}

func init() {
	debCmd.AddCommand(createDebCmd)
	rpmCmd.AddCommand(createRpmCmd)
	pyCmd.AddCommand(createPyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
