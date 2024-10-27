package user

import (
	"fmt"
	"time"

	"github.com/fullstackdev42/mp-emailer/config"
	"github.com/fullstackdev42/mp-emailer/shared"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	RegisterUser(params *RegisterDTO) (*DTO, error)
	LoginUser(params *LoginDTO) (string, error)
	GetUser(params *GetDTO) (*DTO, error)
}

type Service struct {
	repo     RepositoryInterface
	validate *validator.Validate
	cfg      *config.Config
}

type Config struct {
	JWTSecret string
	JWTExpiry time.Duration
}

// Explicitly implement the ServiceInterface
var _ ServiceInterface = (*Service)(nil)

// RegisterUserServiceParams for registering a user
type RegisterUserServiceParams struct {
	Username string
	Email    string
	Password string
}

func (s *Service) RegisterUser(params *RegisterDTO) (*DTO, error) {
	// Validate the DTO
	if err := s.validate.Struct(params); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	// Validate password length
	if len(params.Password) > 72 {
		return nil, fmt.Errorf("password length exceeds 72 bytes")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	// Create the user with the hashed password
	user, err := s.repo.CreateUser(&CreateDTO{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return &DTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Service) LoginUser(params *LoginDTO) (string, error) {
	// Check if user exists and password is correct
	user, err := s.repo.GetUserByUsername(params.Username)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	// Generate JWT token
	expiry, err := time.ParseDuration(s.cfg.JWTExpiry)
	if err != nil {
		return "", fmt.Errorf("invalid JWT expiry duration: %w", err)
	}
	token, err := shared.GenerateToken(params.Username, s.cfg.JWTSecret, int(expiry.Minutes()))
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return token, nil
}

func (s *Service) GetUser(params *GetDTO) (*DTO, error) {
	user, err := s.repo.GetUserByUsername(params.Username)
	if err != nil {
		return nil, err
	}
	return &DTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
