package bTreePlus

import (
	"testing"
	"fmt"
	"unsafe"
	"container/list"
	"strconv"
	"math/rand"
)

func Test(t *testing.T){
	testNum:=7
	for k:=0;k<=testNum;k++{
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
		fmt.Println()
	}

}

func Test1111(t *testing.T){
	ids:=[]int{81,87,47,59,81,18,25,}
	btree:=New(3)

	for i:=0;i<len(ids)-1;i++{
		id:=ids[i]
		fmt.Println("插入: ",id)
		s:=&Student{id,strconv.Itoa(id)}
		btree.Insert(s)
		printAll(btree)
		printLeaf(btree)
		fmt.Println()
	}

}
func TestWithRandInt(t *testing.T){
	testNum:=6
	btree:=New(3)
	fmt.Println("随机数为:")
	for i:=0;i<=testNum ;i++  {
		id:=rand.Intn(100)
		fmt.Print(id,",")
		s:=&Student{id,strconv.Itoa(id)}
		btree.Insert(s)
	}
	//printLeaf(btree)
}

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

func printKey(tn *TreeNode){
	for e:=tn.keys.head.next;e!=nil;e=e.Next(){
		fmt.Print(e.value,",")
	}
}

func printLeaf(bt *BTreePlus){
	ln:=bt.leafHead.next

	for ;ln!=nil ;  ln=ln.next{
		for e:=ln.children.head.next;e!=nil ;e=e.next  {
			fmt.Print(e.value,",")
		}
	}
	fmt.Println()
}

func TestNewLeaf(t *testing.T){
	tn:=newTreeNode(nil)

	//ln:=newLeafNode(tn)
	ln:=&leafNode{*tn,nil,nil}
	tn1:=(*TreeNode)(unsafe.Pointer(ln))
	fmt.Println(ln.children)
	fmt.Println(tn1)
	fmt.Println(tn)

}
//当前测试的阶数为3
//通过测试b+树的6个性质,可以验证是否正确

//性质1:2<=根节点的分支<