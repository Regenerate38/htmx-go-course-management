package model

import "time"

type Quiz struct {
	QuizID    int
	ChapterID int
	StartTime time.Time
	EndTime   time.Time
}
