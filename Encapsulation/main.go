package main

import (
	"fmt"

	"github.com/sagar23sj/Go-OOP/Encapsulation/counter"
)

func main() {

	//Creating variable of Exported type Count
	//Initializing the value of variable to 100
	count := counter.Count(100)

	fmt.Println("Value of the Counter : ", count)
}
