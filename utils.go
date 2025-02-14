package main

import (
	"errors"
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
