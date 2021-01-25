package main

import (
	"fmt"
	"github.com/parvez0/wabacli/pkg/cmd/send"
)

func main() {
	file, err := send.NewFileReader("./assets/test.jpeg")
	if err != nil {
		fmt.Printf("failed to read file : %v", err)
		return
	}
	file.Read()
	fmt.Println(file)
}