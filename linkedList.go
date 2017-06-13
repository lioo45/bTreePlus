package bTreePlus

import "errors"

type linkedList struct {
	head *ListNode
	tail *ListNode
	len int
}

type ListNode struct {
	next *ListNode
	value interface{}

}

func (this *linkedList)Head()*ListNode{
	return this.head
}

func (this *linkedList)Tail()*ListNode{
	return this.tail
}



func (this *ListNode)Next()*ListNode{
	return this.next
}


func NewLinkedList()*linkedList{
	l:=new(linkedList)
	l.tail=new(ListNode)
	l.head=l.tail
	return l
}
func (this *linkedList)Len()int{
	return this.len
}


func (l *linkedList)Push(value interface{}){
	node:=new(ListNode)
	node.value=value
	l.tail.next=node
	if l.head==l.tail{
		l.head.next=node
	}
	l.tail=node
	l.len++
}

//插入元素使得
//第index个元素为 value index从0计数
//如果index<0则插入链头
//如果index超过链表的最大长度 则插入链尾
func (this *linkedList)Insert(value interface{},index int){
	if index<=0{
		next:=this.head.next
		this.head.next=newListNode(value,next)
		if this.head==this.tail{
			this.tail=this.head.next
		}
	} else if index>=this.len{
		this.tail.next=newListNode(value,nil)
		this.tail=this.tail.next
	} else{
		node:=this.getIndexNode(index-1)
		n:=newListNode(value,node.next)
		node.next=n
	}
	this.len++
}


func (this *linkedList)GetIndexValue(index int)interface{}{
	return this.getIndexNode(index).value
}

//返回链表第index个元素的值,index从0开始计算
//inedx如果不合法返回nil
func (this *linkedList)getIndexNode(index int)*ListNode{
	if index<0{
		return nil
	}
	if index>=this.len{
		return this.tail
	}

	node:=this.head
	for i:=0;i<=index;i++{
		node=node.next
	}
	return node
}


func newListNode(value interface{},next *ListNode)*ListNode{
	n:=new(ListNode)
	n.value=value
	n.next=next
	return n
}


func (this *linkedList)SubList(start,end int)*linkedList{

	if start>end{
		return nil
	}
	if start<0{
		start=0
	}
	if end>=this.len{
		end=this.len-1
	}

	list:=NewLinkedList()
	startNode:=this.head
	for i:=0;i<=start ;i++  {
		startNode=startNode.next
	}

	for node:=startNode;start<=end ;start++  {
		list.Push(node.value)
		node=node.next
	}
	return list
}

func (this *linkedList)Fission(fissionPos int)(l1 ,l2 *linkedList){
	if fissionPos<0||fissionPos>=this.len{
		return nil,nil
	}

	return this.SubList(0,fissionPos),this.SubList(fissionPos+1,this.len-1)
}

//index<=0删除链头
//index>=链表长度 抛出异常
func (this *linkedList)Remove(index int)*ListNode{
	if index>=this.len{
		panic(errors.New("index>=链表长度"))
	}
	pre:=this.head
	node:=pre.next
	for i:=0;i<index&&node!=nil ;i++  {
		pre=node
		node=node.next
	}
	pre.next=node.next
	if pre.next==nil{
		this.tail=pre
	}
	if pre==this.head{
		this.tail=this.head
	}
	this.len--
	return node
}

func (this *linkedList)Relace(value interface{},index int){
	node:=this.getIndexNode(index)
	node.value=value
}
