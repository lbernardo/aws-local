package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "awslocal",
	Short:"Start modules awslocal",
}

func init() {

}

func Execute() error {
	return rootCmd.Execute()
}

