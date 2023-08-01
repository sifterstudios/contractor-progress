package data

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	util "github.com/sifterstudios/gontractor/src/util"
)

func ReadDataFromFile() (map[string]Week, Goal, error) {
	weeks := make(map[string]Week)
	goal := Goal{}

	jsonFile, err := os.Open(util.DataFilePath)
	if err != nil {
		return nil, Goal{}, err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, Goal{}, err
	}

	var fileData FileData
	err = json.Unmarshal(jsonData, &fileData)

	if err != nil {
		return nil, goal, errors.New("could not parse data file")
	}

	weeks = fileData.Weeks
	goal = fileData.Goal

	return weeks, goal, nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
