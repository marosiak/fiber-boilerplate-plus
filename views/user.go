package views

import (
	"github.com/gofiber/fiber/v2"
	"project_module/database"
	"project_module/models"
)

func UserListView(c *fiber.Ctx) error {
	users := database.Get()
	return c.Render("index", fiber.Map{
		"users": users,
	})
}

func AddUserView(c *fiber.Ctx) error {
	name := c.FormValue("name")

	if name != "" {
		// TODO
	}
	database.Insert(&models.User{Name: name})

	return c.RedirectToRoute("", fiber.Map{})
}
