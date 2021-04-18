package building

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access buildings from the data source.
type Repository interface {
	// Get returns the building with the specified building ID.
	Get(ctx context.Context, id string) (entity.Building, error)
	// Count returns the number of buildings.
	Count(ctx context.Context) (int, error)
	// Query returns the list of buildings with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Building, error)
	// Create saves a new building in the storage.
	Create(ctx context.Context, building entity.Building) error
	// Update updates the building with given ID in the storage.
	Update(ctx context.Context, building entity.Building) error
	// Delete removes the building with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists buildings in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new building repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the building with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Building, error) {
	var building entity.Building
	err := r.db.With(ctx).Select().Model(id, &building)

	r.db.With(ctx).
		Select("COUNT(*)").
		From("building_access").
		Where(dbx.HashExp{"check_out": nil}).
		AndWhere(dbx.HashExp{"building_id": building.ID}).
		Row(&building.ActiveUsers)

	return building, err
}

// Create saves a new building record in the database.
// It returns the ID of the newly inserted building record.
func (r repository) Create(ctx context.Context, building entity.Building) error {
	return r.db.With(ctx).Model(&building).Insert()
}

// Update saves the changes to an building in the database.
func (r repository) Update(ctx context.Context, building entity.Building) error {
	return r.db.With(ctx).Model(&building).Update()
}

// Delete deletes an building with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	building, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&building).Delete()
}

// Count returns the number of the building records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").Row(&count)
	return count, err
}

// Query retrieves the building records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Building, error) {
	var buildings []entity.Building
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&buildings)

	for i := range buildings {
		r.db.With(ctx).
			Select("COUNT(*)").
			From("building_access").
			Where(dbx.HashExp{"check_out": nil}).
			AndWhere(dbx.HashExp{"building_id": buildings[i].ID}).
			Row(&buildings[i].ActiveUsers)
	}
	return buildings, err
}
