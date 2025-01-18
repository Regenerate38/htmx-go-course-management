package model

import "time"

type Assignment struct {
	AssignmentID   int
	ChapterID      int
	AssignmentType string
	AssignedDate   time.Time
	Deadline       time.Time
}
