package buildingaccess

import (
	"encoding/json"
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
	r.Get("/building-accesses/<id>", res.get)
	r.Get("/building-accesses", res.query)
	r.Post("/building-accesses", res.create)
	r.Put("/building-accesses/<id>", res.update)
	r.Delete("/building-accesses/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	buildingAccess, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(buildingAccess)
}

func (r resource) query(c *routing.Context) error {
	term := c.Query("term")
	filters := make(map[string]interface{})

	// convert JSON string filters to map
	_ = json.Unmarshal([]byte(c.Query("filters")), &filters)

	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	buildingAccesses, err := r.service.Query(ctx, pages.Offset(), pages.Limit(), term, filters)
	if err != nil {
		return err
	}
	pages.Items = buildingAccesses
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateBuildingAccessRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	buildingAccess, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(buildingAccess, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateBuildingAccessRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	buildingAccess, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(buildingAccess)
}

func (r resource) delete(c *routing.Context) error {
	buildingAccess, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(buildingAccess)
}
