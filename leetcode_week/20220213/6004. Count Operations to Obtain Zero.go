package _0220213

// num1 >= num2
// if num1 == 0 || num2 == 0 break
func countOperations(num1 int, num2 int) int {
	count := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 = num1 - num2
		} else {
			num2 = num2 - num1
		}
		count++
	}
	return count
}
