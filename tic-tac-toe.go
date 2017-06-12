package main

import ("fmt")

func main () {
    fmt.Println("")
    fmt.Println("Tic Tac Toe (Winning Move):")
    fmt.Println("---------------------------")    
    tictactoe := [9] int {}
    for i := 0 ; i < len(tictactoe) ; i+=4 {
        tictactoe[i] = 1
    }
    for i := 0 ; i < len(tictactoe) ; i+=3 {
        fmt.Printf("%d %d %d \n", tictactoe[i], tictactoe[i+1], tictactoe[i+2])
    }

}