package main

import (
	"fmt"
	"log"

	"github.com/ThyMakra/gin-boilerplate/backend/models"
	"github.com/ThyMakra/gin-boilerplate/backend/pkg"
	"github.com/ThyMakra/gin-boilerplate/backend/routes"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := setupDatabase()
	app := setupApp()

	routes.NewUserRoute(db, app)

	err := app.Run(fmt.Sprintf(":%s", pkg.GetEnv("APP_PORT")))

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func setupDatabase() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(pkg.GetEnv("DATABASE_URL")),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		},
	)

	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	err = db.AutoMigrate(
		&models.UserModel{},
	)

	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	return db
}

func setupApp() *gin.Engine {
	app := gin.Default()

	if pkg.GetEnv("GO_ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app.Use(helmet.Default())
	app.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true, // Add restrictions e.g. AllowedOrigins: []string{"foo.com"}
			AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
			AllowHeaders:     []string{"Content-Type", "Authorization", "Accept-Encoding"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Length"},
		},
	))

	return app
}
