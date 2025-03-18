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
	query := "SELECT * FROM faculty WHERE faculty_name = ?"
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

func GetResourcesByCourseID(dbPointer *sql.DB, courseID int) ([]model.Resource, error) {
	query := "SELECT * FROM resource WHERE course_id = ?"
	rows, err := dbPointer.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []model.Resource
	for rows.Next() {
		var resource model.Resource
		rowErr := rows.Scan(&resource.ResourceID, &resource.CourseID, &resource.TeacherID, &resource.ResourceName, &resource.ResourceType, &resource.Content)
		if rowErr != nil {
			return nil, rowErr
		}
		resources = append(resources, resource)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return resources, nil
}

func GetChaptersByCourseID(dbPointer *sql.DB, courseID int) ([]model.Chapter, error) {
	query := "SELECT * FROM chapter WHERE course_id = ?"
	rows, err := dbPointer.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chapters []model.Chapter
	for rows.Next() {
		var chapter model.Chapter
		rowErr := rows.Scan(&chapter.ChapterID, &chapter.ChapterNumber, &chapter.ChapterName, &chapter.AllocatedTime, &chapter.CourseID)
		if rowErr != nil {
			return nil, rowErr
		}
		chapters = append(chapters, chapter)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return chapters, nil
}

func GetQuizzesByChapterID(dbPointer *sql.DB, chapterID int) ([]model.Quiz, error) {
	query := "SELECT * FROM quiz WHERE chapter_id = ?"
	rows, err := dbPointer.Query(query, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []model.Quiz
	for rows.Next() {
		var quiz model.Quiz
		rowErr := rows.Scan(&quiz.QuizID, &quiz.ChapterID, &quiz.StartTime, &quiz.EndTime)
		if rowErr != nil {
			return nil, rowErr
		}
		quizzes = append(quizzes, quiz)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return quizzes, nil
}

func GetAssignmentsByChapterID(dbPointer *sql.DB, chapterID int) ([]model.Assignment, error) {
	query := "SELECT * FROM assignment WHERE chapter_id = ?"
	rows, err := dbPointer.Query(query, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assignments []model.Assignment
	for rows.Next() {
		var assignment model.Assignment
		rowErr := rows.Scan(&assignment.AssignmentID, &assignment.ChapterID, &assignment.AssignmentType, &assignment.AssignedDate, &assignment.Deadline)
		if rowErr != nil {
			return nil, rowErr
		}
		assignments = append(assignments, assignment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return assignments, nil
}

func GetTopicsByChapterID(dbPointer *sql.DB, chapterID int) ([]model.Topic, error) {
	query := "SELECT * FROM topic WHERE chapter_id = ?"
	rows, err := dbPointer.Query(query, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topics []model.Topic
	for rows.Next() {
		var topic model.Topic
		rowErr := rows.Scan(&topic.TopicID, &topic.ChapterID, &topic.TopicNumber, &topic.TopicName)
		if rowErr != nil {
			return nil, rowErr
		}
		topics = append(topics, topic)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return topics, nil
}

func GetLecturesByTopicID(dbPointer *sql.DB, topicID int) ([]model.Lecture, error) {
	query := "SELECT * FROM lecture WHERE topic_id = ?"
	rows, err := dbPointer.Query(query, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lectures []model.Lecture
	for rows.Next() {
		var lecture model.Lecture
		rowErr := rows.Scan(&lecture.LectureID, &lecture.TeacherID, &lecture.SectionID, &lecture.ChapterID, &lecture.TopicID, &lecture.Objective, &lecture.StartTime, &lecture.EndTime, &lecture.Room)
		if rowErr != nil {
			return nil, rowErr
		}
		lectures = append(lectures, lecture)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return lectures, nil
}
