package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"hxs/model"
)

type ExamTitle struct {
}

func NewExamTitle() *ExamTitle {
	return &ExamTitle{}
}

func init() {
	orm.RegisterModel(new(model.ExamTitle))
}

func (this *ExamTitle) AddExamTitle(m *model.ExamTitle) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func (this *ExamTitle) GetExamTitleById(id int) (v *model.ExamTitle, err error) {
	o := orm.NewOrm()
	v = &model.ExamTitle{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (this *ExamTitle) UpdateExamTitleById(m *model.ExamTitle) (err error) {
	o := orm.NewOrm()
	v := model.ExamTitle{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func (this *ExamTitle) DeleteExamTitle(id int) (err error) {
	o := orm.NewOrm()
	v := model.ExamTitle{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&model.ExamTitle{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func (this *ExamTitle) GetList() (examTitleLists []*model.ExamTitle, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("SELECT * FROM \"exam_title\";").QueryRows(&examTitleLists)
	return
}
