package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_projecthandler "github.com/nawafilhusnul/me-dashboard-api/src/project/delivery/http"
	_projectrepo "github.com/nawafilhusnul/me-dashboard-api/src/project/repository/firestore"
	_projectusecase "github.com/nawafilhusnul/me-dashboard-api/src/project/usecase"
	utilsfirebase "github.com/nawafilhusnul/me-dashboard-api/utils/firebase"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

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
	err := godotenv.Load(".env")
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
	fs := utilsfirebase.FirestoreClient(context.Background())
	pr := _projectrepo.NewFirestoreProjectRepository(fs)

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
