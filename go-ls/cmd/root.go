package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
	// Parse flags first
	help := flag.Bool("h", false, "show help for go-ls")
	flag.Parse()

	if *help {
		fmt.Println("go-ls: list files in a directory\nUsage: go-ls [directory]")
		return
	}

	// Get directory from args or default to "."
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1] // Note: After flag.Parse(), os.Args might need adjustment
	}

	printFiles(dir)
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
