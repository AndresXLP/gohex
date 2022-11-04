package handler

import (
	"github.com/andresxlp/gohex/internal/app"
)

var (
	BasePath string
	Module   string
)

func StarCreate(fl *app.Service) {
	fl.Start(Module, BasePath)
}
