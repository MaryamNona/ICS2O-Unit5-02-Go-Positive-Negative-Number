// Created by: Maryam Nona
// Createdon: May 2021
//
// This program does basic math

package main

import (
	"fmt"
)

func main () {
	var userGuess int

	// input
	fmt.Println("This program chooses a random number between 1-6 ")
	fmt.Println()
	fmt.Print("Enter number: ")
	fmt.Scanln(&userGuess)

	// output
	if userGuess < 0 {
		fmt.Println("You chose a negative number!")
	} else {
		fmt.Println("You chose a positive number!")
	}
}
