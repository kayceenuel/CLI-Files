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
	for _, file := range os.Args[1:] {
		info, err := os.Stat(file)
		if err != nil {
			fmt.Println("go-cat: ", file, ": ", err)
			continue
		}
		if info.IsDir() {
			fmt.Println("go-cat: ", file, ": is a directory")
			continue
		}
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Can't read files:", os.Args[1])
			return
		}
		os.Stdout.Write(data)
	}
}
