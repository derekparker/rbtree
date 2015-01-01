package rbtree

const (
	RED   = true
	BLACK = false
)

type RedBlackTree struct {
	root *Node
}

type Item interface {
	Less(than Item) bool
	More(than Item) bool
}

type Node struct {
	size        uint64
	left, right *Node
	item        Item
	value       string
	color       bool
}

func New() *RedBlackTree {
	return &RedBlackTree{}
}

func (t *RedBlackTree) Size() uint64 {
	return t.root.Size()
}

func (n *RedBlackTree) Put(key Item) {
	n.root = put(n.root, key)
	n.root.color = BLACK
}

func (n *RedBlackTree) Find(key Item) (Item, bool) {
	return find(n.root, key)
}

func (n *Node) Size() uint64 {
	return n.size
}

func (n *Node) Item() Item {
	return n.item
}

func size(n *Node) uint64 {
	if n == nil {
		return 0
	}

	return n.size
}

func put(n *Node, item Item) *Node {
	switch {
	case n == nil:
		return &Node{item: item, size: 1, color: RED}
	case item.Less(n.item):
		n.left = put(n.left, item)
	case n.item.Less(item):
		n.right = put(n.right, item)
	default:
		n.item = item
	}

	if !isRed(n.left) && isRed(n.right) {
		n = rotateLeft(n)
	}

	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}

	if isRed(n.left) && isRed(n.right) {
		changeColors(n)
	}

	n.size = size(n.left) + size(n.right) + 1
	return n
}

func find(n *Node, item Item) (Item, bool) {
	switch {
	case n == nil:
		return nil, false
	case item.Less(n.item):
		return find(n.left, item)
	case item.More(n.item):
		return find(n.right, item)
	default:
		return n.Item(), true
	}

	return nil, false
}

func rotateLeft(n *Node) *Node {
	x := n.right
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = RED

	x.size = n.size
	n.size = size(n.left) + size(n.right) + 1

	return x
}

func rotateRight(n *Node) *Node {
	x := n.left
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = RED

	x.size = n.size
	n.size = size(n.left) + size(n.right) + 1

	return x
}

func changeColors(n *Node) {
	n.left.color = BLACK
	n.right.color = BLACK
	n.color = RED
}

func isRed(n *Node) bool {
	if n == nil {
		return false
	}

	return n.color
}
