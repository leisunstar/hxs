package cache
import (
    "hxs/model"
    "hxs/dao"
)

var (
    ExamTitleList []*model.ExamTitle
    examTitleDao = &dao.ExamTitle{}
)

func init(){

}

func ReloadExamTitles(){
    examTitleList, err := examTitleDao.GetExamTitleById()
}