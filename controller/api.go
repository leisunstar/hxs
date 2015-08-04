package controller

import (
	"fmt"
	"github.com/martini-contrib/render"
	"net/http"
)

func Index(r render.Render) {

	returnData := make(map[string]interface{})
	titleList, err := appService.GetAllTitles()
	if err != nil {

	}
	returnData["list"] = titleList
	r.HTML(200, "index", returnData)
	return
}

func UploadRecords(req *http.Request, r render.Render) {
	if req.Method == "GET" {
		r.HTML(200, "upload", "")
		return
	}
	req.ParseForm()
	values := req.Form
	titles := values.Get("titles")
	returnData := make(map[string]interface{})
	err := appService.SaveTitles(titles)
	if err != nil {
		returnData["err_msg"] = fmt.Sprintf("err msg %v", err)
		r.HTML(200, "result", returnData)
		return
	}
	returnData["err_msg"] = "成功"
	r.HTML(200, "result", returnData)
	return
}

func Answer(r render.Render) {

	returnData := make(map[string]interface{})
	blankList, choiceList, judgeList := appService.GetTitleId(1, 1, 1)
	returnData["blankList"] = blankList
	returnData["choiceList"] = choiceList
	returnData["judgeList"] = judgeList
	r.HTML(200, "answer", returnData)
	return
}
