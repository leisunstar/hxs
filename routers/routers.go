package routers

import (
	"github.com/go-martini/martini"
	"hxs/controller"
)

func Router(m *martini.ClassicMartini) {
	m.Get("/", controller.Index)
    m.Get("/upload_records", controller.UploadRecords)
    m.Post("/upload_records", controller.UploadRecords)
}
