package input

import (
	"fmt"
)

func SetGoal(goal *float64, newGoal float64) {
	*goal = newGoal
	fmt.Println("Goal set to", newGoal)
}
