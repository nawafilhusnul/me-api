package main

import (
	"context"
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

func main() {
	r := gin.Default()
	fs := utilsfirebase.FirestoreClient(context.Background())
	pr := _projectrepo.NewFirestoreProjectRepository(fs)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _projectusecase.NewProjectUsecase(pr, timeoutContext)
	_projecthandler.NewProjectHandler(r, au)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
