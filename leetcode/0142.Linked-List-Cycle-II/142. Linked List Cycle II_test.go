package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question142 struct {
	para142
	ans142
}

// para 是参数
// one 代表第一个参数
type para142 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans142 struct {
	one bool
}

func Test_Problem142(t *testing.T) {

	qs := []question142{

		{
			para142{[]int{3, 2, 0, -4}},
			ans142{false},
		},

		{
			para142{[]int{1, 2}},
			ans142{false},
		},

		{
			para142{[]int{1}},
			ans142{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 142------------------------\n")

	for _, q := range qs {
		_, p := q.ans142, q.para142
		fmt.Printf("【input】:%v       【output】:%v\n", p, detectCycle(structures.Ints2List(p.one)))
	}
	fmt.Printf("\n\n\n")
}
