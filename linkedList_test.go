package bTreePlus

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T){
	l:=NewLinkedList()
	for i:=0;i<10 ;i++  {
		l.Push(i)
	}
	test2(t,l)
	//插入元素后继续测试
	l.Insert(10,10)
	test2(t,l)

	l=NewLinkedList()
	for i:=0;i<10 ;i++  {
		l.Insert(i,i)
	}
	test2(t,l)


	l=NewLinkedList()
	for i:=0;i<10 ;i++  {
		l.Insert(i+1,i)
	}
	l.Insert(0,0)
	test2(t,l)


	l=NewLinkedList()
	for i:=0;i<10 ;i++  {
		if i>=5{
			l.Insert(i+1,i)
		}else {
			l.Insert(i,i)
		}
	}
	l.Insert(5,5)
	test2(t,l)

}

func TestLinkedList_FissionList(t *testing.T) {
	l:=NewLinkedList()
	for i:=0;i<10 ;i++  {
		l.Push(i)
	}
	l1,l2:=l.FissionList(4)

	for i,e:=0,l1.head.next;i<=4;i++{
		if e.value!=i{
			t.Error(fmt.Sprintf("failure: %d-th' element should be %d,but %d\n",i,i,e.value))
		}
		e=e.Next()
	}

	for i,e:=5,l2.head.next;i<10;i++{
		if e.value!=i{
			t.Error(fmt.Sprintf("failure: %d-th' element should be %d,but %d\n",i-5,i,e.value))
		}
		e=e.Next()
	}
}

func test2(t *testing.T,l *linkedList){
	//测试 getIndexValue
	for i:=0;i<l.Len() ;i++  {
		if(l.GetIndexValue(i)!=i){
			t.Error(fmt.Sprintf("failure: l.GetINdexValue(%d) should be %d,but %d\n",i,i,l.GetIndexValue(i)))
		}
	}
	//从尾到头测试
	for i,e:=l.Len()-1,l.tail; e!=l.head;e=e.pre {
		if e.value!=i{
			t.Error(fmt.Sprintf("failure: from tail to head's %d-th element should be %d,but %d\n",l.len-i,i,e.value))
		}
		i--
	}
	//头到尾测试
	for i,e:=0,l.head.next;i<l.Len();i++{
		if e.value!=i{
			t.Error(fmt.Sprintf("failure: from head to tail's %d-th element should be %d ,but %v\n",i,i,e.value))
		}
		e=e.next
	}
}
