package model

type Chapter struct {
	ChapterID     int     `db:"chapter_id"`
	ChapterNumber int     `db:"chapter_number"`
	ChapterName   string  `db:"chapter_name"`
	AllocatedTime float32 `db:"allocated_time"`
	CourseID      int     `db:"course_id"`
}
