package model

type Section struct {
	SectionID   int    `db:"section_id"`
	SectionName string `db:"section_name"`
	Semester    string `db:"semester"`
	Year        string `db:"year"`
	FacultyName string `db:"faculty"`
}
