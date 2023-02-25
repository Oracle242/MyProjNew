package main


import "fmt"


func main() {
    var number [5]int
    total := 1
    for i := 0; i < 5; i++ {
        fmt.Scanln(&number[i])
        total *= number[i]
    }
    fmt.Println(total)
}


