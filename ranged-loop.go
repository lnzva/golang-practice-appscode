package main

import "fmt"

func main() {
	var x [5]int
	x[1] = 6
	var total = 0
	for _, val := range x {
		total += val
		fmt.Println(val)
	}
	fmt.Println(total)
}
