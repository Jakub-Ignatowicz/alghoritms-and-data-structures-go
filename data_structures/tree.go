package datastructures

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func Insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value, Left: nil, Right: nil}
	}

	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}

	return root
}
