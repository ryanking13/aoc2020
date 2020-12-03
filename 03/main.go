package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	m := make([]string, 10000, 10000)

	h := 0
	for h = 0; ; h = h + 1 {
		_, err := fmt.Fscan(r, &m[h])
		if err != nil {
			break
		}
	}

	w := len(m[0])

	dxs := []int{1, 3, 5, 7, 1}
	dys := []int{1, 1, 1, 1, 2}
	cnts := make([]int, len(dxs), len(dxs))
	for d := 0; d < len(dxs); d += 1 {
		dx := dxs[d]
		dy := dys[d]
		x := 0
		y := 0
		cnt := 0
		for y < h {
			if m[y][x] == '#' {
				cnt += 1
			}
			x = (x + dx) % w
			y = y + dy
		}

		cnts[d] = cnt
	}

	cnt_mul := 1
	fmt.Println(cnts)
	for _, cnt := range cnts {
		cnt_mul *= cnt
	}
	fmt.Println(cnt_mul)
}
