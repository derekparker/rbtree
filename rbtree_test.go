package rbtree

import "testing"

func TestInsertAndRetreivalRedBlack(t *testing.T) {
	tree := New()

	tree.Put(Int(6), "foo")
	if tree.Size() != 1 {
		t.Fatal("Value not inserted correctly")
	}

	tree.Put(Int(7), "bar")

	if tree.Size() != 2 {
		t.Fatal("Value not inserted correctly", tree.Size())
	}

	tree.Put(Int(7), "baz")

	if tree.Size() != 2 {
		t.Fatal("Value not inserted correctly")
	}

	val, ok := tree.Find(Int(6))
	if !ok {
		t.Fatal("Value not found")
	}

	if val != "foo" {
		t.Fatal("Value not retreived correctly")
	}

	val, ok = tree.Find(Int(7))
	if !ok {
		t.Fatal("Value not found")
	}

	if val != "baz" {
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
		tree.Put(Int(i), "foo")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Find(Int(keys[i*3]))
	}
}
