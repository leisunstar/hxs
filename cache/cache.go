package cache

import (
	"github.com/yzw/logs"
	"hxs/dao"
	"hxs/model"
)

var (
	ExamTitleList []*model.ExamTitle
	examTitleDao  = &dao.ExamTitle{}
	ChoiceList    []*model.ExamTitle
	BlankList     []*model.ExamTitle
	JudgeList     []*model.ExamTitle
)

func Init() {
	var err error
	err = ReloadExamTitles()
	if err != nil {
		panic(err)
	}
}

func ReloadExamTitles() (err error) {
	examTitleList, err := examTitleDao.GetList()
	if err != nil {
		logs.Error("examTitleDao.GetList err(%v)", err)
		return
	}
	ChoiceList = make([]*model.ExamTitle, 0)
	BlankList = make([]*model.ExamTitle, 0)
	JudgeList = make([]*model.ExamTitle, 0)
	for _, v := range examTitleList {
		switch v.ExamType {
		case model.TitleModelJudge:
			JudgeList = append(JudgeList, v)
		case model.TitleModelBlank:
			BlankList = append(BlankList, v)
		case model.TitleModelChoice:
			ChoiceList = append(ChoiceList, v)
		}
	}
	return
}
