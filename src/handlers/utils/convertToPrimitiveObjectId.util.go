package utils

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func ConvertToPrimitiveObjectId(context *gin.Context, proformaID string) (error, *primitive.ObjectID) {
	objID, err := primitive.ObjectIDFromHex(proformaID)
	if err != nil {
		context.JSON(http.StatusOK, ResponseInCorrect(proformaID, http.StatusBadRequest, "loggerID", "error", "Invalid proforma ID", "param"))
		return err, nil
	}
	return err, &objID
}
