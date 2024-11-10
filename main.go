package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/fullstackdev42/mp-emailer/api"
	"github.com/fullstackdev42/mp-emailer/campaign"
	"github.com/fullstackdev42/mp-emailer/config"
	"github.com/fullstackdev42/mp-emailer/database"
	"github.com/fullstackdev42/mp-emailer/server"
	"github.com/fullstackdev42/mp-emailer/shared"
	"github.com/fullstackdev42/mp-emailer/user"
	"github.com/gorilla/sessions"
	"github.com/jonesrussell/loggo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	// Check for required configuration before starting the application
	if err := config.CheckRequired(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Initialize application using uber/fx dependency injection
	app := fx.New(
		fx.Options(
			shared.App,
			campaign.Module,
			user.Module,
			server.Module,
			api.Module,
			database.MigrationModule,
			fx.Invoke(registerRoutes, startServer),
		),
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ConsoleLogger{W: os.Stdout}
		}),
	)

	app.Run()
}

// registerRoutes centralizes all route registration for the application
// It takes in all necessary handlers and services via dependency injection
// registerRoutes centralizes all route registration for the application
func registerRoutes(
	e *echo.Echo,
	serverHandler server.HandlerInterface,
	campaignHandler *campaign.Handler,
	userHandler *user.Handler,
	apiHandler *api.Handler,
	renderer shared.TemplateRendererInterface,
	sessionStore sessions.Store,
	cfg *config.Config,
	logger loggo.LoggerInterface,
) {
	// Set custom template renderer for HTML responses
	e.Renderer = renderer

	// Register middleware and route handlers separately for better organization
	registerMiddlewares(e, sessionStore, logger)
	registerHandlers(e, serverHandler, campaignHandler, userHandler, apiHandler, cfg, logger)

	// Serve static files from web/public directory
	e.Static("/static", "web/public")
}

// registerMiddlewares configures all middleware for the application
func registerMiddlewares(e *echo.Echo, sessionStore sessions.Store, logger loggo.LoggerInterface) {
	// Add logger middleware first
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("logger", logger)
			return next(c)
		}
	})

	// Make session store available in all route handlers via context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("store", sessionStore)
			return next(c)
		}
	})

	// Enable request logging for debugging and monitoring
	e.Use(middleware.Logger())

	// Implement rate limiting
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// In your server setup
	e.Use(shared.MethodOverride())
}

// registerHandlers configures all route handlers for the application
func registerHandlers(
	e *echo.Echo,
	serverHandler server.HandlerInterface,
	campaignHandler *campaign.Handler,
	userHandler *user.Handler,
	apiHandler *api.Handler,
	cfg *config.Config,
	logger loggo.LoggerInterface,
) {
	server.RegisterRoutes(serverHandler, e)
	campaign.RegisterRoutes(campaignHandler, e, cfg, logger)
	user.RegisterRoutes(userHandler, e)
	api.RegisterRoutes(apiHandler, e, cfg.JWTSecret)
}

// startServer configures the server and starts it
func startServer(lc fx.Lifecycle, e *echo.Echo, cfg *config.Config, logger loggo.LoggerInterface) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				addr := fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort)
				logger.Info("Starting server", "host", cfg.AppHost, "port", cfg.AppPort)
				if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("Server error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
