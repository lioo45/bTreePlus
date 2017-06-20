package bTreePlus
//
//import (
//	"unsafe"
//	"log"
//	"os"
//	"strconv"
//)
//
//type BTreeValue interface {
//	Key()int
//}
//
//var log1 *log.Logger
////初始化日志文件 用于调试
//func init(){
//	logFile,_:=os.OpenFile("/Users/pro/log/btreeLog.txt",os.O_CREATE|os.O_RDWR,0666)
//	log1=log.New(logFile,"",log.LstdFlags|log.Lshortfile)
//}
//type BTreePlus struct{
//	root *TreeNode
//	degree int
//	leafHead *leafNode
//}
//
////如果当前节点的叶子节点 那么children指向key对应的value 并且当前node的实际类型是leafNode
//type TreeNode struct {
//	isLeaf bool
//	parent *TreeNode
//	keys *linkedList
//	children *linkedList
//}
//
//type leafNode struct {
//	TreeNode
//	pre *leafNode
//	next *leafNode
//}
//
////初始化B+树
//func New(degree int)*BTreePlus{
//	b:=new(BTreePlus)
//	b.degree=degree
//	return b
//}
//
//
////返回B+树的阶数
//func (this *BTreePlus)Degree()int{
//	return this.degree
//}
//
////插入数据
//func (this *BTreePlus)Insert(value BTreeValue){
//	if this.root==nil||this.root.keys.len<=0{
//		this.init(value)
//		return
//	}
//	//查找插入位置和插入节点,
//	node,pos,same:=getPosition(this.root,value.Key())
//	//如果存在相同的key则覆盖掉
//	if same{
//		node.children.Relace(value,pos)
//		return
//	}
//	//插入
//	node.keys.Insert(value.Key(),pos)
//	node.children.Insert(value,pos)
//	//如果value的key>当前treeNode的最大值,递归修改父节点key的最大值.
//	updateFatherMaxKey(pos,value.Key(),node)
//	//判断情况
//	//1.链表长度<=阶数 不用处理
//	//2.链表长度>阶数 需要分裂,可能还需要递归分裂
//	if node.keys.Len()>this.degree{
//		root:=fission(node,this.degree)
//		if root!=nil{
//			this.root=root
//		}
//	}
//}
////节点分裂
//func fission(node *TreeNode,degree int)*TreeNode{
//	if node.keys.Len()<=degree{
//		return nil
//	}
//	fissionPos:=node.keys.Len()/2-1
//	//当前节点分裂
//	kl1,kl2:=node.keys.Fission(fissionPos)
//	cl1,cl2:=node.children.Fission(fissionPos)
//	node.keys=kl2
//	node.children=cl2
//
//	otherNode:=&TreeNode{node.isLeaf,node.parent,kl1,cl1}
//
//	fissionPosKey:=otherNode.keys.tail.value
//	maxKey:=node.keys.tail.value
//
//	if node.isLeaf{
//		//连接叶子节点 如果othernode不重新赋值  那么其父节点引用的是TreeNode而不是leafNode
//		otherNode=connectLeafNode(otherNode,node)
//	}else{
//		//将新分裂出的节点的 所有孩子的父指针 指向新分裂出的节点
//		for e:=cl1.head.next;e!=nil;e=e.next{
//			e.value.(*TreeNode).parent=otherNode
//		}
//	}
//	//处理父节点
//	//如果当前节点为根节点 执行根节点分裂
//	if node.parent==nil{
//		return rootFission(fissionPosKey.(int),maxKey.(int),otherNode,node)
//	}else{
//		pInsertPos,_:=getPos(node.parent,fissionPosKey.(int))
//		node.parent.keys.Insert(fissionPosKey,pInsertPos)
//		node.parent.children.Insert(otherNode,pInsertPos)
//		//递归
//		return fission(node.parent,degree)
//	}
//}
////连接 叶子节点node2分裂成的两个叶子节点 返回第一个叶子节点
//func connectLeafNode(node1,node2 *TreeNode)*TreeNode{
//	ln2:=(*leafNode)(unsafe.Pointer(node2))
//	ln1:=&leafNode{*node1,ln2.pre,ln2}
//	ln2.pre.next=ln1
//	ln2.pre=ln1
//	node1=(*TreeNode)(unsafe.Pointer(ln1))
//	return node1
//}
//
//func updateFatherMaxKey(pos ,key int ,node *TreeNode){
//	if pos>=node.keys.Len()-1{
//		parent:=node.parent
//		for ;parent!=nil;parent=parent.parent{
//			parent.keys.Relace(key,parent.keys.Len()-1)
//		}
//	}
//}
//
//func rootFission(key1,key2 int ,node1,node2 *TreeNode)*TreeNode{
//	root:=newTreeNode(nil)
//	root.keys.Push(key1)
//	root.keys.Push(key2)
//	root.children.Push(node1)
//	root.children.Push(node2)
//	node1.parent=root
//	node2.parent=root
//	return root
//}
//
////如果node==nil 返回nil ,-1
////如果返回的node!=nil int<0为表示 已经含有相同的key了,会替换掉原来的
//func getPosition(node *TreeNode,key int)(*TreeNode,int,bool) {
//	if node==nil{
//		return nil,-1,false
//	}
//	pos,same:=getPos(node,key)
//	if node.isLeaf{
//		return node,pos,same
//	}else{
//		return getPosition(node.children.GetIndexValue(pos).(*TreeNode),key)
//	}
//}
////获取当前小于或者等于key的keys的位置,如果不都小于key则返回keys的长度
//func getPos(node *TreeNode,key int)(pos int,same bool) {
//	if node==nil{
//		return -1,false
//	}
//	pos = 0
//	ln := node.keys.head
//	for ln = ln.next; ln != nil; pos++ {
//		if ln.value.(int)>key{
//			return pos,false
//		}else if ln.value.(int)==key{
//			return pos,true
//		}
//		ln=ln.next
//	}
//	return pos,false
//}
//
////添加B+树的第一个值
//func (this *BTreePlus)init(value BTreeValue){
//	root:=newTreeNode(nil)
//	root.keys.Push(value.Key())
//	root.children.Push(value)
//	root.isLeaf=true
//
//	this.leafHead=&leafNode{TreeNode{},nil,nil}
//	ln:=&leafNode{*root,this.leafHead,nil}
//	this.leafHead.next=ln
//
//	this.root=(*TreeNode)(unsafe.Pointer(ln))
//
//}
//
////初始化树节点
//func newTreeNode(p *TreeNode)*TreeNode{
//	t:=new(TreeNode)
//	t.parent=p
//	t.keys=NewLinkedList()
//	t.children=NewLinkedList()
//	t.isLeaf=false
//	return t
//}
//
////返回一个结点的深度 从1开始计数
//func (this *BTreePlus)getHigh(tn *TreeNode)int {
//	high:=1
//	for n:=tn.parent;n!=nil;n=n.parent{
//		high++
//	}
//	return high
//}
//
////当不存在这个value时返回false
////47,87,
////25,47,  59,87,
////0,18,25,  40,47,  56,59,  81,87,
//func (this *BTreePlus)Remove(value BTreeValue)bool {
//	cNode,_,same:=getPosition(this.root,value.Key())
//	//不存在这个key
//	if !same{
//		return false
//		log1.Println("不存在这个key")
//	}
//	log1.Println("进入remove")
//	return this.remove(cNode,value.Key())
//}
//
////调用它时已经保证这个key已经存在
//func (this *BTreePlus)remove(cNode *TreeNode,key int)bool{
//	log1.Println("要删除的key ",key)
//	//找到要删除的记录key,value的位置
//	pos,_:=getPos(cNode,key)
//	log1.Println("要删除的key的位置位于节点:",cNode.keys.head.next.value," ",cNode.keys.tail.value,"的位置: ",pos)
//	//删除记录
//	maxKey:=cNode.keys.tail.value.(int)
//	deleteKey:=cNode.keys.Remove(pos).value.(int)
//	cNode.children.Remove(pos)
//	if cNode==this.root&&this.root.children.len==0{
//		//删除后 b+树已经没有任何记录
//		return true
//	}
//	//删除记录后 如果这个记录是当前节点的maxKey 需要递归修改父节点的key
//	if deleteKey==maxKey{
//		log1.Println("开始递归更新父节点的key,更新为:",cNode.keys.tail.value.(int))
//		updateFatherKey(deleteKey,cNode.keys.tail.value.(int), cNode.parent)
//	}
//	//删除记录后节点的keys的长度>=度数/2或者当前节点为根节点
//	if cNode.keys.len>=this.degree/2+1||cNode==this.root{
//		log1.Println("删除",key,"完成")
//		if this.root.children.len==1&&!this.root.isLeaf{
//			this.root=this.root.children.head.next.value.(*TreeNode)
//			this.root.parent=nil
//		}
//		return true
//	}else{
//		//<[m/2]+1 则需要借key value 或者合并
//		broNode,broPos:=brotherHasMoreKey(cNode,cNode.keys.tail.value.(int),this.degree)
//		if broPos>=0{
//			//租借这个兄弟节点一个key value
//			log1.Println("向",broNode.keys.head.next.value," ",broNode.keys.tail.value,"租借kv")
//			borrowKVFromBrother(cNode,broNode)
//		}else{
//			log1.Println("合并",cNode.keys.head.next.value," ",cNode.keys.tail.next,"和",
//			broNode.keys.head.next.value," ",broNode.keys.tail.value)
//			//合并 合并后父节点会删除一个kv 所以可能需要递归处理
//			pKey:=this.combineNode(broNode,cNode)
//			//递归处理
//			this.remove(cNode.parent,pKey)
//		}
//	}
//	return false
//}
//
////谁合并进谁的标准是 使得父节点不产生相同的key
//func (this *BTreePlus)combineNode(broNode,cnode *TreeNode)int{
//	broMaxKey:=broNode.keys.tail.value.(int)
//	cnodeMaxKey:=cnode.keys.tail.value.(int)
//
//	//如果broNode是左兄弟节点就把它合并进当前节点
//	if broMaxKey<cnodeMaxKey{
//		cnode.keys.InsertAllToFront(broNode.keys)
//		cnode.children.InsertAllToFront(broNode.children)
//		if broNode.isLeaf{
//			bl:=(*leafNode)(unsafe.Pointer(broNode))
//			cl:=(*leafNode)(unsafe.Pointer(cnode))
//			bl.pre.next=cl
//			cl.pre=bl.pre
//		}else{
//			//将左兄弟的所有孩子的父节点指向当前节点
//			for e:=broNode.children.head.next;e!=nil;e=e.next{
//				e.value.(*TreeNode).parent=cnode
//			}
//		}
//
//
//		return broNode.keys.tail.value.(int)
//	}else{
//		//如果broNode是右兄弟节点,就把当前节点合并进右兄弟节点
//		broNode.keys.InsertAllToFront(cnode.keys)
//		broNode.children.InsertAllToFront(cnode.children)
//		if cnode.isLeaf{
//			bl:=(*leafNode)(unsafe.Pointer(broNode))
//			cl:=(*leafNode)(unsafe.Pointer(cnode))
//			cl.pre.next=bl
//			bl.pre=cl.pre
//		}else{
//			//将当前节点的所有孩子的父指针指向右兄弟节点
//			for e:=cnode.children.head.next;e!=nil;e=e.next{
//				e.value.(*TreeNode).parent=broNode
//			}
//		}
//		return cnode.keys.tail.value.(int)
//	}
//}
//
//
//func borrowKVFromBrother(node,broNode *TreeNode){
//	broMaxKey:=broNode.keys.tail.value.(int)
//	cMaxKey:=node.keys.tail.value.(int)
//	log1.Println("bro key: ",broMaxKey," ckey: ",cMaxKey)
//	if broMaxKey>cMaxKey{
//		//broNode是右兄弟
//		k:=broNode.keys.Remove(0).value.(int)
//		v:=broNode.children.Remove(0).value
//		if !node.isLeaf{
//			v.(*TreeNode).parent=node
//		}
//		originKey:=node.keys.tail.value.(int)
//		node.keys.Push(k)
//		node.children.Push(v)
//		//要递归设置父节点的key
//		updateFatherKey(originKey,k,node.parent)
//	}else if broMaxKey<cMaxKey{
//		//broNode是左兄弟
//		k:=broNode.keys.Remove(broNode.keys.Len()-1).value.(int)
//		v:=broNode.children.Remove(broNode.children.Len()-1).value
//		log1.Println("broNode是左兄弟,租借的k为",k)
//		if !node.isLeaf{
//			v.(*TreeNode).parent=node
//			node.children.Insert(v.(*TreeNode),0)
//		}else{
//			node.children.Insert(v.(*Student),0)
//		}
//		updateKey:=broNode.keys.tail.value.(int)
//		node.keys.Insert(k,0)
//		//更新兄弟节点的父节点的key
//		updateFatherKey(k,updateKey,broNode.parent)
//	}
//
//}
//
//
////存在兄弟节点的key>当前度数/2+1
////返回这个兄弟节点和它在父节点的位置pos>0
////返回一个可以合并的兄弟节点和-1
////如果没有兄弟节点的话 返回nil
//func brotherHasMoreKey(node *TreeNode,key ,degree int)(*TreeNode,int){
//	//获取两个兄弟节点 判断哪个符合条件
//	pNode:=node.parent
//	var b1,b2 *TreeNode
//	pos,_:=getPos(pNode,key)
//	if pos>0{
//		b1=pNode.children.getIndexNode(pos-1).value.(*TreeNode)
//		if b1.keys.len>degree/2+1{
//			return b1,pos-1
//		}
//	}
//	if pos<pNode.children.len-1{
//		b2=pNode.children.getIndexNode(pos+1).value.(*TreeNode)
//		if b2.keys.len>degree/2+1{
//			return b2,pos+1
//		}
//	}
//	//即没有符合条件的兄弟节点 返回一个可以合并的节点
//	var bro *TreeNode
//	if b1!=nil{
//		bro=b1
//	}else if b2!=nil{
//		bro=b2
//	}
//	return bro,-1
//}
//
////参数maxKey为要需要的key值
////hey成修改后的值
////要修改的key所属的节点
//func updateFatherKey(originKey,key int, node *TreeNode) {
//	if node==nil{
//		return
//	}
//	//父节点一定存在这个key
//
//	pos, _ := getPos(node, originKey)
//	log1.Println("父节点",node.keys.head.next.value," ",node.keys.tail.value,
//		"更新了key:",key,"pos: ",pos,"len:",node.keys.Len())
//	node.keys.Relace(key, pos)
//	//当pos的key是当前节点keys的最大值才递归
//	if pos>=node.keys.Len()-1{
//		updateFatherKey(originKey,key,node.parent)
//	}
//}
//
//func (this *TreeNode)toStr()string{
//	return strconv.Itoa(this.keys.head.next.value.(int))+
//			" "+strconv.Itoa(this.keys.tail.value.(int));
//}