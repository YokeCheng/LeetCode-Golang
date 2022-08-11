package leetcode

func searchMatrix240(matrix [][]int, target int) bool {
	// 从最右面开始找，找到大于target的再次向左找
	// if len(matrix) == 0 {
	// 	return false
	// }
	// row, col := 0, len(matrix[0])-1
	// for col >= 0 && row <= len(matrix)-1 {
	// 	if target == matrix[row][col] {
	// 		return true
	// 	} else if target > matrix[row][col] {
	// 		row++
	// 	} else {
	// 		col--
	// 	}
	// }
	// return false

	// 尝试从最大往下找
	if len(matrix) == 0 {
		return false
	}

	row, col := len(matrix)-1, 0
	for col < len(matrix[0]) && row >= 0 {
		if matrix[row][col] == target {
			return true
		} else if target > matrix[row][col] {
			col++
		} else {
			row--
		}
	}
	return false
}

// 解法一 模拟，时间复杂度 O(m+n)
func searchMatrix2401(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row <= len(matrix)-1 {
		a := matrix[row][col]
		if target == a {
			return true
		} else if target > a {
			row++
		} else {
			col--
		}
	}
	return false
}

// 解法二 二分搜索，时间复杂度 O(n log n)
func searchMatrix2402(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for _, row := range matrix {
		low, high := 0, len(matrix[0])-1
		for low <= high {
			mid := low + (high-low)>>1
			if row[mid] > target {
				high = mid - 1
			} else if row[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

// 从最右面开始找，找到大于target的再次向左找
func searchMatrix3(matrix [][]int, target int) bool {
	// 从最右面开始找，找到大于target的再次向左找
	lLen := len(matrix[0])
	rLen := len(matrix)
	for i := 0; i <= rLen-1; i++ {
		if matrix[i][lLen-1] > target {
			for j := lLen - 2; j >= 0; j-- {
				if matrix[i][j] == target {
					return true
				} else if matrix[i][j] < target {
					break
				}
			}
		} else if matrix[i][lLen-1] == target {
			return true
		}
	}
	return false
}
