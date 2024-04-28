package services

import (
	"commerce/handlers/utils"
	"commerce/modules/proforma/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type personData struct {
	FirstName   string `json:"firstName"  binding:"required,min=2,max=20,alpha"`
	LastName    string `json:"lastName"  binding:"required"`
	PhoneNumber string `json:"phoneNumber"  binding:"required,mobile"`
}

type ProformaService struct {
	Collection *mongo.Collection
}

func ProformaServiceClass(collection *mongo.Collection) *ProformaService {
	return &ProformaService{Collection: collection}
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
	// Get the proforma ID from the request params
	id := context.Param("id")

	// Perform a find operation to retrieve the proforma by its ID
	var proforma bson.M
	if err := h.Collection.FindOne(context, bson.M{"_id": id}).Decode(&proforma); err != nil {
		if err == mongo.ErrNoDocuments {
			context.JSON(http.StatusNotFound, gin.H{"error": "Proforma not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the proforma as JSON
	context.JSON(http.StatusOK, proforma)
}

func (h *ProformaService) UpdateProformaService(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func (h *ProformaService) CreateEssentialsProformaService(context *gin.Context) {
	var proforma entity.Proforma

	// Bind JSON body to Proforma struct
	if err := context.BindJSON(&proforma); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default values
	proforma.CreatedAt = time.Now()
	proforma.UpdatedAt = time.Now()

	// Insert Proforma into MongoDB collection
	doc, err := h.Collection.InsertOne(context, proforma)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Find the inserted document by _id
	filter := bson.M{"_id": doc.InsertedID} // Assuming ID is the unique identifier of the document
	cursor, err := h.Collection.Find(context, filter)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created document"})
		return
	}

	// Extract the first document from the cursor
	if cursor.Next(context) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode document"})
			return
		}

		// Return response with proformaData
		context.JSON(http.StatusOK, utils.ResponseCorrect(result, http.StatusOK, "loggerID", "success", "Proforma Essentials successfully"))
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "No document found"})
	}

	cursor.Close(context)
}

func (h *ProformaService) AddProductProformaService(context *gin.Context) {

}

func (h *ProformaService) UploadDocumentProformaService(context *gin.Context) {

}

func (h *ProformaService) FinanceProformaService(context *gin.Context) {

}

func (h *ProformaService) RemoveProformaService(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"queryBinder1": context.Query("username"),
	})
}

func (h *ProformaService) BindQueryArray(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"queryBinderArray1": context.QueryArray("username"),
	})
}

func (h *ProformaService) BindUri(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"uriBinderArray1": context.Param("name"),
	})
}

func (h *ProformaService) BindBForm(context *gin.Context) {
	p := personData{}
	err := context.Bind(&p)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"formBinder1": p,
	})
}

func (h *ProformaService) BindFile(context *gin.Context) {
	file, _ := context.FormFile("file")
	err := context.SaveUploadedFile(file, "file")
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"fileBinder1": file.Filename,
	})
}
