package leetcode

import (
	"fmt"
	"testing"
)

type question238 struct {
	para238
	ans238
}

// para 是参数
// one 代表第一个参数
type para238 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans238 struct {
	one []int
}

func Test_Problem238(t *testing.T) {

	qs := []question238{

		{
			para238{[]int{1, 2, 3, 4}},
			ans238{[]int{24, 12, 8, 6}},
		},

		{
			para238{[]int{-1, 1, 0, -3, 3}},
			ans238{[]int{0, 0, 9, 0, 0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 238------------------------\n")

	for _, q := range qs {
		a, p := q.ans238, q.para238
		fmt.Printf("【input】:%v       【answer】:%v\n       【output】:%v\n", p, a, productExceptSelf(p.one))
	}
	fmt.Printf("\n\n\n")
}
