package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	arr = append(arr[n-1:], arr[:n-1]...)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

}
