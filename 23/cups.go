package main

import (
	"fmt"
	"strconv"
)

// func log(args ...interface{}) {
// 	if debug {
// 		fmt.Println(args...)
// 	}
// }

type node struct {
	val  int
	next *node
}

func makeCups(s string, size int) *node {
	var start, prev, cur *node
	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		cur = &node{v, nil}
		if start == nil {
			start = cur
		}
		if prev != nil {
			prev.next = cur
		}
		prev = cur
	}
	for i := 10; i <= size; i++ {
		cur = &node{i, nil}
		prev.next = cur
		prev = cur
	}
	cur.next = start
	return start
}

func popThree(start *node) *node {
	res := start.next
	start.next = start.next.next.next.next
	return res
}

func fmtCups(start *node) string {
	res := fmt.Sprintf("(%d) ", start.val)
	var cur *node = start
	for cur.next != start {
		cur = cur.next
		res += fmt.Sprintf("%d ", cur.val)
	}
	return res
}

// return val-1, if val-1 is not in three first elements of excluded list
// if val-1 < 0, returns nCups
func getDestVal(val int, excluded *node) int {
	n := val - 1
	if n < 1 {
		n = nCups
	}
	for n == excluded.val || n == excluded.next.val || n == excluded.next.next.val {
		n--
		if n < 1 {
			n = nCups
		}
	}
	return n
}

func getNode(val int, start *node) *node {
	res := start
	for res.val != val {
		res = res.next
	}
	return res
}

func insertThreeAfter(start, three *node) {
	afterStart := start.next
	three.next.next.next = afterStart
	start.next = three
}

func strWithoutOne(start *node) string {
	one := getNode(1, start)
	cur := one.next
	res := ""
	for cur != one {
		res += strconv.Itoa(cur.val)
		cur = cur.next
	}
	return res
}

const nCups = 9
const nMoves = 100

// const nCups = 1_000_000
// const nMoves = 10_000_000

var debug bool = false
var cupCache [nCups + 1]*node

func main() {
	// Demo data
	// cups := makeCups("389125467", nCups)

	cups := makeCups("916438275", nCups)

	for move := 1; move <= nMoves; move++ {
		// log(fmt.Sprintf("-- move %d --\ncups: ", move))
		// log(fmtCups(cups))
		three := popThree(cups)
		// log("pick up:", three.val, three.next.val, three.next.next.val)
		destVal := getDestVal(cups.val, three)
		// log("destination:", destVal)
		destNode := getNode(destVal, cups)
		insertThreeAfter(destNode, three)
		cups = cups.next
		// log()
		if move%1000 == 0 {
			fmt.Println(move)
		}
	}
	if nCups < 10 {
		fmt.Println("Cups:", strWithoutOne(cups))
	}
	one := getNode(1, cups)
	fmt.Println("Numbers after one:", one.next.val, one.next.next.val)
	fmt.Println("Multiple:", one.next.val*one.next.next.val)
}
