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
	opera []string
	nums  [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{[]string{"CQueue", "appendTail", "deleteHead", "deleteHead"}, [][]int{
				{}, {3}, {}, {},
			}},
			ans1{[]int{}},
		},
		{
			para1{[]string{"CQueue", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead"}, [][]int{
				{}, {}, {5}, {2}, {}, {},
			}},
			ans1{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		_, p := q.ans1, q.para1
		var obj CQueue
		var res []interface{}
		for i, s := range p.opera {
			switch s {
			case "CQueue":
				obj = Constructor()
				res = append(res, nil)
			case "appendTail":
				obj.AppendTail(p.nums[i][0])
				res = append(res, nil)
			case "deleteHead":
				res = append(res, obj.DeleteHead())
			}
		}
		fmt.Printf("【input】:%v 【output】:%v\n", p, res)
	}
}
