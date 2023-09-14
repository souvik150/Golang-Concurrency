package main

import (
	"fmt"
	"time"
)

func lisenToChan(ch chan int){
	for{
		// print a got data msg
		i := <- ch
		fmt.Println("Got data", i, "from channel")
		
		// simulate some work
		time.Sleep(time.Second * 1)
	}
}

func main(){
	ch := make(chan int, 100)
	go lisenToChan(ch)

	for i:= 0; i<= 100; i++ {
		fmt.Println("Sending data", i, "to channel")
		ch <- i
		fmt.Println("Sent data", i, "to channel")
	}

	fmt.Println("Done")
	close(ch)
}