package cmd

import (
	"github.com/lbernardo/aws-local/internal/adapters/secondary/yaml"
	//"github.com/lbernardo/aws-local/pkg/apigateway"
	"github.com/lbernardo/aws-local/pkg/apigateway"
	"github.com/spf13/cobra"
)

var apiGatewayCommand = &cobra.Command{
	Use:   "api-gateway",
	Short: "API Gateway with Lambda",
	Run:   ExecuteApiGateway,
}

var serverlessFile string
var volume string
var port string
var host string
var network string
var enviroment string

func init() {
	apiGatewayCommand.PersistentFlags().StringVar(&serverlessFile, "yaml","serverless.yml", "Serverless file yaml")
	apiGatewayCommand.Flags().StringVarP(&volume, "volume", "v", "", "Volume project (Ep: --volume $PWD) (required)")
	apiGatewayCommand.PersistentFlags().StringVar(&port, "port", "3000", "port usage [default 3000]")
	apiGatewayCommand.PersistentFlags().StringVar(&host, "host", "0.0.0.0", "host usage [default 0.0.0.0]")
	apiGatewayCommand.PersistentFlags().StringVar(&enviroment, "env", "", "File for using environment variables other than serverless. Can replace serverless variables")
	apiGatewayCommand.Flags().StringVar(&network, "network", "", "Set network name usage")

	apiGatewayCommand.MarkFlagRequired("volume")

	rootCmd.AddCommand(apiGatewayCommand)
}

func ExecuteApiGateway(cmd *cobra.Command, args []string) {
	params := apigateway.ParamsApiGateway{
		Serverless:     yaml.GetServerlessFramework(serverlessFile, enviroment),
		Volume:         volume,
		Port:           port,
		Host:           host,
		Network:        network,
	}
	apigateway.StartApiGateway(params)
}
