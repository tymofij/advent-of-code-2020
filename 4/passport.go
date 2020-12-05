package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// input.txt contains passport data, separated by empty line.
// count number of passports which have all the required fields
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`(\S{3}:\S+)`)
	// passportProperties := make([]string, 0, 8)
	scanner := bufio.NewScanner(file)

	validPassports := 0
	requiredFields := []string{
		"byr", // Birth Year
		"iyr", // Issue Year
		"eyr", // Expiration Year
		"hgt", // Height
		"hcl", // Hair Color
		"ecl", // Eye Color
		"pid", // Passport ID
		// "cid", // Country ID
	}
	passportSet := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// calculate the validity of current accumulated passport data
			isValid := true
			for _, field := range requiredFields {
				_, present := passportSet[field]
				if !present {
					isValid = false
					break
				}
			}
			if isValid {
				validPassports++
			}

			fmt.Println("")
			// reset accumulator and go on
			for key := range passportSet {
				delete(passportSet, key)
			}
			continue
		}

		for _, chunk := range re.FindAllString(line, -1) {
			splitted := strings.Split(chunk, ":")
			code, val := splitted[0], splitted[1]
			passportSet[code] = val
			fmt.Println(code, val)
		}

	}

	// Final line, calculate the validity of current accumulated passport data again
	isValid := true
	for _, field := range requiredFields {
		_, present := passportSet[field]
		if !present {
			isValid = false
			break
		}
	}
	if isValid {
		validPassports++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("\n===\n", validPassports)

}
