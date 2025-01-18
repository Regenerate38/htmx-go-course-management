package model

import "time"

type Lecture struct {
	LectureID int
	TeacherID int
	SectionID int
	ChapterID int
	TopicID   int
	Objective string
	StartTime time.Time
	EndTime   time.Time
	Room      string
}
