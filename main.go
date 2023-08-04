package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"

	crunching "github.com/sifterstudios/gontractor/src/crunching"
	data "github.com/sifterstudios/gontractor/src/data"
	input "github.com/sifterstudios/gontractor/src/input"
	util "github.com/sifterstudios/gontractor/src/util"
)

var (
	weeks map[string]data.Week
	goal  data.Goal
	stats data.Stats
)

func main() {
	green := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgRed, color.Bold)

	clearScreen()
	fmt.Println("Welcome to Gontractor!")
	fmt.Println("Checking for json file...")

	initalize(&weeks, &goal)
	stats = getStats(&weeks, &goal)
	showStats(stats, green, red)
	fmt.Println("Press enter to continue...")
	fmt.Scanln()

	for shouldContinue := true; shouldContinue; {
		printChoices()
		handleChoice(&shouldContinue)
		data.WriteDataToFile(weeks, goal)
	}
	fmt.Println("Welcome back!")
	data.WriteDataToFile(weeks, goal)
	fmt.Scanln()
}

func printChoices() {
	clearScreen()

	fmt.Println("What would you like to do?")
	fmt.Println("1. Add a week")
	fmt.Println("2. View a week")
	fmt.Println("3. View all weeks")
	fmt.Println("4. Change week")
	fmt.Println("5. Set goal")
	fmt.Println("6. Reset data")
	fmt.Println("7. Exit")
	fmt.Println("")
}

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func handleChoice(shouldContinue *bool) {
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		input.AddWeek(&weeks)
	case 2:
		input.ViewWeek(&weeks)
	case 3:
		input.ViewAllWeeks(&weeks)
	case 4:
		input.ChangeWeek(&weeks)
	case 5:
		// input.SetGoal(&goal.TotalContractHours)
	case 6:
		weeks = make(map[string]data.Week)
		goal = data.Goal{}
		data.WriteDataToFile(weeks, goal)
		main()
	case 7:
		*shouldContinue = false
	case 8:
	// TODO: Backup data
	default:
		fmt.Println("Invalid choice")
	}
}

func initalize(weeks *map[string]data.Week, goal *data.Goal) {
	fileExists := data.FileExists(util.DataFilePath)

	if !fileExists {
		fmt.Println("No json file found. Creating one...")
		data.WriteDataToFile(*weeks, *goal)
	} else {
		fmt.Println("Found json file.")
	}

	err := error(nil)
	*weeks, *goal, err = data.ReadDataFromFile()
	if err != nil {
		fmt.Println("Error reading data file.")
	}

	fmt.Printf("Found %d weeks of data\n", len(*weeks))

	if goal.TotalContractHours == 0 {
		setGoal(&*goal)
		data.WriteDataToFile(*weeks, *goal)
		fmt.Printf("Goal set at %p hours.", &goal.TotalContractHours)
	}
}

func getStats(weeks *map[string]data.Week, goal *data.Goal) data.Stats {
	stats := data.Stats{}

	crunching.GetTotalHours(&*weeks)
	stats.TotalWeeks = len(*weeks)
	stats.TotalHours = crunching.GetTotalHours(weeks)
	stats.TotalSickDays = crunching.GetTotalSickDays(weeks)
	stats.TotalVacationDays = crunching.GetTotalVacationDays(weeks)
	stats.TotalChildcareDays = crunching.GetTotalChildcareDays(weeks)
	stats.TotalNationalHolidays = crunching.GetTotalNationalHolidays(weeks)
	stats.PctCompleted = crunching.GetPercentComplete(&*weeks, &*goal) * 100
	stats.DaysLeft, stats.HoursLeft = crunching.GetTimeLeft(&*weeks, &*goal)
	fmt.Printf("Stats: %+v\n", stats)
	return stats
}

func showStats(stats data.Stats, green *color.Color, red *color.Color) {
	fmt.Printf("Total weeks: %d\n", stats.TotalWeeks)
	fmt.Printf("Total hours: %.1f\n", stats.TotalHours)
	fmt.Printf("Sickdays: %d\n", stats.TotalSickDays)
	fmt.Printf("Vacation-days: %d\n", stats.TotalVacationDays)
	fmt.Printf("Child care: %d\n\n", stats.TotalChildcareDays)
	fmt.Printf("National holidays: %d\n\n", stats.TotalNationalHolidays)

	green.Printf("Percent complete: %.2f\n", stats.PctCompleted)
	red.Printf("This means you have %d days and %.1f hours left to work\n", stats.DaysLeft, stats.HoursLeft)
}

func setGoal(goal *data.Goal) {
	if goal.TotalContractHours == 0 {
		fmt.Println("No goal has been set. How many hours are there in your contract?")
	} else {
		fmt.Printf("Your current goal is %p hours. What would you like to change it to?\n", &goal.TotalContractHours)
	}

	_, err := fmt.Scanln(&goal.TotalContractHours)
	if err != nil {
		fmt.Println("Error reading goal hours.")
	}
}
