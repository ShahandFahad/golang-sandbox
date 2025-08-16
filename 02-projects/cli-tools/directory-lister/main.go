// Example: How to run?
//  1.  Running the go file directly
//      note: you can pass any path
//      a. go run main.go                       # no args
//      b. go run main.go ./                    # only pass path
//      c. go run main.go --show-hidden         # list hidden files
//      d. go run main.go ./ --show-hidden      # pass path and list hidden files also
//
//  2.  Build it and then run it
//      note: you can pass any path, same as above just using the .exe file after building
//      go build main.go                      # first build it
//      a. ./main                             # no args
//      b. ./main ./                          # only pass path
//      c. ./main --show-hidden               # list hidden files
//      d. ./main ./ --show-hidden            # pass path and list hidden files also

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// error handler
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Print file in Tree format
func tree(path string, spacing string, dirs int, files int, showHidden bool) (dirsCount int, filesCount int) {

	// read the current dir via path
	currentDir, err := os.ReadDir(path)
	handleError(err)

	// explore the directory
	for key := range currentDir {

		// get dir item: either a file or directory
		item := currentDir[key]

		itemName := item.Name()

		// if item name starts with "." and is a directory - then skip it
		if !showHidden && strings.Contains(itemName, ".") && item.IsDir() {
			continue
		}

		// check if last item or not
		isLastItem := key == len(currentDir)-1

		// print format
		if isLastItem {
			fmt.Printf("%v└── %v\n", spacing, item.Name())
		} else {
			fmt.Printf("%v├── %v\n", spacing, item.Name())
		}

		// if its a file
		if !item.IsDir() {
			files++
		}

		// if it's a directory
		if item.IsDir() {
			dirs++

			// get last elment of path
			pathLastElement := path[len(path)-1]
			newpath := ""

			// check if the last element is a path seperator then ok, else add path seperator at the end
			if os.IsPathSeparator(pathLastElement) {
				newpath = path + item.Name() + "/"
			} else {
				newpath = path + "/" + item.Name() + "/"
			}

			// spacing format
			newspacing := spacing + "│   "
			if isLastItem {
				newspacing = spacing + "   "
			}

			// RECURSICE CALL: Explore all nested directories
			dirs, files = tree(newpath, newspacing, dirs, files, showHidden)

			// reset spaing after Recursive function call returns
			newspacing = spacing
		}
	}

	// return count of dirs and files
	return dirs, files
}

/* MAIN */
func main() {

	// HOW TO RUN? Instruction are on top

	fmt.Println("+---------------------------------------+")
	fmt.Println("|   Coded By: FAHAD   |  GO FILE TREE   |")
	fmt.Println("+---------------------------------------+")
	fmt.Print("\n.\n")

	// default path is current folder
	path := "./"

	// default spacing
	spacing := ""

	// disable listing hidden files by default
	showHidden := false

	// directories and files count
	dirs := 1
	files := 0

	// coomand line arguments
	argLen := len(os.Args)

	// This is how the switch case below works:
	//  1.  No arguments - (index 0):
	//      Then it will explore the current directory and all nested directories
	//      It will ignore hidden directories (directories which name starts from '.')
	//
	//  2.  One argument - (index 1):
	//      Index 0 will be ignored
	//      Index will either contain 'PATH' or '--show-hidden' argument
	//      These cases are checked seperately:
	//          If it is 'PATH' then explore that
	//          If it is '--show-hidden' then explore the current directory and also list the hidden directories
	//
	//  3.  Two arguments - (index 1, and 2):
	//      Index 0 will be ignored
	//      Index 1 will contain 'PATH'
	//      Index 2 will contain '--show-hidden
	//      So, the 'PATH' will be explored and hidden directories will be listed as well.
	//
	//  4.  Invalid arguments:
	//      Any other args or incorrect path is passed
	//      Will result a default return with and invalid message print

	switch {
	case argLen == 1:
		dirs, files = tree(path, spacing, dirs, files, showHidden)
		break
	case argLen == 2 && os.Args[1] != "--show-hidden":
		path = os.Args[1]
		dirs, files = tree(path, spacing, dirs, files, showHidden)
		break
	case argLen == 2 && os.Args[1] == "--show-hidden":
		showHidden = true
		dirs, files = tree(path, spacing, dirs, files, showHidden)
		break
	case argLen == 3 && os.Args[2] == "--show-hidden":
		path = os.Args[1]
		showHidden = true
		dirs, files = tree(path, spacing, dirs, files, showHidden)
		break
	default:
		fmt.Println("Invalid Arguments")
		return
	}

	// Print dirs and files info
	fmt.Printf("\n%v directories, %v files\n", dirs, files)

}
