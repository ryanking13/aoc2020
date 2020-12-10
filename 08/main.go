package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	inputs, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	ops := strings.Split(string(inputs), "\n")

	// part 1

	visited := make([]bool, len(ops), len(ops))
	idx := 0
	acc := 0
	for {
		if idx == len(ops) {
			break
		}
		if visited[idx] {
			break
		}
		visited[idx] = true

		_op := strings.Split(ops[idx], " ")
		op := _op[0]
		val, err := strconv.Atoi(_op[1])
		if err != nil {
			panic(err)
		}
		switch op {
		case "nop":
			idx += 1
		case "acc":
			idx += 1
			acc += val
		case "jmp":
			idx += val
		}
	}

	fmt.Println(acc)

	// part 2

	opsS := make([]string, len(ops), len(ops))
	opsV := make([]int, len(ops), len(ops))
	for i := 0; i < len(ops); i += 1 {
		_op := strings.Split(ops[i], " ")
		op := _op[0]
		val, err := strconv.Atoi(_op[1])
		if err != nil {
			panic(err)
		}

		opsS[i] = op
		opsV[i] = val
	}

	for i := 0; i < len(ops); i += 1 {
		// fmt.Println("i", i)
		if opsS[i] == "nop" {
			opsS[i] = "jmp"
		} else if opsS[i] == "jmp" {
			opsS[i] = "nop"
		} else {
			continue
		}

		visited := make([]bool, len(ops), len(ops))
		idx := 0
		acc := 0
		for {
			if idx == len(ops) {
				break
			}
			if visited[idx] {
				break
			}
			visited[idx] = true

			switch opsS[idx] {
			case "nop":
				idx += 1
			case "acc":
				acc += opsV[idx]
				idx += 1
			case "jmp":
				idx += opsV[idx]
			}
		}

		// fmt.Println("idx", idx)
		if idx == len(ops) {
			fmt.Println(i, acc)
			break
		}

		if opsS[i] == "nop" {
			opsS[i] = "jmp"
		} else if opsS[i] == "jmp" {
			opsS[i] = "nop"
		}
	}
}
