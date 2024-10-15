package user

import (
	"net/http"

	"github.com/fullstackdev42/mp-emailer/config"
	"github.com/jonesrussell/loggo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service ServiceInterface
	logger  loggo.LoggerInterface
	config  *config.Config
}

func NewHandler(service ServiceInterface, logger loggo.LoggerInterface, config *config.Config) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
		config:  config,
	}
}

func (h *Handler) HandleRegister(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "register.html", nil)
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	err := h.service.RegisterUser(username, email, password)
	if err != nil {
		h.logger.Error("Registration error", err)
		return c.Render(http.StatusBadRequest, "register.html", map[string]interface{}{
			"Error": err.Error(),
		})
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}

func (h *Handler) HandleLogin(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "login.html", nil)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	h.logger.Debug("HandleLogin called with method: POST")
	h.logger.Debug("Login attempt for username: " + username)

	if username == "" || password == "" {
		h.logger.Warn("Login attempt with empty username or password")
		return echo.NewHTTPError(http.StatusBadRequest, "Username and password are required")
	}

	userID, err := h.service.VerifyUser(username, password)
	if err != nil {
		h.logger.Warn("Login failed for user: " + username)
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	sess, err := session.Get(h.config.SessionName, c)
	if err != nil {
		h.logger.Error("Failed to get session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	sess.Values["userID"] = userID
	sess.Values["username"] = username
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		h.logger.Error("Failed to save session", err)
		return c.String(http.StatusInternalServerError, "Failed to save session")
	}

	return c.Redirect(http.StatusSeeOther, "/campaigns")
}

func (h *Handler) HandleLogout(c echo.Context) error {
	h.logger.Debug("Handling logout request")
	sess, err := session.Get(h.config.SessionName, c)
	if err != nil {
		h.logger.Error("Failed to get session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	// Clear session
	sess.Values["userID"] = nil
	sess.Values["username"] = nil
	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		h.logger.Error("Failed to save session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
