package main

import "fmt"

func main() {
	total := soma(2, 3)
	fmt.Println(total)
}

func soma(a int, b int) int {
	return a + b
}
