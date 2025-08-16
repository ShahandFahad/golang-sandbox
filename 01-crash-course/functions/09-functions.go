package main

import (
	"errors"
	"fmt"
)

// void function with params
func greet(name string) {
	fmt.Printf("Hello, %s. I can fix that.\n", name)
}

// returns string and accept string param
func message(description string) string {
    return "New Message: " + description
}

// concatenate two string
func concat(str1, str2 string) string {
    return str1 + str2
}

func doubleIt(num int) int {
    return num * 2
}

// return multiple values from function
func coordinates() (int, int, int) {
    return 2, 4, 6 
}


// naked returns from function and naming the return values also
func createSendToken() (token, expiry string) {
    // token and expiry are initilized with 0 values

    fmt.Printf("I talking from naked return function")

    return // automatically returns token and expiry
}

// return named values from function
func sendInfo() (name string, experience int){
    name = "John Doe"
    experience = 10

    return name, experience
}


// early return - guard claues
func divide(dividend, divisor int) (int, error){

    // guard clause
    if divisor == 0 {
        return 0, errors.New("Can't divide by zero")
    }

    return dividend / divisor, nil
}

func sayHello() (message string) {
    return "Hello, kitty!"
}

func main() {
	greet("Joe")

    msg := message("You look lonely!")
    fmt.Println(msg)

    fmt.Println(concat("You look lonely! ", "I can fix that!"))

    // re-assigning values
    num := 1
    fmt.Println("Num is: ", num)
    num = doubleIt(num)
    fmt.Println("After doubling, Num is: ", num)

    // ignoring the 3rd coordinate, NOTE: if you declare z-cordinate and does not use go will throgh error, as it doesnot allow un-used vars
    x, y, _ := coordinates()
    fmt.Println("Co-ordinates are: x: ", x, ", y: ", y)

    // naked function returns
    fmt.Println(createSendToken())
    
    // named returns from fuction
    name, experience := sendInfo()
    fmt.Println("Name: ", name, ", Experience: ", experience)

    // divide test 1
    result, err := divide(2, 0)
    if result == 0 {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Result: ", result)
    }
    // divide test 2
    result2, err2 := divide(2, 2)
    if result2 == 0 {
        fmt.Println("Error: ", err2)
    } else {
        fmt.Println("Result: ", result2)
    }

    // say hello
    fmt.Println(sayHello())

}
