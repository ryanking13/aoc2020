package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	inputs, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	_numbers := strings.Split(string(inputs), "\n")
	numbers := make([]int, len(_numbers)+2, len(_numbers)+2)

	max := 0
	for i, n := range _numbers {
		number, _ := strconv.Atoi(n)
		numbers[i] = number

		if number > max {
			max = number
		}
	}

	numbers[len(_numbers)] = 0
	numbers[len(_numbers)+1] = max + 3

	sort.Ints(numbers)

	diffs := make([]int, 4, 4)
	for i := 0; i < len(numbers)-1; i += 1 {
		diffs[numbers[i+1]-numbers[i]] += 1
	}

	fmt.Println(diffs[1] * diffs[3])

	// part2

	numPaths := make([]int, len(numbers), len(numbers))
	numPaths[0] = 1
	for i := 0; i < len(numbers); i += 1 {
		for j := i - 3; j < i; j += 1 {
			if j < 0 {
				continue
			}
			if numbers[i]-numbers[j] <= 3 {
				numPaths[i] += numPaths[j]
			}
		}
	}

	// fmt.Println(numPaths)
	fmt.Println(numPaths[len(numbers)-1])
}
