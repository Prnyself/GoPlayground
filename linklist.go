package main

import (
	"fmt"
)

type node struct {
	head bool //标记头结点
	Data string
	Next *node
}

//遍历单链表
func throughLinkedList(Head interface{}) {
	next := Head.(*node)
	for {
		fmt.Printf("data:%s \n", next.Data)
		if next.Next != nil {
			next = next.Next
		} else {
			break
		}
	}
}



//在某个元素前插入一个新元素
func InsertAvalueBeforeAelement(Head, Val, InsertVal interface{}) {
	fmt.Printf("链表操作：在元素:%s 前面插入新元素：%s \n", Val, InsertVal)
	next := Head.(*node)
	last := Head.(*node)

	head_node := Head.(*node)

	for { //只能通过从头到尾遍历找到指定结点以及其前驱结点
		if next.Data == Val.(string) {
			n := new(node)
			n.Data = InsertVal.(string)
			n.Next = next
			if next.head == true { //若指定元素是头结点
				n.head = true
				head_node = n    //新的头结点
				next.head = false //把原头结点head改为false
			} else {
				fmt.Println(last, next, 222)
				//n.Next = &next //把新结点的next指向原节点的next
				//last.Next = n      //原节点的前驱结点的next指向新节点
				last.Next = n
				fmt.Println(last, last.Next,last.Next.Next, 222)
			}
			break
		} else {
			last = next
			fmt.Println(last,111)
			if next.Next != nil {
				next = next.Next
			} else {
				break //尾结点要退出
			}
		}



	}
	//遍历一下来验证
	throughLinkedList(head_node)
}

func main() {

	n1 := node{Data: "A"}
	n2 := node{Data: "B"}
	n3 := node{Data: "C"}
	n4 := node{Data: "D"}
	n1.head = true
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4


	//在某个元素前插入一个新元素
	insert_val := "C1"
	before_B := "B"
	InsertAvalueBeforeAelement(&n1, before_B, insert_val)  // <<<----------这个函数
	//TODO 在头结点“A”之前插入C1是ok的，但在其他任意结点前插入C1就不生效，why？
}