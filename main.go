package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var tasks = make([]Task, 0)
var commands = []Command{
	{
		command: "add",
		example: "add \"Buy Groceries\"",
	},
	{
		command: "update {id} \"{task}\"",
		example: "update 1 \"Buy Groceries and cook dinner\"",
	},
	{
		command: "delete {id}",
		example: "delete 1",
	},
	{
		command: "mark-in-progress {id}",
		example: "mark-in-progress 1",
	},
	{
		command: "mark-in-progress {id}",
		example: "mark-in-progress 1",
	},
	{
		command: "mark-done {id}",
		example: "mark-done 1",
	},
	{
		command: "list",
	},
	{
		command: "list done",
	},
	{
		command: "list todo",
	},
	{
		command: "list in-progress",
	},
	{
		command: "exit",
	},
}

func main() {
	reader := bufio.NewReader(bufio.NewReader(os.Stdin))

	fmt.Println("type command for list command that you can perform")

	for {
		read, err := reader.ReadString('\n') // Reads until newline
		input := strings.TrimSpace(read)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if CheckCommandFormat(input) {
			handleCommand()
			continue
		}

		if CheckAddFormat(input) {
			handleAdd(input)
			continue
		}

		if CheckUpdateFormat(input) {
			handleUpdate(input)
			continue
		}

		if CheckDeleteFormat(input) {
			handleDelete(input)
			continue
		}

		if CheckMarkInProgressFormat(input) {
			handleMarkInProgress(input)
			continue
		}

		if CheckMarkDoneFormat(input) {
			handleMarkDone(input)
			continue
		}

		if CheckListFormat(input) {
			handleList()
			continue
		}

		if CheckListDoneFormat(input) {
			handleListDone()
			continue
		}

		if CheckListTodoFormat(input) {
			handleListTodo()
			continue
		}

		if CheckListInProgressFormat(input) {
			handleListInProgress()
			continue
		}

		// list done

		if input == "exit" {
			return
		} else if input == "" {
			continue
		} else {
			fmt.Println("command not found")
		}
	}
}

func handleUpdate(input string) {
	isContainsComma := strings.Contains(input, " ")

	if !isContainsComma {
		fmt.Println("Your update command is not in correct format!")
		return
	}

	isFullContent := strings.Split(input, " ")

	// indicates that the command is not in correct format
	if len(isFullContent) < 3 {
		fmt.Println("Your update command is not in correct format!")
		return
	}

	// check if id is valid
	id, errId := strconv.Atoi(isFullContent[1])

	if errId != nil {
		fmt.Println("Your id is not valid!")
		return
	}

	// check if task is valid
	task, errQuotes := ExtractDoubleQuotes(input)

	if errQuotes != nil {
		fmt.Println(errQuotes.Error())
		return
	}

	updatedIndex := slices.IndexFunc(tasks, func(t Task) bool {
		return t.Id == id
	})

	if updatedIndex == -1 {
		fmt.Println("Your id is not found!")
		return
	}

	UpdateStruct(&tasks[updatedIndex], Task{
		Description: task,
		UpdatedAt:   time.Now().Local().Format("2006-01-02 15:04:05"),
	})

	fmt.Println("Sucessfuly update task with id: ", id)
	tasks[updatedIndex].ToJson()
}

func handleAdd(input string) {
	if !strings.Contains(input, " ") {
		fmt.Println("You haven't specified the task!")
		return
	}

	input, err := ExtractDoubleQuotes(input)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	now := time.Now().Local().Format("2006-01-02 15:04:05")

	task := Task{
		Id:          len(tasks) + 1,
		Description: input,
		Status:      TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	fmt.Printf("You have been added:")
	task.ToJson()
}

func handleDelete(input string) {
	contentArr := strings.Split(input, " ")
	id, errId := strconv.Atoi(contentArr[1])

	if errId != nil {
		fmt.Println("Your id is not valid!")
		return
	}

	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.Id == id
	})

	if index == -1 {
		fmt.Println("Your id is not found!")
		return
	}

	latestTasks := slices.Delete(tasks, index, index+1)

	tasks = latestTasks

	fmt.Println("Sucessfuly delete task with id: ", id)

	handleList()
}

func handleMarkInProgress(input string) {
	contentArr := strings.Split(input, " ")
	id, errId := strconv.Atoi(contentArr[1])

	if errId != nil {
		fmt.Println("Your id is not valid!")
		return
	}

	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.Id == id
	})

	if index == -1 {
		fmt.Println("Your id is not found!")
		return
	}

	tasks[index].Status = IN_PROGRESS

	fmt.Println("Sucessfuly mark task with id: ", id, " as in progress")
}

func handleMarkDone(input string) {
	contentArr := strings.Split(input, " ")
	id, errId := strconv.Atoi(contentArr[1])

	if errId != nil {
		fmt.Println("Your id is not valid!")
		return
	}

	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.Id == id
	})

	if index == -1 {
		fmt.Println("Your id is not found!")
		return
	}

	tasks[index].Status = DONE

	fmt.Println("Sucessfuly mark task with id: ", id, " as done")
}

func handleList() {
	if len(tasks) == 0 {
		fmt.Println("You have no task!")
		return
	}

	for _, value := range tasks {
		value.ToJson()
	}
}

func handleListDone() {
	doneTasks := FilterSlice(tasks, func(t Task) bool {
		return t.Status == DONE
	})

	if len(doneTasks) == 0 {
		fmt.Println("You have no done task!")
		return
	}

	fmt.Println("Filtering done tasks:")
	for _, value := range doneTasks {
		value.ToJson()
	}
}

func handleListTodo() {
	todoTasks := FilterSlice(tasks, func(t Task) bool {
		return t.Status == TODO
	})

	if len(todoTasks) == 0 {
		fmt.Println("You have no todo task!")
		return
	}

	fmt.Println("Filtering todo tasks:")
	for _, value := range todoTasks {
		value.ToJson()
	}
}

func handleListInProgress() {
	inProgressTask := FilterSlice(tasks, func(t Task) bool {
		return t.Status == IN_PROGRESS
	})

	if len(inProgressTask) == 0 {
		fmt.Println("You have no in-progress task!")
		return
	}

	fmt.Println("Filtering in-progress tasks:")
	for _, value := range inProgressTask {
		value.ToJson()
	}
}

func handleCommand() {
	for index, command := range commands {
		if command.example == "" {
			fmt.Printf("%d. %s\n", index+1, command.command)
		} else {
			fmt.Printf("%d. %s => %s\n", index+1, command.command, command.example)
		}
	}
}
