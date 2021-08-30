package main

import (
	"fmt"

	"github.com/sagar23sj/Go-OOP/Encapsulation/counter"
	"github.com/sagar23sj/Go-OOP/Encapsulation/employee"
)

func main() {

	//Creating variable of Exported type Count
	//Initializing the value of variable to 100
	count := counter.Count(100)

	fmt.Println("Value of the Counter : ", count)

	//Creating Variable of Unexported type internalCounter
	//Initializing the value of variable to 10

	//uncommenting following lines will lead to error in program
	//internalCount := counter.internalCounter(10)
	//fmt.Println("Value of the Internal Count : ", internalCount)

	//Creating Variable of Unexported type internalCounter using
	//Exported NewInternalCounter function from package counter

	internalCount := counter.NewInternalCounter(10)
	fmt.Println("Value of the Internal Count : ", internalCount)

	//Creating an Object of type Employee from employee package.
	//uncommenting following code will result in error
	emp := employee.Employee{
		Name:   "James",
		Age:    33,
		salary: 100000,
	}

	fmt.Printf("\nEmployee Details : %v", emp)

}
