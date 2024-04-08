package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type personData struct {
	FirstName   string `json:"firstName"  binding:"required,min=2,max=20,alpha"`
	LastName    string `json:"lastName"  binding:"required"`
	PhoneNumber string `json:"phoneNumber"  binding:"required,mobile"`
}

type TestHealth struct{}

func NewTestHealth() *TestHealth {
	return &TestHealth{}
}

func (h *TestHealth) Test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func (h *TestHealth) Users(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func (h *TestHealth) UserById(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"id": context.Params.ByName("id"),
	})
}

func (h *TestHealth) UserByUserName(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"username": context.Params.ByName("username"),
	})
}

func (h *TestHealth) AddUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"data": context.ShouldBindJSON(context.Request.Body),
	})
}

func (h *TestHealth) BindHeader(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"headerBinder1": context.GetHeader("user_id"),
		"headerBinder2": context.BindHeader(context.Request.Header),
	})
}

func (h *TestHealth) BindQuery(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"queryBinder1": context.Query("username"),
	})
}

func (h *TestHealth) BindQueryArray(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"queryBinderArray1": context.QueryArray("username"),
	})
}

func (h *TestHealth) BindUri(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"uriBinderArray1": context.Param("name"),
	})
}

func (h *TestHealth) BindBody(context *gin.Context) {
	p := personData{}
	err := context.ShouldBindJSON(&p)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"bodyBinder1": p,
	})
}

func (h *TestHealth) BindBForm(context *gin.Context) {
	p := personData{}
	err := context.Bind(&p)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"formBinder1": p,
	})
}

func (h *TestHealth) BindFile(context *gin.Context) {
	file, _ := context.FormFile("file")
	err := context.SaveUploadedFile(file, "file")
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"fileBinder1": file.Filename,
	})
}
