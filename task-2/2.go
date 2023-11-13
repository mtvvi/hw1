package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%g", &a)
	fmt.Scanf("%g", &b)
	fmt.Scanf("%g", &c)
	maxx := math.Max(math.Max(a, b), c)

	if maxx < a+b+c-maxx {
		fmt.Println("YES")
		return
	fmt.Println("NO")
}
