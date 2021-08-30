package animal

//Exported type struct
type Animal struct {
	Name string
}

//Exported type Dog embedding type Animal
type Dog struct {
	Animal
	NoOfLegs int
}
