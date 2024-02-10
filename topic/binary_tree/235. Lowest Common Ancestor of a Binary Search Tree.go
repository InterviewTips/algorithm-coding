package binary_tree

import "algorithm/template"

// corner case
// https://sichengingermay.com/interviewer/
// todo: 假设这两个员工中，一个已经向另外一个汇报了，那么共同经理是谁？ 返回祖先的父节点
// var lowestCommonAncestor = lowestCommonAncestorBST
func lowestCommonAncestorBST(root, p, q *template.TreeNode) *template.TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorBST(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorBST(root.Right, p, q)
	}

	// 后序遍历 左右中
	left := lowestCommonAncestorBST(root.Left, p, q)
	right := lowestCommonAncestorBST(root.Right, p, q)

	if (left == q && right == p) || (left == p && right == q) {
		return root
	}

	if left == nil && right == nil {
		return nil
	}

	if left == nil { // right 不为空
		return right
	}

	if right == nil { // left 不为空
		return left
	}

	// 理论上走不到这里
	return nil
}
