package utils

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func DecodeSingleDocument(context *gin.Context, doc *mongo.SingleResult) map[string]interface{} {
	var proformaMap map[string]interface{}
	if err := doc.Decode(&proformaMap); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode document"})
		return nil
	}
	return proformaMap
}
