package string

import (
	"math"
	"testing"
)

func updateMatrix(matrix [][]int) [][]int {
	ret := make([][]int, len(matrix))
	for i,lines := range matrix {
		rLine := make([]int, len(matrix[i]))
		for j, item := range lines{
			if item == 0 {
				rLine[j] = 0
			}else {
				rLine[j] = check(matrix,i,j)
			}
		}
		ret[i] = rLine
	}

	return ret
}

func check(matrix [][]int, a,b int) int{
	step := 0
	for i,lines := range matrix {
		for j, item := range lines{
			if item == 0 {
				val := int(math.Abs(float64(a-i)) + math.Abs(float64(b-j)))
				if val < step || step == 0 {
					step = val
				}
			}
		}
	}
	return step
}
/*
func updateMatrixV2(matrix [][]int) [][]int {
	ret := [len(matrix)][]int{}
	for i,lines := range matrix {
		for j, item := range lines{
			if item == 0 {

			}
		}
	}



}
*/
func Test_String(t *testing.T) {
	source := [][]int{
		{0,0,0},
		{0,1,0},
		{1,1,1},
	}
	ret := updateMatrix(source)
	for _, line := range ret{
		t.Log(line)
	}
}