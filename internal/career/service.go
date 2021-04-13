package career

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for careers.
type Service interface {
	Get(ctx context.Context, id string) (Career, error)
	Query(ctx context.Context, offset, limit int) ([]Career, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateCareerRequest) (Career, error)
	Update(ctx context.Context, id string, input UpdateCareerRequest) (Career, error)
	// Delete(ctx context.Context, id string) (Career, error)
}

// Career represents the data about an career.
type Career struct {
	entity.Career
}

// CreateCareerRequest represents an career creation request.
type CreateCareerRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

// Validate validates the CreateCareerRequest fields.
func (m CreateCareerRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateCareerRequest represents an career update request.
type UpdateCareerRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

// Validate validates the CreateCareerRequest fields.
func (m UpdateCareerRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new career service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the career with the specified the career ID.
func (s service) Get(ctx context.Context, id string) (Career, error) {
	career, err := s.repo.Get(ctx, id)
	if err != nil {
		return Career{}, err
	}
	return Career{career}, nil
}

// Create creates a new career.
func (s service) Create(ctx context.Context, req CreateCareerRequest) (Career, error) {
	if err := req.Validate(); err != nil {
		return Career{}, err
	}
	id := entity.GenerateID()
	err := s.repo.Create(ctx, entity.Career{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return Career{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the career with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateCareerRequest) (Career, error) {
	if err := req.Validate(); err != nil {
		return Career{}, err
	}

	career, err := s.Get(ctx, id)
	if err != nil {
		return career, err
	}
	career.Name = req.Name
	career.IsActive = req.IsActive

	if err := s.repo.Update(ctx, career.Career); err != nil {
		return career, err
	}
	return career, nil
}

// Delete deletes the career with the specified ID.
// func (s service) Delete(ctx context.Context, id string) (Career, error) {
// 	career, err := s.Get(ctx, id)
// 	if err != nil {
// 		return Career{}, err
// 	}
// 	if err = s.repo.Delete(ctx, id); err != nil {
// 		return Career{}, err
// 	}
// 	return career, nil
// }

// Count returns the number of careers.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the careers with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Career, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Career{}
	for _, item := range items {
		result = append(result, Career{item})
	}
	return result, nil
}
