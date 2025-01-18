package model

type Teacher struct {
	TeacherID      int    `db:"teacher_id"`
	TeacherName    string `db:"teacher_name"`
	DepartmentName string `db:"department_name"`
}
