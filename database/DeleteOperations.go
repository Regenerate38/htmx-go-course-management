package database

import (
	"database/sql"
	"fmt"
)

func GenericDelete(dbPointer *sql.DB, tableName string, idColumn string, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableName, idColumn)
	stmt, err := dbPointer.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no record found in %s with id %s", tableName, id)
	}
	fmt.Printf("Deleted %d record(s) from %s\n", rowsAffected, tableName)
	return nil
}

func DeleteDept(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "department", "department_name", id)
}

func DeleteCourse(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "course", "course_id", id)
}

func DeleteChapter(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "chapter", "chapter_id", id)
}

func DeleteTopic(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "topic", "topic_id", id)
}

func DeleteLecture(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "lecture", "lecture_id", id)
}

func DeleteResource(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "resource", "resource_id", id)
}

func DeleteSection(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "section", "section_id", id)
}

func DeleteFaculty(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "faculty", "faculty_name", id)
}
func DeleteTeacher(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "teacher", "teacher_id", id)
}

func DeleteQuiz(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "quiz", "quiz_id", id)
}

func DeleteAssignment(dbPointer *sql.DB, id string) error {
	return GenericDelete(dbPointer, "assignment", "assignment_id", id)
}
