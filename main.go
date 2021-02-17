package main

import (
	"fmt"

	"./konovalov"
)

func main() {
	fmt.Println(konovalov.SolutionSquareGenerator(5, 5))
	fmt.Println(konovalov.SolutionBinaryGap(257))
	fmt.Println(konovalov.SolutionRotate([]int{3, 8, 9, 7, 6}, 4))
}
