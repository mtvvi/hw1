package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dict := make(map[string]int)
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ints := strings.Fields(text)

	for i := 0; i < len(ints); i++ {
		_, k := dict[ints[i]]
		if k {
			dict[ints[i]] += 1
		} else {
			dict[ints[i]] = 1
		}
	}
	fmt.Print(len(dict))

}
