package main

import (
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
		Log.Error(fmt.Sprintf("Error initializing the DFA: %v", err))
		return
	}

	// Scan the file
	Log.Info(fmt.Sprintf("Scanner of %s initialized", file.Name()))
	err = scanFile(file, dfa)
	if err != nil {
		Log.Error(fmt.Sprintf("Error scanning file: %v", err))
		return
	}

	// Parsing the tokens
	err = parseTokens(dfa)
	Log.Info(fmt.Sprintf("Parsing tokens of %s initalized", file.Name()))
	if err != nil {
		Log.Error(fmt.Sprintf("Error parsing tokens: %v", err))
		return
	}

	// list := []int{10, 20, 30, 40, 50}
	// for i := 0; i < len(list); i++ {
	// 	fmt.Println(list[i])
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// Synthesize the tokens
	// err = synthesizeTokens(dfa)
	// fmt.Printf("%d lines, %d bytes\n", lineCount, byteCount)
}
