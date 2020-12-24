package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func permutations(s string) []string {
	prev := []string{s}
	res := []string{}
	mod := true
	for mod {
		res = []string{}
		for _, s := range prev {
			mod = false
			if !strings.Contains(s, "X") {
				res = append(res, s)
			} else {
				x := []byte(s)
				for i, c := range x {
					if c == 'X' {
						x[i] = '1'
						res = append(res, string(x))
						x[i] = '0'
						res = append(res, string(x))
						break
					}
				}
				mod = true
			}
		}
		prev = res
	}
	return res
}

func applyMask(val uint64, mask string) string {
	s := fmt.Sprintf("%b", val)
	res := []byte(mask)
	for i := 1; i <= len(s); i++ {
		// i is counted from the end of s
		offsetFromEndOfMask := len(mask) - i
		offsetFromEndOfNum := len(s) - i
		switch mask[offsetFromEndOfMask] {
		case '0':
			res[offsetFromEndOfMask] = s[offsetFromEndOfNum]
		case '1':
			res[offsetFromEndOfMask] = '1'
		}
	}
	return string(res)
}

var xMask string
var re = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
var mem = map[uint64]uint64{}

func main() {
	// fmt.Println(applyMask(0b101, "01X"))
	// fmt.Println(permutations("X0XX"))

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == `mask` {
			xMask = line[7:]
		} else {
			// must be "mem[x] = y"
			parsed := re.FindStringSubmatch(line)[1:]
			key, e := strconv.ParseUint(parsed[0], 10, 64)
			check(e)
			val, e := strconv.ParseUint(parsed[1], 10, 64)
			check(e)
			maskedKey := applyMask(key, xMask)
			// fmt.Println("m", maskedKey)
			for _, p := range permutations(maskedKey) {
				// fmt.Println(p)
				k, e := strconv.ParseUint(p, 2, 64)
				check(e)
				mem[k] = val
			}
		}
	}
	var s uint64
	for _, v := range mem {
		s += v
	}
	fmt.Println(s)
}
