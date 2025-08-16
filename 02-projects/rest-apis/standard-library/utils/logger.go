package utils

import "fmt"

// CUSTOME LOGGER: Log incomming request to the console
func Logger(method, path string) {
    fmt.Printf("A %s request is made on %s\n", method, path)
}

