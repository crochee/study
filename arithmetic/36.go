package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/23 19:02
 */
func IsValidSudoku(board [][]byte) bool {
	var (
		columns [9]map[byte]bool
		boxes   [9]map[byte]bool
	)
	for i := 0; i < 9; i++ {
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		var rows = make(map[byte]bool)
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				if !columns[j][num] {
					columns[j][num] = true
				} else {
					return false
				}
				if !rows[num] {
					rows[num] = true
				} else {
					return false
				}
				if !boxes[(i/3)*3+j/3][num] {
					boxes[(i/3)*3+j/3][num] = true
				} else {
					return false
				}
			}
		}
	}
	return true
}
