package handler

import (
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"
	"htmx-go-course-management/templates"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type DetailsData struct {
	Parent    string
	Value     string
	TableData []model.PassedData
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	contentType := r.URL.Query().Get("contentType")
	fmt.Println(contentType)
	fmt.Println("Reached DetailsHandler")

	courseID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	var passedData []model.PassedData
	var detailsData DetailsData

	switch contentType {
	case "Course":
		course, err := database.GetCourseByID(database.DB, id)
		if err != nil {
			http.Error(w, "Failed to fetch Course details", http.StatusInternalServerError)
			return
		}

		chapters, err := database.GetChapters(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Chapters", http.StatusInternalServerError)
			return
		}

		resources, err := database.GetResource(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Resources", http.StatusInternalServerError)
			return
		}

		var courseChapters []model.Chapter
		for _, chapter := range chapters {
			if chapter.CourseID == courseID {
				courseChapters = append(courseChapters, chapter)
			}
		}

		var courseResources []model.Resource
		for _, resource := range resources {
			if resource.CourseID == courseID {
				courseResources = append(courseResources, resource)
			}
		}

		passedData = append(passedData, model.PassedData{
			Name:   "Chapters",
			Header: extractHeaders(reflect.TypeOf(model.Chapter{})),
			Data:   extractRows(courseChapters),
		})

		passedData = append(passedData, model.PassedData{
			Name:   "Resources",
			Header: extractHeaders(reflect.TypeOf(model.Resource{})),
			Data:   extractRows(courseResources),
		})

		detailsData = DetailsData{
			Parent:    "Course",
			Value:     course.CourseName,
			TableData: passedData,
		}

	case "Chapter":
		chapter, err := database.GetChapterByID(database.DB, id)
		if err != nil {
			http.Error(w, "Failed to fetch Chapter details", http.StatusInternalServerError)
			return
		}

		topics, err := database.GetTopics(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Topics", http.StatusInternalServerError)
			return
		}

		lectures, err := database.GetLectures(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Lectures", http.StatusInternalServerError)
			return
		}

		quizzes, err := database.GetQuiz(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Quizzes", http.StatusInternalServerError)
			return
		}

		assignments, err := database.GetAssignment(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Assignments", http.StatusInternalServerError)
			return
		}

		var chapterTopics []model.Topic
		for _, topic := range topics {
			if topic.ChapterID == chapter.ChapterID {
				chapterTopics = append(chapterTopics, topic)
			}
		}

		var chapterLectures []model.Lecture
		for _, lecture := range lectures {
			if lecture.ChapterID == chapter.ChapterID {
				chapterLectures = append(chapterLectures, lecture)
			}
		}

		var chapterQuizzes []model.Quiz
		for _, quiz := range quizzes {
			if quiz.ChapterID == chapter.ChapterID {
				chapterQuizzes = append(chapterQuizzes, quiz)
			}
		}

		var chapterAssignments []model.Assignment
		for _, assignment := range assignments {
			if assignment.ChapterID == chapter.ChapterID {
				chapterAssignments = append(chapterAssignments, assignment)
			}
		}

		passedData = append(passedData, model.PassedData{
			Name:   "Topics",
			Header: extractHeaders(reflect.TypeOf(model.Topic{})),
			Data:   extractRows(chapterTopics),
		})

		passedData = append(passedData, model.PassedData{
			Name:   "Lectures",
			Header: extractHeaders(reflect.TypeOf(model.Lecture{})),
			Data:   extractRows(chapterLectures),
		})

		passedData = append(passedData, model.PassedData{
			Name:   "Quizzes",
			Header: extractHeaders(reflect.TypeOf(model.Quiz{})),
			Data:   extractRows(chapterQuizzes),
		})

		passedData = append(passedData, model.PassedData{
			Name:   "Assignments",
			Header: extractHeaders(reflect.TypeOf(model.Assignment{})),
			Data:   extractRows(chapterAssignments),
		})

		detailsData = DetailsData{
			Parent:    "Chapter",
			Value:     chapter.ChapterName,
			TableData: passedData,
		}

	case "Topic":
		topic, err := database.GetTopicByID(database.DB, id)
		if err != nil {
			http.Error(w, "Failed to fetch Topic details", http.StatusInternalServerError)
			return
		}

		lectures, err := database.GetLectures(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Lectures", http.StatusInternalServerError)
			return
		}

		var topicLectures []model.Lecture
		for _, lecture := range lectures {
			if lecture.TopicID == topic.TopicID {
				topicLectures = append(topicLectures, lecture)
			}
		}

		passedData = append(passedData, model.PassedData{
			Name:   "Lectures",
			Header: extractHeaders(reflect.TypeOf(model.Lecture{})),
			Data:   extractRows(topicLectures),
		})

		detailsData = DetailsData{
			Parent:    "Topic",
			Value:     topic.TopicName,
			TableData: passedData,
		}

	case "Teacher":
		teacher, err := database.GetTeacherByID(database.DB, id)
		if err != nil {
			http.Error(w, "Failed to fetch Teacher details", http.StatusInternalServerError)
			return
		}

		lectures, err := database.GetLectures(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Lectures", http.StatusInternalServerError)
			return
		}

		resources, err := database.GetResource(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Resources", http.StatusInternalServerError)
			return
		}

		var teacherLectures []model.Lecture
		for _, lecture := range lectures {
			if lecture.TeacherID == teacher.TeacherID {
				teacherLectures = append(teacherLectures, lecture)
			}
		}

		var teacherResources []model.Resource
		for _, resource := range resources {
			if resource.TeacherID == teacher.TeacherID {
				teacherResources = append(teacherResources, resource)
			}
		}

		passedData = append(passedData, model.PassedData{
			Name:   "Lectures",
			Header: extractHeaders(reflect.TypeOf(model.Lecture{})),
			Data:   extractRows(teacherLectures),
		})

		passedData = append(passedData, model.PassedData{
			Name:   "Resources",
			Header: extractHeaders(reflect.TypeOf(model.Resource{})),
			Data:   extractRows(teacherResources),
		})

		detailsData = DetailsData{
			Parent:    "Teacher",
			Value:     teacher.DepartmentName,
			TableData: passedData,
		}

	case "Section":
		section, err := database.GetSectionByID(database.DB, id)
		if err != nil {
			http.Error(w, "Failed to fetch Section details", http.StatusInternalServerError)
			return
		}

		lectures, err := database.GetLectures(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Lectures", http.StatusInternalServerError)
			return
		}

		var sectionLectures []model.Lecture
		for _, lecture := range lectures {
			if lecture.SectionID == section.SectionID {
				sectionLectures = append(sectionLectures, lecture)
			}
		}

		passedData = append(passedData, model.PassedData{
			Name:   "Lectures",
			Header: extractHeaders(reflect.TypeOf(model.Lecture{})),
			Data:   extractRows(sectionLectures),
		})

		detailsData = DetailsData{
			Parent:    "Section",
			Value:     section.SectionName,
			TableData: passedData,
		}

	case "Department":

		department, err := database.GetDepartmentByID(database.DB, id)
		if err != nil {
			http.Error(w, "Failed to fetch Department details", http.StatusInternalServerError)
			return
		}

		faculty, err := database.GetFaculty(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Faculty", http.StatusInternalServerError)
			return
		}

		var departmentFaculty []model.Faculty
		for _, f := range faculty {
			if f.DepartmentName == department.DepartmentName {
				departmentFaculty = append(departmentFaculty, f)
			}
		}

		passedData = append(passedData, model.PassedData{
			Name:   "Faculty",
			Header: extractHeaders(reflect.TypeOf(model.Faculty{})),
			Data:   extractRows(departmentFaculty),
		})

		detailsData = DetailsData{
			Parent:    "Department",
			Value:     department.DepartmentName,
			TableData: passedData,
		}

	case "Faculty":

		faculty, err := database.GetFacultyByID(database.DB, id)
		if err != nil {
			http.Error(w, "Failed to fetch Faculty details", http.StatusInternalServerError)
			return
		}

		courses, err := database.GetCourses(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Courses", http.StatusInternalServerError)
			return
		}

		sections, err := database.GetSection(database.DB)
		if err != nil {
			http.Error(w, "Failed to fetch Sections", http.StatusInternalServerError)
			return
		}

		var facultyCourses []model.Course
		for _, course := range courses {
			if course.FacultyName == faculty.FacultyName {
				facultyCourses = append(facultyCourses, course)
			}
		}

		var facultySections []model.Section
		for _, section := range sections {
			if section.FacultyName == faculty.FacultyName {
				facultySections = append(facultySections, section)
			}
		}

		passedData = append(passedData, model.PassedData{
			Name:   "Courses",
			Header: extractHeaders(reflect.TypeOf(model.Course{})),
			Data:   extractRows(facultyCourses),
		})

		passedData = append(passedData, model.PassedData{
			Name:   "Sections",
			Header: extractHeaders(reflect.TypeOf(model.Section{})),
			Data:   extractRows(facultySections),
		})

		detailsData = DetailsData{
			Parent:    "Faculty",
			Value:     faculty.FacultyName,
			TableData: passedData,
		}

	default:
		http.Error(w, "Content type not implemented for details", http.StatusNotImplemented)
	}

	err = templates.TMPL.ExecuteTemplate(w, "details.html", detailsData)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
