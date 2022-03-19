package leetcode

import "bytes"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// 初始化
	data := make([][]byte, numRows)
	for i := 0; i < len(data); i++ {
		data[i] = make([]byte, 0)
	}

	row := 0
	flag := -1
	for i := 0; i < len(s); i++ {
		// 获取到对应的行
		data[row] = append(data[row], s[i])
		if row == 0 || row == numRows-1 { // 切换 flag 进行反方向
			flag = -flag
		}
		row += flag
	}

	var res bytes.Buffer
	for i := 0; i < len(data); i++ {
		res.Write(data[i])
	}

	return res.String()
}
