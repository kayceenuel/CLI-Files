package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Execute() {
	Readfiles()
}

func ReadFiles() {
	if len(os.Args) > 1 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Cant't read file:", os.Args[1])
		panic(err)
	}
	fmt.Println("File content is:")
	os.Stdout.Write(strings(data))
}
