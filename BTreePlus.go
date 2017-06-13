package bTreePlus

import (
	"unsafe"
	"log"
	"os"
)

type BTreeValue interface {
	Key()int
}

var log1 *log.Logger
func init(){
	logFile,_:=os.OpenFile("/Users/pro/log/btreeLog.txt",os.O_CREATE|os.O_RDWR,0666)
	log1=log.New(logFile,"",log.LstdFlags|log.Lshortfile)
}
type BTreePlus struct{
	root *TreeNode
	degree int
	leafHead *leafNode
}

//如果当前节点的叶子节点 那么children指向key对应的value
type TreeNode struct {
	isLeaf bool
	parent *TreeNode
	keys *linkedList
	children *linkedList
}

type leafNode struct {
	TreeNode
	pre *leafNode
	next *leafNode
}

//初始化B+树
func New(degree int)*BTreePlus{
	b:=new(BTreePlus)
	b.degree=degree
	return b
}

//添加第一个值
func (this *BTreePlus)init(value BTreeValue){
	root:=newTreeNode(nil)
	root.keys.Push(value.Key())
	root.children.Push(value)
	root.isLeaf=true

	this.leafHead=&leafNode{TreeNode{},nil,nil}
	ln:=&leafNode{*root,this.leafHead,nil}
	this.leafHead.next=ln

	this.root=(*TreeNode)(unsafe.Pointer(ln))

}

//初始化树节点
func newTreeNode(p *TreeNode)*TreeNode{
	t:=new(TreeNode)
	t.parent=p
	t.keys=NewLinkedList()
	t.children=NewLinkedList()
	t.isLeaf=false
	return t
}

//返回B+树的阶数
func (this *BTreePlus)Degree()int{
	return this.degree
}

//插入,需要考虑的情况
func (this *BTreePlus)Insert(value BTreeValue){
	if this.root==nil{
		this.init(value)
		return
	}
	//查找
	node,pos,same:=getInsertPosition(this.root,value.Key())

	if value.Key()==94{
		log1.Println("在Insert里选择的node的父节点:",node.parent.keys.head.next.value,"位置 : ",pos)
	}
	//如果存在相同的key则覆盖掉
	//fmt.Println("yundao dao 82")
	if same{
		node.children.Relace(value,pos)
		return
	}
	//插入
	node.keys.Insert(value.Key(),pos)
	node.children.Insert(value,pos)

	//插入的value的key>当前treeNode的最大值,递归修改父节点key的最大值.
	if pos>=node.keys.Len()-1{
		parent:=node.parent
		for ;parent!=nil;parent=parent.parent{
			parent.keys.Relace(value.Key(),parent.keys.Len()-1)
		}
	}

	//判断情况
	if node.keys.Len()<=this.degree{
		//1.链表长度<=阶数 不用处理
		return
	}else{
		//2.链表长度>阶数 需要分裂,可能还需要递归分裂
		root:=fission(node,this.degree)
		if root!=nil{
			this.root=root
		}
	}
}


func fission(node *TreeNode,degree int)*TreeNode{
	if node.keys.Len()<=degree{
		return nil
	}

	fissionPos:=node.keys.Len()/2-1
	//当前节点分裂
	kl1,kl2:=node.keys.Fission(fissionPos)
	cl1,cl2:=node.children.Fission(fissionPos)

	node.keys=kl2
	node.children=cl2

	otherNode:=&TreeNode{node.isLeaf,node.parent,kl1,cl1}

	fissionPosKey:=otherNode.keys.tail.value
	maxKey:=node.keys.tail.value

	if node.isLeaf{
		//连接叶子节点
		ln2:=(*leafNode)(unsafe.Pointer(node))
		ln1:=&leafNode{*otherNode,ln2.pre,ln2}
		ln2.pre.next=ln1
		ln2.pre=ln1
		otherNode=(*TreeNode)(unsafe.Pointer(ln1))
	}else{
		//将otherNode的所有孩子的父指针设置为otherNode
		for e:=cl1.head.next;e!=nil;e=e.next{
			e.value.(*TreeNode).parent=otherNode
		}
	}

	//处理父节点
	//maxValue:=node.children.tail.value
	//如果当前节点为根节点
	if node.parent==nil{
		root:=newTreeNode(nil)
		root.keys.Push(fissionPosKey)
		root.keys.Push(maxKey)
		root.children.Push(otherNode)
		root.children.Push(node)
		node.parent=root
		otherNode.parent=root
		return root
	}else{

		pInsertPos,_:=getInsertPos(node.parent,fissionPosKey.(int))
		if fissionPosKey==11{
			log1.Println("在fission里选择了父节点的插入的位置: ",pInsertPos)
		}
		node.parent.keys.Insert(fissionPosKey,pInsertPos)
		node.parent.children.Insert(otherNode,pInsertPos)
		//递归
		return fission(node.parent,degree)
	}
}

//如果node==nil 返回nil ,-1
//如果返回的node!=nil int<0为表示 已经含有相同的key了,会替换掉原来的
func getInsertPosition(node *TreeNode,key int)(*TreeNode,int,bool) {
	if node==nil{
		return nil,-1,false
	}
	pos,same:=getInsertPos(node,key)
	if key==11{
		log1.Print("node: ",node.keys.head.next.value)
		log1.Println("选择了",pos)
	}
	if node.isLeaf{
		return node,pos,same
	}else{
		return getInsertPosition(node.children.GetIndexValue(pos).(*TreeNode),key)
	}
}
func getInsertPos(node *TreeNode,key int)(r int,same bool) {
	if node==nil{
		return -1,false
	}
	pos := 0
	ln := node.keys.head
	for ln = ln.next; ln != nil; pos++ {
		if ln.value.(int)>key{
			return pos,false
		}else if ln.value.(int)==key{
			return pos,true
		}
		ln=ln.next
	}
	return pos,false
}

//返回一个结点的深度 从1开始计数
func (this *BTreePlus)getHigh(tn *TreeNode)int {
	high:=1
	for n:=tn.parent;n!=nil;n=n.parent{
		high++
	}
	return high
}