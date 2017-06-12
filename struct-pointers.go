package main

import ("fmt")

func (C *cordinates) increment () {
    C.X++
    C.Y++
    C.Y++    
}

//func increment(C *cordinates) {
//    C.X++
//    C.Y++
//}

type cordinates struct {
    X int
    Y int
}

func main () {
    fmt.Println("")
    fmt.Println("Struct & Pointers:")
    fmt.Println("------------------")    
    //F := new(cordinates)
    //F := cordinates {}
    //F := cordinates {1,2}
    F := cordinates {X:1,Y:2}
    fmt.Println(F)
    F.increment()
    fmt.Println(F)    

}