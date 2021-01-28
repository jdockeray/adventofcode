package tree

// A Tree is a binary tree with integer values.
type Tree struct {
	Value    int
	children map[int]*Tree
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{
			v, nil,
		}
	}
	if v > t.Value {
		if v-t.Value <= 3 {
			if t.children == nil {
				t.children = make(map[int]*Tree)
			}
			_, exists := t.children[v]
			if !exists {
				t.children[v] = insert(nil, v)
			}
		}
		for _, element := range t.children {
			insert(element, v)
		}
	}
	return t
}

func BuildTree(numbs []int) *Tree {
	tree := insert(nil, numbs[0])
	for _, numb := range numbs {
		insert(tree, numb)
	}
	return tree
}

func (tree *Tree) isLeaf() bool {
	return len(tree.children) == 0
}

func CountLeaves(t *Tree) int {
	if t.isLeaf() {
		return 1
	}
	var sum int
	for _, value := range t.children {
		sum = sum + CountLeaves(value)
	}
	return sum

}
