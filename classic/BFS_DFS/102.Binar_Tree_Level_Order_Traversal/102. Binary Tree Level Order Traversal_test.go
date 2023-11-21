package _02_Binar_Tree_Level_Order_Traversal

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question102 struct {
	para102
	ans102
}

// para 是参数
// one 代表第一个参数
type para102 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans102 struct {
	one [][]int
}

func Test_Problem102(t *testing.T) {

	qs := []question102{
		{
			para102{[]int{3, 9, 20, 5, 6, 15, 7}},
			ans102{[][]int{{3}, {9, 20}, {5, 6, 15, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 102------------------------\n")

	for _, q := range qs {
		_, p := q.ans102, q.para102
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", levelOrder(root))
	}
	fmt.Printf("\n\n\n")
}

func Test_Problem1023(t *testing.T) {
	fmt.Printf("【output】:%v      \n", isStraight([]int{9, 4, 2, 5, 6}))
}
