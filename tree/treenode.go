package tree

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{}
}
