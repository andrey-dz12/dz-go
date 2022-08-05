package main

import (
    "fmt"
    "strconv"
)

func main() {
    var a, b, s float64
    var result, input string
    var err error

    fmt.Println("Программа вычисления площади прямоугольника")

    fmt.Print("Введите длину:  ")
    fmt.Scanln(&input)
    a, err = strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Print("Введите ширину: ")
    fmt.Scanln(&input)
    b, err = strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println(err)
        return
    }

    s = a * b
    result = fmt.Sprintf("%.2f", s)
    fmt.Printf("Результат равен: %s\n", result)
}