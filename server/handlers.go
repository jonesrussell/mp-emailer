package server

import (
	"net/http"

	"github.com/fullstackdev42/mp-emailer/campaign"
	"github.com/fullstackdev42/mp-emailer/email"
	"github.com/fullstackdev42/mp-emailer/shared"
	"github.com/fullstackdev42/mp-emailer/user"
	"github.com/gorilla/sessions"
	"github.com/jonesrussell/loggo"
	"github.com/labstack/echo/v4"
)

// Handler struct
type Handler struct {
	Logger          loggo.LoggerInterface
	Store           sessions.Store
	emailService    email.Service
	templateManager *TemplateManager
	userService     user.ServiceInterface
	campaignService campaign.ServiceInterface
	errorHandler    *shared.ErrorHandler
}

// NewHandler creates a new Handler instance
func NewHandler(
	logger loggo.LoggerInterface,
	emailService email.Service,
	tmplManager *TemplateManager,
	userService user.ServiceInterface,
	campaignService campaign.ServiceInterface,
) *Handler {
	return &Handler{
		Logger:          logger,
		emailService:    emailService,
		templateManager: tmplManager,
		userService:     userService,
		campaignService: campaignService,
		errorHandler:    shared.NewErrorHandler(logger),
	}
}

// Page data struct
type PageData struct {
	Campaigns []campaign.Campaign
}

// Home page handler
func (h *Handler) HandleIndex(c echo.Context) error {
	campaigns, err := h.campaignService.GetAllCampaigns()
	if err != nil {
		return h.errorHandler.HandleError(c, err, http.StatusInternalServerError, "Error fetching campaigns")
	}

	pageData := PageData{Campaigns: campaigns}
	return c.Render(http.StatusOK, "index.html", pageData)
}
