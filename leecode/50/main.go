package main

import "fmt"

func main() {
	fmt.Println(myPow(2.00000, -2))
}

//方法一：快速幂 + 递归
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)

	if n%2 >0 {
		return y*y*x
	}
	return y*y
}

//方法一：快速幂 + 迭代
func myPow2(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)

	if n%2 >0 {
		return y*y*x
	}
	return y*y
}

