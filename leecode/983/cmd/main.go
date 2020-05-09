package main

import (
	"fmt"
	"math"
)

type mem struct {
	days map[int]bool
	memo []int
	costs []int
}

func main() {
	days := []int{1,2,3,4,5,6,7,8,9,10,30,31}
	costs := []int{2,7,15}
	fmt.Println(mincostTickets(days, costs))
}

func mincostTickets(days []int, costs []int) int {
	m := new(mem)
	m.init(days,costs)
	return m.dp(1)
}

func (t *mem) init(days []int, costs []int){
	t.days = make(map[int]bool)
	for _,d := range days{
		t.days[d] = true
	}
	t.memo = make([]int, 366)
	t.costs = costs
}

func  (t *mem) dp(i int) int{
	if i > 365{
		return 0
	}

	if t.memo[i] != 0{
		return t.memo[i]
	}

	if _,ok := t.days[i];ok{
		t.memo[i] = int(math.Min(math.Min(float64(t.costs[0] + t.dp(i + 1)), float64(t.costs[1] + t.dp(i + 7))), float64(t.costs[2] + t.dp(i + 30))))
	} else {
		t.memo[i] = t.dp(i + 1)
	}

	return t.memo[i]
}