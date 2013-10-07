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

type StringKey string

func (k *StringKey) Equals(other interface{}) bool {
  return string(*k) == string(*(other.(*StringKey)))
}

func (k *StringKey) Less(other interface{}) bool {
  return string(*k) < string(*(other.(*StringKey)))
}

func (k *StringKey) String() string {
  return string(*k)
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
  string_key := (*StringKey)(&key)
  node := t.find_insertion_leaf(string_key, value)
  if node == nil {
    node = NewNode(string_key, value)
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

func (t *Tree) find_insertion_leaf(key Key, value []byte) *Node {
  if t.Root == nil {
    return nil
  }
  panic("unimplemented")
}
