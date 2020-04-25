package sample

type department struct {
	deptName string
	emp      employee
}

type departmentHandler interface {
	getDepartmentDetails() string
}

func (dept department) getDepartmentDetails() string {
	return dept.deptName + " :: " + dept.emp.getEmployeeName()
}
