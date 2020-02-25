package arithmetic

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/23 20:30
 */

func Rotate(matrix [][]int) {
	var (
		n    = len(matrix)
		temp int
	)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp = matrix[i][j]                     // 存第0行数
			matrix[i][j] = matrix[n-j-1][i]         // 第0行换成第0列
			matrix[n-j-1][i] = matrix[n-1-i][n-j-1] // 第0列换成第n行
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i] // 第n行换成第n列
			matrix[j][n-1-i] = temp                 // 第n列换成第0行
		}
	}
}
