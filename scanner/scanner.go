package scanner

import (
	"bufio"
	"os"
	"strings"
	"todoConsoleApp/todo"
)

type Scanner struct {
	Events   []Event
	TodoList *todo.List
	bufio    *bufio.Scanner
}

func NewScanner(bufioScanner *bufio.Scanner, todoList *todo.List, events []Event) *Scanner {
	return &Scanner{
		bufio:    bufioScanner,
		TodoList: todoList,
		Events:   events,
	}
}

func Start() {
	bufioScanner := bufio.NewScanner(os.Stdin)
	todoList := todo.NewList()
	events := []Event{}

	scanner := NewScanner(bufioScanner, todoList, events)

	for {
		printPromt()
		ok := scanner.bufio.Scan()
		if !ok {
			break
		}

		userInput := scanner.bufio.Text()
		result := scanner.process(userInput)

		if result != "" {
			if result == exitCommand {
				break
			}
			scanner.Events = append(scanner.Events, *newEvent(result, userInput))

		} else {

			scanner.Events = append(scanner.Events, *newEvent("", userInput))
		}
	}
}

func (s *Scanner) process(userInput string) string {
	fields := strings.Fields(userInput)

	if len(fields) < 1 {
		return emptyField
	}

	cmd := fields[0]

	if cmd == "help" {
		return s.cmdHelp(fields)
	}

	if cmd == "add" {
		return s.cmdAdd(fields)
	}

	if cmd == "delete" {
		return s.cmdDelete(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}

	if cmd == "events" {
		return s.cmdEvents(fields)
	}

	if cmd == "done" {
		return s.cmdDone(fields)
	}

	if cmd == "exit" {
		return s.cmdExit(fields)
	}

	return unknownCommand
}

func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		printError()

		return wrongArgs
	}
	printHelp()

	return ""
}

func (s *Scanner) cmdExit(fields []string) string {
	if len(fields) != 1 {
		printError()

		return wrongArgs
	}
	printExit()

	return exitCommand
}

func (s *Scanner) cmdAdd(fields []string) string {
	name := fields[1]
	text := ""

	if len(fields) < 3 {
		printError()

		return wrongArgs
	}
	// 10 массив 2 3 4 5 6 7(8) 8(9) 9(10)
	for i := 2; i < len(fields); i++ {
		text += fields[i]
		if i > len(fields)-1 {
			break
		}
		text += " "
	}
	printAdd(name)

	return s.TodoList.AddTask(todo.NewTask(name, text))
}

func (s *Scanner) cmdDelete(fields []string) string {
	if len(fields) != 2 {
		printError()

		return wrongArgs
	}
	result := s.TodoList.DeleteTask(fields[1])

	if result == "" {
		printDelete(fields[1])
	}

	return result
}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		printError()

		return wrongArgs
	}

	tasks := s.TodoList.ListTasks()
	printList(tasks)

	return ""
}

func (s *Scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		printError()

		return wrongArgs
	}

	result := s.TodoList.DoneTask(fields[1])

	if result == "" {
		printDone(fields[1])
	}

	return result
}

func (s *Scanner) cmdEvents(fields []string) string {
	if len(fields) != 1 {
		printError()

		return wrongArgs
	}

	printEvents(s.Events)

	return ""
}
