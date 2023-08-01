package data

import (
	"encoding/json"
	"errors"
	"os"

	util "github.com/sifterstudios/gontractor/src/util"
)

func writeDataToFile(weeks map[string]Week, goal Goal) error {
	jsonData, err := json.Marshal(FileData{weeks, goal})
	if err != nil {
		return errors.New("could not parse data file")
	}

	err = os.WriteFile(util.DataFilePath, jsonData, 0644)
	if err != nil {
		return errors.New("could not write to data file")
	}

	return nil
}
