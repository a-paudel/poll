package routes

import (
	"poll/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterQuestionRoutes(app *echo.Echo) {
	var routes = app.Group("/api/questions")

	routes.POST("", createQuestion)
	routes.GET("/:id", getQuestion)
	routes.POST("/vote", vote)

}

type createInput struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

func createQuestion(c echo.Context) error {
	var input createInput
	var err = c.Bind(&input)
	if err != nil {
		return c.String(400, "Invalid input")
	}

	var question = models.Question{}
	err = question.Create(input.Question)
	if err != nil {
		return c.String(500, "Error creating question")
	}

	for _, answer := range input.Answers {
		var ans = models.Answer{}
		err = ans.Create(answer, question.ID)
		if err != nil {
			return c.String(500, "Error creating answer "+answer)
		}
	}
	question.Get(question.ID)
	return c.JSON(200, question)
}

func getQuestion(c echo.Context) error {
	var idString = c.Param("id")
	var id, err = strconv.Atoi(idString)
	if err != nil {
		return c.String(400, "Invalid id")
	}
	var question = models.Question{}
	err = question.Get(id)
	if err != nil {
		return c.String(404, "Question not found")
	}
	return c.JSON(200, question)
}

type voteInput struct {
	ID int `json:"id"`
}

func vote(c echo.Context) error {
	var input voteInput
	var err = c.Bind(&input)
	if err != nil {
		return c.String(400, "Invalid input")
	}
	var ans = models.Answer{}
	err = ans.Vote(input.ID)
	if err != nil {
		return c.String(500, "Error voting")
	}

	return c.JSON(200, ans)
}
