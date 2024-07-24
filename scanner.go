package main

import (
	"strings"
	"unicode"
)

// Represents the possible states of the DFA.
type State int

// Represents the possible states of the DFA.
const (
	Initial State = iota
	Identifier
	DotIdentifier
	Register
	Zero
	Decimal
	Hexadecimal
	Comma
	LParen
	RParen
	LabelDef
	Comment
)

// String method to convert the current state to a string.
func (s State) String() string {
	return [...]string{"Initial", "Identifier", "DotIdentifier", "Register", "Zero", "Decimal", "Hexadecimal", "Comma", "LParen", "RParen", "LabelDef", "Comment"}[s]
}

// Represents a token in the scanner.
type Token struct {
	tokenType  string
	tokenValue string
}

// Represents the scanning Deterministic Finite Automaton(DFA) for the scanner.
type DFA struct {
	state        State
	currentToken string
	tokens       []Token
}

// Constructor to initialize memory for the DFA.
func NewDFA() (*DFA, error) {
	dfa := &DFA{}
	dfa.state = Initial
	dfa.currentToken = ""
	dfa.tokens = []Token{}
	return dfa, nil
}

// Add a token to the list of tokens.
func (dfa *DFA) AddToken(tokenType string, tokenValue string) {
	dfa.tokens = append(dfa.tokens, Token{tokenType, tokenValue})
}

// Store the current state of the DFA.
func (dfa *DFA) Store() {
	if dfa.state != Initial && dfa.state != Comment {
		dfa.AddToken(dfa.state.String(), dfa.currentToken)
	}
	dfa.Reset()
}

// Reset the DFA to its initial state.
func (dfa *DFA) Reset() {
	dfa.state = Initial
	dfa.currentToken = ""
	// dfa.tokens = []string{}
}

// Transition the DFA to a new state based on the input.
func (dfa *DFA) Transition(input rune) {
	// fmt.Printf("State: %s, Rune: '%c'\n", dfa.state.String(), input)
	switch dfa.state {
	case Initial:
		if input == '.' {
			dfa.state = DotIdentifier
		} else if input == '$' {
			dfa.state = Register
		} else if input == '0' {
			dfa.state = Zero
		} else if unicode.IsDigit(input) || input == '-' {
			dfa.currentToken = string(input)
			dfa.state = Decimal
		} else if input == ',' {
			dfa.AddToken(dfa.state.String(), ",")
			dfa.Reset()
		} else if input == '(' {
			dfa.AddToken(dfa.state.String(), "(")
			dfa.Reset()
		} else if input == ')' {
			dfa.AddToken(dfa.state.String(), ")")
			dfa.Reset()
		} else if input == '#' {
			dfa.Store()
			dfa.state = Comment
		} else if !unicode.IsSpace(input) {
			dfa.currentToken = string(input)
			dfa.state = Identifier
		}
	case Identifier:
		if input == ':' {
			dfa.state = LabelDef
			dfa.Store()
		} else if !unicode.IsSpace(input) {
			dfa.currentToken += string(input)
		} else {
			dfa.Store()
		}
	case DotIdentifier:
		if unicode.IsLetter(input) {
			dfa.currentToken += string(input)
		} else {
			dfa.Store()
		}
	case Register:
		if input == ',' {
			dfa.Store()
			dfa.AddToken(dfa.state.String(), ",")
		} else if unicode.IsDigit(input) || unicode.IsLetter(input) {
			dfa.currentToken += string(input)
		} else {
			dfa.Store()
		}
	case Zero:
		if input == 'x' {
			dfa.state = Hexadecimal
		} else if unicode.IsDigit(input) {
			dfa.currentToken = string(input)
			dfa.state = Decimal
		} else {
			dfa.Store()
		}
	case Decimal:
		if unicode.IsDigit(input) {
			dfa.currentToken += string(input)
		} else {
			dfa.Store()
		}
	case Hexadecimal:
		if strings.ContainsAny(string(input), "0123456789abcdefABCDEF") {
			dfa.currentToken += string(input)
		} else {
			dfa.Store()
		}
	}
}
