package btree

import (
  "fmt"
)

type Key interface {
  Equals(other interface{}) bool
  Less(other interface{}) bool
  String() string
}

type Node struct {
  Parent            *Node
  Left, Right       Key
  LeftVal, RightVal []byte
  Children          [3]*Node
}

type Tree struct {
  Root   *Node
  Size   uint
  Height uint
}

func NewNode() *Node {
  return &Node{}
}

func (node *Node) String() string {
  if node == nil {
    return "<nil>"
  }
  var left, right string
  if node.Left != nil {
    left = node.Left.String()
  } else {
    left = "<nil>"
  }

  if node.Right != nil {
    right = node.Right.String()
  } else {
    right = "<nil>"
  }

  return fmt.Sprintf("[ %s | %s ]", left, right)
}

func NewTree() *Tree {
  return &Tree{}
}

func (t *Tree) String() string {
  return t.Root.String()
}

func (t *Tree) Put(key Key, value []byte) (err error) {
  node := t.find_insertion_leaf(key, value)
  if node == nil {
    node = NewNode()
    node.Left = key
    node.LeftVal = value
    t.Root = node
    t.Height += uint(1)
    t.Size += uint(1)
    return
  }
  panic("unimplemented")
}

func (t *Tree) Get(key string) (value []byte, err error) {
  panic("unimplemented")
}

func (t *Tree) GetSize() uint {
  return t.Size
}

func (t *Tree) GetHeight() uint {
  return t.Height
}

const (
  LEFT = iota
  MID
  RIGHT
)

func Compare(keyA, keyB, other Key) int {
  if other.Less(keyA) {
    return LEFT
  } else if other.Less(keyB) {
    return MID
  }
  return RIGHT
}

func (t *Tree) find_insertion_leaf(key Key, value []byte) (node *Node) {
  node = t.Root
  for node != nil {
    if node.Right == nil {
      return
    }
    cmp := Compare(node.Left, node.Right, key)
    node = node.Children[cmp]
  }
  return
}
