package model

type Resource struct {
	ResourceID   int    `db:"resource_id"`
	CourseID     int    `db:"course_id"`
	TeacherID    int    `db:"teacher_id"`
	ResourceName string `db:"resource_name"`
	ResourceType string `db:"resource_type"`
	Content      string `db:"content"`
}
