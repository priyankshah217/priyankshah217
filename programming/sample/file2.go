package sample

type Department struct {
	DeptName string
	Emp      Employee
}

func (dept Department) GetDepartmentDetails() string {
	return dept.DeptName + " :: " + dept.Emp.getEmployeeName()
}
