package campaign_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/jonesrussell/mp-emailer/campaign"
	"github.com/jonesrussell/mp-emailer/internal/testutil"
	"github.com/jonesrussell/mp-emailer/middleware"
	campaignmocks "github.com/jonesrussell/mp-emailer/mocks/campaign"
	sharedmocks "github.com/jonesrussell/mp-emailer/mocks/shared"
	"github.com/jonesrussell/mp-emailer/shared"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	testutil.BaseTestSuite
	handler          *campaign.Handler
	CampaignService  *campaignmocks.MockServiceInterface
	TemplateRenderer *sharedmocks.MockTemplateRendererInterface
	ErrorHandler     *sharedmocks.MockErrorHandlerInterface
}

func (s *HandlerTestSuite) SetupTest() {
	s.BaseTestSuite.SetupTest()

	// Register renderer with Echo
	s.Echo.Renderer = s.TemplateRenderer

	params := campaign.HandlerParams{
		BaseHandlerParams: shared.BaseHandlerParams{
			Logger:           s.Logger,
			ErrorHandler:     s.ErrorHandler,
			TemplateRenderer: s.TemplateRenderer,
			Store:            s.Store,
			Config:           s.Config,
		},
		Service:                     s.CampaignService,
		RepresentativeLookupService: s.RepresentativeLookupService,
		EmailService:                s.EmailService,
		Client:                      s.CampaignClient,
	}

	result, err := campaign.NewHandler(params)
	s.NoError(err)
	s.handler = result.Handler
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (s *HandlerTestSuite) TestCampaignGET() {
	testCases := []struct {
		name           string
		setupMocks     func(c echo.Context)
		campaignID     string
		expectedStatus int
		expectedError  error
	}{
		{
			name: "successful_campaign_fetch",
			setupMocks: func(c echo.Context) {
				// Mock middleware manager
				middlewareManager := &middleware.Manager{}
				c.Set("middleware_manager", middlewareManager)

				// Setup campaign service mock
				campaignData := &campaign.Campaign{
					BaseModel: shared.BaseModel{
						ID: uuid.MustParse("b8568959-70eb-42f9-bde6-57250faced25"),
					},
					Name:        "Test Campaign",
					Description: "Test Description",
					Template:    "Test Template",
					OwnerID:     uuid.New(),
				}

				// Setup campaign service expectations using the correct method name
				s.CampaignService.EXPECT().
					GetCampaignByID(mock.Anything, campaign.GetCampaignParams{
						ID: campaignData.ID,
					}).Return(campaignData, nil)

				// Setup template renderer expectations
				s.TemplateRenderer.EXPECT().
					Render(
						mock.Anything,
						"campaign/show.html",
						"layouts/main",
						mock.AnythingOfType("map[string]interface{}"),
					).Return(nil)

				// Setup error handler expectations
				s.ErrorHandler.EXPECT().
					HandleHTTPError(
						mock.AnythingOfType("*echo.context"),
						mock.AnythingOfType("*errors.errorString"),
						"Internal server error",
						http.StatusInternalServerError,
					).Return(echo.NewHTTPError(http.StatusInternalServerError)).Maybe()
			},
			campaignID:     "b8568959-70eb-42f9-bde6-57250faced25",
			expectedStatus: http.StatusOK,
		},
		// ... other test cases ...
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/campaign/"+tc.campaignID, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/campaign/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.campaignID)

			// Setup mocks
			if tc.setupMocks != nil {
				tc.setupMocks(c)
			}

			// Execute
			err := s.handler.CampaignGET(c)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}
			s.Equal(tc.expectedStatus, rec.Code)
		})
	}
}

func (s *HandlerTestSuite) TestGetCampaigns() {
	campaigns := []campaign.Campaign{
		{
			Name: "Campaign 1",
			BaseModel: shared.BaseModel{
				ID: uuid.New(),
			},
		},
		{
			Name: "Campaign 2",
			BaseModel: shared.BaseModel{
				ID: uuid.New(),
			},
		},
	}

	tests := []struct {
		name           string
		setupMocks     func()
		expectedStatus int
		expectedError  string
	}{
		{
			name: "successful campaigns fetch",
			setupMocks: func() {
				s.Logger.EXPECT().Debug("Handling GetCampaigns request")
				s.Logger.EXPECT().Debug("Rendering all campaigns", "count", len(campaigns))

				s.CampaignService.EXPECT().GetCampaigns(
					mock.Anything,
				).Return(campaigns, nil)

				s.TemplateRenderer.EXPECT().Render(
					mock.Anything,
					"campaigns",
					mock.MatchedBy(func(data shared.Data) bool {
						content, ok := data.Content.(map[string]interface{})
						return ok && data.Title == "Campaigns" &&
							data.PageName == "campaigns" &&
							len(content["Campaigns"].([]campaign.Campaign)) == len(campaigns)
					}),
					mock.Anything,
				).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "service error",
			setupMocks: func() {
				s.Logger.EXPECT().Debug("Handling GetCampaigns request")

				dbErr := errors.New("database error")
				s.CampaignService.EXPECT().GetCampaigns(
					mock.Anything,
				).Return(nil, dbErr)

				// Create the HTTP error that will be returned
				httpError := echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")

				s.ErrorHandler.EXPECT().HandleHTTPError(
					mock.Anything, // echo.Context
					dbErr,         // original error
					"Internal server error",
					http.StatusInternalServerError,
				).Return(httpError).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "Internal server error",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Clear any previous mock calls
			s.SetupTest()

			tt.setupMocks()

			c := s.NewContext(http.MethodGet, "/campaigns", nil)
			err := s.handler.GetCampaigns(c)

			if tt.expectedError != "" {
				s.Error(err)
				httpErr, ok := err.(*echo.HTTPError)
				s.True(ok)
				s.Equal(tt.expectedStatus, httpErr.Code)
				s.Equal(tt.expectedError, httpErr.Message)
			} else {
				s.NoError(err)
				s.Equal(tt.expectedStatus, s.Recorder.Code)
			}
		})
	}
}

func (s *HandlerTestSuite) TestCreateCampaignForm() {
	s.Run("successful form render", func() {
		s.Logger.EXPECT().Debug("Handling CreateCampaignForm request")

		s.TemplateRenderer.EXPECT().Render(
			mock.Anything,
			"campaign_create",
			mock.MatchedBy(func(data shared.Data) bool {
				return data.Title == "Create Campaign" &&
					data.PageName == "campaign_create" &&
					data.Content == nil
			}),
			mock.Anything,
		).Return(nil)

		c := s.NewContext(http.MethodGet, "/campaign/create", nil)
		err := s.handler.CreateCampaignForm(c)

		s.NoError(err)
		s.Equal(http.StatusOK, s.Recorder.Code)
	})
}
