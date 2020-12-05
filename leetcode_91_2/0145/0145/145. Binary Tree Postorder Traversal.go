package _145

import "github.com/InterviewTips/algorithm-coding/guns"

type TreeNode = guns.TreeNode

// var prev *TreeNode
//    for root != nil || len(stk) > 0 {
//        for root != nil {
//            stk = append(stk, root)
//            root = root.Left
//        }
//        root = stk[len(stk)-1]
//        stk = stk[:len(stk)-1]
//        if root.Right == nil || root.Right == prev {
//            res = append(res, root.Val)
//            prev = root
//            root = nil
//        } else {
//            stk = append(stk, root)
//            root = root.Right
//        }
//    }
//    return
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/er-cha-shu-de-hou-xu-bian-li-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func postorderTraversal(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			res = append(res, node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return res
}
