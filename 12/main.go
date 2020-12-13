package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Direction struct {
	x int
	y int
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	r := bufio.NewReader(os.Stdin)
	input, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	cmds := strings.Split(string(input), "\n")
	dirs := []Direction{
		Direction{1, 0},  // east
		Direction{0, -1}, // south
		Direction{-1, 0}, // west
		Direction{0, 1},  // north
	}

	x := 0
	y := 0
	dir := 0

	for _, _cmd := range cmds {
		cmd := _cmd[0]
		val, _ := strconv.Atoi(_cmd[1:])

		switch cmd {
		case 'N':
			y += val
		case 'S':
			y -= val
		case 'E':
			x += val
		case 'W':
			x -= val
		case 'L':
			val /= 90
			dir -= val
			if dir < 0 {
				dir += 4
			}
		case 'R':
			val /= 90
			dir += val
			dir %= 4
		case 'F':
			d := dirs[dir]
			x += d.x * val
			y += d.y * val
		}
	}

	fmt.Println(abs(x) + abs(y))

	// part2

	wx := 10
	wy := 1
	x = 0
	y = 0

	rotR := func(x int, y int) (int, int) {
		return y, -x
	}

	for _, _cmd := range cmds {
		cmd := _cmd[0]
		val, _ := strconv.Atoi(_cmd[1:])

		switch cmd {
		case 'N':
			wy += val
		case 'S':
			wy -= val
		case 'E':
			wx += val
		case 'W':
			wx -= val
		case 'L':
			val /= 90
			for i := 0; i < val*3; i += 1 {
				wx, wy = rotR(wx, wy)
			}
		case 'R':
			val /= 90
			for i := 0; i < val; i += 1 {
				wx, wy = rotR(wx, wy)
			}
		case 'F':
			x += wx * val
			y += wy * val
		}
	}

	fmt.Println(abs(x) + abs(y))
}
