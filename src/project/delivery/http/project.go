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

func NewProjectHandler(g *gin.Engine, ps domain.ProjectUsecase) {
	h := &ProjectHandler{
		PUsecase: ps,
	}
	v1 := g.Group("v1")
	{
		v1.GET("/projects", h.Find)
		v1.POST("/projects", h.Create)
		v1.PUT("/projects/:id", h.Update)
		v1.DELETE("/projects/:id", h.Delete)
	}
}

// @Description create a new project
// @Accept  json
// @Produce  json
// @Tags	Projects
// @Param	body	body		domain.ProjectRequest	true	"Project detail"
// @Success 200 {object} domain.MetaResponse
// @Failure 400 {object} domain.MetaResponse
// @Failure 404 {object} domain.MetaResponse
// @Router /projects [post]
func (a *ProjectHandler) Create(c *gin.Context) {
	var req domain.ProjectRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.MetaResponse{Message: err.Error()})
		return
	}

	p := new(domain.Project)
	p.BindFromReq(req)

	ctx := c.Request.Context()
	err = a.PUsecase.Create(ctx, p)
	if err != nil {
		c.JSON(global.GetStatusCode(err), domain.MetaResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.MetaResponse{Message: domain.POSTSuccess})
}

// @Description update project detail by ID
// @Accept  json
// @Produce  json
// @Tags	Projects
// @Param   id     path    string     true        "Project ID"
// @Param	body	body		domain.ProjectRequest	true	"Project detail"
// @Success 200 {object} domain.MetaResponse
// @Failure 400 {object} domain.MetaResponse
// @Failure 404 {object} domain.MetaResponse
// @Router /projects/{id} [put]
func (a *ProjectHandler) Update(c *gin.Context) {
	var req domain.ProjectRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.MetaResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.MetaResponse{Message: err.Error()})
		return
	}

	p := new(domain.Project)
	p.BindFromReq(req)
	p.ID = c.Param("id")

	ctx := c.Request.Context()
	err = a.PUsecase.Update(ctx, p)
	if err != nil {
		c.JSON(global.GetStatusCode(err), domain.MetaResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.MetaResponse{Message: domain.GETSuccess})
}

// @Description delete project detail by ID
// @Accept  json
// @Produce  json
// @Tags	Projects
// @Param   id     path    string     true        "Project ID"
// @Success 200 {object} domain.MetaResponse
// @Failure 400 {object} domain.MetaResponse
// @Failure 404 {object} domain.MetaResponse
// @Router /projects/{id} [delete]
func (a *ProjectHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	err := a.PUsecase.Delete(ctx, id)
	if err != nil {
		c.JSON(global.GetStatusCode(err), domain.MetaResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, domain.MetaResponse{Message: domain.DELETESuccess})
}

// @Description create a new project
// @Accept  json
// @Produce  json
// @Tags	Projects
// @Success 200 {object} domain.ProjectListResponse
// @Failure 400 {object} domain.MetaResponse
// @Failure 404 {object} domain.MetaResponse
// @Router /projects/ [get]
func (a *ProjectHandler) Find(c *gin.Context) {
	ctx := c.Request.Context()
	pl, err := a.PUsecase.Find(ctx)
	if err != nil {
		c.JSON(global.GetStatusCode(err), domain.MetaResponse{Message: err.Error()})
		return
	}

	res := make([]domain.ProjectDTO, len(pl))
	for idx := range pl {
		pl[idx].BindToRes(&res[idx])
	}

	c.JSON(http.StatusOK, domain.ProjectListResponse{
		Meta: domain.MetaResponse{
			Message: domain.GETSuccess,
		},
		Data: res,
	})
}
