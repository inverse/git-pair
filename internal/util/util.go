package util

import (
	"errors"
	"log"
	"os"
)

func CheckFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func GetHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}

func UniqueStrings(input []string) []string {
	uniqueMap := make(map[string]struct{})

	for _, str := range input {
		uniqueMap[str] = struct{}{}
	}

	uniqueSlice := make([]string, 0, len(uniqueMap))
	for str := range uniqueMap {
		uniqueSlice = append(uniqueSlice, str)
	}

	return uniqueSlice
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
