package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"hxs/model"
)

type TitleRecord struct {
}

func NewTitleRecord() *TitleRecord {
	return &TitleRecord{}
}

func init() {
	orm.RegisterModel(new(model.TitleRecord))
}

func (this *TitleRecord) AddTitleRecord(m *model.TitleRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func (this *TitleRecord) GetTitleRecordById(id int) (v *model.TitleRecord, err error) {
	o := orm.NewOrm()
	v = &model.TitleRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (this *TitleRecord) UpdateTitleRecordById(m *model.TitleRecord) (err error) {
	o := orm.NewOrm()
	v := model.TitleRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func (this *TitleRecord) DeleteTitleRecord(id int) (err error) {
	o := orm.NewOrm()
	v := model.TitleRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&model.TitleRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
