package app

import (
	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/http/handler"
	"github.com/coderkamlesh/hypershop_go/internal/repository"
	"github.com/coderkamlesh/hypershop_go/internal/service"
)

// Container holds all application handlers
type Container struct {
	AuthHandler *handler.AuthHandler
	AuthService service.AuthService

	// UserHandler    *handler.UserHandler    // Future
	// ProductHandler *handler.ProductHandler // Future
}

// NewContainer initializes all handlers with dependency injection
func NewContainer() *Container {

	db := config.DB

	// Repositories
	userRepo := repository.NewUserRepository(db)
	regOtpRepo := repository.NewRegistrationOtpRepository(db)
	userOtpRepo := repository.NewUserOtpRepository(db)
	userSessionRepo := repository.NewUserSessionRepository(db)

	// Services
	authService := service.NewAuthService(
		userRepo,
		regOtpRepo,
		userOtpRepo,
		userSessionRepo,
	)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)

	return &Container{
		AuthHandler: authHandler,
		AuthService: authService, // ✅ IMPORTANT
	}
}
