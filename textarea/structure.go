package textarea

type cursor struct {
	line  *node
	start int
	len   int
}

type node struct {
	data     string
	cursor   *cursor
	previous *node
	next     *node
}

func newNode(data string) *node {
	return &node{
		previous: nil,
		next:     nil,
		data:     data,
	}
}

func appendNode(head *node, data string) {
	tmp := head.next
	head.next = newNode(data)
	head.next.previous = head
	head.next.next = tmp
}

func deletePrevious(n *node) {
	if n.previous != nil {
		n.previous = n.previous.previous
		n.previous.previous.next = n
	}
}

func deleteNext(n *node) {
	if n.next != nil {
		n.next = n.next.next
		n.next.next.previous = n
	}
}
