package sample

type Employee struct {
	firstName string
	lastName  string
}

type employeeHandler interface {
	getEmployeeName() string
}

func (emp Employee) getEmployeeName() string {
	return emp.firstName + " " + emp.lastName
}
