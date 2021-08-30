package main

import (
	"fmt"

	"github.com/sagar23sj/Go-OOP/Encapsulation/animal"
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
	// emp := employee.Employee{
	// 	Name:   "James",
	// 	Age:    33,
	// 	salary: 100000,
	// }

	//Creating an Object of type Employee from employee package.
	emp := employee.Employee{
		Name: "James",
		Age:  33,
	}

	//Setting unexported field salary using an exported method
	emp.SetSalary(1111111)
	fmt.Printf("Employee Details : %v", emp)

	//Creating an object of type Dog from animal package
	//to make below commented code work, change type animal to Animal
	// dog := animal.Dog{
	// 	Animal:   animal.Animal{Name: "Bruno"},
	// 	NoOfLegs: 4,
	// }

	dog := animal.Dog{
		NoOfLegs: 4,
	}
	dog.Name = "Maggie"

	fmt.Printf("\nDetails of Dog : %+v\n", dog)
}
