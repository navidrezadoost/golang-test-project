package main

import (
	"commerce/db"
	"commerce/handlers/config"
	"commerce/handlers/validations"
	"commerce/modules"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	cfg := config.GetConfig()
	clint, err := db.SetupMongoDB(cfg)
	if err != nil {
		log.Fatalf("Error setting up MongoDB: %v", err)
	}
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("uuid", validations.IsUUID, true)
		if err != nil {
			return
		}
	}
	modules.SetupServer(clint, cfg)
}
