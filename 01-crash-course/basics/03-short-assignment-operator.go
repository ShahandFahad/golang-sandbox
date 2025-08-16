package main
import "fmt"

func main(){
    // Dynamically assign type does not neet to declare type with each
    smsSendingLimit := 0 
    costPerSMS := 0.00
    hasPermission := false
    username := ""

    fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
