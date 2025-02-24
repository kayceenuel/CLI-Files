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
}
