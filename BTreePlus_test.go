package bTreePlus

import (
	"testing"
	"fmt"
	"unsafe"
	"container/list"
	"strconv"
	"math/rand"
)

//当前测试的阶数为3
//通过测试b+树的6个性质,可以验证是否正确
func TestAll(t *testing.T){
	testNum:=100000
	btree:=New(5)
	fmt.Println("随机数为:")
	for i:=0;i<=testNum ;i++  {
		id:=rand.Intn(100000)
		fmt.Print(id,",")
		s:=&Student{id,strconv.Itoa(id)}
		btree.Insert(s)
	}
	testAllNature(btree,t)
	//printAll(btree)
}

func testAllNature(bt *BTreePlus,t *testing.T){
	testNature123(bt,t)
	testNature4(bt,t)
	testNature5(bt,t)
	testNature6(bt,t)

}
//(1)根结点只有1个，分支数量范围[2,m]。这个就不用测了..
//(2)除根以外的非叶子结点，每个结点包含分支数范围[[m/2],m]，其中[m/2]表示取大于m/2的最小整数。
//(3)所有非叶子节点的关键字数目等于它的分支数量。
//按层序遍历所有节点,测试节点是否满足这两个性质
func testNature123(bt *BTreePlus,t *testing.T){
	tn:=bt.root
	l1:=list.New()
	l2:=list.New()
	l1.PushBack(tn)
	min:=bt.degree/2+1
	max:=bt.degree
	for ;l1.Len()>0 ;  {
		l1,l2=l2,l1
		for ; l2.Len()>0;  {
			e:=l2.Remove(l2.Front()).(*TreeNode)
			if bt.root!=e{
				if e.keys.Len()<min||e.keys.Len()>max||e.children.Len()<min||e.children.Len()>max{
					t.Error(fmt.Sprintf("不满足性质2. 节点:",e,"最小值:",e.keys.head.next.value,"\n"))

				}
				if e.keys.Len()!=e.children.Len(){
					t.Error(fmt.Sprintf("不满足性质3. 节点:",e,"最小值:",e.keys.head.next.value,"\n"))
				}
			}

			if !e.isLeaf{
				for c:=e.children.head.next;c!=nil;c=c.next {
					l1.PushBack(c.value.(*TreeNode))
				}
			}
		}
	}
}
//(4) 所有叶子节点都在同一层，且关键字数目范围是[[m/2],m]，其中[m/2]表示取大于m/2的最小整数。
//我实现的b+树中有一个指针 指向叶子节点的头节点,根据这个指针可以遍历到所有的叶子节点
//那么我们只要求每个叶子节点的深度就可以验证这个性质的第一点
func testNature4(bt *BTreePlus,t *testing.T){
	ln:=bt.leafHead.next
	high:=bt.getHigh((*TreeNode)(unsafe.Pointer(ln)))
	for ln=ln.next;ln!=nil;ln=ln.next{
		if high!=bt.getHigh((*TreeNode)(unsafe.Pointer(ln))){
			t.Error(fmt.Sprintf("不满足性质4. 节点:",ln,"最小值:",ln.keys.head.next.value,"\n"))
		}
	}
}

//性质5:所有非叶子节点的关键字可以看成是索引部分，这些索引等于其子树（根结点）中的最大（或最小）关键字。
// 例如一个非叶子节点包含信息: (n，A0,K0, A1,K1,……,Kn,An)，其中Ki为关键字，Ai为指向子树根结点的指针，
// n表示关键字个数。即Ai所指子树中的关键字均小于或等于Ki，而Ai+1所指的关键字均大于Ki（i=1，2，……，n）。
//测试:按层序从左到右测试
func testNature5(bt *BTreePlus,t *testing.T) {
	tn:=bt.root
	l1:=list.New()
	l2:=list.New()
	l1.PushBack(tn)
	for ;l1.Len()>0 ;  {
		l1,l2=l2,l1
		for ; l2.Len()>0;  {
			e:=l2.Remove(l2.Front()).(*TreeNode)

			if !e.isLeaf{
				for c:=e.children.head.next;c!=nil;c=c.next {
					l1.PushBack(c.value.(*TreeNode))
				}
				//测试是否 keyi>=[ai]
				k,c:=e.keys.head.next,e.children.head.next
				for ;k!=nil&&c!=nil;k,c=k.next,c.next{
					isOrdered(k.value.(int),c.value.(*TreeNode),t)
				}
			}
		}
	}
}

