package cmd

import (
	"fmt"
	"log"
	"os"
)

func Execute() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
