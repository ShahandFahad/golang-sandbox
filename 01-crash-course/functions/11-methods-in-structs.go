package main

import "fmt"

// rectangle struct
type rectt struct {
    width int
    height int
}

// rectangle struct method
func (r rectt) area() int {
    return r.width * r.height
}

// user auth info
type authenticationInfo struct {
    username string
    password string
}
// authenticationInfo struct method
func (authInfo authenticationInfo) getBasicAuth() string { 
    return fmt.Sprintf(
        "Authorization: Basic %s:%s",
        authInfo.username,
        authInfo.password,
    )
}


func main() {
    small_rectangle := rectt{ width: 10, height: 9 }
    fmt.Printf("Small Rectangle which has width %v, heigh %v, and area is %v\n", small_rectangle.width, small_rectangle.height, small_rectangle.area())

    test_user := authenticationInfo{ username: "Diddy", password: "ilikminr" }
    fmt.Println(test_user.getBasicAuth())
}
