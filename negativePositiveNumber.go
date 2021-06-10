// Created by: Maryam Nona
// Createdon: May 2021
//
// This program does basic math

package main

import (
	"fmt"
)

func main () {
	var userInteger int

	// input
	fmt.Println("This programs shows if user chose negative or positive number.")
	fmt.Println()
	fmt.Println("Enter Number (-/+): ")
	fmt.Scanln(&userInteger)

	// output
	if userInteger < 0 {
		fmt.Println("You chose a negative number.")
  } else {
		fmt.Println("You chose a positive number.")
	}
}
