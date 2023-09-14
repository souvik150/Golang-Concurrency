package main

import (
	"fmt"
	"time"
)


func server1(ch chan string){
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string){
	for {
		time.Sleep(2 * time.Second)
		ch <- "This is from server 2"
	}
}

func main(){
	fmt.Println("Select with channels");
	fmt.Println("--------------------");

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
			case message1 := <- channel1:
				fmt.Println("Case 1: " ,message1)
			case message2 := <- channel1:
				fmt.Println("Case 2: " ,message2)
			case message3 := <- channel2:
				fmt.Println("Case 3: " ,message3)
			case message4 := <- channel2:
				fmt.Println("Case 4: " ,message4)
			// default:
			// 	fmt.Println("Default case")
		}
	}
}