package bTreePlus

import "strconv"

type student struct {
	id int//学号
	name string //姓名
}


func (this *student)Key()string{
	return strconv.Itoa(this.id)
}

