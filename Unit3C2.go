/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-05
 * @fileoverview This program displays the multiplication table of 9 from 0 to 12.
*/

package main

import "fmt"

func main() {
	// initialize number as a constant
	const NUMBER int = 9

	// INPUT - none

	// PROCESS & OUTPUT
	// display multiplication table
	fmt.Println("0 X", NUMBER, "=", 0*NUMBER)
	fmt.Println("1 X", NUMBER, "=", 1*NUMBER)
	fmt.Println("2 X", NUMBER, "=", 2*NUMBER)
	fmt.Println("3 X", NUMBER, "=", 3*NUMBER)
	fmt.Println("4 X", NUMBER, "=", 4*NUMBER)
	fmt.Println("5 X", NUMBER, "=", 5*NUMBER)
	fmt.Println("6 X", NUMBER, "=", 6*NUMBER)
	fmt.Println("7 X", NUMBER, "=", 7*NUMBER)
	fmt.Println("8 X", NUMBER, "=", 8*NUMBER)
	fmt.Println("9 X", NUMBER, "=", 9*NUMBER)
	fmt.Println("10 X", NUMBER, "=", 10*NUMBER)
	fmt.Println("11 X", NUMBER, "=", 11*NUMBER)
	fmt.Println("12 X", NUMBER, "=", 12*NUMBER)

	fmt.Println("\nDone.")
}
