package bTreePlus


type Student struct {
	ID int//学号
	Name string //姓名
}


func (this *Student)Key()int{
	return this.ID
}

