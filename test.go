package main

import (
	"fmt"
)


func main() {
	var matrix [2][3]int
matrix[0] = [3]int{1,2,3}
matrix[1] = [3]int{4,5,6}

fmt.Println(matrix)

fmt.Println(matrix[1][2])

}