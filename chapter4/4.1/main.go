package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	foo := sha256.Sum256([]byte("foo"))
	bar := sha256.Sum256([]byte("bar"))
	fmt.Println(countBitsDiff(foo, bar))
}

func countBitsDiff(x, y [32]byte) int {
	count := 0
	for i := range x {
		for j := 0; j < 7; j++ {
			xBit := (x[i] >> j) & 1
			yBit := (y[i] >> j) & 1

			if xBit != yBit {
				count++
			}
		}
	}

	return count
}
