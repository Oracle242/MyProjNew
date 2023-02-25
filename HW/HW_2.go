package main


import "fmt"


func main() {
    var number [5]int
    for i := 0; i < 5; i++ {
        fmt.Scanln(&number[i])
        if number[i] > 0 {
            fmt.Println("Ваше положительое число", number[i])
            fmt.Println("Эллемент числа", i)
	     return
        }
    }
}
