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
	r.PATCH("/updateProforma/:id", proformaService.UpdateProformaService)
	r.POST("/createProforma/essentials", proformaService.CreateEssentialsProformaService)
	r.POST("/createProforma/addProduct/:proformaId", proformaService.AddProductProformaService)
	r.POST("/createProforma/uploadDocuments/:proformaId", proformaService.UploadDocumentProformaService)
	r.POST("/createProforma/finance/:proformaId", proformaService.FinanceProformaService)
	r.DELETE("/removeProforma/:id", proformaService.RemoveProformaService)
}
