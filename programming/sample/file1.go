package sample

type Employee struct {
	FirstName string
	LastName  string
}

func (emp Employee) getEmployeeName() string {
	return emp.FirstName + " " + emp.LastName
}
