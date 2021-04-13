package career

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/careers/<id>", res.get)
	r.Get("/careers", res.query)
	r.Post("/careers", res.create)
	r.Put("/careers/<id>", res.update)
	// r.Delete("/careers/<id>", res.delete)

}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	career, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(career)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	careers, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = careers
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateCareerRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	career, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(career, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateCareerRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	career, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(career)
}

// func (r resource) delete(c *routing.Context) error {
// 	career, err := r.service.Delete(c.Request.Context(), c.Param("id"))
// 	if err != nil {
// 		return err
// 	}

// 	return c.Write(career)
// }
