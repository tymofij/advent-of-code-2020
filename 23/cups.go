package main

import (
	"fmt"
	"strconv"
)

var debug bool

func log(args ...interface{}) {
	if debug {
		fmt.Println(args...)
	}
}

func pop(i int, arr []int) (popped int, newArr []int) {
	if i == len(arr) {
		popped = arr[0]
		arr = arr[1:]
	}
	if i == len(arr)-1 {
		popped = arr[i]
		arr = arr[:i]
	}
	if i < len(arr)-1 {
		popped = arr[i]
		copy(arr[i:], arr[i+1:])
		arr = arr[:len(arr)-1]
	}
	return popped, arr
}

func fmtCaps(curCapPos int, arr []int) string {
	res := ""
	for i, cup := range arr {
		if i == curCapPos {
			res += fmt.Sprintf("(%d) ", cup)
		} else {
			res += fmt.Sprintf("%d ", cup)
		}
	}
	return res
}

func find(arr []int, x int) int {
	for i, elem := range arr {
		if x == elem {
			return i
		}
	}
	return -1
}

func insert(arr []int, index int, value int) []int {
	newArr := append(arr[:index+1], arr[index:]...)
	newArr[index] = value
	return newArr
}

func makeStr(arr []int) string {
	res := ""
	oneIdx := find(arr, 1)
	for i := oneIdx + 1; i < len(arr); i++ {
		res += strconv.Itoa(arr[i])
	}
	for i := 0; i < oneIdx; i++ {
		res += strconv.Itoa(arr[i])
	}
	return res
}

func getDest(val int, arr []int) (destPos int) {
	destVal := val - 1
	for find(arr, destVal) == -1 {
		destVal--
		if destVal < 0 {
			for _, elem := range arr {
				if destVal < elem {
					destVal = elem
				}
			}
		}
	}
	return find(arr, destVal)
}

func main() {
	// Demo data
	// cups := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

	cups := []int{9, 1, 6, 4, 3, 8, 2, 7, 5}
	var a, b, c int // pick ups
	curIdx := 0
	for move := 1; move <= 100; move++ {
		curVal := cups[curIdx]
		log(fmt.Sprintf("-- move %d --\ncups: ", move))
		log(fmtCaps(curIdx, cups))
		log()
		a, cups = pop(curIdx+1, cups)
		curIdx = find(cups, curVal)
		b, cups = pop(curIdx+1, cups)
		curIdx = find(cups, curVal)
		c, cups = pop(curIdx+1, cups)
		log("pick up:", a, b, c)
		destIdx := getDest(curVal, cups)
		log("destination:", cups[destIdx])
		cups = insert(cups, destIdx+1, a)
		cups = insert(cups, destIdx+2, b)
		cups = insert(cups, destIdx+3, c)
		curIdx = find(cups, curVal)
		curIdx++
		curIdx %= len(cups)
		log()
	}
	fmt.Println(makeStr(cups))
}
