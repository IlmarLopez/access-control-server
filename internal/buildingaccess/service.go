package buildingaccess

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for buildingAccesses.
type Service interface {
	Get(ctx context.Context, id string) (BuildingAccess, error)
	Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]BuildingAccess, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateBuildingAccessRequest) (BuildingAccess, error)
	Update(ctx context.Context, id string, input UpdateBuildingAccessRequest) (BuildingAccess, error)
	Delete(ctx context.Context, id string) (BuildingAccess, error)
}

// BuildingAccess represents the data about an buildingAccess.
type BuildingAccess struct {
	entity.BuildingAccess
}

// CreateBuildingAccessRequest represents an buildingAccess creation request.
type CreateBuildingAccessRequest struct {
	BuildingID string    `json:"building_id"`
	UserID     string    `json:"user_id"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
}

// Validate validates the CreateBuildingAccessRequest fields.
func (m CreateBuildingAccessRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.BuildingID, validation.Required, validation.Length(36, 36)),
		validation.Field(&m.UserID, validation.Required, validation.Length(36, 36)),
		validation.Field(&m.CheckIn, validation.Required),
	)
}

// UpdateBuildingAccessRequest represents an buildingAccess update request.
type UpdateBuildingAccessRequest struct {
	BuildingID string    `json:"building_id"`
	UserID     string    `json:"user_id"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
}

// Validate validates the CreateBuildingAccessRequest fields.
func (m UpdateBuildingAccessRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.BuildingID, validation.Required, validation.Length(36, 36)),
		validation.Field(&m.UserID, validation.Required, validation.Length(36, 36)),
		validation.Field(&m.CheckIn, validation.Required),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new buildingAccess service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the buildingAccess with the specified the buildingAccess ID.
func (s service) Get(ctx context.Context, id string) (BuildingAccess, error) {
	buildingAccess, err := s.repo.Get(ctx, id)
	if err != nil {
		return BuildingAccess{}, err
	}
	return BuildingAccess{buildingAccess}, nil
}

// Create creates a new buildingAccess.
func (s service) Create(ctx context.Context, req CreateBuildingAccessRequest) (BuildingAccess, error) {
	if err := req.Validate(); err != nil {
		return BuildingAccess{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.BuildingAccess{
		ID:         id,
		BuildingID: req.BuildingID,
		UserID:     req.UserID,
		CheckIn:    req.CheckIn,
		CreatedAt:  now,
	})
	if err != nil {
		return BuildingAccess{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the buildingAccess with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateBuildingAccessRequest) (BuildingAccess, error) {
	if err := req.Validate(); err != nil {
		return BuildingAccess{}, err
	}

	buildingAccess, err := s.Get(ctx, id)
	if err != nil {
		return buildingAccess, err
	}

	now := time.Now()

	buildingAccess.CheckIn = req.CheckIn
	buildingAccess.CheckOut = &req.CheckOut
	buildingAccess.UpdatedAt = &now

	if err := s.repo.Update(ctx, buildingAccess.BuildingAccess); err != nil {
		return buildingAccess, err
	}
	return buildingAccess, nil
}

// Delete deletes the buildingAccess with the specified ID.
func (s service) Delete(ctx context.Context, id string) (BuildingAccess, error) {
	buildingAccess, err := s.Get(ctx, id)
	if err != nil {
		return BuildingAccess{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return BuildingAccess{}, err
	}
	return buildingAccess, nil
}

// Count returns the number of buildingAccesses.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the buildingAccesses with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]BuildingAccess, error) {
	items, err := s.repo.Query(ctx, offset, limit, term, filters)
	if err != nil {
		return nil, err
	}
	result := []BuildingAccess{}
	for _, item := range items {
		result = append(result, BuildingAccess{item})
	}
	return result, nil
}
