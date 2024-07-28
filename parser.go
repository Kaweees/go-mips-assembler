package main

import "fmt"

// Symbol lookup table
var symbolTable = map[string]int32{}

func parseTokens(dfa *DFA) error {
	lineCount := 0
	for _, token := range dfa.totalTokens {
		for _, t := range token {
			if t.tokenType == "DotIdentifier" {
				if _, ok := symbolTable[t.tokenValue]; ok {
					return fmt.Errorf("duplicate label definition: %s", t.tokenValue)
				} else {
					symbolTable[t.tokenValue] = int32(lineCount)
				}
			}
		}
		lineCount += 1
	}
	return nil
}