func isOrdered(key int,tn *TreeNode,t *testing.T){

	for e:=tn.keys.head.next;e!=nil;e=e.next{
		if e.value.(int)>key{
			t.Error(fmt.Sprintf("不满足性质5. 节点:",tn,"最小值:",tn.keys.head.next.value,"\n"))
		}
	}
}

//(6)叶子节点包含全部关键字的信息(非叶子节点只包含索引)，且叶子结点中的所有关键字依照大小顺序链接
//(所以一个B+树通常有两个头指针，一个是指向根节点的root，另一个是指向最小关键字的sqt)。
func testNature6(bt *BTreePlus,t *testing.T) {
	//直接测试是否有序即可
	ln:=bt.leafHead.next
	for ;ln!=nil ;  ln=ln.next{
		for e:=ln.children.head.next;e.next!=nil ;e=e.next  {
			if e.value.(*Student).Key()>e.next.value.(*Student).Key(){
				t.Error(fmt.Sprintf("不满足性质6. 节点:",ln,"最小值:",ln.keys.head.next.value,"\n"))
			}
		}
	}
	fmt.Println()
}


func Test1111(t *testing.T){
	ids:=[]int{81,87,47,59,81,18,25,}
	btree:=New(3)
	for i:=0;i<len(ids);i++{
		id:=ids[i]
		fmt.Println("插入: ",id)
		s:=&Student{id,strconv.Itoa(id)}
		btree.Insert(s)
		printAll(btree)
		printLeaf(btree)
		fmt.Println()
	}
	fmt.Println(btree.leafHead)
	fmt.Printf("%p",(*leafNode)(unsafe.Pointer(btree.root.children.GetIndexValue(0).(*TreeNode))))
}

func TestWithRandInt(t *testing.T){
	testNum:=100
	btree:=New(3)
	fmt.Println("随机数为:")
	for i:=0;i<=testNum ;i++  {
		id:=rand.Intn(100000)
		fmt.Print(id,",")
		s:=&Student{id,strconv.Itoa(id)}
		btree.Insert(s)
	}
	fmt.Println()
	fmt.Println()
	printAll(btree)
	printLeaf(btree)
	n1:=btree.root.children.head.next.value.(*TreeNode)
	n2:=n1.children.head.next.value.(*TreeNode)
	fmt.Println(n2.keys.tail.value)

}
//按层序输出B+树
func printAll(bt *BTreePlus){
	tn:=bt.root
	l1:=list.New()
	l2:=list.New()
	l1.PushBack(tn)
	for ;l1.Len()>0 ;  {
		l1,l2=l2,l1
		for ; l2.Len()>0;  {
			e:=l2.Remove(l2.Front()).(*TreeNode)
			printKey(e)

			if !e.isLeaf{
				for c:=e.children.head.next;c!=nil;c=c.next {
					l1.PushBack(c.value.(*TreeNode))
				}
			}
			fmt.Print("  ")
		}
		fmt.Println()
	}
}
//输出一个节点所包含的所有关键字
func printKey(tn *TreeNode){
	for e:=tn.keys.head.next;e!=nil;e=e.Next(){
		fmt.Print(e.value,",")
	}
}

//输出所有叶子节点
func printLeaf(bt *BTreePlus){
	ln:=bt.leafHead.next

	for ;ln!=nil ;  ln=ln.next{
		for e:=ln.children.head.next;e!=nil ;e=e.next  {
			fmt.Print(e.value,",")
		}

	}
	fmt.Println()
}

func Test(t *testing.T){
	testNum:=4
	for k:=testNum;k<=testNum;k++{
		btree:=New(3)
		fmt.Println("测试:",k)
		for i:=0;i<= k;  i++{
			s:=new(Student)
			s.ID=i
			s.Name=strconv.Itoa(i)
			btree.Insert(s)
		}
		printAll(btree)
		printLeaf(btree)


		fmt.Println(btree.leafHead)
		fmt.Printf("%p",(*leafNode)(unsafe.Pointer(btree.root.children.GetIndexValue(0).(*TreeNode))))
	}

}

func verDouble(btree *BTreePlus){
	for pre,node:=btree.leafHead,btree.leafHead.next;node!=nil;{
		if !(pre==node.pre){
			fmt.Println("不是双链表-----------")
			return
		}
		node=node.next
		pre=pre.next
	}
}
