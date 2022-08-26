/*
链接：https://ac.nowcoder.com/acm/contest/5657/I
来源：牛客网

输入描述:
多个测试用例，每个测试用例一行。

每行通过空格隔开，有n个字符，n＜100
输出描述:
对于每组测试用例，输出一行排序过的字符串，每个字符串通过空格隔开

示例1
输入
a c bb
f dddd
nowcoder

输出
a bb c
dddd f
nowcoder
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		line := inputs.Text()
		words := strings.Split(line, " ")
		sort.Strings(words)
		fmt.Println(strings.Join(words, " ")) // 重要 记得需要先 join 而不是直接输出数组
	}
}
