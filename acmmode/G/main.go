/*
链接：https://ac.nowcoder.com/acm/contest/5657/G
来源：牛客网

输入描述:
输入数据有多组, 每行表示一组输入数据。

每行不定有n个整数，空格隔开。(1 <= n <= 100)。

输出描述:
每组数据输出求和的结果

示例1

输入
1 2 3
4 5
0 0 0 0 0

输出
6
9
0
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := bufio.NewScanner(os.Stdin) // 分割行
	for inputs.Scan() {                  // 每次读入一行
		data := strings.Split(inputs.Text(), " ") // 空格分割
		var sum int
		for i := 0; i < len(data); i++ {
			num, _ := strconv.Atoi(data[i])
			sum += num
		}
		fmt.Println(sum)
	}
}
