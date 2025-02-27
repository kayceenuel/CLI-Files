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
	// loop over arguments
	for _, file := range os.Args[1:] {
		info, err := os.Args(file)
		if err != nil {
			fmt.Println(file.Name())
			return
		}
		if !info.IsDir() {
			fmt.Println("go-cat: <filePath>: is a directory", os.Args[1])
			return
		}
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read files:", os.Args[1])
		return
	}
	os.Stderr.Write(data)
}
