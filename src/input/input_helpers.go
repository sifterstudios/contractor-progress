package input

import (
	"fmt"
	"os"
)

func promptInt(message string) int {
	var input int
	fmt.Print(message)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input. please try again.")
		os.Exit(1)
	}
	return input
}
func promptFloat(message string) float64 {
	var input float64
	fmt.Print(message)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input. please try again.")
		os.Exit(1)
	}
	return input
}
