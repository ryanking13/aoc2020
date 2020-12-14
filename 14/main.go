package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func pow(n uint64, e int) uint64 {
	var m uint64 = 1
	for i := 0; i < e; i += 1 {
		m *= n
	}
	return m
}

func main() {
	r := bufio.NewReader(os.Stdin)
	inputs, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	cmds := strings.Split(string(inputs), "\n")

	mask := ""
	mem := make([]uint64, 65535)
	for _, cmd := range cmds {
		if cmd[0:4] == "mask" {
			mask = cmd[7:]
		} else {
			regex := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")
			matches := regex.FindStringSubmatch(cmd)
			offset, _ := strconv.Atoi(matches[1])
			val, _ := strconv.Atoi(matches[2])
			valBinary := strconv.FormatInt(int64(val), 2)

			for len(valBinary) < len(mask) {
				valBinary = "0" + valBinary
			}

			for i := range mask {
				if mask[i] == 'X' {
					continue
				}
				valBinary = valBinary[:i] + string(mask[i]) + valBinary[i+1:]
			}

			valInt, err := strconv.ParseUint(valBinary, 2, 64)
			if err != nil {
				panic(err)
			}
			mem[offset] = valInt
		}
	}

	var answer uint64
	for _, v := range mem {
		answer += v
	}
	fmt.Println(answer)

	// part 2

	mask = ""
	mem2 := map[uint64]uint64{}
	for _, cmd := range cmds {
		if cmd[0:4] == "mask" {
			mask = cmd[7:]
		} else {
			regex := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")
			matches := regex.FindStringSubmatch(cmd)
			offset, _ := strconv.Atoi(matches[1])
			val, _ := strconv.Atoi(matches[2])
			valInt := uint64(val)

			offsetBinary := strconv.FormatInt(int64(offset), 2)

			for len(offsetBinary) < len(mask) {
				offsetBinary = "0" + offsetBinary
			}

			floatings := make([]uint64, 0)
			for i := range mask {
				if mask[i] == 'X' {
					floatings = append(floatings, pow(2, 35-i))
					offsetBinary = offsetBinary[:i] + "0" + offsetBinary[i+1:]
				} else if mask[i] == '1' {
					offsetBinary = offsetBinary[:i] + "1" + offsetBinary[i+1:]
				}
			}
			// fmt.Println(floatings)

			offsetInt, err := strconv.ParseUint(offsetBinary, 2, 64)
			if err != nil {
				panic(err)
			}

			floatingsSize := int(pow(2, len(floatings)))
			for i := 0; i < floatingsSize; i += 1 {
				offsetInt_ := offsetInt
				for j := 0; j < len(floatings); j += 1 {
					floatingsMask := (i >> j) & 1
					if floatingsMask == 1 {
						offsetInt_ += floatings[j]
					}
				}
				mem2[offsetInt_] = valInt
			}
		}
	}

	var answer2 uint64
	for _, v := range mem2 {
		// fmt.Println(k, v)
		answer2 += v
	}
	fmt.Println(answer2)

}
