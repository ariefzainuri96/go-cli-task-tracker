package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

func ExtractDoubleQuotes(input string) (string, error) {
	re := regexp.MustCompile(`"(.*?)"`)
	matches := re.FindStringSubmatch(input)

	if len(matches) > 1 {
		return matches[1], nil
	} else {
		return "", errors.New("you have specified more than 1 task")
	}
}

// UpdateStruct updates only non-zero fields in the new object
func UpdateStruct(existing interface{}, updates interface{}) {
	existingVal := reflect.ValueOf(existing).Elem()
	updatesVal := reflect.ValueOf(updates)

	for i := 0; i < updatesVal.NumField(); i++ {
		field := updatesVal.Field(i)

		// Skip zero values (default values)
		if !field.IsZero() {
			existingField := existingVal.Field(i)

			// Ensure the field can be set (ignores unexported fields)
			if existingField.CanSet() {
				existingField.Set(field)
			}
		}
	}
}

func FilterSlice[T any](items []T, filterFunc func(T) bool) []T {
	result := make([]T, 0, len(items)) // Preallocate memory

	for _, item := range items {
		if filterFunc(item) {
			result = append(result, item)
		}
	}
	return result
}

func LoadJsonData(filePath string, tasks *[]Task) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&tasks)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Successfully loaded previous saved tasks!")
}

func SaveToJson(tasks []Task) {
	// Convert struct slice to JSON
	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Save JSON to a file
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func CheckAddFormat(input string) bool {
	re := regexp.MustCompile(`^add\s+"[^"]+"$`)
	return re.MatchString(input)
}

func CheckUpdateFormat(input string) bool {
	re := regexp.MustCompile(`^update\s+\d+\s+"[^"]+"$`)
	return re.MatchString(input)
}

func CheckDeleteFormat(input string) bool {
	re := regexp.MustCompile(`^delete\s+\d$`)
	return re.MatchString(input)
}

func CheckMarkInProgressFormat(input string) bool {
	re := regexp.MustCompile(`^mark-in-progress\s+\d$`)
	return re.MatchString(input)
}

func CheckMarkDoneFormat(input string) bool {
	re := regexp.MustCompile(`^mark-done\s+\d$`)
	return re.MatchString(input)
}

func CheckListFormat(input string) bool {
	return input == "list"
}

func CheckListDoneFormat(input string) bool {
	return input == "list done"
}

func CheckListTodoFormat(input string) bool {
	return input == "list todo"
}

func CheckListInProgressFormat(input string) bool {
	return input == "list in-progress"
}

func CheckCommandFormat(input string) bool {
	return input == "command"
}
