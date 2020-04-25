package sample

type employee struct {
	firstName string
	lastName  string
}

type employeeHandler interface {
	getEmployeeName() string
}

func (emp employee) getEmployeeName() string {
	return emp.firstName + " " + emp.lastName
}
