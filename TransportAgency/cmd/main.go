package cmd

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"transportAgency/internal/agency/service"
	adapter "transportAgency/pkg/adapters/storage"
)

func main() {
	// Initialize the repository and service
	agencyRepo := adapter.NewInMemoryAgencyRepo()
	createAgencyService := &service.CreateAgencyService{
		Repo: agencyRepo,
	}

	// Create a new Fiber app
	app := fiber.New()

	// Define the HTTP handler for creating an agency
	app.Post("/create-agency", func(c *fiber.Ctx) error {
		var command service.CreateAgencyCommand

		// Parse the request body into CreateAgencyCommand
		if err := c.BodyParser(&command); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}

		// Execute the service to create the agency
		agencyID, err := createAgencyService.Execute(command)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error creating agency",
			})
		}

		// Send the response back with the created agency ID
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id": agencyID,
		})
	})

	// Start the server
	log.Println("Starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
