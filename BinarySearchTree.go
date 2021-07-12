package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	left *node
	right *node
}

func (n node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type bst struct {
	root *node
	len int
}

func (b bst) String() string {
	var sb strings.Builder
	b.traversal(&sb, b.root)
	return sb.String()
}

func (b bst) traversal(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.traversal(sb,root.left)
	sb.WriteString(fmt.Sprintf("%v ", root))
	b.traversal(sb, root.right)

}

func (b *bst) add(value int) {
	b.root = b.addTraversal(b.root, value)
	b.len++
}

func (b *bst) addTraversal(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	if value < root.value {
		root.left = b.addTraversal(root.left, value)
	} else {
		root.right = b.addTraversal(root.right, value)
	}
	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchTraversal(b.root, value)
}

func (b bst) searchTraversal(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}
	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchTraversal(root.left, value)
	} else {
		return b.searchTraversal(root.right, value)
	}
}

func (b *bst) remove(value int) {
	b.root = b.removeTraversal(b.root, value)
	b.len--
}

func (b *bst) removeTraversal(root *node, value int) *node {
	if root == nil {
		return root
	}
	if value < root.value {
		root.left = b.removeTraversal(root.left, value)
	} else if value > root.value {
		root.right = b.removeTraversal(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		}
		temp := root.left
		for temp.right != nil {
			temp = temp.right
		}
		root.value = temp.value
		root.left = b.removeTraversal(root.left,temp.value)
	}
	return root
}

func main() {
	b := new(bst)
	b.add(2)
	b.add(5)
	b.add(0)
	b.add(1)
	b.add(8)
	b.add(39)
	b.add(6)
	b.add(17)
	b.add(3)
	fmt.Println(b)
	fmt.Println(b.search(5))
	fmt.Println(b.search(10))
	b.remove(2)
	fmt.Println(b)
	b.remove(5)
	fmt.Println(b)
}
