package cmd

import (
	"github.com/lbernardo/aws-local/pkg/s3"
	"github.com/spf13/cobra"
)

var s3Command = &cobra.Command{
	Use:   "s3",
	Short: "S3",
	Run:   Executes3Command,
}
func init() {
	s3Command.Flags().StringVarP(&volumeS3, "volume", "v", "", "Volume  for storage S3")
	s3Command.PersistentFlags().StringVar(&portS3, "port", "3002", "port usage [default 3002]")
	s3Command.PersistentFlags().StringVar(&hostS3, "host", "0.0.0.0", "host usage [default 0.0.0.0]")
	s3Command.Flags().StringVar(&networkS3, "network", "", "Set network name usage")

	s3Command.MarkFlagRequired("volume")

	rootCmd.AddCommand(s3Command)
}

var volumeS3 string
var portS3 string
var hostS3 string
var networkS3 string


func Executes3Command(cmd *cobra.Command, args []string) {

	s3.NewS3StorageLocal(s3.ParamsS3{
		Volume: volumeS3,
		Host: hostS3,
		Port: portS3,
		Network: networkS3,
	}).StartS3Storage()

}
