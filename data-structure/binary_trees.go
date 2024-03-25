package data_structure

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	Left  *Node[T] `json:"left,omitempty"`
	Right *Node[T] `json:"right,omitempty"`
	Value *T
}

type BinaryTrees[T cmp.Ordered] struct {
	HeadNode Node[T]
}

type NodePosition int

const (
	Left NodePosition = iota
	Right
)

func NewBinaryTrees[T cmp.Ordered](value T) BinaryTrees[T] {
	return BinaryTrees[T]{
		HeadNode: Node[T]{
			Left:  nil,
			Right: nil,
			Value: &value,
		},
	}
}

func (t *BinaryTrees[T]) Add(value T) error {
	if t.HeadNode.Value == nil {
		errors.New("head is nil")
	}

	if value <= *t.HeadNode.Value && t.HeadNode.Left == nil {
		t.HeadNode.Left = &Node[T]{
			Left:  nil,
			Right: nil,
			Value: &value,
		}
		return nil
	}

	isKeepLookup := true
	currentNode := &t.HeadNode
	for isKeepLookup {
		if value <= *t.HeadNode.Value {
			notEmptyNode := t.LookupEmptyNode(currentNode, Left)
			if notEmptyNode == nil {
				currentNode.Left = &Node[T]{
					Left:  nil,
					Right: nil,
					Value: &value,
				}
				isKeepLookup = false
			}
			currentNode = notEmptyNode
		} else {
			notEmptyNode := t.LookupEmptyNode(currentNode, Right)
			if notEmptyNode == nil {
				currentNode.Right = &Node[T]{
					Left:  nil,
					Right: nil,
					Value: &value,
				}
				isKeepLookup = false
			}
			currentNode = notEmptyNode
		}
	}

	return nil
}

func (t BinaryTrees[T]) Search(value T) T {
	return value
}

func (t BinaryTrees[T]) LookupEmptyNode(node *Node[T], position NodePosition) *Node[T] {
	switch true {
	case position == Left && node.Left == nil:
		return nil
	case position == Right && node.Right == nil:
		return nil
	}

	if position == Left {
		return node.Left
	} else {
		return node.Right
	}
}

func (t BinaryTrees[T]) Print() {
	res, _ := json.Marshal(t.HeadNode)
	fmt.Println(string(res))
}
