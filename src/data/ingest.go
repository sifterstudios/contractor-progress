package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	util "github.com/sifterstudios/gontractor/src/util"
)

func readDataFromFile() (map[string]Week, Goal, error) {
	weeks := make(map[string]Week)
	goal := Goal()

	if !fileExists(util.DataFilePath) {
		return weeks, goal, errors.New("Data file not found!")
		// TODO: create file if it doesn't exist
	}

	jsonData, err := ioutil.ReadFile(util.DataFilePath)
	if err != nil {
		return nil, Goal{}, err
	}

	var fileData FileData
	err = json.Unmarshal(jsonData, &fileData)

	if err != nil {
		return nil, goal, errors.New("Could not parse data file!")
	}

	weeks = fileData.Weeks
	goal = fileData.Goal

	return weeks, goal, nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
