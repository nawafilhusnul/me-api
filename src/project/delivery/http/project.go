package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nawafilhusnul/me-dashboard-api/global"
	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
)

type ProjectHandler struct {
	PUsecase domain.ProjectUsecase
}

type MetaResponse struct {
	Message string `json:"message"`
}

func NewProjectHandler(g *gin.Engine, ps domain.ProjectUsecase) {
	h := &ProjectHandler{
		PUsecase: ps,
	}
	v1 := g.Group("v1")
	{
		v1.POST("/projects", h.Create)
		v1.PUT("/projects/:id", h.Update)
		v1.DELETE("/projects/:id", h.Delete)
	}
}

func (a *ProjectHandler) Create(c *gin.Context) {
	var req domain.ProjectRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, MetaResponse{Message: err.Error()})
		return
	}

	p := new(domain.Project)
	p.BindFromReq(req)

	ctx := c.Request.Context()
	err = a.PUsecase.Create(ctx, p)
	if err != nil {
		c.JSON(global.GetStatusCode(err), MetaResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, MetaResponse{Message: domain.POSTSuccess})
}

// @Description update project detail by ID
// @Accept  json
// @Produce  json
// @Tags	Projects
// @Param   id     path    string     true        "Project ID"
// @Success 200 {object} MetaResponse
// @Failure 400 {object} MetaResponse
// @Failure 404 {object} MetaResponse
// @Router /projects/{id} [put]
func (a *ProjectHandler) Update(c *gin.Context) {
	var req domain.ProjectRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, MetaResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, MetaResponse{Message: err.Error()})
		return
	}

	p := new(domain.Project)
	p.BindFromReq(req)
	p.ID = c.Param("id")

	ctx := c.Request.Context()
	err = a.PUsecase.Update(ctx, p)
	if err != nil {
		c.JSON(global.GetStatusCode(err), MetaResponse{Message: err.Error()})
		return
	}
	res := new(domain.ProjectResponse)
	p.BindToRes(res)
	c.JSON(http.StatusCreated, MetaResponse{Message: domain.PUTSuccess})
}

// @Description delete project detail by ID
// @Accept  json
// @Produce  json
// @Tags	Projects
// @Param   id     path    string     true        "Project ID"
// @Success 200 {object} MetaResponse
// @Failure 400 {object} MetaResponse
// @Failure 404 {object} MetaResponse
// @Router /projects/{id} [delete]
func (a *ProjectHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	err := a.PUsecase.Delete(ctx, id)
	if err != nil {
		c.JSON(global.GetStatusCode(err), MetaResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, MetaResponse{Message: domain.DELETESuccess})
}
