package bTreePlus

import (
	"unsafe"
)

type BTreeValue interface {
	Key()int
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
	node,pos:=getInsertPosition(this.root,value.Key(),true)

	//存在相同的key 覆盖掉
	if node!=nil&&pos<0{
		pos=-pos
		node.children.Relace(value,pos)
	}

	//插入
	node.keys.Insert(value.Key(),pos)
	node.children.Insert(value,pos)
	//if value.Key()=="10"{
	//	fmt.Println("调试信息,node的最大值",node.keys.tail.value,"位置:",pos)
	//}

	if pos>=node.keys.Len()-1{
		//插入的value的key>当前treeNode的最大值,递归修改父节点key的最大值.
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

	if node.isLeaf{
		//连接叶子节点
		ln2:=(*leafNode)(unsafe.Pointer(node))
		ln1:=&leafNode{*otherNode,ln2.pre,ln2}
		ln2.pre.next=ln1
		ln2.pre=ln1
	}


	fissionPosKey:=otherNode.keys.tail.value
	maxKey:=node.keys.tail.value

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
		parent,pInsertPos:=getInsertPosition(node.parent,maxKey.(int),false)
		parent.keys.Insert(fissionPosKey,pInsertPos)
		parent.children.Insert(otherNode,pInsertPos)
		//递归
		return fission(node.parent,degree)
	}
}

//func judgeSituation()


//如果node==nil 返回nil ,-1
//如果返回的node!=nil int<0为表示 已经含有相同的key了,会替换掉原来的
func getInsertPosition(node *TreeNode,key int,rec bool)(*TreeNode,int){
	if node==nil{
		return nil,-1
	}
	pos:=0
	ln:=node.keys.head
	for ln=ln.next;ln!=nil;pos++{
		if ln.value.(int)>key{
			if node.isLeaf||!rec{
				return node,pos
			}else {
				return getInsertPosition(node.children.getIndexNode(pos).value.(*TreeNode),key,rec)
			}
		}else if ln.value.(int)==key{
			if !rec{
				return node,pos
			}
			if node.isLeaf{
				return  node,-pos
			}else{
				return getInsertPosition(node.children.getIndexNode(pos).value.(*TreeNode),key,rec)
			}

		}
		ln=ln.next
	}
	if node.isLeaf||!rec{
		return node,pos
	}else{
		return getInsertPosition(node.children.getIndexNode(pos).value.(*TreeNode),key,rec)
	}
}

