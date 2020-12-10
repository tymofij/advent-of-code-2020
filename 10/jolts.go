package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var numbers = []int{0}

var cache = map[int]int64{}

func countWays(nums []int) int64 {
	// going from right end of the array to the left
	if len(nums) <= 2 {
		return 1
	}
	l := len(nums) - 1
	last := nums[l]
	res, ok := cache[last]
	if ok {
		return res
	}
	if last-nums[l-1] <= 3 {
		res += countWays(nums[:l])
	}
	if len(nums) >= 3 && last-nums[l-2] <= 3 {
		res += countWays(nums[:l-1])
	}
	if len(nums) >= 4 && last-nums[l-3] <= 3 {
		res += countWays(nums[:l-2])
	}
	cache[last] = res
	return res
}

func countWaysNonRecursive(nums []int) int64 {
	// going from the left to the right
	var cache = make([]int64, len(nums))
	var res int64
	for i, last := range nums {
		res = 0
		if i == 0 {
			res = 1
		}
		if i >= 1 && last-nums[i-1] <= 3 {
			res += cache[i-1]
		}
		if i >= 2 && last-nums[i-2] <= 3 {
			res += cache[i-2]
		}
		if i >= 3 && last-nums[i-3] <= 3 {
			res += cache[i-3]
		}
		cache[i] = res
	}
	return res
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	for _, s := range strings.Split(text, "\n") {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	sort.Ints(numbers)
	// fmt.Println(countWays(numbers))
	fmt.Println(countWaysNonRecursive(numbers))
}
