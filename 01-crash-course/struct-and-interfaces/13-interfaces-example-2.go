package main

import "fmt"

type employee interface {
    getName() string
    getSalary() int
}

type contractor struct {
    name            string
    hourlyPay       int
    hoursPerYear    int
}

// implements emplyee interface 
func (c contractor) getName() string {
    return c.name
}

func (c contractor) getSalary() int {
    return c.hourlyPay * c.hoursPerYear
}

func main() {
    c := contractor { name: "Sibidi", hourlyPay: 10, hoursPerYear: 1200 }
    fmt.Println("Contractor: ", c.getName())
    fmt.Println("Salary: ", c.getSalary())
}
