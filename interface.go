package main

import ("fmt")

type Dog struct {}
type Cat struct {}
type Animal interface {
    Speak() string
}
func (this Dog) Speak() string {
    return "woof"
}

func (this *Cat) Speak() string {
    return "meow"
}

func (this *Cat) Listen() string {
    return "hmmmmmmmm"
}

func Type_Checker(this interface{}) string {
    var datatype string
    switch this.(type) {
        case int:
            datatype = "int"
        case string:
            datatype = "string"
        case float64:
            datatype = "float64"        
        default:
            datatype = "Unknown"
    }
    return datatype
}

func main () {
    fmt.Println("")
    fmt.Println("Interfaces:")
    fmt.Println("-----------")
    
    animals := [] Animal {
        &Dog{},
        &Cat{},
    }
    
    fmt.Println("Animals (Speak)")
    fmt.Println("---------------")    
    for _,v := range animals {
        fmt.Println(v.Speak())
    }
    
    fmt.Println("")
    fmt.Println("Cat (Speak & Listen)")
    fmt.Println("--------------------")    
    animalObj := animals[1]    
    MyCat := animalObj.(*Cat)
    fmt.Println(MyCat.Speak())
    fmt.Println(MyCat.Listen())
    
    
    fmt.Println("")
    fmt.Println("Interfaces (Type Assertions)):")
    fmt.Println("------------------------------")
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s)

    s, okay := i.(string)
    fmt.Println(s, okay)

    fl, okay := i.(float64)
    fmt.Println(okay)
    fmt.Println(fl)

    //onlyVal := i.(float64)
    //fmt.Println(onlyVal)

    var teststring interface {} = "test"
    teststringValue, teststringExist := teststring.(string)
    fmt.Println("Test String (Exist & Value) : ", teststringExist, teststringValue)
    
    var testnum interface {} = 1
    testnumValue, testnumExist := testnum.(int)
    fmt.Println("Test Num    (Exist & Value) : ", testnumExist, testnumValue)
    
    TestVal := true
    fmt.Println("")
    fmt.Println("TestVal   : ",TestVal)
    fmt.Println("Data Type : ",Type_Checker(TestVal))
}