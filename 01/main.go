package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	a := make([]int, 1000, 1000)

	r := bufio.NewReader(os.Stdin)
	var n int
	for n = 0; n < 1000; n = n + 1 {
		_, err := fmt.Fscan(r, &a[n])
		if err != nil {
			break
		}
	}

	for i := 0; i < n; i = i + 1 {
		for j := i + 1; j < n; j = j + 1 {
			if a[i]+a[j] == 2020 {
				fmt.Println(a[i] * a[j])
			}
		}
	}
}
