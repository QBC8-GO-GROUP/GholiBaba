package cmd

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"transportAgency/internal/agency/service"
	adapter "transportAgency/pkg/adapters/storage"
)

func main() {
	agencyRepo := adapter.NewInMemoryAgencyRepo()
	createAgencyService := &service.CreateAgencyService{
		Repo: agencyRepo,
	}

	app := fiber.New()

	app.Post("/create-agency", func(c *fiber.Ctx) error {
		var command service.CreateAgencyCommand

		if err := c.BodyParser(&command); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}

		agencyID, err := createAgencyService.Execute(command)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error creating agency",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id": agencyID,
		})
	})

	log.Println("Starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
