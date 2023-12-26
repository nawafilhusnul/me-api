package http

import (
	"github.com/gin-gonic/gin"
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

	res, err := a.PUsecase.Find(c.Request.Context())
	if err != nil {
		c.JSON(500, "error")
	}

	c.JSON(200, res)
}
