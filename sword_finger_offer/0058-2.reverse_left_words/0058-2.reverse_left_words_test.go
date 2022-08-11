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
	num int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one string
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{"We are happy.", 2},
			ans1{" are happy.We"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		_, p := q.ans1, q.para1

		fmt.Printf("【input】:%v 【output】:%v\n", p, reverseLeftWords(p.str, p.num))
	}
}
