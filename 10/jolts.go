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

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	for _, s := range strings.Split(text, "\n") {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	sort.Ints(numbers)
	fmt.Println(countWays(numbers))
}
