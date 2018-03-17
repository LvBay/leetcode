package array

import (
	"github.com/astaxie/beego"
)

// 1 2 3
// 4 5 6
// 7 8 9

// 1 4 7
// 2 5 8
// 3 6 9

// 7 4 1
// 8 5 2
// 9 6 3

// 5,1,9,11
// 2,4,8,10
// 13,3,6,7
// 15,14,12,16

// 5 2 13 15
// 1 4 3 14
// 9 8 6 12
// 11 10 7 16

func rotateImage(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	beego.Info(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= (len(matrix[i])-1)/2; j++ {
			beego.Info("j:", j, len(matrix[i])-j-1)
			matrix[i][j], matrix[i][len(matrix[i])-j-1] = matrix[i][len(matrix[i])-j-1], matrix[i][j]
		}
	}
}
