package service

import (
	"errors"
	"github.com/yzw/logs"
	"hxs/cache"
	"hxs/common"
	"hxs/model"
	"strings"
)

type App struct{}

func (this *App) GetTitleId(blankNum, choiceNum, judgeNum int) (blankList, choiceList, judgeList []*model.ExamTitle) {
	blankList = common.RandList(blankNum, cache.BlankList)
	choiceList = common.RandList(choiceNum, cache.ChoiceList)
	judgeList = common.RandList(judgeNum, cache.JudgeList)
	return
}

func (this *App) GetAllTitles() (titleList []*model.ExamTitle, err error) {
	titleList, err = examTitleDao.GetList()
	if err != nil {
		logs.Error("examTitleDao.GetList err (%v)", err)
	}
	return
}

func (this *App) SaveTitles(titles string) (err error) {
	titleStrList := strings.Split(titles, "\n")
	if len(titleStrList) == 0 {
		err = errors.New("传点东西上来")
		return
	}
	for _, v := range titleStrList {
		tmpList := strings.Split(v, "|")
		tmpListLen := len(tmpList)
		if tmpListLen < 3 {
			logs.Warning("%s, 少于三个", v)
			continue
		}
		t := &model.ExamTitle{}
		switch tmpList[0] {
		case "选择":
			if tmpListLen != 4 {
				logs.Warning("%s, no 4", v)
				continue
			}
			t.ExamType = model.TitleModelChoice
			t.Answer = tmpList[3]
			t.Content = tmpList[2]
		case "填空":
			if tmpListLen != 3 {
				continue
			}
			t.ExamType = model.TitleModelBlank
			t.Answer = tmpList[2]
		case "判断":
			if tmpListLen != 3 {
				continue
			}
			t.ExamType = model.TitleModelJudge
			t.Answer = tmpList[2]
		default:
			continue
		}
		t.Title = tmpList[1]
		_, err = examTitleDao.AddExamTitle(t)
		if err != nil {
			logs.Error("examTitleDao.AddExamTitle err(%v)", err)
			continue
		}
	}
	return nil
}
