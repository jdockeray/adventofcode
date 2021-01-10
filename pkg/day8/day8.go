package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Instruction struct {
	command string
	unit    int
	symbol  string
}

func ParseInstruction(str string) Instruction {
	strs := strings.Split(str, " ")
	command := strs[0]
	symbol := strs[1][0:1]
	unit, _ := strconv.Atoi(strs[1][1:len(strs[1])])
	return Instruction{
		command: command,
		unit:    unit,
		symbol:  symbol,
	}
}

func plusMinus(acc int, symbol string, unit int) int {
	if symbol == "+" {
		return acc + unit
	}
	return acc - unit
}

func readInstruction(index int, acc int, instructions []Instruction, completed map[int]bool) (int, bool) {
	if index >= len(instructions) || index < 0 {
		return acc, true
	}
	instruction := instructions[index]
	complete := completed[index]
	if complete {
		return acc, false
	} else {
		completed[index] = true
	}

	if instruction.command == "jmp" {
		index = plusMinus(index, instruction.symbol, instruction.unit)
		return readInstruction(index, acc, instructions, completed)
	}
	if instruction.command == "acc" {
		acc = plusMinus(acc, instruction.symbol, instruction.unit)
	}
	return readInstruction(index+1, acc, instructions, completed)
}

func ReadInstructions(instructions []Instruction) (int, bool) {
	return readInstruction(0, 0, instructions, make(map[int]bool))
}

func FindPath(instructions []Instruction) (int, bool) {
	acc, finished := ReadInstructions(instructions)
	if finished {
		return acc, finished
	}

	for index, instruction := range instructions {
		instructions[index] = Instruction{
			command: "nop",
			unit:    0,
			symbol:  "+",
		}
		acc, finished := ReadInstructions(instructions)
		if finished {
			return acc, finished
		} else {
			instructions[index] = instruction
		}
	}
	return 0, false
}

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/data/day8.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var instructions []Instruction

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		str := scanner.Text()
		instruction := ParseInstruction(str)

		if instructions == nil {
			instructions = []Instruction{instruction}
		} else {
			instructions = append(instructions, instruction)
		}
	}
	fmt.Println(FindPath(instructions))
}
