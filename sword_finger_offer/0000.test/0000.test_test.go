package leetcode

import (
	"fmt"
	"testing"
)

type question66 struct {
	para66
	ans66
}

// para 是参数
// one 代表第一个参数
type para66 struct {
	nums []string
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans66 struct {
	one []int
}

func Test_Problem66(t *testing.T) {

	qs := []question66{

		{
			para66{[]string{"A..", "AA.", "..."}, 3},
			ans66{[]int{5, 6, 7, 1, 2, 3, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 66------------------------\n")

	for _, q := range qs {
		_, p := q.ans66, q.para66
		fmt.Printf("【input】:%v ", p)

		fmt.Printf("【output】:%v\n", findContinuousSequence(9))
	}
	fmt.Printf("\n\n\n")
}
