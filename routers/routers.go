package routers

import (
	"github.com/go-martini/martini"
	"hxs/controller"
)

func Router(m *martini.ClassicMartini) {
	m.Get("/", controller.Index)
    m.Get("/upload_record", controller.Upload)
    m.Post("/upload_record", controller.Upload)
}
