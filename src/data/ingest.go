package data

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	util "github.com/sifterstudios/gontractor/src/util"
)

func ReadDataFromFile() (map[string]Week, Goal, error) {
	goal := Goal{}

	jsonFile, err := os.Open(util.DataFilePath)
	if err != nil {
		return nil, Goal{}, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, Goal{}, err
	}
	createBackup(jsonData)

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

func createBackup(data []byte) error {
	t := time.Now()
	timeString := t.Format("2006-01-02_15:04:05")
	fileName := "backup_" + timeString + ".json"

	os.WriteFile(fileName, data, 0644)
	return nil
}
