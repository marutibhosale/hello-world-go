package main

import (
	"bufio"
	"fmt"
	test "myapp/doctore" // import all packages from doctore location or directore
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := test.Intro() // intro is fun from test package

	fmt.Println(whatToSay)

	//userInput, _ := reader.ReadString('\n') // '\n' its taking input up to presss enter

	//fmt.Println(userInput)

	for {
		fmt.Printf("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) // for windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // fot linux and mac
		if userInput == "quit" {
			break
		} else {
			response := test.Response(userInput)
			fmt.Println(response)
		}

	}

}
