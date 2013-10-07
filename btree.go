package btree

import (
  "fmt"
  "math"
)

const order = 3 // maximum number of children
var min_non_leaf_order = int(math.Ceil(float64(order) / 2.0))

type Key interface {
  Equals(other interface{}) bool
  Less(other interface{}) bool
  String() string
}

type Node struct {
  Key      Key
  Value    []byte
  Children []*Node
}

type Tree struct {
  Root   *Node
  Size   uint
  Height uint
}

func NewNode(key Key, value []byte) *Node {
  return &Node{Key: key, Value: value, Children: []*Node{}}
}

func NewTree() *Tree {
  return &Tree{}
}

func (t *Tree) String() string {
  return fmt.Sprintf("%+v", t)
}

func (t *Tree) Put(key string, value []byte) (err error) {
  panic("unimplemented")
}

func (t *Tree) GetSize() uint {
  return t.Size
}

func (t *Tree) GetHeight() uint {
  return t.Height
}
