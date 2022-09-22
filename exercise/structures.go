package exercise

import "github.com/halfrost/LeetCode-Go/structures"

func Ints2TreeNode() {
	tree := []int{3, 9, 20, 5, 6, 15, 7}
	root := structures.Ints2TreeNode(tree)
	println(root)
}

func Ints2List() {
	ln := structures.Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println(ln)
}
