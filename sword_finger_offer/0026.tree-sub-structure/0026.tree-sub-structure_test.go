package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type question32 struct {
	para32
	ans32
}

// para 是参数
// one 代表第一个参数
type para32 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans32 struct {
	one bool
}

func Test_Problem26(t *testing.T) {

	qs := []question32{
		{
			para32{[]int{}, []int{}},
			ans32{false},
		},
		/*{
			para32{[]int{1, 2, 3}, []int{3, 1}},
			ans32{false},
		},
		{
			para32{[]int{3, 4, 5, 1, 2}, []int{4, 1}},
			ans32{true},
		},*/
	}

	fmt.Printf("------------------------Leetcode Problem 32------------------------\n")

	for _, q := range qs {
		_, p := q.ans32, q.para32
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		b := structures.Ints2TreeNode(p.two)
		fmt.Printf("【output】:%v      \n", isSubStructure(root, b))
	}
	fmt.Printf("\n\n\n")
}
