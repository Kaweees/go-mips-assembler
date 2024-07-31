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
	opcode uint8
	format InstructionType
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
	op          uint8  // The opcode of the instruction
	rs          string // The source register
	rt          string // The target register
	rd          string // The destination register
	shamt       uint8  // The shift amount
	funct       uint8  // The function code
	imm         int16  // The immediate value
	addr        int32  // The address value for jump instructions
}

// Synthesize an R-type instruction
func synthesizeRType(asm AssemblyInstruction, instruction Instruction) (uint32, error) {
	source, ok := registerMap[asm.rs]
	if !ok {
		return 0, fmt.Errorf("invalid source register: %s", asm.rs)
	}

	target, ok := registerMap[asm.rt]
	if !ok {
		return 0, fmt.Errorf("invalid target register: %s", asm.rt)
	}

	destination, ok := registerMap[asm.rd]
	if !ok {
		return 0, fmt.Errorf("invalid destination register: %s", asm.rd)
	}
	encodedInstruction := uint32(instruction.opcode) << 26
	encodedInstruction |= uint32(source) << 21
	encodedInstruction |= uint32(target) << 16
	encodedInstruction |= uint32(destination) << 11
	encodedInstruction |= uint32(asm.shamt) << 6
	encodedInstruction |= uint32(asm.funct)
	return encodedInstruction, nil
}

// Synthesize an I-type instruction
func synthesizeIType(asm AssemblyInstruction, instruction Instruction) (uint32, error) {
	source, ok := registerMap[asm.rs]
	if !ok {
		return 0, fmt.Errorf("invalid source register: %s", asm.rs)
	}

	target, ok := registerMap[asm.rt]
	if !ok {
		return 0, fmt.Errorf("invalid target register: %s", asm.rt)
	}
	encodedInstruction := uint32(instruction.opcode) << 26
	encodedInstruction |= uint32(source) << 21
	encodedInstruction |= uint32(target) << 16
	encodedInstruction |= (uint32(asm.imm) & 0xFFFF)
	return encodedInstruction, nil
}

// Synthesize a J-type instruction
func synthesizeJType(asm AssemblyInstruction, instruction Instruction) (uint32, error) {
	encodedInstruction := uint32(instruction.opcode) << 26
	encodedInstruction |= uint32(asm.addr) & 0x3FFFFFF
	return encodedInstruction, nil
}

func synthesize(asm AssemblyInstruction) (uint32, error) {
	instruction, ok := instructionSet[asm.instruction]
	if !ok {
		return 0, fmt.Errorf("invalid instruction: %s", asm.instruction)
	}

	switch instruction.format {
	case R_TYPE:
		encodedInstruction, err := synthesizeRType(asm, instruction)
		if err != nil {
			return 0, fmt.Errorf("error synthesizing R-type instruction: %s", err)
		} else {
			return encodedInstruction, nil
		}
	case I_TYPE:
		encodedInstruction, err := synthesizeIType(asm, instruction)
		if err != nil {
			return 0, fmt.Errorf("error synthesizing I-type instruction: %s", err)
		} else {
			return encodedInstruction, nil
		}
	case J_TYPE:
		encodedInstruction, err := synthesizeJType(asm, instruction)
		if err != nil {
			return 0, fmt.Errorf("error synthesizing J-type instruction: %s", err)
		} else {
			return encodedInstruction, nil
		}
	default:
		return 0, fmt.Errorf("unknown instruction type: %d", instruction.format)
	}
}
