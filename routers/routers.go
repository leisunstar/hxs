package routers

import (
	"github.com/go-martini/martini"
	"hxs/controller"
)

func Router(m *martini.ClassicMartini) {
	m.Get("/", controller.Index)
	m.Get("/upload_records", controller.UploadRecords)
	m.Post("/upload_records", controller.UploadRecords)
	m.Get("/choose/exam", controller.ChooseExam)
	//答题
	m.Get("/answer", controller.Answer)
	m.Post("/answer", controller.Answer)
}
