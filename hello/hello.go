package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, Go! 안녕!!")
	fmt.Println("Happy", math.Pi, "day")
	fmt.Printf("Now  you have %g problems.\n", math.Nextafter(10.2, 30))
	fmt.Printf("Now  you have %.30f problems.", math.Nextafter(2, 1.0))
	fmt.Printf("Now  you have %.30f problems.", math.Nextafter(2, 2.0000000000000004))
}
