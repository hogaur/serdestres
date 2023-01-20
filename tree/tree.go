package tree

type Tree struct {
	Root *TreeNode
}

func NewTree(root int) *Tree {
	rootNode := NewTreeNode(root)
	return &Tree{Root: rootNode}
}
