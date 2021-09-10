package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/xamenyap/the-go-programming-language-exercises/chapter2/2.2/lengthconv"
)

// Convert length from feet to meter
func main() {
	for _, arg := range os.Args[1:] {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal("cannot parse arg", f)
		}

		fmt.Printf("%.2f Feet = %.2f Meter\n", f, lengthconv.FeetToMeter(lengthconv.Feet(f)))
	}
}
