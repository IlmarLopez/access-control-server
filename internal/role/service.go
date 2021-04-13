package role

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for roles.
type Service interface {
	Get(ctx context.Context, id string) (Role, error)
	Query(ctx context.Context, offset, limit int) ([]Role, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateRoleRequest) (Role, error)
	Update(ctx context.Context, id string, input UpdateRoleRequest) (Role, error)
	Delete(ctx context.Context, id string) (Role, error)
}

// Role represents the data about an role.
type Role struct {
	entity.Role
}

// CreateRoleRequest represents an role creation request.
type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate validates the CreateRoleRequest fields.
func (m CreateRoleRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateRoleRequest represents an role update request.
type UpdateRoleRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// Validate validates the CreateRoleRequest fields.
func (m UpdateRoleRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new role service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the role with the specified the role ID.
func (s service) Get(ctx context.Context, id string) (Role, error) {
	role, err := s.repo.Get(ctx, id)
	if err != nil {
		return Role{}, err
	}
	return Role{role}, nil
}

// Create creates a new role.
func (s service) Create(ctx context.Context, req CreateRoleRequest) (Role, error) {
	if err := req.Validate(); err != nil {
		return Role{}, err
	}
	id := entity.GenerateID()
	err := s.repo.Create(ctx, entity.Role{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return Role{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the role with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateRoleRequest) (Role, error) {
	if err := req.Validate(); err != nil {
		return Role{}, err
	}

	role, err := s.Get(ctx, id)
	if err != nil {
		return role, err
	}
	role.Name = req.Name
	role.Description = req.Description

	if err := s.repo.Update(ctx, role.Role); err != nil {
		return role, err
	}
	return role, nil
}

// Delete deletes the role with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Role, error) {
	role, err := s.Get(ctx, id)
	if err != nil {
		return Role{}, err
	}

	if err = s.repo.Delete(ctx, id); err != nil {
		return Role{}, err
	}
	return role, nil
}

// Count returns the number of roles.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the roles with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Role, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Role{}
	for _, item := range items {
		result = append(result, Role{item})
	}
	return result, nil
}
