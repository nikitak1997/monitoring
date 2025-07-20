package main

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Print(*x2, *x1)
}
func main() {
	x1, x2 := 2, 4
	test(&x1, &x2)
}

// dfgdfgfdgdfgfgdgdgdfg
// puk //
// 2 proverka//
