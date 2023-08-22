package handlers

import (
	"github.com/cornevanstraten/trivia-api/database"
	"github.com/cornevanstraten/trivia-api/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	// return c.Status(200).JSON(facts)
	return c.Render("index", fiber.Map{
		"Title":    "Trivia Time!",
		"Subtitle": "Facts for fun times with friends.",
		"Facts":    facts,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	database.DB.Db.Where("id = ?", id).First(&fact)

	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}
