package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	op string
	n  int
}

const ACC, NOP, JMP = "acc", "nop", "jmp"

func run(code *[]instruction) (acc int, looped bool) {
	i := 0
	visited := map[int]bool{}
	for i < len(*code) {
		_, ok := visited[i]
		if ok {
			return acc, true
		}
		visited[i] = true
		instr := (*code)[i]
		switch instr.op {
		case ACC:
			acc += instr.n
			i++
		case NOP:
			i++
		case JMP:
			i += instr.n
		}
	}
	return acc, false
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.Split(strings.TrimSpace(string(data)), "\n")
	code := make([]instruction, 0, 700)
	for _, line := range text {
		line := strings.Fields(line)
		n, _ := strconv.Atoi(line[1])
		code = append(code, instruction{line[0], n})
	}

	acc, _ := run(&code)
	fmt.Println("Accumulator at first loop:", acc)

	for i := 0; i < len(code); i++ {
		instr := code[i]
		if instr.op == ACC {
			continue
		}
		if instr.op == NOP {
			code[i].op = JMP
		} else {
			code[i].op = NOP
		}
		acc, looped := run(&code)
		if !looped {
			fmt.Println("Final accumulator in fixed code:", acc)
			break
		}
		code[i].op = instr.op
	}
}
