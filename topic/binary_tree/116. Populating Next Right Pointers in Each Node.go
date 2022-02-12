package binary_tree

type PerfectNode struct {
	Val   int
	Left  *PerfectNode
	Right *PerfectNode
	Next  *PerfectNode
}

// Node -> PerfectNode 跟原先 NTreeNode 冲突，所以修改下 提交的时候要改回来 -> type PerfectNode = Node
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
// 层次遍历的变种
func connect(root *PerfectNode) *PerfectNode {
	if root == nil {
		return nil
	}
	treeQueue := make([]*PerfectNode, 0)
	// push root to queue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		size := len(treeQueue)
		levelNodes := make([]*PerfectNode, 0)
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			if node.Left != nil {
				treeQueue = append(treeQueue, node.Left)
			}
			if node.Right != nil {
				treeQueue = append(treeQueue, node.Right)
			}
			levelNodes = append(levelNodes, node)
		}
		// next 指向
		for i := 0; i < len(levelNodes); i++ {
			if i > 0 {
				levelNodes[i-1].Next = levelNodes[i]
			}
		}
	}

	return root
}

// 创建完美二叉树
func createPerfectTree(data []int) *PerfectNode {
	// 类似普通的二叉树创建
	if len(data) == 0 {
		return nil
	}
	q := make([]*PerfectNode, 0)
	head := &PerfectNode{Val: data[0]}
	q = append(q, head)
	index := 1 // 从下标索引 1 开始
	for index < len(data) {
		tmp := q[0]
		q = q[1:]
		if data[index] != Null { // -> left
			node := &PerfectNode{Val: data[index]}
			tmp.Left = node
			q = append(q, node) // add node
		}
		index++
		if index < len(data) && data[index] != Null { // -> right
			node := &PerfectNode{Val: data[index]}
			tmp.Right = node
			q = append(q, node) // add node
		}
		index++
	}

	return head
}

// 层次遍历，不过是通过 next 指针来实现 不通过常规实现
func levelOrderPerfectNode(root *PerfectNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	treeQueue := make([]*PerfectNode, 0)
	// push root to queue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		size := len(treeQueue)
		levelRes := make([]int, 0)
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			if node.Left != nil {
				treeQueue = append(treeQueue, node.Left)
			}
			if node.Right != nil {
				treeQueue = append(treeQueue, node.Right)
			}
			// 只有当每层的首个元素才需要遍历
			if i == 0 {
				for node != nil { // 遍历
					levelRes = append(levelRes, node.Val)
					node = node.Next
				}
				// append nil to levelRes
				levelRes = append(levelRes, Null)
			}
		}
		// push levelRes to res
		res = append(res, levelRes...)
	}

	return res
}
