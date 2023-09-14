package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printsomething(t *testing.T){
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printSomething("hello world", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "hello world") {
		t.Errorf("Expected \"hello world\", got %s", output)
	}
}

func Test_updateMessage(t *testing.T){
	wg.Add(1)

	go updateMessage("Hello, universe!")

	wg.Wait()

	if msg != "Hello, universe!" {
		t.Errorf("Expected \"Hello, universe!\", got %s", msg)
	}

	wg.Add(1)

	go updateMessage("Hello, cosmos!")

	wg.Wait()

	if msg != "Hello, cosmos!" {

		t.Errorf("Expected \"Hello, cosmos!\", got %s", msg)
	}

}


func Test_printMessage(t *testing.T){
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected \"epsilon\", got %s", output)
	}
}