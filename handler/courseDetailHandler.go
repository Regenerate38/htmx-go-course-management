package handler

import (
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"
	"htmx-go-course-management/templates"
	"net/http"
	"strconv"
)

func CourseDetailsHandler(w http.ResponseWriter, r *http.Request) {
	courseID := r.URL.Query().Get("courseID")
	if courseID == "" {
		http.Error(w, "Course ID is required", http.StatusBadRequest)
		return
	}

	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	resources, err := database.GetResourcesByCourseID(database.DB, courseIDInt)
	if err != nil {
		http.Error(w, "Failed to fetch resources", http.StatusInternalServerError)
		return
	}

	fmt.Println(resources)
	chapters, err := database.GetChaptersByCourseID(database.DB, courseIDInt)
	if err != nil {
		http.Error(w, "Failed to fetch chapters", http.StatusInternalServerError)
		return
	}

	var courseDetails struct {
		Resources []model.Resource
		Chapters  []struct {
			model.Chapter
			Quizzes     []model.Quiz
			Assignments []model.Assignment
			Topics      []struct {
				model.Topic
				Lectures []model.Lecture
			}
		}
	}

	courseDetails.Resources = resources
	for _, chapter := range chapters {
		var chapterDetails struct {
			model.Chapter
			Quizzes     []model.Quiz
			Assignments []model.Assignment
			Topics      []struct {
				model.Topic
				Lectures []model.Lecture
			}
		}

		chapterDetails.Chapter = chapter

		chapterDetails.Quizzes, err = database.GetQuizzesByChapterID(database.DB, chapter.ChapterID)
		if err != nil {
			http.Error(w, "Failed to fetch quizzes", http.StatusInternalServerError)
			return
		}

		chapterDetails.Assignments, err = database.GetAssignmentsByChapterID(database.DB, chapter.ChapterID)
		if err != nil {
			http.Error(w, "Failed to fetch assignments", http.StatusInternalServerError)
			return
		}

		topics, err := database.GetTopicsByChapterID(database.DB, chapter.ChapterID)
		if err != nil {
			http.Error(w, "Failed to fetch topics", http.StatusInternalServerError)
			return
		}

		for _, topic := range topics {
			var topicDetails struct {
				model.Topic
				Lectures []model.Lecture
			}

			topicDetails.Topic = topic

			topicDetails.Lectures, err = database.GetLecturesByTopicID(database.DB, topic.TopicID)
			if err != nil {
				http.Error(w, "Failed to fetch lectures", http.StatusInternalServerError)
				return
			}

			chapterDetails.Topics = append(chapterDetails.Topics, topicDetails)
		}

		courseDetails.Chapters = append(courseDetails.Chapters, chapterDetails)
	}

	err = templates.TMPL.ExecuteTemplate(w, "courseDetail.html", courseDetails)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
