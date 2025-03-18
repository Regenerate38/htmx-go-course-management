package handler

import (
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"
	"net/http"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

func ExportCourseDetailsToPDF(response http.ResponseWriter, r *http.Request) {
	courseID := r.URL.Query().Get("courseID")
	if courseID == "" {
		http.Error(response, "Course ID is required", http.StatusBadRequest)
		return
	}

	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		http.Error(response, "Invalid course ID", http.StatusBadRequest)
		return
	}

	resources, err := database.GetResourcesByCourseID(database.DB, courseIDInt)
	if err != nil {
		http.Error(response, "Failed to fetch resources", http.StatusInternalServerError)
		return
	}

	chapters, err := database.GetChaptersByCourseID(database.DB, courseIDInt)
	if err != nil {
		http.Error(response, "Failed to fetch chapters", http.StatusInternalServerError)
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
			http.Error(response, "Failed to fetch quizzes", http.StatusInternalServerError)
			return
		}

		chapterDetails.Assignments, err = database.GetAssignmentsByChapterID(database.DB, chapter.ChapterID)
		if err != nil {
			http.Error(response, "Failed to fetch assignments", http.StatusInternalServerError)
			return
		}

		topics, err := database.GetTopicsByChapterID(database.DB, chapter.ChapterID)
		if err != nil {
			http.Error(response, "Failed to fetch topics", http.StatusInternalServerError)
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
				http.Error(response, "Failed to fetch lectures", http.StatusInternalServerError)
				return
			}

			chapterDetails.Topics = append(chapterDetails.Topics, topicDetails)
		}

		courseDetails.Chapters = append(courseDetails.Chapters, chapterDetails)
	}

	pdf := gofpdf.New("L", "mm", "A3", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 18)
	pdf.Cell(0, 10, "Course Details")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Resources:")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	for _, resource := range courseDetails.Resources {
		pdf.Cell(0, 10, fmt.Sprintf("%s (%s)", resource.ResourceName, resource.ResourceType))
		pdf.Ln(10)
	}

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Chapters:")
	pdf.Ln(15)

	// Table header
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(20, 10, "SN")
	pdf.Cell(60, 10, "Chapter")
	pdf.Cell(30, 10, "Quiz")
	pdf.Cell(60, 10, "Assignment")
	pdf.Cell(40, 10, "Topics")
	pdf.Cell(40, 10, "Lecture")
	pdf.Cell(30, 10, "Room")
	pdf.Cell(60, 10, "Time")
	pdf.Ln(10)

	// Draw table border for header
	pdf.SetLineWidth(0.5)
	pdf.SetDrawColor(0, 0, 0)
	x := pdf.GetX()
	y := pdf.GetY() - 10
	w := 350.0
	h := 10.0
	pdf.Rect(x, y, w, h, "D")

	sn := 1
	for _, chapter := range courseDetails.Chapters {
		for _, quiz := range chapter.Quizzes {
			for _, assignment := range chapter.Assignments {
				for _, topic := range chapter.Topics {
					for _, lecture := range topic.Lectures {
						pdf.SetFont("Arial", "", 10)
						pdf.Cell(20, 10, strconv.Itoa(sn))
						pdf.Cell(60, 10, chapter.ChapterName)
						pdf.Cell(30, 10, fmt.Sprintf("Quiz %d", quiz.QuizID))
						pdf.Cell(60, 10, assignment.AssignmentType)
						pdf.Cell(40, 10, topic.TopicName)
						pdf.Cell(40, 10, fmt.Sprintf("Lecture %d", lecture.LectureID))
						pdf.Cell(30, 10, lecture.Room)
						pdf.Cell(60, 10, fmt.Sprintf("%s - %s", lecture.StartTime.Format("2006-01-02 15:04:05"), lecture.EndTime.Format("2006-01-02 15:04:05")))
						pdf.Ln(10)

						// Draw table border for row
						pdf.SetLineWidth(0.5)
						pdf.SetDrawColor(0, 0, 0)
						x = pdf.GetX()
						y = pdf.GetY() - 10
						w = 350
						h = 10
						pdf.Rect(x, y, w, h, "D")
						sn++
					}
				}
			}
		}
	}

	err = pdf.Output(response)
	if err != nil {
		http.Error(response, "Failed to output PDF", http.StatusInternalServerError)
	}

}
