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
		fmt.Printf("%d ->", current.data)
		current = current.nextNode
	}
	fmt.Println("nil")
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

func (l *linkedList) prependNode(n *node) {
	oldHead := l.head
	l.head = n
	l.head.nextNode = oldHead
}

func (l *linkedList) deleteNode(n *node) {

}

func generateNodes() []*node {
	var nodes []*node
	for i := 2; i < 100; i += 2 {
		nodes = append(nodes, &node{
			data: i,
		})
	}
	return nodes
}

func main() {
	myList := linkedList{}
	nodes := generateNodes()
	for i := 0; i < len(nodes); i++ {
		myList.prependNode(nodes[i])
	}
	myList.printList()
}
