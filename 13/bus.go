package main

import "fmt"

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func isValid(n int64, buses map[int64]int64) bool {
	var sinceDepart, toNext int64
	for k, v := range buses {
		sinceDepart = n % v
		toNext = 0
		if sinceDepart != 0 {
			toNext = v - sinceDepart
		}
		if toNext != k {
			return false
		}
	}
	return true
}

func getOffsetAndPeriod(data map[int64]int64, startOffset, startPeriod int64) (offset int64, period int64) {
	x := startOffset
	for period == 0 {
		if isValid(x, data) {
			if offset == 0 {
				offset = x
			} else {
				period = x - offset
			}
		}
		x += startPeriod
	}
	return offset, period
}

// func getGcdOffsetAndPeriod(prevOffset, prevPeriod, reqOffset, reqPeriod int64) (newOffset, newPeriod int64) {

// }

func main() {
	// data := map[int64]int64{0: 17, 2: 13, 3: 19}
	data := map[int64]int64{0: 41, 35: 37, 41: 431, 49: 23, 54: 13, 58: 17, 60: 19, 72: 863, 101: 29}
	buses := map[int64]int64{}
	var offset, period int64 = 1, 1
	for k, v := range data {
		fmt.Println("offset", offset, "period", period, "kv", k, v)
		buses[k] = v
		offset, period = getOffsetAndPeriod(buses, offset, period)
	}
	fmt.Println("!!", offset, period)
}
