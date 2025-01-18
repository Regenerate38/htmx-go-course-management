package model

type Topic struct {
	TopicID     int    `db:"topic_id"`
	ChapterID   int    `db:"chapter_id"`
	TopicNumber int    `db:"topic_number"`
	TopicName   string `db:"topic_name"`
}
