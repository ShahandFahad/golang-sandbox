package main

import "fmt"


// expense inteface
type expense interface {
    cost() float64
}

// printer interface
type printer interface {
    print()
}

// email struct
type email struct {
    isSubscribed bool
    body         string
    toAddress    string
}

// email struct
type sms struct {
    isSubscribed  bool
    body          string
    toPhoneNumber string
}

// implementing the interfaces
func (e email) cost() float64 {
    // subscribed cost
    if e.isSubscribed {
        return float64(len(e.body)) * 0.01
    }

    // not subscribed
    return float64(len(e.body)) * 0.05
}

func (e email) print() {
    fmt.Printf("Message: %s\n", e.body)
}



// implementing the expense interface for SMS struct
func (e sms) cost() float64 {
    // subscribed cost
    if e.isSubscribed {
        return float64(len(e.body)) * 0.01
    }

    // not subscribed
    return float64(len(e.body)) * 0.05
}

// type assertion here
func getExpenseReport(e expense) (string, float64) {
    em, ok := e.(email)
    if ok {
        return em.toAddress, em.cost()
    }

    s, ok := e.(sms)
    if ok {
        return s.toPhoneNumber, s.cost()
    }

    return "", 0.0
}

// using type switches for interface type checking
func getExpenseRepostViaSwitchCase(exp expense) (string, float64) {
    switch e := exp.(type) {
    case email:
        return e.toAddress, e.cost()
    case sms:
        return e.toPhoneNumber, e.cost()
    default:
        return "", 0.0
    }
}

func test1(e expense){
    address, cost := getExpenseReport(e)
    switch e.(type){
    case email:
        fmt.Printf("The email is gong to %v. Will cost %v\n", address, cost)
        fmt.Println("------------------------------------------------------")
    case sms:
        fmt.Printf("The sms is gong to %v. Will cost %v\n", address, cost)
        fmt.Println("------------------------------------------------------")
    default:
        fmt.Println("Invalid Expense")
        fmt.Println("------------------------------------------------------")
    }
}

// testing get expesnse via swithc
func test2(e expense){
    address, cost := getExpenseRepostViaSwitchCase(e)
    switch e.(type){
    case email:
        fmt.Printf("The email is gong to %v. Will cost %v\n", address, cost)
        fmt.Println("------------------------------------------------------")
    case sms:
        fmt.Printf("The sms is gong to %v. Will cost %v\n", address, cost)
        fmt.Println("------------------------------------------------------")
    default:
        fmt.Println("Invalid Expense")
        fmt.Println("------------------------------------------------------")
    }
}
// MAIN
func main() {
    mail := email{ isSubscribed: false, body: "Sedhe sedhe bol, Degi k nahi", toAddress: "test_1@email.com"}
    msg := sms{ isSubscribed: false, body: "Sedhe sedhe bol, Degi k nahi", toPhoneNumber: "12345"}

    test1(mail)
    test1(msg)


    test2(mail)
    test2(msg)


}

