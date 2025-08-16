package main

import "fmt"

type car struct {
    name    string
    model   int
}



// Anonymous nested structed
type Tour struct {
    name string
    duration int
    destination string

    Guide struct {
        name string
        role string
    }
}


// embedded struct: Color and vechicle
type Colorr struct {
    paint string
}

// embedded struct: Color and vechicle
type Vehicle struct {
    name string
    Colorr // embedded struct
}

// MAIN
func main() {

    // define bmw with struct
    bmw := car{ name : "BMW-i8", model: 2016}
    fmt.Printf("Car info\n\tName: %v, Model: %v\n", bmw.name, bmw.model)

    // honda is not defined
    honda := car{}
    if honda.name != "" && honda.model >= 0 {
        fmt.Printf("Car info\n\tName: %v, Model: %v\n", honda.name, honda.model)
    } else {
        fmt.Println("Car not defined")
    }

    // Anonymous struct - only exist once and bound to user
    user := struct { 
        name string
        role string
        active bool
    } {
        name: "John Doe",
        role: "User",
        active: true,
    }
    fmt.Println("Anonymous Struct: ", user)


    // Nested Structs & NOTE: Nested struct can NOT be used as composite literals & can not be accessed at top level
    tour_one := Tour{} 

    tour_one.name= "The adventure" 
    tour_one.duration= 7 
    tour_one.destination= "Swat" 
    tour_one.Guide.name= "Gul Khan"
    tour_one.Guide.role= "Lead Guide"

    fmt.Println(tour_one)

    // Embedded Struct & can be defined as composite literals and fields are accessed at top level unlike nested struct
    cart := Vehicle{ name: "Donkey Cart", Color: Colorr{ paint: "Red" } }
    fmt.Println("Embedded Struct: ", cart)
    fmt.Printf("Embedded Struct: NAME: %v, COLOR: %v\n", cart.name, cart.paint)

}
