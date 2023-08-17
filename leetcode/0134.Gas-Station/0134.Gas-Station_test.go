package leetcode

import (
	"fmt"
	"testing"
)

type question135 struct {
	para135
	ans135
}

// para 是参数
// one 代表第一个参数
type para135 struct {
	gas  []int
	cost []int
}

// ans 是答案
// one 代表第一个答案
type ans135 struct {
	one int
}

func Test_Problem135(t *testing.T) {

	qs := []question135{

		{
			para135{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
			ans135{3},
		},

		{
			para135{[]int{2, 3, 4}, []int{3, 4, 3}},
			ans135{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 135------------------------\n")

	for _, q := range qs {
		a, p := q.ans135, q.para135
		fmt.Printf("【input】:%v   【answer】:%v   【output】:%v\n", p, a, canCompleteCircuit(p.gas, p.cost))
	}
	fmt.Printf("\n\n\n")
}
