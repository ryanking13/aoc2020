package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	cnt := 0
	for i := 0; i < 999999; i = i + 1 {
		s, err := r.ReadString('\n')

		r := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]+): ([a-z]+)")
		matches := r.FindStringSubmatch(s)
		// fmt.Printf("%q\n", matches)

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		c := matches[3][0]
		str := matches[4]

		cCnt := 0
		for i := 0; i < len(str); i++ {
			if c == str[i] {
				cCnt = cCnt + 1
			}
		}

		if cCnt >= min && cCnt <= max {
			cnt = cnt + 1
		}

		if err != nil {
			// fmt.Println(s)
			break
		}
	}

	fmt.Println(cnt)

}
