package console

import (
	"commerce/handlers/config"
	"github.com/mattn/go-colorable"
	"log"
)

func UpServer(cfg *config.Config) {
	logger := log.New(colorable.NewColorableStdout(), "", log.LstdFlags)
	logger.Printf("\033[1;34mThe server is running successfully and opens to port %v\033[0m\n", cfg.SERVER.PORT)
}
