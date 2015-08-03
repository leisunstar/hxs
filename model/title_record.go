package model

type TitleRecord struct {
	Id          int
	ExamId      int    `orm:"column(exam_id)"`
	ExamTitleId int    `orm:"column(exam_title_id)"`
	Answer      string `orm:"column(answer)"`
	Right       bool   `orm:"column(right)"`
}
