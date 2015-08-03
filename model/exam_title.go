package model

type ExamTitle struct {
	Id       int
	ExamType int    `orm:"column(exam_type)"`
	Title    string `orm:"column(title)"`
	Content  string `orm:"column(content)"`
	Answer   string `orm:"column(answer)"`
}
