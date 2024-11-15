package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fullstackdev42/mp-emailer/internal/testutil"
	"github.com/fullstackdev42/mp-emailer/shared"
	"github.com/fullstackdev42/mp-emailer/user"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	mocksUser "github.com/fullstackdev42/mp-emailer/mocks/user"
)

type HandlerTestSuite struct {
	testutil.BaseTestSuite
	handler        *user.Handler
	UserService    *mocksUser.MockServiceInterface
	UserRepo       *mocksUser.MockRepositoryInterface
	SessionManager *mocksUser.MockSessionManager
}

func (s *HandlerTestSuite) SetupTest() {
	s.BaseTestSuite.SetupTest()

	s.UserService = mocksUser.NewMockServiceInterface(s.T())
	s.UserRepo = mocksUser.NewMockRepositoryInterface(s.T())
	s.SessionManager = mocksUser.NewMockSessionManager(s.T())

	s.Config.SessionName = "test_session"

	flashHandler := shared.NewFlashHandler(shared.FlashHandlerParams{
		Store:        s.Store,
		Config:       s.Config,
		Logger:       s.Logger,
		ErrorHandler: s.ErrorHandler,
	})

	s.Echo.Renderer = s.TemplateRenderer

	params := user.HandlerParams{
		BaseHandlerParams: shared.BaseHandlerParams{
			Logger:           s.Logger,
			ErrorHandler:     s.ErrorHandler,
			TemplateRenderer: s.TemplateRenderer,
			Store:            s.Store,
			Config:           s.Config,
		},
		Service:        s.UserService,
		FlashHandler:   flashHandler,
		Repo:           s.UserRepo,
		SessionManager: s.SessionManager,
	}

	s.handler = user.NewHandler(params)
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (s *HandlerTestSuite) debugResponse(rec *httptest.ResponseRecorder) {
	s.T().Logf("Response Status: %d", rec.Code)
	s.T().Logf("Response Headers: %v", rec.Header())
	s.T().Logf("Response Body: %s", rec.Body.String())
}

func (s *HandlerTestSuite) TestLoginPOST() {
	tests := []struct {
		name    string
		payload string

		setupMocks     func() *sessions.Session
		expectedStatus int
		expectedPath   string
	}{
		{
			name:    "Successful login",
			payload: `{"username": "testuser", "password": "password123"}`,
			setupMocks: func() *sessions.Session {
				sess := sessions.NewSession(s.Store, "test_session")
				sess.Values = make(map[interface{}]interface{})

				testUser := &user.User{
					BaseModel: shared.BaseModel{
						ID: uuid.New(),
					},
					Username: "testuser",
				}

				s.SessionManager.On("GetSession", mock.Anything).Return(sess, nil)
				s.SessionManager.On("SetSessionValues", sess, testUser)
				s.SessionManager.On("SaveSession", mock.Anything, sess).Return(nil)

				s.UserService.EXPECT().AuthenticateUser("testuser", "password123").
					Return(true, testUser, nil)

				return sess
			},
			expectedStatus: http.StatusSeeOther,
			expectedPath:   "/",
		},
		{
			name:    "Session store error",
			payload: `{"username": "testuser", "password": "password123"}`,
			setupMocks: func() *sessions.Session {
				s.SessionManager.On("GetSession", mock.Anything).
					Return(nil, errors.New("session store error"))

				s.ErrorHandler.On("HandleHTTPError",
					mock.MatchedBy(func(c echo.Context) bool { return true }),
					mock.MatchedBy(func(err error) bool { return err.Error() == "session store error" }),
					"Error getting session",
					http.StatusInternalServerError,
				).Return(echo.NewHTTPError(http.StatusInternalServerError))

				return nil
			},
			expectedStatus: http.StatusInternalServerError,
			expectedPath:   "",
		},
		{
			name:    "Invalid credentials",
			payload: `{"username": "wronguser", "password": "wrongpass"}`,
			setupMocks: func() *sessions.Session {
				sess := sessions.NewSession(s.Store, "test_session")
				sess.Values = make(map[interface{}]interface{})

				s.SessionManager.On("GetSession", mock.Anything).Return(sess, nil)

				s.TemplateRenderer.On("Render",
					mock.AnythingOfType("*bytes.Buffer"),
					"login",
					mock.AnythingOfType("*shared.Data"),
					mock.AnythingOfType("*echo.context"),
				).Return(nil)

				s.UserService.EXPECT().AuthenticateUser("wronguser", "wrongpass").
					Return(false, nil, nil)

				return sess
			},
			expectedStatus: http.StatusUnauthorized,
			expectedPath:   "",
		},
		{
			name:    "Session save error",
			payload: `{"username": "testuser", "password": "password123"}`,
			setupMocks: func() *sessions.Session {
				sess := sessions.NewSession(s.Store, "test_session")
				sess.Values = make(map[interface{}]interface{})

				testUser := &user.User{Username: "testuser"}

				s.SessionManager.On("GetSession", mock.Anything).Return(sess, nil)
				s.SessionManager.On("SetSessionValues", sess, testUser)
				s.SessionManager.On("SaveSession", mock.Anything, sess).
					Return(errors.New("failed to save session"))

				s.UserService.EXPECT().AuthenticateUser("testuser", "password123").
					Return(true, testUser, nil)

				s.ErrorHandler.On("HandleHTTPError",
					mock.MatchedBy(func(c echo.Context) bool {
						s.T().Logf("Matching context: %+v", c)
						return true
					}),
					mock.MatchedBy(func(err error) bool {
						s.T().Logf("Matching error: %v", err)
						return true
					}),
					mock.AnythingOfType("string"),
					http.StatusInternalServerError,
				).Run(func(args mock.Arguments) {
					// Set the status code on the response
					c := args.Get(0).(echo.Context)
					c.Response().Status = http.StatusInternalServerError
				}).Return(echo.NewHTTPError(http.StatusInternalServerError))

				return sess
			},
			expectedStatus: http.StatusInternalServerError,
			expectedPath:   "",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.SetupTest() // Reset mocks for each test case
			sess := tt.setupMocks()

			// Create request with JSON payload
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			// Create Echo context
			c := s.Echo.NewContext(req, rec)
			if sess != nil {
				c.Set("session", sess)
			}

			// Execute handler
			err := s.handler.LoginPOST(c)
			if err != nil {
				s.T().Logf("Handler returned error: %v", err)
				if he, ok := err.(*echo.HTTPError); ok {
					s.T().Logf("HTTP Error: code=%d, message=%v", he.Code, he.Message)
					rec.Code = he.Code // Set the status code from the error
				}
			}

			// Debug response
			s.debugResponse(rec)

			// Verify response
			s.Equal(tt.expectedStatus, rec.Code)
			if tt.expectedPath != "" {
				s.Equal(tt.expectedPath, rec.Header().Get("Location"))
			}

			// Verify all mocks
			s.SessionManager.AssertExpectations(s.T())
		})
	}
}
