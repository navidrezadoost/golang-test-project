package services

import (
	"commerce/common"
	"commerce/handlers/enum"
	"commerce/handlers/utils"
	"commerce/modules/proforma/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProformaService struct {
	Collection  *mongo.Collection
	HTTPService *common.HTTPClient
}

func ProformaServiceClass(collection *mongo.Collection) *ProformaService {
	httpClient := common.NewHTTPClient()
	return &ProformaService{Collection: collection, HTTPService: httpClient}
}

func (h *ProformaService) FindAllProformasService(context *gin.Context) {
	// Define a slice to store the results
	var proformas []bson.M

	// Perform a find operation to retrieve all documents from the collection
	cursor, err := h.Collection.Find(context, bson.M{})
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.ResponseCorrect(err.Error(), http.StatusInternalServerError, "loggerID", "error", "Failed to retrieve proformas"))
		return
	}
	defer cursor.Close(context)

	// Iterate over the cursor and decode each document into the slice
	if err := cursor.All(context, &proformas); err != nil {
		context.JSON(http.StatusInternalServerError, utils.ResponseCorrect(err.Error(), http.StatusInternalServerError, "loggerID", "error", "Failed to decode proformas"))
		return
	}

	// Return the results as JSON
	context.JSON(http.StatusOK, utils.ResponseCorrect(proformas, http.StatusOK, "loggerID", "success", "Proformas retrieved successfully"))
}

func (h *ProformaService) FindByIdProformaService(context *gin.Context) {
	id := context.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid proforma ID"})
		return
	}
	// Find the inserted document by _id
	filter := bson.M{"_id": objectID} // Assuming ID is already set by MongoDB
	result := h.Collection.FindOne(context, filter)
	if result.Err() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created document"})
		return
	}
	// Return response with proformaData as JSON
	context.JSON(http.StatusOK, utils.ResponseCorrect(utils.DecodeSingleDocument(context, result), http.StatusOK, "loggerID", "success", "Proforma Essentials successfully"))
}

func (h *ProformaService) InitEssentialsProformaService(context *gin.Context) {
	var proforma entity.Proforma

	// Bind JSON body to Pro forma struct
	if err := context.BindJSON(&proforma); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default values
	proforma.CreatedAt = time.Now()
	proforma.Documents = []entity.File{}
	proforma.UpdatedAt = proforma.CreatedAt // UpdatedAt should be the same as CreatedAt initially
	proforma.Status = enum.Draft            // UpdatedAt should be the same as CreatedAt initially

	// Insert Proforma into MongoDB collection
	req, err := h.Collection.InsertOne(context, proforma)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Find the inserted document by _id
	filter := bson.M{"_id": req.InsertedID} // Assuming ID is already set by MongoDB
	result := h.Collection.FindOne(context, filter)
	if result.Err() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created document"})
		return
	}

	// Decode the document to a map
	var proformaMap map[string]interface{}
	if err := result.Decode(&proformaMap); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode document"})
		return
	}

	// Return response with proformaData as JSON
	context.JSON(http.StatusOK, utils.ResponseCorrect(proformaMap, http.StatusOK, "loggerID", "success", "Proforma Essentials successfully"))
}

func (h *ProformaService) UpdateEssentialsProformaService(context *gin.Context) {
	// Get the proforma ID from the URL parameter
	proformaID := context.Param("id")

	// Parse the proforma ID into an ObjectID
	objID, err := primitive.ObjectIDFromHex(proformaID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid proforma ID"})
		return
	}

	// Parse the request body into a map or a struct
	var updateData map[string]interface{}
	if err := context.BindJSON(&updateData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the map to BSON format
	update := bson.M{"$set": updateData}

	// Perform the update operation in MongoDB
	filter := bson.M{"_id": objID}
	_, err = h.Collection.UpdateOne(context, filter, update)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update proforma"})
		return
	}

	// Find the inserted document by _id
	result := h.Collection.FindOne(context, filter)
	if result.Err() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created document"})
		return
	}

	// Decode the document to a map
	var proformaMap map[string]interface{}
	if err := result.Decode(&proformaMap); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode document"})
		return
	}

	// Return the updated proforma in the response
	context.JSON(http.StatusOK, utils.ResponseCorrect(proformaMap, http.StatusOK, "loggerID", "success", "Proforma updated successfully"))

}

func (h *ProformaService) AddProductProformaService(context *gin.Context) {

}

func (h *ProformaService) UploadDocumentProformaService(context *gin.Context) {
	// Get proforma ID from URL parameter
	proformaID := context.Param("id")

	// Parse proforma ID into ObjectID
	err, objID := utils.ConvertToPrimitiveObjectId(context, proformaID)

	// Bind the request body to a file object
	var file entity.File
	if err := context.BindJSON(&file); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define the filter to find the proforma document
	filter := bson.M{"_id": objID}
	// Execute the update operation
	_, err = h.Collection.UpdateOne(context, filter, bson.M{"$push": bson.M{"documents": file}})
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update proforma"})
		return
	}
	// Assuming ID is already set by MongoDB
	result := h.Collection.FindOne(context, filter)
	if result.Err() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created document"})
		return
	}

	// Decode the document to a map
	var proformaMap map[string]interface{}
	if err := result.Decode(&proformaMap); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode document"})
		return
	}

	context.JSON(http.StatusOK, utils.ResponseCorrect(proformaMap, http.StatusOK, "loggerID", "success", "Document uploaded successfully"))
}

func (h *ProformaService) FinanceProformaService(context *gin.Context) {

}

func (h *ProformaService) RemoveProformaService(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"queryBinder1": context.Query("username"),
	})
}
