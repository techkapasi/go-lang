package main

import ("fmt")

func printSlice(slice[] int) {
    fmt.Printf("len=%d cap=%d array=%v \n", len(slice), cap(slice), slice)
}

type MyFloat float64
func (this MyFloat) Abs () float64 {
    if this < 0 {
        return float64(-this)
    }
    return float64(this)
}

func main () {
    fmt.Println("")
    fmt.Println("Extending/Binding Function to Type:")
    fmt.Println("-----------------------------------") 
    f := MyFloat(-6.0)
    fmt.Println("Original Number : ", f)
    fmt.Println("Absolute Number : ", f.Abs())    

    fmt.Println("")
    fmt.Println("Array & Slices:")
    fmt.Println("---------------")
    arr := [4] int {1,2,3,4}
    fmt.Println("Array ==> ", arr)
    slice := arr[:]
    fmt.Println("Slice ==> ", slice)    
    dynamicarr := []int{}    
    for i := 1 ; i <= 10 ; i++ {
        dynamicarr = append(dynamicarr, i)
        printSlice(dynamicarr)
    }

}