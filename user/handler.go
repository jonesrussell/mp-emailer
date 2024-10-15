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

func NewHandler(
	service ServiceInterface,
	logger loggo.LoggerInterface,
	config *config.Config,
) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
		config:  config,
	}
}

func (h *Handler) RegisterGET(c echo.Context) error {
	sess, err := session.Get(h.config.SessionName, c)
	if err != nil {
		h.logger.Error("Failed to get session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	data := map[string]interface{}{}
	flashes := sess.Flashes()
	if len(flashes) > 0 {
		data["Error"] = flashes[0]
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		h.logger.Error("Failed to save session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.Render(http.StatusOK, "register.html", data)
}

func (h *Handler) RegisterPOST(c echo.Context) error {
	sess, err := session.Get(h.config.SessionName, c)
	if err != nil {
		h.logger.Error("Failed to get session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	if username == "" || password == "" || email == "" {
		h.logger.Warn("Missing required fields", "username", username, "email", email, "password", "***")
		sess.AddFlash("Username, password, and email are required")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			h.logger.Error("Failed to save session", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return c.Redirect(http.StatusSeeOther, "/register")
	}

	if err := h.service.RegisterUser(username, password, email); err != nil {
		h.logger.Error("Failed to register user", err)
		sess.AddFlash("Failed to register user. Please try again.")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			h.logger.Error("Failed to save session", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return c.Redirect(http.StatusSeeOther, "/register")
	}

	h.logger.Debug("User registered successfully")
	return c.Redirect(http.StatusSeeOther, "/login")
}

// Handler for GET requests
func (h *Handler) LoginGET(c echo.Context) error {
	sess, err := session.Get(h.config.SessionName, c)
	if err != nil {
		h.logger.Error("Failed to get session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"Error": sess.Flashes("error"),
	})
}

// Handler for POST requests
func (h *Handler) LoginPOST(c echo.Context) error {
	sess, err := session.Get(h.config.SessionName, c)
	if err != nil {
		h.logger.Error("Failed to get session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	username := c.FormValue("username")
	password := c.FormValue("password")

	h.logger.Debug("Login called with method: POST")
	h.logger.Debug("Login attempt for username: " + username)

	if username == "" || password == "" {
		h.logger.Warn("Login attempt with empty username or password")
		sess.AddFlash("Username and password are required", "error")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			h.logger.Error("Failed to save session", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	userID, err := h.service.VerifyUser(username, password)
	if err != nil {
		h.logger.Warn("Login failed for user: " + username)
		sess.AddFlash("Invalid username or password", "error")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			h.logger.Error("Failed to save session", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	sess.Values["userID"] = userID
	sess.Values["username"] = username
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		h.logger.Error("Failed to save session", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
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
