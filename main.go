package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var dfa *DFA
	var err error

	// Initialize the logger
	initalizeLogger()

	file, err := os.Open("demo.asm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize the dfa
	dfa, err = NewDFA()
	if err != nil {
		Log.Error(fmt.Sprintf("Error initializing the request handler: %v", err))
		return
	}

	lineCount := 0
	byteCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dfa.Reset()
		for i, r := range scanner.Text() {
			dfa.Transition(r)
			// fmt.Print(i, r)
			fmt.Printf("Index: %d, Rune: %c\n", i, r)
			byteCount += 1
		}
		lineCount += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d lines, %d bytes\n", lineCount, byteCount)
}
