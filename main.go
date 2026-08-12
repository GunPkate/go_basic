package main

import "fmt"

func main() {
	const name string = "GP"
	// name = "GunP"
	var age int = 28
	var salary float32 = 28000.00
	var aboveAvg bool = true
	fmt.Println("Hello World")
	fmt.Println(name, age)
	fmt.Println("salary ", salary)
	fmt.Println("above avg ", aboveAvg)
	fmt.Printf("value %v\n", salary)
	fmt.Printf("Data Type %T\n", salary)

	var num1, num2 int = 10, 2
	fmt.Printf("%v + %v = %v\n", num1, num2, num1+num2)
	fmt.Printf("%v - %v = %v\n", num1, num2, num1-num2)
	fmt.Printf("%v * %v = %v\n", num1, num2, num1*num2)
	fmt.Printf("%v / %v = %v\n", num1, num2, num1/num2)

	fmt.Printf("%v == %v = %v\n", num1, num2, num1 == num2)
	fmt.Printf("%v != %v = %v\n", num1, num2, num1 != num2)
	fmt.Printf("%v > %v = %v\n", num1, num2, num1 > num2)
	fmt.Printf("%v >= %v = %v\n", num1, num2, num1 >= num2)
	fmt.Printf("%v < %v = %v\n", num1, num2, num1 < num2)
	fmt.Printf("%v <= %v = %v\n", num1, num2, num1 <= num2)

}
