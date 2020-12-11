package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	b := bufio.NewReader(os.Stdin)
	inputs, err := ioutil.ReadAll(b)

	if err != nil {
		panic(err)
	}

	seatsRaw := strings.Split(string(inputs), "\n")

	seats := make([][]int, len(seatsRaw))
	floor := 0
	empty := 1
	occupied := 2
	for i := range seats {
		seats[i] = make([]int, len(seatsRaw[i]))
		for j, s := range seatsRaw[i] {
			if s == '.' {
				seats[i][j] = floor
			} else if s == 'L' {
				seats[i][j] = empty
			} else if s == '#' {
				seats[i][j] = occupied
			}
		}
	}

	seatsOriginal := seats

	dx := []int{0, 0, -1, -1, -1, 1, 1, 1}
	dy := []int{-1, 1, 0, -1, 1, 0, -1, 1}
	for {
		changed := false
		newSeats := make([][]int, len(seats))
		for i := range seats {
			newSeats[i] = make([]int, len(seats[i]))
		}

		for x := range seats {
			for y := range seats[x] {
				numOccupied := 0
				for d := range dx {
					xx := x + dx[d]
					yy := y + dy[d]

					if xx < 0 || xx >= len(seats) || yy < 0 || yy >= len(seats[x]) {
						continue
					}

					if seats[xx][yy] == occupied {
						numOccupied += 1
					}
				}

				if seats[x][y] == empty && numOccupied == 0 {
					newSeats[x][y] = occupied
					changed = true
				} else if seats[x][y] == occupied && numOccupied >= 4 {
					newSeats[x][y] = empty
					changed = true
				} else {
					newSeats[x][y] = seats[x][y]
				}
			}
		}

		seats = newSeats

		if !changed {
			break
		}
	}

	cnt := 0
	for x := range seats {
		for y := range seats[x] {
			if seats[x][y] == occupied {
				cnt += 1
			}
		}
	}

	fmt.Println(cnt)

	// part 2
	seats = seatsOriginal
	for {
		changed := false
		newSeats := make([][]int, len(seats))
		for i := range seats {
			newSeats[i] = make([]int, len(seats[i]))
		}

		for x := range seats {
			for y := range seats[x] {
				numOccupied := 0
				for d := range dx {
					xx := x
					yy := y
					for {
						xx = xx + dx[d]
						yy = yy + dy[d]

						if xx < 0 || xx >= len(seats) || yy < 0 || yy >= len(seats[x]) {
							break
						}

						if seats[xx][yy] == occupied {
							numOccupied += 1
							break
						}

						if seats[xx][yy] == empty {
							break
						}
					}
				}

				if seats[x][y] == empty && numOccupied == 0 {
					newSeats[x][y] = occupied
					changed = true
				} else if seats[x][y] == occupied && numOccupied >= 5 {
					newSeats[x][y] = empty
					changed = true
				} else {
					newSeats[x][y] = seats[x][y]
				}
			}
		}

		seats = newSeats

		if !changed {
			break
		}
	}

	cnt2 := 0
	for x := range seats {
		for y := range seats[x] {
			if seats[x][y] == occupied {
				cnt2 += 1
			}
		}
	}

	fmt.Println(cnt2)

}
