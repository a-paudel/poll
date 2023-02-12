package routes

import (
	"poll/models"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
)

func RegisterQuestionRoutes(app *echo.Echo) {
	var router = app.Group("/api/questions")

	router.POST("", createQuestion)

	router.GET("/:id", getQuestion)
	router.GET("/:id/listen", listenQuestion)

	router.POST("/:id/vote", vote)
}

type createInput struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

func createQuestion(c echo.Context) error {
	var input createInput
	var err = c.Bind(&input)
	if err != nil {
		return c.String(400, "Invalid request body")
	}

	var answers []models.Answer
	for _, answer := range input.Answers {
		var a models.Answer
		a.Answer = answer
		answers = append(answers, a)
	}

	var question models.Question
	question.Question = input.Question
	question.Answers = answers

	models.DB.Create(&question)

	return c.JSON(200, question)
}

func getQuestion(c echo.Context) error {
	var id = c.Param("id")
	var question models.Question
	models.DB.Preload("Answers").First(&question, id)

	if question.ID == 0 {
		return c.String(404, "Question not found")
	}
	return c.JSON(200, question)
}
func listenQuestion(c echo.Context) error {
	var id = c.Param("id")
	var question models.Question
	models.DB.Preload("Answers").First(&question, id)
	if question.ID == 0 {
		return c.String(404, "Question not found")
	}

	handler := websocket.Handler(func(ws *websocket.Conn) {
		websocket.Message.Send(ws, question.ToJson())
		for {
			var newQuestion models.Question
			models.DB.Preload("Answers").First(&newQuestion, id)
			if !reflect.DeepEqual(question, newQuestion) {
				question = newQuestion
				websocket.Message.Send(ws, question.ToJson())
			}

			time.Sleep(1 * time.Second)
		}
	})

	handler.ServeHTTP(c.Response(), c.Request())
	return nil
}

type voteInput struct {
	AnswerID uint `json:"answerId"`
}

func vote(c echo.Context) error {
	var id = c.Param("id")
	var input voteInput
	var err = c.Bind(&input)
	if err != nil {
		return c.String(400, "Invalid request body")
	}

	var answer models.Answer
	models.DB.Where("id = ? AND question_id = ?", input.AnswerID, id).First(&answer)
	if answer.ID == 0 {
		return c.String(404, "Answer not found")
	}
	models.DB.Model(&answer).Update("votes", gorm.Expr("votes + 1"))

	var question models.Question
	models.DB.Preload("Answers").First(&question, id)
	return c.JSON(200, question)
}
