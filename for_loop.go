package main

import "fmt"

func main() {

	// ===== for loop in form of while loop
	a := 1
	for a < 100 {
		a += a
		if a == 8 {
			continue // will skip 8
		}
		fmt.Println(a)
		if a == 32 {
			break // get out of iteration when condition met
		}
	}
	fmt.Println(a)
}
