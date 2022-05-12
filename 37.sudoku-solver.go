/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

/**
 * 先遍历整个表，建立x, y, z 三个维度的bitmap
 * bitmap的x为每行包含的数字bit集, y为每列, z为每个小9宫格
 * 例如x[0] = 001001010 表示该行包含的数字为2,4,7
 * 在遍历表的同时，将所有未填写的cell保存到数组中并返回
 * 遍历所有空cell，根据bitmap来获取可以填的数字
 * 按照每个cell可选数字的个数排序，结果记为emptyCells
 * 从emptyCells[0]开始，遍历该cell所有可能的数字，在bitmap中测试是否有效
 * 如果有效，设置该位置对应的bitmap, 调用solve(emptyCells, i+1)递归求后续的解
 * 如果无效, 清除该位置对应的bitmap, 循环该cell下一个数字
 */
// @lc code=start
type EmptyCell struct {
	X        int
	Y        int
	Possible []int
	cur      int
}
type BitMap struct {
	x [9]int
	y [9]int
	z [9]int
}

func (bm *BitMap) IsValid(x, y, n int) bool {
	z := x/3*3 + y/3
	bit := 1 << n
	allBit := bm.x[x] | bm.y[y] | bm.z[z]
	return allBit&bit == 0
}
func (bm *BitMap) AddBit(x, y, n int) {
	z := x/3*3 + y/3
	bit := 1 << n
	bm.x[x] |= bit
	bm.y[y] |= bit
	bm.z[z] |= bit
}
func (bm *BitMap) RemoveBit(x, y, n int) {
	z := x/3*3 + y/3
	bit := 1 << n
	bm.x[x] ^= bit
	bm.y[y] ^= bit
	bm.z[z] ^= bit
}
func (bm *BitMap) SetPossible(cell *EmptyCell) {
	//对所有bit求并,所有可能的数字对应的bit为0
	z := cell.X/3*3 + cell.Y/3
	allBit := bm.x[cell.X] | bm.y[cell.Y] | bm.z[z]
	possible := []int{}
	for n := 0; n < 9; n++ {
		if allBit&(1<<n) == 0 {
			possible = append(possible, n)
		}
	}
	cell.Possible = possible
}

func solveSudoku(board [][]byte) {
	bitMap, emptyCells := getBitMap(board)
	for i := 0; i < len(emptyCells); i++ {
		bitMap.SetPossible(emptyCells[i])
	}
	//对所有可能的位置数组排序
	sort.Slice(emptyCells, func(i, j int) bool {
		return len(emptyCells[i].Possible) < len(emptyCells[j].Possible)
	})
	res := solve(board, emptyCells, 0, bitMap)
	if res {
		//对board进行真正填充
		for i := 0; i < len(emptyCells); i++ {
			p := emptyCells[i]
			x, y := p.X, p.Y
			n := p.Possible[p.cur]
			board[x][y] = byte(n + 49)
		}
	}
}

//从最小可能长度的位置开始填充
func solve(board [][]byte, possibleList []*EmptyCell, i int, bm *BitMap) bool {
	if i >= len(possibleList) {
		return true
	}
	p := possibleList[i]
	for j := 0; j < len(p.Possible); j++ {
		if bm.IsValid(p.X, p.Y, p.Possible[j]) {
			p.cur = j
			bm.AddBit(p.X, p.Y, p.Possible[j])
			if solve(board, possibleList, i+1, bm) {
				return true
			}
			bm.RemoveBit(p.X, p.Y, p.Possible[j])
		}
	}
	return false
}

//返回填充后的坐标和空的位置
func getBitMap(board [][]byte) (*BitMap, []*EmptyCell) {
	bitMap := &BitMap{}
	empty := []*EmptyCell{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n := board[i][j]
			if n == '.' {
				empty = append(empty, &EmptyCell{X: i, Y: j})
				continue
			}
			bitMap.AddBit(i, j, int(n)-49)
		}
	}
	return bitMap, empty
}