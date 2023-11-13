package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i <= (n+1)/2; i++ {
		for j := 0; j <= (n+1)/2; j++ {
			if matrix[i][j] != matrix[j][i] {
				fmt.Print("no")
				return
			}
		}
	}
	fmt.Print("yes")

}
