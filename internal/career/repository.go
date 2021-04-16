package career

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access careers from the data source.
type Repository interface {
	// Get returns the career with the specified career ID.
	Get(ctx context.Context, id string) (entity.Career, error)
	// Count returns the number of careers.
	Count(ctx context.Context) (int, error)
	// Query returns the list of careers with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Career, error)
	// Create saves a new career in the storage.
	Create(ctx context.Context, career entity.Career) error
	// Update updates the career with given ID in the storage.
	Update(ctx context.Context, career entity.Career) error
	// Delete removes the career with given ID from the storage.
	// Delete(ctx context.Context, id string) error
}

// repository persists careers in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new career repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the career with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Career, error) {
	var career entity.Career
	err := r.db.With(ctx).Select().Model(id, &career)
	return career, err
}

// Create saves a new career record in the database.
// It returns the ID of the newly inserted career record.
func (r repository) Create(ctx context.Context, career entity.Career) error {
	return r.db.With(ctx).Model(&career).Insert()
}

// Update saves the changes to an career in the database.
func (r repository) Update(ctx context.Context, career entity.Career) error {
	return r.db.With(ctx).Model(&career).Update()
}

// Delete deletes an career with the specified ID from the database.
// func (r repository) Delete(ctx context.Context, id string) error {
// 	career, err := r.Get(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	return r.db.With(ctx).Model(&career).Delete()
// }

// Count returns the number of the career records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").Row(&count)
	return count, err
}

// Query retrieves the career records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Career, error) {
	var careers []entity.Career
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&careers)

	for i := range careers {
		r.db.With(ctx).
			Select().
			From("groups").
			Where(dbx.HashExp{"career_id": careers[i].ID}).
			All(&careers[i].Groups)
	}
	return careers, err
}
