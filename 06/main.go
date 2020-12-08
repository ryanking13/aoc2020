package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getMap() map[rune]int {
	return map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	s, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	groups := strings.Split(string(s), "\n\n")
	anySum := 0
	allSum := 0
	for _, group := range groups {
		persons := strings.Split(group, "\n")
		yesMap := getMap()
		for _, p := range persons {
			for _, yes := range p {
				yesMap[yes] += 1
			}
		}

		any := 0
		all := 0
		for _, v := range yesMap {
			if v > 0 {
				any += 1
			}
			if v == len(persons) {
				all += 1
			}
		}
		anySum += any
		allSum += all
	}
	fmt.Println(anySum)
	fmt.Println(allSum)
}
