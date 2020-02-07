package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"strings"
)

var s3Command = &cobra.Command{
	Use:   "s3",
	Short: "S3",
	Run:   Executes3Command,
}
func init() {
	rootCmd.AddCommand(s3Command)
}

func Executes3Command(cmd *cobra.Command, args []string) {
	//route := mux.NewRouter()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "PUT" {
			ss := strings.Split(request.RequestURI,"/")
			filename := ss[len(ss)-1]
			fmt.Println("filename",filename)

		}
		writer.WriteHeader(200)
		writer.Write([]byte("OK"))
	});

	log.Fatalln(http.ListenAndServe("0.0.0.0:3002",nil))
}
