package main

import ("fmt";"time")


func sum(s []int, c chan int) {
	// implement a method that sums all elements of `s` slice and
	// sends the result to `c` channel
	time.Sleep(time.Second)
	c <- (s[0] + s[1] + s[2])
}

func main() {
	s := []int{1, 5, -2, 7, 2, 8, -9, 4, 0}
	c := make(chan int)
	totalsum  := 0
	
	for i := 0 ; i < len(s) ; i+=3 {
		var slice [] int = s[i:i+3]
		var chan_element int
        fmt.Print(slice, " ==> ")
        
		go sum(slice,c)
		chan_element = <- c
		fmt.Println(chan_element)

		totalsum += chan_element
	}
	fmt.Println("totalsum ==> ", totalsum)
	
	// sum `s` slice by taking three elements each time
	// print sum of each time and the total result
}