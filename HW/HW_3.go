package main


import "fmt"


func main() {
    total := 0
    numberPl := 0
    var number [5]int
    for i := 0; i < 5; i++ {
        fmt.Scanln(&number[i])
        if number[i] > 0 {
            total += number[i]
            numberPl++
        }
    }
    fmt.Println("Сумма положительных чисел", total)
    fmt.Println("Введено положительных чисел", numberPl)
}


