package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func isPositiveDigit(input string) bool {
	digit, err := strconv.ParseInt(input, 10, 64)
	status := err == nil && !math.Signbit(float64(digit))
	if !status {
		fmt.Print("Enter 81 digit sudoku, denote empty spaces as 0 \n")
	}
	return status
}

func getInput() string {
	var input string
	for !isPositiveDigit(input) {
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
	}

	return input
}

func validateInput(input string) bool {
	if len(input) < 81 {
		fmt.Printf("Given input is too short being %d characters instead of 81 \n", len(input))
		return false
	} else if len(input) > 81 {
		fmt.Printf("Given input is too long being %d characters instead of 81 \n", len(input))
		return false
	} else if _, e := strconv.Atoi(input); e != nil {
		fmt.Printf("Given input should only contain digits \n %s", e)
		return false
	}
	return true
}
