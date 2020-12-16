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

type class struct {
	name string
	s1   int
	e1   int
	s2   int
	e2   int
}

func isValid(classes *[]class, value int) bool {
	for _, c := range *classes {
		if (c.s1 <= value && value <= c.e1) || (c.s2 <= value && value <= c.e2) {
			return true
		}
	}
	return false
}

func isValidAll(classes *[]class, values []int) bool {
	for i := range *classes {
		c := (*classes)[i]
		v := values[i]
		if !((c.s1 <= v && v <= c.e1) || (c.s2 <= v && v <= c.e2)) {
			return false
		}
	}
	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	inputs, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	inputs_ := strings.Split(string(inputs), "\n\n")
	classes_ := strings.Split(inputs_[0], "\n")
	myTicket := strings.Split(strings.Split(inputs_[1], "\n")[1], ",")
	nearbyTickets := strings.Split(inputs_[2], "\n")[1:]

	classes := make([]class, 0)
	classRegex := regexp.MustCompile("(.+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
	for _, c_ := range classes_ {
		matches := classRegex.FindStringSubmatch(c_)
		// fmt.Println(matches)
		name := matches[1]
		s1, _ := strconv.Atoi(matches[2])
		e1, _ := strconv.Atoi(matches[3])
		s2, _ := strconv.Atoi(matches[4])
		e2, _ := strconv.Atoi(matches[5])

		c := class{
			name,
			s1,
			e1,
			s2,
			e2,
		}
		classes = append(classes, c)
	}

	errorRate := 0
	for _, ticket_ := range nearbyTickets {
		ticket := strings.Split(ticket_, ",")
		for _, val_ := range ticket {
			val, _ := strconv.Atoi(val_)
			if !isValid(&classes, val) {
				errorRate += val
			}
		}
	}
	fmt.Println(errorRate)

	// part2

	validTickets := make([][]int, 0)
	for _, ticket_ := range nearbyTickets {
		ticket := strings.Split(ticket_, ",")
		valid := true
		for _, val_ := range ticket {
			val, _ := strconv.Atoi(val_)
			if !isValid(&classes, val) {
				valid = false
				break
			}
		}

		if valid {
			validTickets = append(validTickets, make([]int, len(ticket)))
			l := len(validTickets)
			for i, val_ := range ticket {
				val, _ := strconv.Atoi(val_)
				validTickets[l-1][i] = val
			}
		}
	}

	validTickets = append(validTickets, make([]int, len(myTicket)))
	l := len(validTickets)
	for i, val_ := range myTicket {
		val, _ := strconv.Atoi(val_)
		validTickets[l-1][i] = val
	}

	// positions := permutations(classes)
	// for _, p := range positions {
	// 	fmt.Println(p)
	// 	valid := true
	// 	for _, ticket := range validTickets {
	// 		if !isValidAll(&p, ticket) {
	// 			valid = false
	// 			break
	// 		}
	// 	}
	// 	if valid {
	// 		fmt.Println(p)
	// 	}
	// }

	ticketAvailablePositions := make([][]int, 0)
	for i := 0; i < len(validTickets[0]); i += 1 {

		availablePositions := make([]int, len(classes))
		for j := 0; j < len(classes); j += 1 {
			availablePositions[j] = j
		}

		for j := 0; j < len(validTickets); j += 1 {
			val := validTickets[j][i]

			_availablePositions := make([]int, 0)
			for _, p := range availablePositions {
				c := classes[p]
				if (c.s1 <= val && val <= c.e1) || (c.s2 <= val && val <= c.e2) {
					_availablePositions = append(_availablePositions, p)
				}
			}
			availablePositions = _availablePositions
		}

		// fmt.Println(i, availablePositions)
		ticketAvailablePositions = append(ticketAvailablePositions, availablePositions)
	}

	answer := 1
	for l := 0; l < len(ticketAvailablePositions); l += 1 {
		for i := 0; i < len(ticketAvailablePositions); i += 1 {
			if len(ticketAvailablePositions[i]) == 1 {
				position := ticketAvailablePositions[i][0]
				// fmt.Println(i, position)
				if strings.HasPrefix(classes[position].name, "departure") {
					answer *= validTickets[len(validTickets)-1][i]
				}
				for ii := range ticketAvailablePositions {
					pj := -1
					for j, p := range ticketAvailablePositions[ii] {
						if p == position {
							pj = j
							break
						}
					}

					if pj == -1 {
						continue
					}

					ticketAvailablePositions[ii] = append(ticketAvailablePositions[ii][:pj], ticketAvailablePositions[ii][pj+1:]...)
				}
				break
			}
		}
	}
	fmt.Println(answer)
}
