package main


import "fmt"


func main() {
    var number [5]int
    for i := 0; i < 5; i++ {
        fmt.Scanln(&number[i])
    }


    min := number[0]
    imin := 0


    for i := 0; i < 5; i++ {
        if number[i] < min {
            min = number[i]
            imin = i
        }
    }
    fmt.Println(min)
    fmt.Println(imin)
}
