package main

import (
	"fmt"
	"log"
	"os"
)


func handleError(err error)  {
    if err != nil {
        log.Fatal(err);
        return
    }
}

// Print file in Tree format
func tree(path string, isNested bool, nestedLevel int, spacing string) {
    // fmt.Println("Level: ", nestedLevel, " Path: ", path)                          

    currentDir, err := os.ReadDir(path)
    // if err != nil {
      //  log.Fatal(err);
       // return
    //}
    handleError(err)

    // dirCount := 0
    // fileCount := 0

    //fmt.Print("\n")
    for key := range currentDir {
        // get dir item: either a file or directory
        item := currentDir[key]

        // TODO: Use this
        isLastItem := key == len(currentDir) - 1 

        // fmt.Printf("%v└── %v\n", spacing, item.Name())
        // fmt.Printf("%v├── %v\n", spacing, item.Name())
        // fmt.Printf("%v└── %v\n", spacing, item.Name())
//        fmt.Printf("%v%v\n", spacing, item.Name())

            // if key == len(currentDir) - 1 {
            if isLastItem {
                fmt.Printf("%v└── %v\n", spacing, item.Name())
            } else {
                fmt.Printf("%v├── %v\n", spacing, item.Name())
            }
        // if its a file
        if !item.IsDir() {

            // fmt.Printf("%v└── %v\n", spacing, item.Name())
            // fmt.Println(item.Name())
            // fmt.Printf("%v├──%v\n", spacing, item.Name())
        }

        // if its a dir
        if item.IsDir() {
            // fmt.Printf("%v├── %v\n", spacing, item.Name())
            // fmt.Printf("/%v\n", item.Name())
            // fmt.Printf("%v├──%v\n", spacing, item.Name())

            isNested = true
            nestedLevel++

            // make a recursive call 
            newpath := path + item.Name() + "/"
            // spacing += "│   "
            newspacing := spacing + "│   "
            if isLastItem {
                newspacing = spacing + "   "
        }
            // path += item.Name() + "/"
            err := os.Chdir(newpath)
            //err := os.Chdir(path)
            handleError(err)
            tree("./", isNested, nestedLevel, newspacing)
            err = os.Chdir("..")
            isNested = false
            nestedLevel--
            handleError(err)
            newspacing = spacing // remove the addded spacing
        }

    }
    return
}


/* MAIN */
func main() {

    fmt.Println("+-------------------+")
    fmt.Println("|   GO FILE TREE    |")
    fmt.Println("+-------------------+")

    /* 
     * params: path, isNested, nestedLevel
     **/
    if os.Args[1] != "" {
        tree(os.Args[1], false, 1, "")
    } else {
        tree("./", false, 1, "")
    }
    fmt.Println("ARG 1: ", os.Args[0])
    fmt.Println("ARG 2: ", os.Args[1])
//    fmt.Println("ARG 3: ", os.Args[2])

}

