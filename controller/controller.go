package controller

import (
	"hxs/service"
)

var (
    appService  *service.App
)

func Init() (err error) {
	//初始化自己
	//初始化service
    appService = &service.App{}

	err = service.Init()
	if err != nil {
		return
	}
	return
}
