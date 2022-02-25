package _0220225

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// 实部+虚部i
func complexNumberMultiply(num1 string, num2 string) string {
	real1, imaginary1 := getComplexNumber(num1)
	real2, imaginary2 := getComplexNumber(num2)
	// 运算
	calcReal := real1*real2 + (-imaginary1 * imaginary2)
	calcImaginary := real1*imaginary2 + real2*imaginary1

	log.Println(calcReal, calcImaginary)

	return fmt.Sprintf("%v+%vi", calcReal, calcImaginary)
}

func getComplexNumber(num string) (int, int) {
	num = strings.TrimSuffix(num, "i")
	nums := strings.Split(num, "+")
	if len(nums) != 2 {
		panic(fmt.Sprintf("err num %v", num))
	}
	realPart, _ := strconv.Atoi(nums[0])
	imaginaryPart, _ := strconv.Atoi(nums[1])

	log.Println("num", realPart, imaginaryPart)
	return realPart, imaginaryPart
}
