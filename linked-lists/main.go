package main

import "fmt"

type node struct {
	data     int
	nextNode *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) printList() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.nextNode
	}
}

func (l *linkedList) appendNode(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.nextNode != nil {
			current = current.nextNode
		}
		current.nextNode = n
	}
}

func (l *linkedList) prependNode() {

}

func main() {
	myList := linkedList{}

	myList.appendNode(&node{
		data: 100,
	})
	myList.appendNode(&node{
		data: 10,
	})
	myList.printList()
}
