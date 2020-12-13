package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mulInv(a, b int64) int64 {
	b0 := b
	var x0, x1 int64 = 0, 1
	if b == 1 {
		return 1
	}

	for a > 1 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1-q*x0, x0
	}

	if x1 < 0 {
		x1 += b0
	}
	return x1
}

func crt(num []int64, mod []int64) int64 {
	var prod int64 = 1
	for _, n := range num {
		prod *= n
	}

	var sum int64 = 0
	for i := range num {
		p := prod / num[i]
		sum += mod[i] * mulInv(p, num[i]) * p
	}
	return sum % prod
}

// var one = big.NewInt(1)

// func crt(n, a []*big.Int) (*big.Int, error) {
// 	p := new(big.Int).Set(n[0])
// 	for _, n1 := range n[1:] {
// 		p.Mul(p, n1)
// 	}
// 	var x, q, s, z big.Int
// 	for i, n1 := range n {
// 		q.Div(p, n1)
// 		z.GCD(nil, &s, n1, &q)
// 		if z.Cmp(one) != 0 {
// 			return nil, fmt.Errorf("%d not coprime", n1)
// 		}
// 		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
// 	}
// 	return x.Mod(&x, p), nil
// }

func main() {
	var arriveTime int
	var busRaw string

	fmt.Scan(&arriveTime, &busRaw)

	bus := make([]int, 0)

	for _, b := range strings.Split(busRaw, ",") {
		if b == "x" {
			continue
		}

		t, _ := strconv.Atoi(b)
		bus = append(bus, t)
	}

	minTime := 1987654321
	minBus := -1
	for _, b := range bus {
		th := (arriveTime + b - 1) / b
		time := b * th
		if time >= arriveTime && time < minTime {
			minTime = time
			minBus = b
		}
	}

	waitTime := minTime - arriveTime

	fmt.Println(waitTime, minBus, waitTime*minBus)

	// part2
	mods := make([]int64, 0)
	times := make([]int64, 0)
	// mods := make([]*big.Int, 0)
	// times := make([]*big.Int, 0)
	for i, b := range strings.Split(busRaw, ",") {
		if b == "x" {
			continue
		}

		t, _ := strconv.Atoi(b)
		mods = append(mods, int64((t-i)%t))
		times = append(times, int64(t))
		// mods = append(mods, big.NewInt(int64(i)))
		// times = append(times, big.NewInt(int64(t)))
	}

	// fmt.Println(mods)
	// fmt.Println(times)

	fmt.Println(crt(times, mods))
}
