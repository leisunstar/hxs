package controller

import (
    "github.com/martini-contrib/render"
    "net/http"
)

func Index(r render.Render) {

    returnData := make(map[string]interface{})
    titleList, err := appService.GetAllTitles()
    if err != nil{

    }
    returnData["list"] = titleList
    r.HTML(200, "index", returnData)
    return
}


func UploadRecords(req *http.Request,r render.Render) {
    if req.Method == "GET"{
        r.HTML(200, "upload", "")
        return
    }
    returnData := make(map[string]interface{})
    titleList, err := appService.GetAllTitles()
    if err != nil{

    }
    returnData["list"] = titleList
    r.HTML(200, "index", returnData)
    return
}
