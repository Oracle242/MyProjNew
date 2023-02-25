package main


import "fmt"


func main() {
    var number [5]int
    for i := 0; i < 5; i++ {
        fmt.Scanln(&number[i])
    }


    max := number[0]
    imax := 0


    for i := 0; i < 5; i++ {
        if number[i] > max {
            max = number[i]
            imax = i
        }
    }
    fmt.Println(max)
    fmt.Println(imax)
}
