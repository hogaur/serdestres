package tree

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{Data: data, Left: nil, Right: nil}
}
