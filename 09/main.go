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
	input, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	_numbers := strings.Split(string(input), "\n")
	numbers := make([]int, len(_numbers), len(_numbers))
	for i, n := range _numbers {
		_n, _ := strconv.Atoi(n)
		numbers[i] = _n
	}

	prevSize := 25
	invalidNumber := -1
	for i := prevSize; i < len(numbers); i += 1 {
		valid := false
		for j := i - 25; j < i; j += 1 {
			for k := j + 1; k < i; k += 1 {
				if numbers[j]+numbers[k] == numbers[i] {
					valid = true
					break
				}
			}

			if valid {
				break
			}
		}

		if !valid {
			fmt.Println(i, numbers[i])
			invalidNumber = numbers[i]
			break
		}
	}

	sums := make([]int, len(numbers)+1, len(numbers)+1)
	for i := 1; i < len(numbers)+1; i += 1 {
		sums[i] = sums[i-1] + numbers[i-1]
	}

	for i := 0; i < len(numbers); i += 1 {
		for j := i; j < len(numbers)+1; j += 1 {
			if sums[j]-sums[i] != invalidNumber {
				continue
			}

			min := 987654321
			max := 0
			for k := i + 1; k <= j; k += 1 {
				if min > numbers[k] {
					min = numbers[k]
				}
				if max < numbers[k] {
					max = numbers[k]
				}
			}

			fmt.Println(sums[i], sums[j], i, j, min, max, min+max)
		}
	}
}
