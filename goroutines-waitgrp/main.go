package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}


func main() {
	
	// words := []string{
	// 	"hello", 
	// 	"world",
	// 	"alice",
	// 	"bob",
	// 	"charlie",
	// 	"eve",
	// 	"mallory",
	// }

	// wg.Add(len(words))

	// for _, word := range words {
	// 	go printSomething(fmt.Sprint(word), &wg)
	// }

	// wg.Wait()

	// wg.Add(1)
	// printSomething("This is the second line of hello world", &wg)

	msg = "Hello, world!"

	
 	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()
	
	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()
}

// package main

// import (
// 	"fmt"
// )

// var msg string

// func updateMessage(s string) {
// 	msg = s
// }

// func printMessage() {
// 	fmt.Println(msg)
// }

// func main() {

// 	 challenge: modify this code so that the calls to updateMessage() on lines
// 	28, 30, and 33 run as goroutines, and implement wait groups so that
// 	the program runs properly, and prints out three different messages.
// 	Then, write a test for all three functions in this program: updateMessage(),
//  printMessage(), and main().

// 	msg = "Hello, world!"

// 	updateMessage("Hello, universe!")
// 	printMessage()

// 	updateMessage("Hello, cosmos!")
// 	printMessage()

// 	updateMessage("Hello, world!")

// 	printMessage()
// }
