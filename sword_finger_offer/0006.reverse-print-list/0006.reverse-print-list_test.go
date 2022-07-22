package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{1, 2, 3}},
			ans1{[]int{3, 2, 1}},
		},
		{

			para1{[]int{4, 5, 6}},
			ans1{[]int{6, 5, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		_, p := q.ans1, q.para1

		fmt.Printf("【input】:%v 【output】:%v\n", p, reversePrint(structures.Ints2List(p.nums)))
	}
}
