package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type question32 struct {
	para32
	ans32
	ans32_2
}

// para 是参数
// one 代表第一个参数
type para32 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans32 struct {
	one []int
}

type ans32_2 struct {
	one [][]int
}

func Test_Problem32_1(t *testing.T) {

	qs := []question32{
		{
			para32{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans32{[]int{3, 9, 20, 15, 7}},
			ans32_2{},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 32------------------------\n")

	for _, q := range qs {
		_, p := q.ans32, q.para32
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", levelOrder(root))
	}
	fmt.Printf("\n\n\n")
}

func Test_Problem32_2(t *testing.T) {

	qs := []question32{
		{
			para32{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans32{},
			ans32_2{[][]int{{3}, {9, 20}, {15, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 32------------------------\n")

	for _, q := range qs {
		_, p := q.ans32, q.para32
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", levelOrder32_2(root))
	}
	fmt.Printf("\n\n\n")
}
