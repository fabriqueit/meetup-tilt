package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"backend/controllers"
	"backend/docs"
	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func defaultRouter() http.Handler {
	router := gin.New()

	// Add middlewares to the router
	router.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Register routes
	controllers.RegisterRoutes(router)

	return router
}

// @title           Backend API
// @version         1.0
// @description     This is the backend api for a meetup.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /

func main() {

	gin.SetMode(gin.ReleaseMode)

	pgHost := GetEnvironmentVariable("DB_HOST", "localhost")
	pgUser := GetEnvironmentVariable("DB_USER", "admin")
	pgPassword := GetEnvironmentVariable("DB_PASSWORD", "admin")
	dbName := GetEnvironmentVariable("DB_NAME", "backend")
	pgPort := GetEnvironmentVariable("DB_PORT", "5432")
	pgSslMode := GetEnvironmentVariable("DB_SSL_MODE", "disable")

	dsn := "host=" + pgHost + " user=" + pgUser + " password=" + pgPassword + " dbname=" + dbName + " port=" + pgPort + " sslmode=" + pgSslMode + " TimeZone=Europe/Paris"
	models.ConnectDatabase(dsn, dbName, pgUser)

	httpPort := GetEnvironmentVariable("HTTP_PORT", "8080")
	httpAddress := GetEnvironmentVariable("HTTP_ADDRESS", "0.0.0.0")
	defaultRouter := &http.Server{
		Addr:         httpAddress + ":" + httpPort,
		Handler:      defaultRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initializing each server in a goroutine so they won't block the graceful shutdown handling below
	go func() {
		if err := defaultRouter.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Printf("HTTP server is running on http://" + httpAddress + ":" + httpPort)
	}()

	// wait for termination signal and register database & http server clean-up operations
	wait := utils.GracefulShutdown(context.Background(), 5*time.Second, map[string]utils.Operation{
		"http-server": func(ctx context.Context) error {
			return defaultRouter.Shutdown(ctx)
		},
		"database": func(ctx context.Context) error {
			return models.Shutdown()
		},
		// Add other cleanup operations here
	})

	<-wait
}
