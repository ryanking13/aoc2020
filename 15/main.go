package main

import "fmt"

func main() {
	turns := map[int]int{
		12: 1,
		1:  2,
		16: 3,
		3:  4,
		11: 5,
	}

	number := 0
	// for i := 6; i <= 2020; i += 1 {
	for i := 6; i <= 30000000; i += 1 {
		val, ok := turns[number]
		turns[number] = i
		if !ok {
			number = 0
		} else {
			number = i - val
		}
	}

	for k, v := range turns {
		// fmt.Println(k, v)
		if v == 2020 {
			fmt.Println("2020:", k)
		}
		if v == 30000000 {
			fmt.Println("30000000:", k)
		}
	}
}
