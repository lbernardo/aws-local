package cmd

import (
	"github.com/lbernardo/aws-local/pkg/ssm"
	"github.com/spf13/cobra"
)

var ssmCommand = &cobra.Command{
	Use:   "ssm",
	Short: "Security Secret Manager (SSM)",
	Run:   ExecuteSsm,
}

var portSSM string
var hostSSM string
var enviromentSSM string

func init() {
	ssmCommand.PersistentFlags().StringVar(&portSSM, "port", "3003", "port usage [default 3003]")
	ssmCommand.PersistentFlags().StringVar(&hostSSM, "host", "0.0.0.0", "host usage [default 0.0.0.0]")
	ssmCommand.Flags().StringVarP(&enviromentSSM, "values", "v", "","File using for SSM values")
	ssmCommand.MarkFlagRequired("values")
	rootCmd.AddCommand(ssmCommand)
}

func ExecuteSsm(cmd *cobra.Command, args []string) {
	p := ssm.Params{
		Host: hostSSM,
		Port: portSSM,
		EnvFile: enviromentSSM,
	}
	ssmLocal := ssm.NewSSMLocal(p)
	ssmLocal.Start()
}
