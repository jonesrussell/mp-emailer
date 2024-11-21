package session

import (
	"context"
	"time"

	"github.com/gorilla/sessions"
	"github.com/jonesrussell/loggo"
	"github.com/labstack/echo/v4"
)

type manager struct {
	store   Store
	logger  loggo.LoggerInterface
	options Options
	cleaner *Cleaner
}

func NewManager(store Store, logger loggo.LoggerInterface, options Options) Manager {
	m := &manager{
		store:   store,
		logger:  logger,
		options: options,
	}

	// Configure store with security options
	store.SetSecure(options.Secure)
	store.SetSameSite(options.SameSite)
	store.SetOptions(&sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
		SameSite: options.SameSite,
	})

	// Initialize cleaner
	m.cleaner = NewCleaner(store, options.CleanupInterval, options.MaxAge, logger)

	return m
}

func (m *manager) GetSession(c echo.Context, name string) (*sessions.Session, error) {
	session, err := m.store.Get(c.Request(), name)
	if err != nil {
		m.logger.Error("Failed to get session", err)
		return nil, ErrSessionNotFound
	}

	// Update last accessed time
	session.Values["last_accessed"] = time.Now()
	return session, nil
}

func (m *manager) SaveSession(c echo.Context, session *sessions.Session) error {
	if err := session.Save(c.Request(), c.Response().Writer); err != nil {
		m.logger.Error("Failed to save session", err)
		return ErrSessionStoreFailed
	}
	return nil
}

func (m *manager) ClearSession(c echo.Context, name string) error {
	session, err := m.GetSession(c, name)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return m.SaveSession(c, session)
}

func (m *manager) RegenerateSession(c echo.Context, name string) (*sessions.Session, error) {
	oldSession, err := m.GetSession(c, name)
	if err != nil {
		return nil, err
	}

	// Store old values
	values := oldSession.Values

	// Clear old session
	if err := m.ClearSession(c, name); err != nil {
		return nil, err
	}

	// Create new session
	newSession, err := m.store.Get(c.Request(), name)
	if err != nil {
		return nil, err
	}

	// Copy old values to new session
	for k, v := range values {
		newSession.Values[k] = v
	}

	// Update creation time
	newSession.Values["created_at"] = time.Now()

	if err := m.SaveSession(c, newSession); err != nil {
		return nil, err
	}

	return newSession, nil
}

func (m *manager) SetSessionValues(sess *sessions.Session, userData interface{}) {
	if userData == nil {
		return
	}
	if ud, ok := userData.(UserData); ok {
		sessionData := Data{
			UserID:          ud.GetID(),
			Username:        ud.GetUsername(),
			LastAccessed:    time.Now(),
			CreatedAt:       time.Now(),
			IsAuthenticated: true,
			CustomData:      ud.GetCustomData(),
		}

		sess.Values["user_id"] = sessionData.UserID
		sess.Values["username"] = sessionData.Username
		sess.Values["last_accessed"] = sessionData.LastAccessed
		sess.Values["created_at"] = sessionData.CreatedAt
		sess.Values["is_authenticated"] = sessionData.IsAuthenticated
		sess.Values["custom_data"] = sessionData.CustomData
	}
}

func (m *manager) GetSessionValue(sess *sessions.Session, key string) interface{} {
	if value, ok := sess.Values[key]; ok {
		return value
	}
	return nil
}

func (m *manager) DeleteSessionValue(sess *sessions.Session, key string) {
	delete(sess.Values, key)
}

func (m *manager) IsAuthenticated(c echo.Context) bool {
	session, err := m.GetSession(c, m.options.CookieName)
	if err != nil {
		return false
	}

	auth, ok := session.Values["is_authenticated"].(bool)
	return ok && auth
}

func (m *manager) SetAuthenticated(c echo.Context, authenticated bool) error {
	session, err := m.GetSession(c, m.options.CookieName)
	if err != nil {
		return err
	}

	session.Values["is_authenticated"] = authenticated
	return m.SaveSession(c, session)
}

func (m *manager) ValidateSession(name string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, err := m.GetSession(c, name)
			if err != nil {
				return ErrSessionNotFound
			}

			if session.IsNew {
				return ErrSessionNotFound
			}

			// Check session expiration
			lastAccessed, ok := session.Values["last_accessed"].(time.Time)
			if !ok || time.Since(lastAccessed) > time.Duration(m.options.MaxAge)*time.Second {
				return ErrSessionExpired
			}

			return next(c)
		}
	}
}

func (m *manager) StartCleanup(ctx context.Context) {
	m.cleaner.StartCleanup(ctx)
}

func (m *manager) StopCleanup() error {
	return m.cleaner.StopCleanup()
}