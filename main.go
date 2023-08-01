package main

import (
	"fmt"

	data "github.com/sifterstudios/gontractor/src/data"
	util "github.com/sifterstudios/gontractor/src/util"
)

var (
	weeks map[string]data.Week
	goal  data.Goal
)

func main() {
	fmt.Println("Welcome to Gontractor!")
	fmt.Println("Checking for json file...")

	fileExists := data.FileExists(util.DataFilePath)

	if !fileExists {
		fmt.Println("No json file found. Creating one...")
		data.WriteDataToFile(weeks, goal)
	} else {
		fmt.Println("Found json file.")
	}
}
