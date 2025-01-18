package model

import "time"

type Quiz struct {
	QuizID    int       `db:"quiz_id"`
	ChapterID int       `db:"chapter_id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}
