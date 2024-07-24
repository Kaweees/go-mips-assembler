package main

type Instruction struct {
	Opcode uint8
	Format string
}

var instructionSet = map[string]Instruction{
	// Register (R-format) instructions
	"add": {0b100000, "R"},
	"sub": {0b100010, "R"},
	"and": {0b100100, "R"},
	"or":  {0b100101, "R"},
	"xor": {0b100110, "R"},
	"nor": {0b100111, "R"},
	"slt": {0b101010, "R"},
	// Immediate (I-format) instructions

	// Jump (J-format) instructions
	"j":   {0b000010, "J"},
	"jal": {0b000011, "J"},
}

var symbolTable = map[string]any{}
