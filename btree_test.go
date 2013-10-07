package btree

import (
  "testing"
)

func TestEmpty(t *testing.T) {
  tree := NewTree()
  if tree.GetSize() != 0 || tree.GetHeight() != 0 {
    panic("New tree is not empty")
  }
}

func TestPutGet(t *testing.T) {
  tree := NewTree()
  err := tree.Put("one", []byte{1})
  if err != nil {
    t.Fatal(err)
  }

  value, err := tree.Get("one")
  if err != nil {
    t.Fatal(err)
  }

  if len(value) != 1 || value[0] != 1 {
    t.Fail()
  }
}
