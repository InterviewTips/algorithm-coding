package binary_tree

func verifySequenceOfBST(sequence []int) bool {
	if len(sequence) <= 1 { // 小于等于 1 表示真
		return true
	}

	// 获取根节点
	last := len(sequence) - 1
	mid := sequence[last]
	// 先找到第一个大于根节点的索引
	index := 0
	for sequence[index] < mid && index < len(sequence) {
		index++
	}
	splitN := index // 切分左右子树
	for sequence[index] > mid && index < len(sequence) {
		index++
	}

	// [0, splitN)
	left := verifySequenceOfBST(sequence[0:splitN])
	// [splitN, last)
	right := verifySequenceOfBST(sequence[splitN:last])

	// index == last 表示是正常的二叉搜索树
	return index == last && left && right
}
