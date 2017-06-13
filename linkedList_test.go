package bTreePlus

import (
	"testing"
	"fmt"
)

func TestLinkedList(t *testing.T) {
	testNum:=10
	list:=NewLinkedList()
	for i:=0;i<testNum;i++{
		list.Push(i)
	}
	delete:=5
	v:=list.Remove(delete)
	if v.value!=delete{
		t.Error("删除的元素应该是3,但实际是",v.value)
	}


	list.Insert(delete*2,delete)
	list.Relace(5,5)
	//test(list,t)

	for i,node:=0,list.head.next;i<list.Len()&&node!=nil ;i++  {
		fmt.Println(i,":",node.value)
		node=node.next
	}




}


func test(list *linkedList,t *testing.T){

	for i:=0;i<list.Len();i++{
		if list.GetIndexValue(i)!=i{
			t.Error(fmt.Sprintf("第 %d个元素应该是%d,但是实际是%d",i,i,list.GetIndexValue(i)))
		}
	}
}
