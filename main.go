package main

import (
	"fmt"
	"net/http"

	"embed"

	"github.com/fullstackdev42/mp-emailer/campaign"
	"github.com/fullstackdev42/mp-emailer/config"
	"github.com/fullstackdev42/mp-emailer/email"
	"github.com/fullstackdev42/mp-emailer/internal/database"
	"github.com/fullstackdev42/mp-emailer/pkg/handlers"
	"github.com/fullstackdev42/mp-emailer/pkg/server"
	"github.com/fullstackdev42/mp-emailer/routes"
	"github.com/fullstackdev42/mp-emailer/user"
	"github.com/gorilla/sessions"
	"github.com/jonesrussell/loggo"
)

//go:embed web/templates/* web/templates/partials/*
var templateFS embed.FS

const migrationsPath string = "./migrations"

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	logger, err := loggo.NewLogger("mp-emailer.log", config.GetLogLevel())
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		return
	}

	db, err := database.NewDB(config.DatabaseDSN(), logger, migrationsPath)
	if err != nil {
		logger.Error("Error connecting to database", err)
		return
	}
	defer db.SQL.Close()

	emailService := email.NewEmailService(config)

	tmplManager, err := server.NewTemplateManager(templateFS)
	if err != nil {
		logger.Error("Error initializing templates", err)
		return
	}

	// Log the current log level
	logger.Info(fmt.Sprintf("Application started with log level: %v", config.GetLogLevel()))

	// Create a session store (you need to import and configure this)
	store := sessions.NewCookieStore([]byte(config.SessionSecret))

	handler := handlers.NewHandler(
		logger,
		store,
		emailService,
		tmplManager,
	)

	campaignRepo := campaign.NewRepository(db.SQL)
	campaignService := campaign.NewService(campaignRepo)
	representativeLookupService := campaign.NewRepresentativeLookupService(logger)
	defaultClient := campaign.NewDefaultClient(logger)
	campaignHandler := campaign.NewHandler(campaignService, logger, representativeLookupService, emailService, defaultClient)
	userRepo := user.NewRepository(db.SQL)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService, logger)

	e := server.New(config, logger.(*loggo.Logger), tmplManager)
	routes.RegisterRoutes(e, handler, campaignHandler, userHandler)

	logger.Info(fmt.Sprintf("Attempting to start server on :%s", config.AppPort))
	if err := e.Start(":" + config.AppPort); err != http.ErrServerClosed {
		logger.Error("Error starting server", err)
	}
}
