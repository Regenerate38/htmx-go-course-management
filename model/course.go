package model

type Course struct {
	CourseId      int     `db:"course_id"`
	CourseName    string  `db:"course_name"`
	AllocatedTime float32 `db:"allocated_time"`
	FacultyName   string  `db:"faculty_name"`
}
