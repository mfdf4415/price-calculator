package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFileName  string
	OutputFileName string
}

func (manager FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(manager.InputFileName)
	if err != nil {
		return nil, errors.New("file not open baby")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("scan not completed baby")
	}

	file.Close()
	return lines, nil

}

func (manager FileManager) WriteJSONToFile(data any) error {
	json, err := json.Marshal(data)

	if err != nil {
		return errors.New("cant marshal json baby")
	}

	return os.WriteFile(manager.OutputFileName, json, 0644)

}

func New(inputName, outputName string) FileManager {
	return FileManager{
		inputName,
		outputName,
	}
}
