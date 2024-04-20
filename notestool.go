package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	GetMessage(7)

	var action int

	args := os.Args[1:]
	// args := []string{"coding_ideas.txt"}

	if len(args) != 1 || args[0] == "help" {
		GetMessage(1)
		return
	}

	for {
		action = GetOperation()

		switch action {
		case 1:

		case 2:

		case 3:

		case 4:
			{
				GetMessage(2)
				return
			}
		default:
			GetMessage(3)
		}
	}
}

func GetInput() int {
	return 0
}

func GetOperation() int {
	in := bufio.NewReader(os.Stdin)

	GetMessage(0)
	operation, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return -1
	}
	operation = strings.TrimSpace(operation)

	nOperation, err := strconv.Atoi(operation)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return nOperation
}

func ReadMessage() string {
	return ""
}

func ReadFile(filename string) int {
	return 0
}

func DeleteANote(filename string, lineNum int) {
}

func AddANote(s string, filename string) {
	return
}

func GetMessage(n int) {

	var s string

	switch n {
	case 0:
		s = "\nSelect operation: \n1. Show notes.\n2. Add a note.\n3. Delete a note.\n4. Exit."
	case 1:
		s =
			`
=======================================
Please type database name as an argument
Usage: ./notestool [file name]
========================================

`
	case 2:
		s = "\nGoodbye!\n"
	case 3:
		s = "\n[You have entered unsuitable selection for operation. Please select again.]"
	case 4:
		s = "\n>Enter the note or 0 to cancel:"
	case 5:
		s = "[You have entered no message. Please try again.]"
	case 6:
		s = "\n>Enter the number of note to remove or 0 to cancel:"
	case 7:
		s = "\nWelcome to the notes tool!"
	case 8:
		s = "\nNotes:"
	}

	fmt.Println(s)
}