package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if newColor == color {
		return image
	}
	image[sr][sc] = newColor
	imgLen := len(image)
	// 判断前后左右
	for i := 0; i < 4; i++ {
		// 判断是否超出边界
		//左变右不变，右边做不变
		newSr := sr + dir[i][0]
		newSc := sc + dir[i][1]
		if (newSr >= 0 && newSr < imgLen) && (newSc >= 0 && newSc < imgLen) && image[newSr][newSc] == color {
			floodFill(image, newSr, newSc, newColor)
		}
	}

	return image
}

func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if newColor == color {
		return image
	}
	dfs733(image, sr, sc, newColor)
	return image
}

func dfs733(image [][]int, x, y int, newColor int) {
	if image[x][y] == newColor {
		return
	}
	//先更改目标位置
	oldColor := image[x][y]
	image[x][y] = newColor
	// 目标位置下，最多上下左右，4个方向遍历
	for i := 0; i < 4; i++ {
		// 深度优先遍历，判断当前是否溢出
		if (x+dir[i][0] >= 0 && x+dir[i][0] < len(image)) && (y+dir[i][1] >= 0 && y+dir[i][1] < len(image[0])) && image[x+dir[i][0]][y+dir[i][1]] == oldColor {
			dfs733(image, x+dir[i][0], y+dir[i][1], newColor)
		}
	}
}

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	currColor := image[sr][sc]
	if currColor == newColor {
		return image
	}
	n, m := len(image), len(image[0])
	// 通过队列形成广度优先的队列，防止重复进入
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor
	for i := 0; i < len(queue); i++ {
		cell := queue[i]
		for j := 0; j < 4; j++ {
			mx, my := cell[0]+dx[j], cell[1]+dy[j]
			if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == currColor {
				queue = append(queue, []int{mx, my})
				image[mx][my] = newColor
			}
		}
	}
	return image
}
