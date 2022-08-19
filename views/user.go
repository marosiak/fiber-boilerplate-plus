package views

import (
	"project_module/sessioncontext"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"project_module/database"
	"project_module/models"
)

type UserViews struct {
	sessionCtx *sessioncontext.SessionContext
	log        *log.Entry
}

func NewUserViews(sessionCtx *sessioncontext.SessionContext, logger *log.Entry) *UserViews {
	return &UserViews{sessionCtx: sessionCtx, log: logger}
}

const ErrorSessionKey = "error"

func (u UserViews) UserListView(c *fiber.Ctx) error {
	users := database.Get()

	errorMsg := u.sessionCtx.Get(c, ErrorSessionKey, "")
	if errorMsg != nil {
		// We will render this error and clear it from storage, so it won't appear forever
		u.sessionCtx.Set(c, ErrorSessionKey, "")
	}

	return c.Render("index", fiber.Map{
		"error": errorMsg.(string),
		"users": users,
	})
}

func (u UserViews) AddUserView(c *fiber.Ctx) error {
	name := c.FormValue("name")

	if name == "" {
		u.log.Debug("name is empty")
		u.sessionCtx.Set(c, ErrorSessionKey, "The name cannot be empty")
		return c.Redirect("/")
	} else {
		database.Insert(&models.User{Name: name})
	}

	return c.RedirectToRoute("", fiber.Map{"error": "abc"})
}
