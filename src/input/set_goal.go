package input

import (
	"fmt"
)

func SetGoal(goal *float64) {
	newGoal := 0.0
	fmt.Printf("Current goal is %.0f. What would you like to change it to? \n", *goal)
	fmt.Scanln(&newGoal)
	*goal = newGoal
	fmt.Println("Goal set to", newGoal)
}
