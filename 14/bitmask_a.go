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

var andMask, orMask uint64
var re = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
var mem = map[int]uint64{}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == `mask` {
			mask := line[7:]
			andMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, `X`, `1`), 2, 64)
			orMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, `X`, `0`), 2, 64)
			andMask += 0xFFFFFFF0_00000000
		} else {
			// must be "mem[x] = y"
			parsed := re.FindStringSubmatch(line)[1:]
			key, e := strconv.Atoi(parsed[0])
			check(e)
			val, e := strconv.ParseUint(parsed[1], 10, 64)
			check(e)
			val |= orMask
			val &= andMask
			mem[key] = val
		}
	}
	var s uint64
	for _, v := range mem {
		s += v
	}
	fmt.Println(s)
}
