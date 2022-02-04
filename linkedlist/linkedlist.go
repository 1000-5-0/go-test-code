package main

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) First() *Node {
	return l.head
}
func (l *LinkedList) Last() *Node {
	return l.tail
}
func (l *LinkedList) Push(value int) {
	node := &Node{value: value}
	if l.head == nil { // head가 없을 경우 : head를 추가시키고 head의 previous를 nil로 만든다.
		l.head = node
		l.head.previous = nil
	} else {
		node.previous = l.tail
		l.tail.next = node // 리스트의 맨 마지막에 현재 노드를 추가
	}
	l.tail = node
}

type Node struct {
	value    interface{}
	next     *Node
	previous *Node
}

func (n *Node) Previous() *Node {
	return n.previous
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *LinkedList) Shift() *Node {
	head := l.head
	newHead := l.head.Next()
	l.head = newHead
	l.head.previous = nil
	return head
}

func (l *LinkedList) Remove(index int) {
	n := l.First()
	i := 0
	for {
		if i == index {
			back := n.Next() // 지워야 할 노드의 뒤를 잡음
			if index == 0 {
				back.previous = nil
				l.head = back
			} else {
				front := n.Previous() // 지워야 할 노드의 앞을 잡음 지금 있는 인덱스 제외
				front.next = back
				back.previous = front
			}
			break
		}
		n = n.Next()
		if n == nil {
			break
		}
		i++
	}
}

func main() {
	l := &LinkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	n := l.First()
	println("--------")
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	println("--------")
	head := l.Shift()
	println("head : ", head.value.(int))
	println("now head : ", l.head.value.(int))
	println("--------add")
	l.Push(4)
	l.Push(5)
	l.Push(6)
	println("--------")
	n = l.First()
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	println("--------remove")
	l.Remove(0)
	println("--------")
	n = l.First()
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	l.Remove(2)
	println("--------")
	n = l.Last()
	for {
		println(n.value.(int))
		n = n.Previous()
		if n == nil {
			break
		}
	}
}
