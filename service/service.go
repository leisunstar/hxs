package service

import (
	"hxs/dao"
)

var (
	examDao        *dao.Exam
	examTitleDao   *dao.ExamTitle
	titleRecordDao *dao.TitleRecord
)

func Init() (err error) {
	//初始化dao
	examDao = dao.NewExam()
	examTitleDao = dao.NewExamTitle()
	titleRecordDao = dao.NewTitleRecord()

	err = dao.Init()
	if err != nil {
		return
	}
	return
}
