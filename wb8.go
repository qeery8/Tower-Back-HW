package main

import "fmt"

func SetBit(n int64, i uint, value bool) int64 {
	if value {
		return n | (1 << i)
	} else {
		return n & ^(1 << i)
	}
}

func main() {
	var num int64 = 10
	var bitIndex uint = 1

	num = SetBit(num, bitIndex, true)
	fmt.Printf("After setting bit %d to 1: %064b\n", bitIndex, num)

	num = SetBit(num, bitIndex, false)
	fmt.Printf("After setting bit %d to 0: %064b\n", bitIndex, num)
}
