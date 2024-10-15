package server

import (
	"net/http"

	"github.com/fullstackdev42/mp-emailer/email"
	"github.com/fullstackdev42/mp-emailer/user"
	"github.com/gorilla/sessions"
	"github.com/jonesrussell/loggo"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Logger          *loggo.Logger
	Store           sessions.Store
	emailService    email.Service
	templateManager *TemplateManager
	userService     user.ServiceInterface
}

func NewHandler(
	logger *loggo.Logger,
	store sessions.Store,
	emailService email.Service,
	tmplManager *TemplateManager,
	userService user.ServiceInterface,
) *Handler {
	return &Handler{
		Logger:          logger,
		Store:           store,
		emailService:    emailService,
		templateManager: tmplManager,
		userService:     userService,
	}
}

func (h *Handler) HandleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (h *Handler) handleError(err error, statusCode int, message string) error {
	h.Logger.Error(message, err)
	return echo.NewHTTPError(statusCode, message)
}

func (h *Handler) HandleEcho(c echo.Context) error {
	type EchoRequest struct {
		Message string `json:"message"`
	}

	req := new(EchoRequest)
	if err := c.Bind(req); err != nil {
		return h.handleError(err, http.StatusBadRequest, "Error binding request")
	}

	return c.JSON(http.StatusOK, req)
}