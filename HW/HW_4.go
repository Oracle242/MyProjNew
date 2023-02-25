func main() {
    whataNumb := 0
    fmt.Println("Введи число для поиска:")
    fmt.Scanln(&whataNumb)


    var numb [5]int
    for i := 0; i < 5; i++ {
        fmt.Println("Введи число в массив")
        fmt.Scanln(&numb[i])
    }


    for i := 0; i < 5; i++ {
        if numb[i] == whataNumb {
            fmt.Println("Индекс числа", i)
        }
    }
}
