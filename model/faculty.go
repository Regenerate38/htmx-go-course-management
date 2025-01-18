package model

type Faculty struct {
	FacultyName    string `db:"faculty_name"`
	DepartmentName string `db:"department_name"`
}
