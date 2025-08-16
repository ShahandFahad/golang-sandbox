package main
import "fmt"

func main() {
    isLoggedIn := false

    if isLoggedIn {
        fmt.Println("You're logged in!")
    } else {
        fmt.Println("You're not logged in!")
    }


    // or declare a condition in there and check
    if max := 10; max <= 10 {
        fmt.Printf("Max is %d. Limit reached.\n", max)
    }

    if min := 0; min >= 0 {
        fmt.Printf("Min is %d. Limit reached.\n", min)
    }
}

