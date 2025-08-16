package main
import "fmt"

func main() {

    // NOTE: Printf() - returns the formatted string to std out

    // %v (default) can be used for any value
    fmt.Printf("Hi, I am %v.\n", "IQSF")

    // %s (Interpolate string)
    fmt.Printf("I am %s.\n", "JOBLESS")

    // %d (Interpolate integer in decimal form)
    fmt.Printf("I like %d.\n", 9)

    // %f (Interpolate a decimal)
    fmt.Printf("I like %f.\n", 69.69)

    // NOTE: Sprints() - return a formatted string 
    MESSAGE := fmt.Sprintf("I am %s dev and I like %d and also %f.", "MERN", 69, 69.69)
    fmt.Println(MESSAGE)
}
