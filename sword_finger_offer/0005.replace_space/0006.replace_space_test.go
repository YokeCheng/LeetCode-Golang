package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	str string
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one string
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{"We are happy."},
			ans1{"We%20are%20happy."},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		_, p := q.ans1, q.para1

		fmt.Printf("【input】:%v 【output】:%v\n", p, replaceSpace2(p.str))
	}
}
