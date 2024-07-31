package main

import "fmt"

// Represents the instruction types in MIPS
type InstructionType int

// MIPS instruction types
const (
	R_TYPE InstructionType = iota // Register (R-format) instructions
	I_TYPE                        // Immediate (I-format) instructions
	J_TYPE                        // Jump (J-format) instructions
)

// Represents an instruction in MIPS
type Instruction struct {
	Opcode uint8
	Format InstructionType
}

// Instruction set
var instructionSet = map[string]Instruction{
	"add": {0b100000, R_TYPE},
	"sub": {0b100010, R_TYPE},
	"and": {0b100100, R_TYPE},
	"or":  {0b100101, R_TYPE},
	"xor": {0b100110, R_TYPE},
	"nor": {0b100111, R_TYPE},
	"slt": {0b101010, R_TYPE},
	"j":   {0b000010, J_TYPE},
	"jal": {0b000011, J_TYPE},
}

// Register map
var registerMap = map[string]int{
	"$zero": 0, "$at": 1, "$v0": 2, "$v1": 3,
	"$a0": 4, "$a1": 5, "$a2": 6, "$a3": 7,
	"$t0": 8, "$t1": 9, "$t2": 10, "$t3": 11,
	"$t4": 12, "$t5": 13, "$t6": 14, "$t7": 15,
	"$s0": 16, "$s1": 17, "$s2": 18, "$s3": 19,
	"$s4": 20, "$s5": 21, "$s6": 22, "$s7": 23,
	"$t8": 24, "$t9": 25, "$k0": 26, "$k1": 27,
	"$gp": 28, "$sp": 29, "$fp": 30, "$ra": 31,
}

// Represents a MIPS assembly instruction
type AssemblyInstruction struct {
	instruction string // The instruction name
	source      string // The source register
	target      string // The target register
	destination string // The destination register
	shift       uint8  // The shift amount
	function    uint8  // The function code
	immediate   int16  // The immediate value
	address     uint32 // The address value for jump instructions
}

func synthesize(asm AssemblyInstruction) (uint32, error) {
	var encodedInstruction = uint32(0)
	instruction, ok := instructionSet[asm.instruction]
	if !ok {
		return 0, fmt.Errorf("invalid instruction: %s", asm.instruction)
	}

	switch instruction.Format {
	case R_TYPE:
		encodedInstruction |= uint32(instruction.Opcode) << 26
	case I_TYPE:
		encodedInstruction |= uint32(instruction.Opcode) << 26
	case J_TYPE:
		encodedInstruction |= uint32(instruction.Opcode) << 26
	default:
		return 0, fmt.Errorf("unknown instruction type: %d", instruction.Format)
	}
	return encodedInstruction, nil
}
