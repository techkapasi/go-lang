package main

import ("fmt")

func main () {
    fmt.Println("")
    fmt.Println("Panic & Defer:")
    fmt.Println("--------------")
    fmt.Println("testing before calling panic")


    defer func() {
        err := recover()
        fmt.Println("first defer  : ", err)

        defer func() {
            err := recover()
            fmt.Println("second defer : ", err)
            //goto AfterPanic  //this is label is not defined yet
        }()
        panic("panic2")
    }()
    //AfterPanic:
    panic("panic1")


    fmt.Println("testing after calling panic")
    //panic("panic1 called")
    //recover()

	goto AfterPanic
AfterPanic:
	var a int
	func () { a = 1     		; fmt.Println(a)}()
	defer   func () { a = a+1   ; fmt.Println(a)}()
	func () { a = 3     		; fmt.Println(a)}()
	defer   func () { a = 4     ; fmt.Println(a)}()
	func () { a = 5     		; fmt.Println(a)}()
	defer   func () { a = 6     ; fmt.Println(a)}()

}