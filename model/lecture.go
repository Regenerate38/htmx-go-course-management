package model

import "time"

type Lecture struct {
	LectureID int       `db:"lecture_id"`
	TeacherID int       `db:"teacher_id"`
	SectionID int       `db:"section_id"`
	ChapterID int       `db:"chapter_id"`
	TopicID   int       `db:"topic_id"`
	Objective string    `db:"objective"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
	Room      string    `db:"room"`
}
