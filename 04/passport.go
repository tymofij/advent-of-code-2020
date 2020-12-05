package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type validator func(string) bool

var requiredFields = map[string]validator{
	"byr": func(val string) bool { // Birth Year
		year, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return 1920 <= year && year <= 2020
	},
	"iyr": func(val string) bool { // Issue Year
		year, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return 2010 <= year && year <= 2020
	},
	"eyr": func(val string) bool { // Expiration Year
		year, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return 2020 <= year && year <= 2030
	},
	"hgt": func(val string) bool { // Height
		units := val[len(val)-2:]
		n, err := strconv.Atoi(val[:len(val)-2])
		if err != nil {
			return false
		}
		switch units {
		case "cm":
			return 150 <= n && n <= 193
		case "in":
			return 59 <= n && n <= 76
		}
		return false
	},
	"hcl": func(val string) bool { // Hair Color
		match, _ := regexp.MatchString(`^#[\da-f]{6}$`, val)
		return match
	},
	"ecl": func(val string) bool { // Eye Color
		switch val {
		case
			"amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		}
		return false
	},
	"pid": func(val string) bool { // Passport ID
		match, _ := regexp.MatchString(`^\d{9}$`, val)
		return match
	},
	// "cid": isTrue // Country ID
}

func isValidPassport(data map[string]string) bool {
	for fieldName, validatorFunc := range requiredFields {
		val, ok := data[fieldName]
		if !ok {
			return false
		}
		if !validatorFunc(val) {
			return false
		}
	}
	return true
}

// input.txt contains passport data, separated by empty line.
// count number of passports which have all the required fields
// and those fields satisfy given conditions
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`(\S{3}:\S+)`)
	scanner := bufio.NewScanner(file)
	validPassports := 0

	passportData := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// calculate the validity of current accumulated passport data
			if isValidPassport(passportData) {
				validPassports++
			}
			// reset the passportData and go on
			for key := range passportData {
				delete(passportData, key)
			}
			continue
		}

		for _, chunk := range re.FindAllString(line, -1) {
			splitted := strings.Split(chunk, ":")
			code, val := splitted[0], splitted[1]
			passportData[code] = val
		}
	}
	// Final line, calculate the validity of current accumulated passport data again
	if isValidPassport(passportData) {
		validPassports++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(validPassports)
}
