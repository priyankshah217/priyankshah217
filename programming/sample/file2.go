package sample

type Department struct {
	deptName string
	emp      Employee
}

type departmentHandler interface {
	getDepartmentDetails() string
}

func (dept Department) getDepartmentDetails() string {
	return dept.deptName + " :: " + dept.emp.getEmployeeName()
}
