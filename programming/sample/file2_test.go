package sample

import "testing"

func Test_department_getDepartmentDetails(t *testing.T) {
	type fields struct {
		deptName string
		emp      Employee
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
				emp: Employee{
					FirstName: "Rahul",
					LastName:  "Dravid",
				},
			},
			"Mechanical :: Rahul Dravid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dept := Department{
				DeptName: tt.fields.deptName,
				Emp:      tt.fields.emp,
			}
			if got := dept.GetDepartmentDetails(); got != tt.want {
				t.Errorf("department.getDepartmentDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
