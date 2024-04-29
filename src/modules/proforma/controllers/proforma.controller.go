package controllers

import (
	"commerce/handlers/config"
	"commerce/modules/proforma/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupProformaRoutes(r *gin.RouterGroup, cfg *config.Config, client *mongo.Client) {
	collection := client.Database(cfg.MONGODB.DATABASE_NAME).Collection("proforma")

	// Create a new ProformaService instance with the initialized mongo.Collection
	proformaService := services.ProformaServiceClass(collection)
	r.GET("/findAllProformas", proformaService.FindAllProformasService)
	r.GET("/findByIdProforma/:id", proformaService.FindByIdProformaService)
	r.POST("/essentials/:id", proformaService.UpdateEssentialsProformaService)
	r.POST("/essentials", proformaService.InitEssentialsProformaService)
	r.POST("/addProduct/:id", proformaService.AddProductProformaService)
	r.POST("/uploadDocuments/:id", proformaService.UploadDocumentProformaService)
	r.POST("/finance/:id", proformaService.FinanceProformaService)
	r.DELETE("/removeProforma/:id", proformaService.RemoveProformaService)
}
