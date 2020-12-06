package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	passportsRaw := string(input)
	passports := strings.Split(passportsRaw, "\n\n")
	// fmt.Printf("%q", passports)

	r := regexp.MustCompile("(byr|iyr|eyr|hgt|hcl|ecl|pid|cid):(\\S+)\\s*")
	cnt := 0
	for _, passport := range passports {
		matches := r.FindAllStringSubmatch(passport, -1)
		fields := map[string]string{
			"byr": "",
			"iyr": "",
			"eyr": "",
			"hgt": "",
			"hcl": "",
			"ecl": "",
			"pid": "",
			"cid": "",
		}
		for _, match := range matches {
			// fmt.Printf("%q\n", match)
			field := match[1]
			value := match[2]
			fields[field] = value
		}

		valid := true
		for k, v := range fields {
			if v == "" && k != "cid" {
				valid = false
				break
			}
			if k == "byr" {
				v_, err := strconv.Atoi(v)
				if err != nil || v_ < 1920 || v_ > 2002 {
					valid = false
					break
				}
			} else if k == "iyr" {
				v_, err := strconv.Atoi(v)
				if err != nil || v_ < 2010 || v_ > 2020 {
					valid = false
					break
				}

			} else if k == "eyr" {
				v_, err := strconv.Atoi(v)
				if err != nil || v_ < 2020 || v_ > 2030 {
					valid = false
					break
				}
			} else if k == "hgt" {
				hgtRegex := regexp.MustCompile("(\\d+)(cm|in)")
				match := hgtRegex.FindStringSubmatch(v)
				if len(match) < 3 {
					valid = false
					break
				}
				hgt, err := strconv.Atoi(match[1])
				hgtType := match[2]
				if err != nil {
					valid = false
					break
				}
				if hgtType == "cm" && (hgt < 150 || hgt > 193) {
					valid = false
					break
				}
				if hgtType == "in" && (hgt < 59 || hgt > 76) {
					valid = false
					break
				}
			} else if k == "ecl" {
				switch v {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					break // do nothing
				default:
					valid = false
				}
			} else if k == "hcl" {

				if len(v) != 7 {
					valid = false
					break
				}

				if v[0] != '#' {
					valid = false
					break
				}
				for i := 1; i <= 6; i += 1 {
					vi := int(v[i])
					if !((vi >= int('0') && vi <= int('9')) || (vi >= int('a') && vi <= int('f'))) {
						valid = false
						break
					}
				}
			} else if k == "pid" {
				if len(v) != 9 {
					valid = false
					break
				}
				for _, v_ := range v {
					vi := int(v_)
					if !(vi >= int('0') && vi <= int('9')) {
						valid = false
						break
					}
				}
			}
		}
		if valid {
			cnt += 1
		}
	}

	fmt.Println(cnt)
}
