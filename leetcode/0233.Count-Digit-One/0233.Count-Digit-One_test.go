package leetcode

import (
	"fmt"
	"testing"
)

type question233 struct {
	para233
	ans233
}

// para 是参数
// one 代表第一个参数
type para233 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans233 struct {
	one int
}

func Test_Problem233(t *testing.T) {

	qs := []question233{

		{
			para233{13},
			ans233{6},
		},

		{
			para233{0},
			ans233{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 233------------------------\n")

	for _, q := range qs {
		_, p := q.ans233, q.para233
		fmt.Printf("【input】:%v       【output】:%v\n", p, countDigitOne(p.one))
	}
	fmt.Printf("\n\n\n")
}
