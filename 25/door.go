package main

import "fmt"

func transform(subject int, loops int) int {
	res := 1
	for i := 1; i <= loops; i++ {
		res *= subject
		res %= 20201227
	}
	return res
}

// loop 7177897  == pubkey 19072108
// loop 7779516  == pubkey 1965712
// encryption key == 16881444

func main() {
	fmt.Println(transform(1965712, 7177897))
	fmt.Println(transform(19072108, 7779516))
}
