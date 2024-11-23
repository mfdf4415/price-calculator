package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
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

func WriteJSONToFile(fileName string, data any) error {
	json, err := json.Marshal(data)

	if err != nil {
		return errors.New("cant marshal json baby")
	}
	
	return os.WriteFile(fileName, json, 0644)

}
