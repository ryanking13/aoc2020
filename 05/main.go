package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	seats := strings.Split(string(b), "\n")
	if err != nil {
		panic(err)
	}

	maxSeat := -1
	seatsFilled := make([]bool, 127*8+8, 127*8+8)
	for _, seat := range seats {
		minR := 0
		maxR := 127
		minC := 0
		maxC := 7
		for r := 0; r < 7; r += 1 {
			row := seat[r]
			midR := (minR + maxR) / 2
			if row == 'F' {
				maxR = midR
			} else if row == 'B' {
				minR = midR + 1
			} else {
				panic(fmt.Errorf("Invalid char: %c", row))
			}
		}
		for c := 7; c < 10; c += 1 {
			col := seat[c]
			midC := (minC + maxC) / 2
			if col == 'L' {
				maxC = midC
			} else if col == 'R' {
				minC = midC + 1
			} else {
				panic(fmt.Errorf("Invalid char: %c", col))
			}
		}

		// fmt.Printf("%d %d / %d %d\n", minR, maxR, minC, maxC)
		seatID := minR*8 + minC
		if seatID > maxSeat {
			maxSeat = seatID
		}
		seatsFilled[seatID] = true
	}
	fmt.Println("max seat:", maxSeat)

	// for idx, s := range seatsFilled {
	// 	if !s {
	// 		fmt.Println(idx)
	// 	}
	// }
	for idx := 1; idx < len(seatsFilled)-1; idx += 1 {
		if !seatsFilled[idx] && seatsFilled[idx-1] && seatsFilled[idx+1] {
			fmt.Println(idx)
		}
	}
}
