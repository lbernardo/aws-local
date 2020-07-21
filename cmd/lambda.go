package cmd

import (
	"os"

	"github.com/lbernardo/aws-local/internal/adapters/secondary/env"
	"github.com/lbernardo/aws-local/internal/helpers"
	"github.com/lbernardo/aws-local/pkg/lambda"
	"github.com/spf13/cobra"
)

var lambdaCommand = &cobra.Command{
	Use:   "lambda [bin]",
	Short: "Run Lambda event",
	Run:   ExecuteLambda,
}
var fileEnv string
var runtime string
var volumeLambda string
var bodyExecute string

func init() {
	lambdaCommand.PersistentFlags().StringVar(&fileEnv, "env", "", "File for using environment variables")
	lambdaCommand.Flags().StringVarP(&volumeLambda, "volume", "v", "", "Volume project (Ep: --volume $PWD) (required)")
	lambdaCommand.PersistentFlags().StringVar(&runtime, "runtime", "go1.x", "Runtime lambda")
	lambdaCommand.PersistentFlags().StringVar(&bodyExecute, "body", "", "Content body")
	rootCmd.AddCommand(lambdaCommand)
}

func ExecuteLambda(cmd *cobra.Command, args []string) {
	lambda.StartLambda(lambda.ParamsLambda{
		Env:     getEnviroment(),
		Bin:     args[0],
		Runtime: runtime,
		Volume:  volumeLambda,
		Body:    bodyExecute,
	})
}

func getEnviroment() map[string]string {
	if fileEnv != "" {
		file, err := os.Open(fileEnv)
		if err != nil {
			helpers.PrintError(err)
		}
		defer file.Close()
		envMap, err := env.Parse(file)
		if err != nil {
			helpers.PrintError(err)
		}
		contentEnv := make(map[string]string, 0)
		for envName, envValue := range envMap {
			contentEnv[envName] = envValue
		}
		return contentEnv
	}
	return map[string]string{}
}
