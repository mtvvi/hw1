package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	for i := 0; i < len(n); i++ {
		if n[i] == '1' {
			fmt.Print("one")
		} else {
			fmt.Print(string(n[i]))
		}
	}

}
