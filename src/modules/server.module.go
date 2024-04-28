package modules

import (
	"commerce/handlers/config"
	"commerce/handlers/console"
	"commerce/handlers/middlewares"
	"commerce/modules/proforma/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func SetupServer(db *mongo.Client, cfg *config.Config) {
	app := gin.New()
	app.Use(middlewares.ServerMiddleware())
	setupRoutes(app, cfg, db)
	console.UpServer(cfg)
	err := app.Run(fmt.Sprintf(":%d", cfg.SERVER.PORT))
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func setupRoutes(app *gin.Engine, cfg *config.Config, db *mongo.Client) {
	controllers.SetupProformaRoutes(app.Group(cfg.SERVER.CONTROLLER_PREFIX).Group("/proforma"), cfg, db)
}
