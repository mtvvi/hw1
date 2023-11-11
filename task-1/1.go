package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Scanf("%g\n", &a)
	fmt.Scanf("%g\n", &b)
	ans := math.Pow(math.Pow(a, 2)+math.Pow(b, 2), 0.5)
	fmt.Println(ans)
}
