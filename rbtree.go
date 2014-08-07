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
	key         Item
	value       string
	color       bool
}

func New() *RedBlackTree {
	return &RedBlackTree{}
}

func (t *RedBlackTree) Size() uint64 {
	return t.root.Size()
}

func (n *RedBlackTree) Put(key Item, value string) {
	n.root = put(n.root, key, value)
	n.root.color = BLACK
}

func (n *RedBlackTree) Find(key Item) (string, bool) {
	return find(n.root, key)
}

func (n *Node) Size() uint64 {
	return n.size
}

func (n *Node) Value() string {
	return n.value
}

func (n *Node) Key() Item {
	return n.key
}

func size(n *Node) uint64 {
	if n == nil {
		return 0
	}

	return n.size
}

func put(n *Node, key Item, value string) *Node {
	switch {
	case n == nil:
		node := Node{key: key, value: value, size: 1, color: RED}
		return &node
	case key.Less(n.key):
		n.left = put(n.left, key, value)
	case n.key.Less(key):
		n.right = put(n.right, key, value)
	default:
		n.value = value
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

func find(n *Node, key Item) (string, bool) {
	switch {
	case n == nil:
		return "", false
	case n.key == key:
		return n.Value(), true
	case key.Less(n.key):
		return find(n.left, key)
	case n.key.Less(key):
		return find(n.right, key)
	}

	return "", false
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

type Int int

func (i Int) Less(item Item) bool {
	return i < item.(Int)
}

func (i Int) More(item Item) bool {
	return i > item.(Int)
}
