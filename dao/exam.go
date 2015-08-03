package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"hxs/model"
)

func init() {
	orm.RegisterModel(new(model.Exam))
}

type Exam struct {
}

func NewExam() *Exam {
	return &Exam{}
}

func (this *Exam) AddExam(m *model.Exam) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func (this *Exam) GetExamById(id int) (v *model.Exam, err error) {
	o := orm.NewOrm()
	v = &model.Exam{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (this *Exam) UpdateExamById(m *model.Exam) (err error) {
	o := orm.NewOrm()
	v := model.Exam{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func (this *Exam) DeleteExam(id int) (err error) {
	o := orm.NewOrm()
	v := model.Exam{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&model.Exam{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
