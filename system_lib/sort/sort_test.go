package sort

import (
	"math"
	"sort"
	"testing"
)

type array []int

func (cls array) Len() int {
	return len(cls)
}

func (cls array) Less(i, j int) bool {
	return cls[i] < cls[j]
}

func (cls array) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

//-------------------------------------------//
func Test_SortString(t *testing.T) {
	ids := []string{"2", "3", "1"}
	sort.Strings(ids)
	t.Log(ids)
}


func Test_SortDef(t *testing.T) {
	a := array{5,6,1,2,10,3,4,8,7,9}
	sort.Sort(a)
	t.Log(a)
}

func Test_Float64s(t *testing.T) {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	t.Log(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s)
	t.Log(s)
}

func Test_Float64sAreSorted(t *testing.T) {
	s := []float64{0.7, 1.3, 2.6, 3.8, 5.2} // sorted ascending
	t.Log(sort.Float64sAreSorted(s))

	s = []float64{5.2, 3.8, 2.6, 1.3, 0.7} // sorted descending
	t.Log(sort.Float64sAreSorted(s))

	s = []float64{5.2, 1.3, 0.7, 3.8, 2.6} // unsorted
	t.Log(sort.Float64sAreSorted(s))
}