// init funtion example : 
package main

import "fmt"

var count int

func main() {

	fmt.Printf("I am printing at last i think")
}

func init() {
	count++
	fmt.Printf("I am exeuting first my number is %d\n", count)
}

func init() {
	count++
	fmt.Printf("I am exeuting second my number is %d\n", count)
}
