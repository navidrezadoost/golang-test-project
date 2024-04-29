package common

import (
	"commerce/handlers/config"
	"fmt"
)

type UrlService struct {
	cfg *config.Config
}

func (h *UrlService) UrlProduct(variable string) string {
	return fmt.Sprintf("%s%d%s/%s",
		h.cfg.SERVICE.ProductService.BaseURL,
		h.cfg.SERVICE.ProductService.Port, h.cfg.SERVICE.ProductService.Path, variable)
}
