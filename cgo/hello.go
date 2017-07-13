package main

/*
#include <stdio.h>
#include <stdlib.h>

char* greeting = "Hello, World ~10Pearls";

*/
import "C"
import ("fmt")


func main() {
	//stdio.Stdout.WriteString(stdio.Greeting + "\n")
	var Greeting = C.GoString(C.greeting)
	fmt.Println(Greeting)
	//C.printf(C.greeting)
}
