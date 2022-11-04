package handler

import (
	"github.com/andresxlp/gohex/internal/app"
)

func StartGoModule(fl *app.Service) {
	fl.InitGoModule(Module)
}
