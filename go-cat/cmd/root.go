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
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read files:", os.Args[1])
		return
	}
	os.Stderr.Write(data)
}
