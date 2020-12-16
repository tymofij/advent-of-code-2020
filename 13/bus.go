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

// offsets calculated for Full data
// ----------------
// offset 1 period 1 kv 72 863
// offset 791 period 863 kv 0 41
// offset 30996 period 35383 kv 35 37
// offset 1021720 period 1309171 kv 41 431
// offset 354497890 period 564252701 kv 49 23

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

// Takes a lifetime on the full dataset :(
func main() {
	data := map[int64]int64{0: 17, 2: 13, 3: 19}
	// data := map[int64]int64{0: 41, 35: 37, 41: 431, 49: 23, 54: 13, 58: 17, 60: 19, 72: 863, 101: 29}
	buses := map[int64]int64{}
	var offset, period int64 = 0, 1

	for k, v := range data {
		buses[k] = v
		offset, period = getOffsetAndPeriod(buses, offset, period)
		fmt.Println("offset", offset, "period", period, "| kv", k, v)
	}
	fmt.Println("!!", offset, period)
}
