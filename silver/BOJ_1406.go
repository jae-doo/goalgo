package main

type Cursor struct {
	cur    *Node
	isLeft bool
}

type Node struct {
	prev *Node
	next *Node
	val  rune
}

func (cursor *Cursor) add(c rune) {
	newNode := &Node{val: c}
	cursor.isLeft = false
	if cursor.cur == nil {
		cursor.cur = newNode
	} else {
		if cursor.cur.next != nil {
			cursor.cur.next.prev = newNode
			newNode.next = cursor.cur.next
		}
		cursor.cur.next = newNode
		newNode.prev = cursor.cur
	}
}

func (cursor *Cursor) remove() {
	if !cursor.isLeft {
		cursor.cur.next.prev = cursor.cur.prev
		if cursor.cur.prev != nil {
			cursor.cur.prev.next = cursor.cur.next
			cursor.cur = cursor.cur.prev
		} else {
			cursor.cur = cursor.cur.next
			cursor.isLeft = true
			cursor.cur.prev = nil
		}
	}
}
