package main
import (
"fmt"
"os"
"math"

)


func main() {
var a, b, res float64
var op string




fmt.Print("Введите арифметическую операцию (+, -, *, /, **, fact): ")
fmt.Scanln(&op)
fmt.Print("Введите первое число: ")
fmt.Scanln(&a)
fmt.Print("Введите второе число: ")
fmt.Scanln(&b)

switch op {
case "+":
res = a + b
case "-":
res = a - b
case "*":
res = a * b
case "/":
res = a / b
case "**":
res = math.Pow(a, b)
case "fact":
res = math.Fact()
	default:
	fmt.Println("Операция выбрана неверно")
	os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f", res)
	}
	