package bTreePlus


type BTreeValue interface {
	Key()string
}

type BTreePlus struct{
	root *TreeNode
	degree int
}

//如果当前节点的叶子节点 那么children指向key对应的value
type TreeNode struct {
	isLeaf bool
	parent *TreeNode
	keys *linkedList
	children *linkedList
}

//初始化B+树
func New(degree int)*BTreePlus{
	b:=new(BTreePlus)
	b.degree=degree
	return b
}

//添加第一个值
func (this *BTreePlus)init(value BTreeValue){
	this.root=this.makeTreeNode(nil)
	this.root.keys.Push(value.Key())
	this.root.children.Push(value)
	this.root.isLeaf=true
}

//初始化树节点
func (this *BTreePlus)makeTreeNode(p *TreeNode)*TreeNode{
	t:=new(TreeNode)
	t.parent=p
	t.keys=NewLinkedList()
	t.children=NewLinkedList()
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
	node,pos:=getInsertPosition(this.root,value)
	//插入
	node.keys.Insert(value.Key(),pos)
	node.children.Insert(value,pos)


	if pos>=node.keys.Len(){
		//插入的value的key>当前treeNode的最大值,递归修改父节点key的最大值.
		parent:=node.parent
		for ;parent!=nil;parent=parent.parent{
			parent.keys.Push(value.Key())
		}
	}
	//判断情况
	if node.keys.Len()<=this.degree{
		//1.链表长度<=阶数 不用处理
		return
	}else{
		//2.链表长度>阶数 需要分裂,可能还需要递归分裂
		//kl1,kl2:=node.keys.FissionList(pos)
		//vl1,vl2:=node.children.FissionList(pos)
		
	}
}


func getInsertPosition(node *TreeNode,v BTreeValue)(*TreeNode,int){
	key:=v.Key()
	pos:=0
	ln:=node.keys.head
	for ln=ln.next;ln!=nil;pos++{
		if ln.value.(BTreeValue).Key()>key{
			if node.isLeaf{
				return node,pos
			}else {
				return getInsertPosition(node.children.getIndexNode(pos).value.(*TreeNode),v)
			}
		}
	}
	if node.isLeaf{
		return node,pos
	}else{
		return getInsertPosition(node.children.getIndexNode(pos-1).value.(*TreeNode),v)
	}
}

