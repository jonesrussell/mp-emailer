package shared

import "github.com/jonesrussell/mp-emailer/logger"

// LoggableService defines the basic logging operations that a service should support
type LoggableService interface {
	Info(message string, params ...interface{})
	Warn(message string, params ...interface{})
	Error(message string, err error, params ...interface{})
}

// ServiceWithLogging combines LoggableService with common service operations
type ServiceWithLogging interface {
	LoggableService
}

// GenericLoggingDecorator adds logging functionality to any LoggableService
type GenericLoggingDecorator[T ServiceWithLogging] struct {
	Service T
	Logger  logger.Interface
}

// NewGenericLoggingDecorator creates a new instance of GenericLoggingDecorator
func NewGenericLoggingDecorator[T ServiceWithLogging](service T, log logger.Interface) *GenericLoggingDecorator[T] {
	return &GenericLoggingDecorator[T]{
		Service: service,
		Logger:  log,
	}
}

// Info logs an info message with the given parameters
func (d *GenericLoggingDecorator[T]) Info(message string, params ...interface{}) {
	d.Logger.Info(message, params...)
	d.Service.Info(message, params...)
}

// Warn logs a warning message with the given parameters
func (d *GenericLoggingDecorator[T]) Warn(message string, params ...interface{}) {
	d.Logger.Warn(message, params...)
	d.Service.Warn(message, params...)
}

// Error logs an error message with the given parameters
func (d *GenericLoggingDecorator[T]) Error(message string, err error, params ...interface{}) {
	d.Logger.Error(message, err, params...)
	d.Service.Error(message, err, params...)
}
