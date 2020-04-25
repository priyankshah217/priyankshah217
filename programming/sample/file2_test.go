package sample

import "testing"

func Test_department_getDepartmentDetails(t *testing.T) {
	type fields struct {
		deptName string
		emp      employee
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
				emp: employee{
					firstName: "Rahul",
					lastName:  "Dravid",
				},
			},
			"Mechanical :: Rahul Dravid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dept := department{
				deptName: tt.fields.deptName,
				emp:      tt.fields.emp,
			}
			if got := dept.getDepartmentDetails(); got != tt.want {
				t.Errorf("department.getDepartmentDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
