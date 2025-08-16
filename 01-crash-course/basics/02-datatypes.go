package main
import "fmt"
/*
bool

string

int int8 int16 int32 int64
unit unit8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represent a Unicode point

float32 float64

complex64 complex12B
*/

func main() {
    var smsSendingLimit int
    var costPerSMS float64
    var hasPermission bool
    var username string

    fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
