package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	replacement := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		replacement.Parent.Left = rplc
	} else {
		replacement.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = replacement.Parent
	}

	return root
}
