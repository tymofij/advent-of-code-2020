package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// find the two entries that sum to 2020 and then multiply those two numbers together
func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	nums := make([]int, 0, len(lines))
	for _, v := range lines {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	for _, v := range nums {
		for _, w := range nums {
			if v+w == 2020 {
				fmt.Println(v, w, v*w)
				return
			}
		}
	}
}
