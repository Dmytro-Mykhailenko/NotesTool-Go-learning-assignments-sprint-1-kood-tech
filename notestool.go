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

	if len(args) != 1 {
		GetMessage(1)
		return
	} else {

		if args[0] == "help" {
			GetMessage(9)
			return
		}
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
	var noteN int

	in := bufio.NewReader(os.Stdin)

	strN, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return -1
	}

	strN = strings.TrimSpace(strN)
	noteN, err = strconv.Atoi(strN)

	if err != nil {
		fmt.Println(err)
		return -1
	}
	return noteN
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
	in := bufio.NewReader(os.Stdin)

	message, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	message = strings.TrimSpace(message)

	return message

}

func ReadFile(filename string) int {
	GetMessage(8)
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer file.Close()

	in := bufio.NewReader(file)
	x := 1
	for {
		note, err := in.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("%.3d - %s", x, note)
		x++
	}
	return x
}

func DeleteANote(filename string, lineNum int) {
}

func AddANote(s string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
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
	case 9:
		s =
			`
- Detailed Help:
-
- Use the following commands after starting the program with a file name:
-
-	1. Show notes: Display all current notes from the specified file.
-	2. Add a note: Prompt you to enter a note text, which is then added to the file.
-	3. Delete a note: Prompt you to select a note to delete based on its number.
-	4. Exit: Quit the program.
-
- Usage Examples:
-
- To manage notes in a file named 'mynotes.txt':
-
-	$ ./notestool mynotes
-	After starting, follow prompts to manage your notes.
-
- To get this help message:
-
-	$ ./notestool help
`
	}

	fmt.Println(s)
}
