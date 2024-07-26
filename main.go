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

	// Scan the file to generate tokens
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			dfa.Transition(r)
			// fmt.Print(i, r)
			// fmt.Printf("Index: %d, Rune: %c\n", i, r)
			byteCount += 1
		}
		dfa.Store()
		dfa.StoreLine()
	}

	// Parsing the tokens
	for _, token := range dfa.totalTokens {
		for _, t := range token {
			if t.tokenType == "LabelDef" {
				if _, ok := symbolTable[t.tokenValue]; ok {
					// fmt.Errorf("Duplicate label definition: %s", t.tokenValue)
				} else {
					symbolTable[t.tokenValue] = int32(lineCount)
				}
			}
			fmt.Printf("Token Type: %s, Value: %s | ", t.tokenType, t.tokenValue)
		}
		fmt.Println()
		lineCount += 1
	}
	// list := []int{10, 20, 30, 40, 50}
	// for i := 0; i < len(list); i++ {
	// 	fmt.Println(list[i])
	// }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d lines, %d bytes\n", lineCount, byteCount)
}
