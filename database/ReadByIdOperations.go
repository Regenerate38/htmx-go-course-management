package database

import (
	"database/sql"
	"fmt"
	"htmx-go-course-management/model"
	"reflect"
)

func GenericGetByID(dbPointer *sql.DB, query string, id interface{}, dest interface{}) error {
	stmt, err := dbPointer.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("dest must be a non-nil pointer to a struct")
	}

	fieldPointers := make([]interface{}, val.Elem().NumField())
	for i := 0; i < val.Elem().NumField(); i++ {
		fieldPointers[i] = val.Elem().Field(i).Addr().Interface() // Get pointers to each field
	}

	if err := row.Scan(fieldPointers...); err != nil {
		return err
	}

	return nil
}

// GetDepartmentByID retrieves a department by its ID.
func GetDepartmentByID(dbPointer *sql.DB, id string) (model.Department, error) {
	var dept model.Department
	query := "SELECT * FROM `department` WHERE department_name = ?"
	if err := GenericGetByID(dbPointer, query, id, &dept); err != nil {
		return dept, err
	}
	return dept, nil
}

func GetCourseByID(dbPointer *sql.DB, id string) (model.Course, error) { // Change id type to int
	var course model.Course

	query := "SELECT * FROM course WHERE course_id = ?"
	fmt.Println(query, course, id)

	if err := GenericGetByID(dbPointer, query, id, &course); err != nil {

		return course, err
	}
	fmt.Println(query, course, id)

	return course, nil
}

// GetChapterByID retrieves a chapter by its ID.
func GetChapterByID(dbPointer *sql.DB, id string) (model.Chapter, error) {
	var chapter model.Chapter
	query := "SELECT * FROM chapter WHERE chapter_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &chapter); err != nil {
		return chapter, err
	}
	return chapter, nil
}
func GetTeacherByID(dbPointer *sql.DB, id string) (model.Teacher, error) {
	var teacher model.Teacher
	query := "SELECT * FROM teacher WHERE teacher_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &teacher); err != nil {
		return teacher, err
	}
	return teacher, nil
}

// GetTopicByID retrieves a topic by its ID.
func GetTopicByID(dbPointer *sql.DB, id string) (model.Topic, error) {
	var topic model.Topic
	query := "SELECT * FROM topic WHERE topic_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &topic); err != nil {
		return topic, err
	}
	return topic, nil
}

// GetLectureByID retrieves a lecture by its ID.
func GetLectureByID(dbPointer *sql.DB, id string) (model.Lecture, error) {
	var lecture model.Lecture
	query := "SELECT * FROM lecture WHERE lecture_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &lecture); err != nil {
		return lecture, err
	}
	return lecture, nil
}

// GetResourceByID retrieves a resource by its ID.
func GetResourceByID(dbPointer *sql.DB, id string) (model.Resource, error) {
	var resource model.Resource
	query := "SELECT * FROM resource WHERE resource_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}

// GetSectionByID retrieves a section by its ID.
func GetSectionByID(dbPointer *sql.DB, id string) (model.Section, error) {
	var section model.Section
	query := "SELECT * FROM section WHERE section_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &section); err != nil {
		return section, err
	}
	return section, nil
}

// GetFacultyByID retrieves a faculty member by their ID.
func GetFacultyByID(dbPointer *sql.DB, id string) (model.Faculty, error) {
	var faculty model.Faculty
	query := "SELECT * FROM faculty WHERE faculty_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &faculty); err != nil {
		return faculty, err
	}
	return faculty, nil
}

// GetQuizByID retrieves a quiz by its ID.
func GetQuizByID(dbPointer *sql.DB, id string) (model.Quiz, error) {
	var quiz model.Quiz
	query := "SELECT * FROM quiz WHERE quiz_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &quiz); err != nil {
		return quiz, err
	}
	return quiz, nil
}

// GetAssignmentByID retrieves an assignment by its ID.
func GetAssignmentByID(dbPointer *sql.DB, id string) (model.Assignment, error) {
	var assignment model.Assignment
	query := "SELECT * FROM assignment WHERE assignment_id = ?"
	if err := GenericGetByID(dbPointer, query, id, &assignment); err != nil {
		return assignment, err
	}
	return assignment, nil
}
