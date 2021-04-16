package user

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates usecase logic for users.
type Service interface {
	Me(ctx context.Context) (User, error)
	Get(ctx context.Context, id string) (User, error)
	Query(ctx context.Context, offset, limit int) ([]User, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateUserRequest) (User, error)
	Update(ctx context.Context, id string, input UpdateUserRequest) (User, error)
	Delete(ctx context.Context, id string) (User, error)
}

// User represents the data about an user.
type User struct {
	entity.User
}

// CreateUserRequest represents an user creation request.
type CreateUserRequest struct {
	Username           string  `json:"username"`
	Password           string  `json:"password"`
	RoleID             string  `json:"role_id"`
	FirstName          string  `json:"first_name"`
	LastName           string  `json:"last_name"`
	Email              string  `json:"email"`
	RegistrationNumber *string `json:"registration_number,omitempty"`
	CareerID           string  `json:"career_id"`
	GroupID            string  `json:"group_id"`
}

// Validate validates the CreateUserRequest fields.
func (m CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.Password, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.RoleID, validation.Required, validation.Length(0, 36)),
		validation.Field(&m.FirstName, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.LastName, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.Email, is.Email, validation.Length(3, 50)),
		validation.Field(&m.RegistrationNumber, validation.Length(0, 10)),
	)
}

// UpdateUserRequest represents an user update request.
type UpdateUserRequest struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	RoleID             string `json:"role_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	IsActive           bool   `json:"is_active"`
	Email              string `json:"email"`
	RegistrationNumber string `json:"registration_number,omitempty"`
	CareerID           string `json:"career_id"`
	GroupID            string `json:"group_id"`
}

// Validate validates the CreateUserRequest fields.
func (m UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.FirstName, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.LastName, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.Username, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.RoleID, validation.Required, validation.Length(0, 36)),
		validation.Field(&m.Password, validation.Length(0, 50)),
		validation.Field(&m.Email, is.Email, validation.Length(3, 50)),
		validation.Field(&m.RegistrationNumber, validation.Length(0, 10)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new user service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the user with the specified the user ID.
func (s service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}

// Create creates a new user.
func (s service) Create(ctx context.Context, req CreateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}
	id := entity.GenerateID()
	now := time.Now()

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	err = s.repo.Create(ctx, entity.User{
		ID:        id,
		Username:  req.Username,
		Password:  string(hash),
		CreatedAt: now,
		RoleID:    req.RoleID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
		CareerID:  &req.CareerID,
		GroupID:   &req.GroupID,
	})
	if err != nil {
		return User{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the user with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}

	user, err := s.Get(ctx, id)
	if err != nil {
		return user, err
	}

	now := time.Now()
	if len(req.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}
		user.Password = string(hash)
	}

	user.Username = req.Username
	user.RoleID = req.RoleID
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.UpdatedAt = &now
	user.IsActive = req.IsActive
	user.GroupID = &req.GroupID
	user.CareerID = &req.CareerID

	if err := s.repo.Update(ctx, user.User); err != nil {
		return user, err
	}
	return user, nil
}

// Delete deletes the user with the specified ID.
func (s service) Delete(ctx context.Context, id string) (User, error) {
	user, err := s.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return User{}, err
	}
	return user, nil
}

// Count returns the number of users.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the users with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]User, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []User{}
	for _, item := range items {
		result = append(result, User{item})
	}
	return result, nil
}

func (s service) Me(ctx context.Context) (User, error) {
	user, err := s.repo.Me(ctx)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}
