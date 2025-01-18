package model

import "time"

type Assignment struct {
	AssignmentID   int       `db:"assignment_id"`
	ChapterID      int       `db:"chapter_id"`
	AssignmentType string    `db:"assignment_type"`
	AssignedDate   time.Time `db:"assigned_date"`
	Deadline       time.Time `db:"deadline"`
}
