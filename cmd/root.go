package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "awslocal",
	Short:"Start modules awslocal",
	Long:"Start modules aws for local developer or test.\nCreated by Lucas Bernardo <lbernardo.brito@gmail.com>",
}

func init() {

}

func Execute() error {
	return rootCmd.Execute()
}

