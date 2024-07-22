package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("demo.asm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineCount := 0
	byteCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount += 1
		byteCount += len(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d lines, %d bytes", lineCount, byteCount)
}
