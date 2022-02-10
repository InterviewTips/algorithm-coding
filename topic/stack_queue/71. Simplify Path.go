package stack_queue

func simplifyPath(path string) string {
	tmp := make([]byte, 0)
	res := stack{
		nums: make([]byte, 0),
	}
	newPath := make([]byte, 0)
	for i := 0; i < len(path); i++ {
		newPath = append(newPath, path[i])
	}
	newPath = append(newPath, '/')
	path = string(newPath)
	for i := 0; i < len(path); i++ {
		item := path[i]
		switch item {
		case '/':
			if string(tmp) == "." {
				res.pop()
				tmp = make([]byte, 0) // 重置 tmp
				continue
			}
			if string(tmp) == ".." { // 取出上一级目录 需要注意有可能无法取出上级
				// pop 三次先 /..
				res.pop()
				res.pop()
				res.pop()
				for !res.empty() && res.top() != '/' {
					res.pop()
				}
				if res.empty() { // 根目录
					res.push('/')
				}
				tmp = make([]byte, 0) // 重置 tmp
				continue
			}
			// ... 目录情况
			tmp = make([]byte, 0) // 重置 tmp
			// 正常情况
			if i != len(path)-1 { // 不是最后一个
				res.push(item)
			}
		default: // 正常 push
			tmp = append(tmp, item)
			res.push(item)
		}
	}

	if res.top() == '/' && res.len() != 1 {
		res.pop()
	}

	return string(res.nums)
}

type stack struct {
	nums []byte
}

func (s *stack) len() int {
	return len(s.nums)
}

func (s *stack) empty() bool {
	return len(s.nums) == 0
}

func (s *stack) push(x byte) {
	if !s.empty() && s.top() == '/' && x == '/' { // 多个 / 的情况放这里
		return
	}
	s.nums = append(s.nums, x)
}

func (s *stack) pop() byte {
	x := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return x
}

func (s *stack) top() byte {
	x := s.nums[len(s.nums)-1]
	return x
}
