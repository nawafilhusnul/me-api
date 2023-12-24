package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_projecthandler "github.com/nawafilhusnul/me-dashboard-api/src/project/delivery/http"
	_projectrepo "github.com/nawafilhusnul/me-dashboard-api/src/project/repository/mysql"
	_projectusecase "github.com/nawafilhusnul/me-dashboard-api/src/project/usecase"
	utils "github.com/nawafilhusnul/me-dashboard-api/utils/mysql"
)

//
// @title Me Dashboard API
// @version 1.0
// @description This is the APIs for dashboard used to manage me UI.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8080
// @BasePath /v1
func main() {
	pathDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load(filepath.Join(pathDir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	gormDb := utils.GormClient()
	pr := _projectrepo.NewProjectRepository(gormDb)

	timeout, err := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	if err != nil {
		log.Fatal("Error get context timeout" + err.Error())
	}
	timeoutContext := time.Duration(timeout) * time.Second
	au := _projectusecase.NewProjectUsecase(pr, timeoutContext)
	_projecthandler.NewProjectHandler(r, au)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.StaticFile("/docs", "./docs/public/index.html")
	r.StaticFile("/swagger.json", "./docs/public/swagger.json")

	r.Run()
}
