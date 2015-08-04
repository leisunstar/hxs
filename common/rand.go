package common

import (
	"hxs/model"
	"math/rand"
	"time"
)

func RandList(num int, l []*model.ExamTitle) (tl []*model.ExamTitle) {
	m := make(map[int]interface{})
	total := len(l)
	if total <= num {
		tl = l
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := 0
	for i < num {
		k := r.Intn(total)
		if _, ok := m[k]; !ok {
			if k >= total {
				k = total - 1
			}
			m[k] = 1
			tl = append(tl, l[k])
			i++
		}
	}

	return
}
