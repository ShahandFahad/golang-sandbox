package main

import (
	"fmt"
	"log"
	"os"
)

// read file in bytes
func readFiles(){
    // open  a file
    file, err := os.Open("./simple-server.go")
    if err != nil {
        log.Fatal(err)
    }

    // make 100 byte slice
    data := make([]byte, 100)
    count, err := file.Read(data) // read 100 bytes from file into slice
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("read %d bytes: %q\n", count, data[:count]) // print info
}

// FIX: print the host name in color ascii
// get host name
func greetHost(){
    host, err := os.Hostname()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Hello, ", host)
}

// read file to stdout
func readIt(filename string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    os.Stdout.Write(data)
}

// List files and dirs of current directory
func fileLister(dirPath string, isNested bool) {

    // read current dir
    files, err := os.ReadDir(dirPath) 
    

    if err != nil {
        log.Fatal(err)
    }


    dirCount :=     0
    fileCount :=    0
    // isNested := false

    // fmt.Println(".\n|")
    fmt.Println("")

   
    // iterate over directory
    for fileKey := range files {

        // file via index
        file := files[fileKey]

        // file name
        // fmt.Println(file.Name())

        // file count
        if !file.IsDir() {

            // FIX: This duplicate logic for nesting checking
            // fmt.Print("├── ")
            if isNested {

                // FIX: Remove this hard code spacing
                fmt.Printf("│   │   ├── %v\n", file.Name())

            } else {

                fmt.Printf("├── %v\n", file.Name())
            }
            fileCount++

        }

        // directory count
        if file.IsDir() {
	
           // fmt.Printf("├──/%v", file.Name())



           // FIX: This also
            if isNested {

                fmt.Printf("│   ├──/%v", file.Name())

            } else {
                fmt.Printf("├──/%v", file.Name())
                // fmt.Print("├──/")
                // fmt.Printf("\033[34m%v", file.Name())
                // fmt.Print("\033[97m")
            }

            if file.Name() == ".git" {
                fmt.Println("")
                continue
            }
            // cd into nested directory
            err := os.Chdir(file.Name())
            if err != nil {
                fmt.Println("Nested Dir Err: ", err)
                continue
            }

            // NOTE: document this step well, and pass space automatically from here
            isNested = true // inside nested
            fileLister("./", isNested) // explore nested dir
            os.Chdir("..") // step out of nested dir after function returns
            isNested = false // not inside nested

            fmt.Print("│   ")
            dirCount++
        }
    }

    // fmt.Printf("\n%v directories, %v files\n", dirCount, fileCount)
    // isNested = false
    // os.Chdir("../")
    return
}

// MAIN
func main() {

    greetHost()
   // readFiles()
   // pass file name to read
   // readIt("./16-errors.go")

   // NOTE: make the path pass from std::in
   // make it to start looking from $HOME and forward


   // NOTE: Take the dir from command line and also handle the empty routes etc
   fileLister("./", false)

}
