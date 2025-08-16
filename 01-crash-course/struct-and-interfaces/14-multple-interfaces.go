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
}

// implementing the interfaces
func (e email) cost() float64 {
    
    // subscribed cost
    if e.isSubscribed {
        return float64(len(e.body)) * float64(0.01)
    }

    // not subscribed
    return float64(len(e.body)) * 0.05
}

func (e email) print() {
    fmt.Printf("Message: %s\n", e.body)
}


/////////////////////////////////////////////////////////////////////////
// General - Naming interface example
type Copier interface {
    Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
/////////////////////////////////////////////////////////////////////

func main() {
    email_1 := email{ isSubscribed: false, body: "Sedhe sedhe bol, Degi k nahi" }
    email_1.print()
    fmt.Printf("Subscribed: %v, Cost: %v\n", email_1.isSubscribed, email_1.cost())


    email_2 := email{ isSubscribed: true, body: "Sedhe sedhe bol, Degi k nahi" }
    email_2.print()
    fmt.Printf("Subscribed: %v, Cost: %v\n", email_2.isSubscribed, email_2.cost())
}
