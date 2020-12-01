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

// find the three entries that sum to 2020 and then multiply those three numbers together
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
			for _, y := range nums {
				if v+w+y == 2020 {
					fmt.Println(v, w, y, v*w*y)
					return
				}
			}
		}
	}
}
