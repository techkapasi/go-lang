package main

/*
#include <stdio.h>

char* greeting = "hello, world";

*/
import "C"
import ("fmt")

func main() {
	//stdio.Stdout.WriteString(stdio.Greeting + "\n")
	var Greeting = C.GoString(C.greeting)
	fmt.Println(Greeting)
	//C.printf(C.greeting)
}
