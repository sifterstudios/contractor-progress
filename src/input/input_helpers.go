package input

import (
	"fmt"
)

func PromptInt(message string) (int, error) {
	var input int
	fmt.Print(message)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input. please try again.")
	}
	return input, err
}

func PromptFloat(message string) (float64, error) {
	var input float64
	fmt.Print(message)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input. please try again.")
	}
	return input, err
}
