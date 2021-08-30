package animal

//Exported type struct
type animal struct {
	Name string
}

//Exported type Dog embedding type Animal
type Dog struct {
	animal
	NoOfLegs int
}
