package buildingaccess

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access buildingAccesses from the data source.
type Repository interface {
	// Get returns the buildingAccess with the specified buildingAccess ID.
	Get(ctx context.Context, id string) (entity.BuildingAccess, error)
	// Count returns the number of buildingAccesses.
	Count(ctx context.Context) (int, error)
	// Query returns the list of buildingAccesses with the given offset and limit.
	Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]entity.BuildingAccess, error)
	// Create saves a new buildingAccess in the storage.
	Create(ctx context.Context, buildingAccess entity.BuildingAccess) error
	// Update updates the buildingAccess with given ID in the storage.
	Update(ctx context.Context, buildingAccess entity.BuildingAccess) error
	// Delete removes the buildingAccess with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists buildingAccesses in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new buildingAccess repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the buildingAccess with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.BuildingAccess, error) {
	var buildingAccess entity.BuildingAccess
	err := r.db.With(ctx).Select().Model(id, &buildingAccess)

	r.db.With(ctx).
		Select().
		From("users").
		Where(dbx.HashExp{"id": buildingAccess.UserID}).
		One(&buildingAccess.User)

	return buildingAccess, err
}

// Create saves a new buildingAccess record in the database.
// It returns the ID of the newly inserted buildingAccess record.
func (r repository) Create(ctx context.Context, buildingAccess entity.BuildingAccess) error {
	return r.db.With(ctx).Model(&buildingAccess).Insert()
}

// Update saves the changes to an buildingAccess in the database.
func (r repository) Update(ctx context.Context, buildingAccess entity.BuildingAccess) error {
	return r.db.With(ctx).Model(&buildingAccess).Update()
}

// Delete deletes an buildingAccess with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	buildingAccess, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&buildingAccess).Delete()
}

// Count returns the number of the buildingAccess records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").Row(&count)
	return count, err
}

// Query retrieves the buildingAccess records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]entity.BuildingAccess, error) {
	var buildingAccesses []entity.BuildingAccess

	switch term {
	case "by-building-and-not-check-out":
		if err := r.db.With(ctx).
			Select().
			Where(dbx.HashExp{"building_id": filters["building_id"]}).
			Where(dbx.HashExp{"check_out": nil}).
			OrderBy("id").
			Offset(int64(offset)).
			Limit(int64(limit)).
			All(&buildingAccesses); err != nil {
			return buildingAccesses, err
		}
	default:
		if err := r.db.With(ctx).
			Select().
			OrderBy("id").
			Offset(int64(offset)).
			Limit(int64(limit)).
			All(&buildingAccesses); err != nil {
			return buildingAccesses, err
		}
	}

	for i := range buildingAccesses {
		r.db.With(ctx).
			Select("id", "username", "role_id", "first_name", "last_name", "created_at", "updated_at", "is_active", "(select name from roles where id = role_id) as role_name").
			From("users").
			Where(dbx.HashExp{"id": buildingAccesses[i].UserID}).
			One(&buildingAccesses[i].User)
	}

	return buildingAccesses, nil
}
