package database

import (
	"database/sql"
	"htmx-go-course-management/model"
)

func GetDept(dbPointer *sql.DB) ([]model.Department, error) {

	query := "SELECT * FROM `department` WHERE 1"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Department
	for rows.Next() {
		var dept model.Department
		rowErr := rows.Scan(&dept.DepartmentName)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}

func GetCourses(dbPointer *sql.DB) ([]model.Course, error) {

	query := "SELECT * FROM course"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Course
	for rows.Next() {
		var dept model.Course
		rowErr := rows.Scan(&dept.CourseId, &dept.CourseName, &dept.AllocatedTime, &dept.FacultyName)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
func GetTeachers(dbPointer *sql.DB) ([]model.Teacher, error) {

	query := "SELECT * FROM teacher"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []model.Teacher
	for rows.Next() {
		var teacher model.Teacher
		rowErr := rows.Scan(&teacher.TeacherID, &teacher.TeacherName, &teacher.DepartmentName)
		if rowErr != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teachers, nil
}

func GetChapters(dbPointer *sql.DB) ([]model.Chapter, error) {

	query := "SELECT * FROM chapter"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Chapter
	for rows.Next() {
		var dept model.Chapter
		rowErr := rows.Scan(&dept.ChapterID, &dept.ChapterNumber, &dept.ChapterName, &dept.AllocatedTime, &dept.CourseID)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}

func GetTopics(dbPointer *sql.DB) ([]model.Topic, error) {

	query := "SELECT * FROM topic"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Topic
	for rows.Next() {
		var dept model.Topic
		rowErr := rows.Scan(&dept.TopicID, &dept.ChapterID, &dept.TopicNumber, &dept.TopicName)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
func GetLectures(dbPointer *sql.DB) ([]model.Lecture, error) {

	query := "SELECT * FROM lecture"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Lecture
	for rows.Next() {
		var dept model.Lecture
		rowErr := rows.Scan(&dept.LectureID, &dept.TeacherID, &dept.SectionID, &dept.ChapterID, &dept.TopicID, &dept.Objective, &dept.StartTime, &dept.EndTime, &dept.Room)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
func GetResource(dbPointer *sql.DB) ([]model.Resource, error) {

	query := "SELECT * FROM resource"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Resource
	for rows.Next() {
		var dept model.Resource
		rowErr := rows.Scan(&dept.ResourceID, &dept.CourseID, &dept.TeacherID, &dept.ResourceName, &dept.ResourceType, &dept.Content)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
func GetSection(dbPointer *sql.DB) ([]model.Section, error) {

	query := "SELECT * FROM section"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Section
	for rows.Next() {
		var dept model.Section
		rowErr := rows.Scan(&dept.SectionID, &dept.SectionName, &dept.Semester, &dept.Year, &dept.FacultyName)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
func GetFaculty(dbPointer *sql.DB) ([]model.Faculty, error) {

	query := "SELECT * FROM faculty"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Faculty
	for rows.Next() {
		var dept model.Faculty
		rowErr := rows.Scan(&dept.FacultyName, &dept.DepartmentName)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
func GetQuiz(dbPointer *sql.DB) ([]model.Quiz, error) {

	query := "SELECT * FROM quiz"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Quiz
	for rows.Next() {
		var dept model.Quiz
		rowErr := rows.Scan(&dept.QuizID, &dept.ChapterID, &dept.StartTime, &dept.EndTime)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
func GetAssignment(dbPointer *sql.DB) ([]model.Assignment, error) {

	query := "SELECT * FROM assignment"
	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []model.Assignment
	for rows.Next() {
		var dept model.Assignment
		rowErr := rows.Scan(&dept.AssignmentID, &dept.ChapterID, &dept.AssignmentType, &dept.AssignedDate, &dept.Deadline)
		if rowErr != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return depts, nil
}
