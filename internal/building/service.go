package building

import (
	"context"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for buildings.
type Service interface {
	Get(ctx context.Context, id string) (Building, error)
	Query(ctx context.Context, offset, limit int) ([]Building, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateBuildingRequest) (Building, error)
	Update(ctx context.Context, id string, input UpdateBuildingRequest) (Building, error)
	Delete(ctx context.Context, id string) (Building, error)
}

// Building represents the data about an building.
type Building struct {
	entity.Building
}

// CreateBuildingRequest represents an building creation request.
type CreateBuildingRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserLimit   int    `json:"user_limit"`
	IsActive    bool   `json:"is_active"`
}

// Validate validates the CreateBuildingRequest fields.
func (m CreateBuildingRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.Description, validation.Required, validation.Length(0, 150)),
	)
}

// UpdateBuildingRequest represents an building update request.
type UpdateBuildingRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserLimit   int    `json:"user_limit"`
	IsActive    bool   `json:"is_active"`
}

// Validate validates the CreateBuildingRequest fields.
func (m UpdateBuildingRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 50)),
		validation.Field(&m.Description, validation.Required, validation.Length(0, 150)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new building service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the building with the specified the building ID.
func (s service) Get(ctx context.Context, id string) (Building, error) {
	building, err := s.repo.Get(ctx, id)
	if err != nil {
		return Building{}, err
	}
	return Building{building}, nil
}

// Create creates a new building.
func (s service) Create(ctx context.Context, req CreateBuildingRequest) (Building, error) {
	if err := req.Validate(); err != nil {
		return Building{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Building{
		ID:          id,
		Name:        strings.TrimSpace(req.Name),
		Description: &req.Description,
		UserLimit:   req.UserLimit,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   &now,
	})
	if err != nil {
		return Building{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the building with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateBuildingRequest) (Building, error) {
	if err := req.Validate(); err != nil {
		return Building{}, err
	}

	building, err := s.Get(ctx, id)
	if err != nil {
		return building, err
	}

	now := time.Now()

	building.Name = strings.TrimSpace(req.Name)
	building.Description = &req.Description
	building.UserLimit = req.UserLimit
	building.IsActive = req.IsActive
	building.UpdatedAt = &now

	if err := s.repo.Update(ctx, building.Building); err != nil {
		return building, err
	}
	return building, nil
}

// Delete deletes the building with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Building, error) {
	building, err := s.Get(ctx, id)
	if err != nil {
		return Building{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Building{}, err
	}
	return building, nil
}

// Count returns the number of buildings.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the buildings with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Building, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Building{}
	for _, item := range items {
		result = append(result, Building{item})
	}
	return result, nil
}
