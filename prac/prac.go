package main

import (
	"fmt"
)

func first(x int) string {
	if x <= 100 {
		return "E"
	} else if x <= 200 {
		return "D"
	} else if x <= 300 {
		return "C"
	} else if x <= 400 {
		return "B"
	} else {
		return "A"
	}
}
func isEven(x int) bool {
	return x&1 == 0
}

func cal(num1, num2 float64, op string) float64 {
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return -1

	}
}
func sum(n int) int {
	return n * (n + 1) / 2
}
func swap(x, y *int) {
	*x, *y = *y, *x
}

type vehicle struct {
	model string
}
type bank struct {
	Name          string
	AccountNumber string
	IFSCcode      string
	PhoneNumber   string
}

func (b *bank) customPrint() {
	fmt.Println(b.Name, b.AccountNumber, b.IFSCcode, b.PhoneNumber)
}
func (g *vehicle) third() {
	switch g.model {
	case "car":

		fmt.Println("It's a car")
	case "bike":
		fmt.Println("It's a bike")
	case "truck":
		fmt.Println("It's a truck")
	default:
		fmt.Println("It's a vehicle")
	}
}

func second(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else if b > a {
		if b > c {
			return b
		} else {
			return c
		}
	} else {
		return c
	}
}

func main() {
	//fmt.Println("Grade:", first(450))
	//fmt.Println(second(2, 3, 2))
	//fmt.Println("isEven:", isEven(3))
	//fmt.Println(cal(2, 3, "/"))
	//fmt.Println(sum(10))
	//a := 10
	//b := 20
	//println("Before swapping:", a, b)
	//swap(&a, &b)
	//println("After swapping:", a, b)
	//account := bank{"Jeevanjot", "1234567890", "JAKA0MIRAAN", "7006999999"}
	//account.custom_print()
	//gadi := vehicle{"truck"}
	//gadi.third()
}
