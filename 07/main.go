package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type bagCount struct {
	color string
	count int
}

var containsShiny = map[string]bool{}

func dfs(bags map[string][]bagCount, color string) bool {

	if exists, ok := containsShiny[color]; ok {
		return exists
	} else {
		containsShiny[color] = false
		for _, bag := range bags[color] {
			if bag.color == "shiny gold" {
				containsShiny[color] = true
				break
			} else {
				containsShiny[color] = containsShiny[color] || dfs(bags, bag.color)
			}
		}
		return containsShiny[color]
	}
}

func dfs2(bags map[string][]bagCount, color string) int {
	sum := 0
	for _, bag := range bags[color] {
		sum += bag.count * (dfs2(bags, bag.color) + 1)
	}
	return sum
}

func main() {
	r := bufio.NewReader(os.Stdin)
	input, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rules := strings.Split(string(input), "\n")
	bags := map[string][]bagCount{}
	for _, rule := range rules {
		words := strings.Split(rule, " ")
		color := words[0] + " " + words[1]
		// fmt.Println(color)

		childs := make([]bagCount, 0, 0)

		idx := 4
		for {
			if words[idx] == "no" {
				break
			}

			cnt, _ := strconv.Atoi(words[idx])
			childColor := words[idx+1] + " " + words[idx+2]
			childs = append(childs, bagCount{childColor, cnt})
			// if sentence ends
			if words[idx+3][len(words[idx+3])-1] == '.' {
				break
			}

			idx += 4
		}

		bags[color] = childs
	}

	answer := 0
	for color, _ := range bags {
		if dfs(bags, color) {
			answer += 1
		}
	}

	fmt.Println("# bags containing shiny gold:", answer)

	///////////////////////////
	answer2 := dfs2(bags, "shiny gold")
	fmt.Println("# bags in shiny gold:", answer2)
}
