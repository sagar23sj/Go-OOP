package employee

//Exported struct type Employee with employee details
type Employee struct {
	Name   string
	Age    int
	salary float64
}

func (e *Employee) SetSalary(val float64) {
	e.salary = val
}
