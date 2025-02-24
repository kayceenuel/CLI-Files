package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	printFiles(".")
}

func printFiles(dir string) {
	// check if it's a file or directory
	info, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if !info.IsDir() {
		fmt.Println(dir) // just print file name
		return
	}
	// it's a directory, so list content
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error in accessing directory:", err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name()) // print each file name.
	}
}
