# Notes Tool

---
#### Team

1. Dmytro Mykhailenko
2. Bert Kallas
3. Natalja Kravtsenko

---

This tool allows users to manage short single-line notes. Using this application, users can create collections of notes, open and view them, add new notes, or remove existing notes.
- Display notes from the collection
- Add a note to the collection
- Remove a note from the collection
- Exit the program

---

## How to use

**To use the Notes Tool, follow these steps:**

1. Run the tool using command go run notestool.go

`$ go run notestool.go`

2. Upon running you will be greeted with a welcome message and asked to select desired operation by pressing 1, 2, 3 or 4, which is show notes, add a note, delete a note and exit.

```
Welcome to the notes tool!
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
```

3. After operation 1 is selected you are asked to choose a note
```
Notes:
001 - note one
002 - note two
```
4. After operation 2 is selected you are asked to add a note
```
Enter the note text:
note three

Notes:
001 - note one
002 - note two
003 - note three
```
5. After operation 3 is selected you are asked to delete a note
```
Enter the number of note to remove or 0 to cancel:
3

Notes:
001 - note one
002 - note two
```
6. After operation 4 is selected you are asked to exit
```
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
4
```
---

## Data storage
 
 - For each collection, a separate database is created as a plain text file with the same name as the collection.
 - Notes are stored in separate rows within the database file.
 - Notes persist between runs of the tool.


## Resources

 Go packages:
`"bufio"`
`"fmt"`
`"io"`
`"os"`
`"strconv"`
`"strings"`
[Repository](https://gitea.kood.tech/dmytromykhailenko/notes)

