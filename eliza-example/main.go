package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

	// Intro for Eliza
	intro := doctor.Intro()
	fmt.Println("Introduction: ", intro)

	// Reader for taking input from user
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("--> ")
		userInput, _ := reader.ReadString('\n')

		// For windows
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		// For Mac/BSD/Linux
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println("response: ", doctor.Response(userInput))
		}
	}
}
