package role

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access roles from the data source.
type Repository interface {
	// Get returns the role with the specified role ID.
	Get(ctx context.Context, id string) (entity.Role, error)
	// Count returns the number of roles.
	Count(ctx context.Context) (int, error)
	// Query returns the list of roles with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Role, error)
	// Create saves a new role in the storage.
	Create(ctx context.Context, role entity.Role) error
	// Update updates the role with given ID in the storage.
	Update(ctx context.Context, role entity.Role) error
	// Delete removes the role with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists roles in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new role repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the role with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Role, error) {
	var role entity.Role
	err := r.db.With(ctx).Select().Model(id, &role)
	return role, err
}

// Create saves a new role record in the database.
// It returns the ID of the newly inserted role record.
func (r repository) Create(ctx context.Context, role entity.Role) error {
	return r.db.With(ctx).Model(&role).Insert()
}

// Update saves the changes to an role in the database.
func (r repository) Update(ctx context.Context, role entity.Role) error {
	return r.db.With(ctx).Model(&role).Update()
}

// Delete deletes an role with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	role, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&role).Delete()
}

// Count returns the number of the role records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").Row(&count)
	return count, err
}

// Query retrieves the role records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&roles)
	return roles, err
}
