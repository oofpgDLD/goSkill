package main

import (
	"fmt"
	"math"
)

func main() {
	/*a := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}*/

	a := [][]byte{
		{'1'},
	}

	fmt.Println(maximalSquare(a))
}

func maximalSquare(matrix [][]byte) int{
	maxSide := 0
	dp := make([][]int, len(matrix))
	for i,line := range matrix{
		dp[i] = make([]int, len(matrix[i]))
		for j,v := range line{
			dp[i][j] = int(v - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i:=1; i < len(matrix); i++{
		for j := 1; j < len(matrix[i]); j++{
			if dp[i][j] == 1{
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])), float64(dp[i-1][j-1]))) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}