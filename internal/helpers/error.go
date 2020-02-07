package helpers

import (
	"fmt"
	"os"
)

func PrintError(err error) {
	if os.Getenv("DEBUG") == "true" {
		panic(err)
	}
	fmt.Printf("\033[01;31m%v\033[0m\n",err.Error())
	os.Exit(1)
}