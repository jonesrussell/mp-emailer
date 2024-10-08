package handlers

import (
	"net/http"

	"github.com/fullstackdev42/mp-emailer/pkg/api"
	"github.com/fullstackdev42/mp-emailer/pkg/database"
	"github.com/fullstackdev42/mp-emailer/pkg/services"
	"github.com/gorilla/sessions"
	"github.com/jonesrussell/loggo"
	"github.com/labstack/echo/v4"
)

const (
	adminEmail = "admin@example.com" // Replace with your actual admin email
)

type Handler struct {
	logger       loggo.LoggerInterface
	client       api.ClientInterface
	store        sessions.Store
	db           *database.DB
	emailService services.EmailService
}

type TemplateData struct {
	IsAuthenticated bool
	// Add other fields as needed
}

func NewHandler(logger loggo.LoggerInterface, client api.ClientInterface, sessionSecret string, db *database.DB, emailService services.EmailService) *Handler {
	store := sessions.NewCookieStore([]byte(sessionSecret))
	return &Handler{
		logger:       logger,
		client:       client,
		store:        store,
		db:           db,
		emailService: emailService,
	}
}

func (h *Handler) getSession(c echo.Context) (*sessions.Session, error) {
	return h.store.Get(c.Request(), "session")
}

func (h *Handler) saveSession(session *sessions.Session, c echo.Context) error {
	return session.Save(c.Request(), c.Response().Writer)
}

func (h *Handler) handleError(err error, statusCode int, message string) error {
	h.logger.Error(message, err)
	return echo.NewHTTPError(statusCode, message)
}

func (h *Handler) HandleLogin(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "login.html", nil)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	valid, err := h.db.VerifyUser(username, password)
	if err != nil {
		return h.handleError(err, http.StatusInternalServerError, "Error verifying user")
	}

	if !valid {
		data := map[string]interface{}{
			"Error": "Invalid username or password",
		}
		return c.Render(http.StatusUnauthorized, "login.html", data)
	}

	session, err := h.getSession(c)
	if err != nil {
		return h.handleError(err, http.StatusInternalServerError, "Failed to get session")
	}

	session.Values["user"] = username
	if err := h.saveSession(session, c); err != nil {
		return h.handleError(err, http.StatusInternalServerError, "Failed to save session")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) HandleLogout(c echo.Context) error {
	session, err := h.getSession(c)
	if err != nil {
		return h.handleError(err, http.StatusInternalServerError, "Failed to get session")
	}

	session.Values["user"] = nil
	if err := h.saveSession(session, c); err != nil {
		return h.handleError(err, http.StatusInternalServerError, "Failed to save session during logout")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := h.store.Get(c.Request(), "session")
		user := session.Values["user"]
		c.Set("isAuthenticated", user != nil)
		return next(c)
	}
}
