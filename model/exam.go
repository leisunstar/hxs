package model

type Exam struct {
	Id    int
	Name  string `orm:"column(name)"`
	Score int    `orm:"column(score)"`
	CTime int64  `orm:"column(ctime)"`
}
