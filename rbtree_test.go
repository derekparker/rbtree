package rbtree

import "testing"

type Thing struct {
	Key Int
	Val string
}

func (t Thing) Less(item Item) bool {
	return t.Key < item.(Thing).Key
}

func (t Thing) More(item Item) bool {
	return t.Key > item.(Thing).Key
}

type Int int

func (i Int) Less(item Item) bool {
	return i < item.(Thing).Key
}

func (i Int) More(item Item) bool {
	return i > item.(Thing).Key
}

func TestInsertAndRetreivalRedBlack(t *testing.T) {
	tree := New()

	tree.Put(Thing{Int(6), "foo"})
	if tree.Size() != 1 {
		t.Fatal("Value not inserted correctly")
	}

	tree.Put(Thing{Int(7), "bar"})

	if tree.Size() != 2 {
		t.Fatal("Value not inserted correctly", tree.Size())
	}

	tree.Put(Thing{Int(7), "baz"})

	if tree.Size() != 2 {
		t.Fatal("Value not inserted correctly")
	}

	item, ok := tree.Find(Int(6))
	if !ok {
		t.Fatal("Value not found")
	}

	if item.(Thing).Val != "foo" {
		t.Fatal("Value not retreived correctly")
	}

	item, ok = tree.Find(Int(7))
	if !ok {
		t.Fatal("Value not found")
	}

	if item.(Thing).Val != "baz" {
		t.Fatal("Value not retreived correctly")
	}
}

func BenchmarkRedBlack(b *testing.B) {
	var (
		keys = make([]int, 0, b.N)
		tree = New()
	)

	for i := 0; i < b.N*3; i++ {
		keys = append(keys, i)
		tree.Put(Thing{Int(i), "foo"})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Find(Int(keys[i*3]))
	}
}
