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
