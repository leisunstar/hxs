package service
import (
    "hxs/model"
    "github.com/yzw/logs"
)

type App struct {}

func (this *App) GetTitleId(blankNum, choiceNum, judgeNum int)(blankList, choiceList, judgeList []*model.ExamTitle, err error){


    return
}

func (this *App) GetAllTitles()(titleList []*model.ExamTitle, err error){
    titleList, err = examTitleDao.GetList()
    if err != nil {
        logs.Error("examTitleDao.GetList err (%v)", err)
    }
    return
}