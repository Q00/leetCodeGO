package array

/*
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
Example 2:

Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/

// Rotate90 array 90 degree => 대칭으로 해결
func Rotate90(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			// a_n,n을 중심으로 대칭, 행 열 자체가 바뀜
			temp := matrix[j][i]
			// temp2:= matrix[len(matrix)-1-i][j]
			// temp3:= matrix[len(matrix)-1-i][len(matrix)-1-i]
			// temp4:= matrix[j][len(matrix)-1-i]
			// fmt.Println(temp1)
			// fmt.Println(temp2)
			// fmt.Println(temp3)
			// fmt.Println(temp4)
			// fmt.Println("~~~~~~~~~~~")
			matrix[j][i] = matrix[i][j]
			// matrix[j][len(matrix)-1-i] = temp1
			// matrix[len(matrix)-1-i][len(matrix)-1-i] = temp4
			// matrix[len(matrix)-1-i][j] = temp3
			matrix[i][j] = temp

		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			// n*n 행렬일때 a_i,j => a_i,n-j-1 같은 행안에서 열만 바뀜
			temp := matrix[i][len(matrix)-1-j]
			matrix[i][len(matrix)-1-j] = matrix[i][j]
			matrix[i][j] = temp
		}
	}
}
