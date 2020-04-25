package integration_tests

import (
	"github.com/priyankshah217/programming/sample"
	"testing"
)

func Test_department_integration_test(t *testing.T) {
	type fields struct {
		deptName string
		emp      sample.Employee
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Sample Test",
			fields{
				deptName: "Mechanical",
				emp: sample.Employee{
					FirstName: "Rahul",
					LastName:  "Dravid",
				},
			},
			"Mechanical :: Rahul Dravid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dept := sample.Department{
				DeptName: tt.fields.deptName,
				Emp:      tt.fields.emp,
			}
			if got := dept.GetDepartmentDetails(); got != tt.want {
				t.Errorf("department.GetDepartmentDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
