package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	runn := true

	for runn {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(strings.Trim(text, "\r\n"))

		if strings.Trim(text, "\r\n") == "x" {
			runn = false
		}
	}

}
