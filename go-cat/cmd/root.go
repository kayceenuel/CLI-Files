package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	ReadFiles()
}

func ReadFiles() {
	if len(os.Args) <= 1 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
}
