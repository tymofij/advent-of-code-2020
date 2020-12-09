package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const prefix = 25

var numbers = make([]int64, 0, 1024)

func isValid(nums []int64, n int64) bool {
	for _, a := range nums {
		for _, b := range nums {
			if a+b == n {
				return true
			}
		}
	}
	return false

}

func sum(nums []int64) int64 {
	var res int64
	for _, x := range nums {
		res += x
	}
	return res
}

func min(nums []int64) int64 {
	var res int64 = nums[0]
	for _, x := range nums {
		if x < res {
			res = x
		}
	}
	return res
}

func max(nums []int64) int64 {
	var res int64 = nums[0]
	for _, x := range nums {
		if x > res {
			res = x
		}
	}
	return res
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	for _, s := range strings.Split(text, "\n") {
		n, _ := strconv.ParseInt(s, 10, 64)
		numbers = append(numbers, n)
	}

	var x int64
	for i := prefix; i < len(numbers); i++ {
		if !isValid(numbers[i-prefix:i], numbers[i]) {
			x = numbers[i]
			break
		}
	}
	fmt.Println(x)

	for i := range numbers {
		for j := i + 2; j < len(numbers); j++ {
			if sum(numbers[i:j]) == x {
				r := numbers[i:j]
				fmt.Println(min(r) + max(r))
			}
		}
	}

}
