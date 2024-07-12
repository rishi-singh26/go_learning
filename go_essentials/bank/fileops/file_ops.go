package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	valueTxt := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueTxt), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	data, readErr := os.ReadFile(filename)
	if readErr != nil {
		return 0, errors.New("failed to find file")
	}
	txt := string(data)
	value, parseErr := strconv.ParseFloat(txt, 64)
	if parseErr != nil {
		return 0, errors.New("failed to parse file content")
	}
	return value, nil
}
