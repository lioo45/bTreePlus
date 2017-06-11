package bTreePlus

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
	if index<0||index>=this.len{
		return nil
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
	return this.SubList(0,fissionPos),this.SubList(fissionPos+1,this.len-1)
}



func (this *linkedList)FissionList(fissionPos int)(l1 ,l2 *linkedList){
	if this==nil||fissionPos<0||fissionPos>=this.Len(){
		return nil,nil
	}
	e:=this.head
	for i:=0;i<=fissionPos;i++{
		e=e.next
	}
	en:=e.next

	e.next=nil
	l1=NewLinkedList()
	l1.len=fissionPos+1
	l1.head=this.head
	l1.tail=e

	l2=NewLinkedList()
	l2.len=this.len-l1.len
	l2.head.next=en
	for ;en.next!=nil;{
		en=en.next
	}
	l2.tail=en

	return
}