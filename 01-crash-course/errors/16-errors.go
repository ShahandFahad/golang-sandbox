package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor int) (int, error) {

    if divisor == 0 {
        return 0, errors.New("Err: Infinity")
    }

    return dividend / divisor, nil
}

/*

// this is actual error interface and can be implemented in multiple ways
type error interface {
    Error() string
}

*/

func main() {
    result, err := divide(100, 10)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Result: ", result)

    result, err = divide(69, 0)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Result: ", result)

}
