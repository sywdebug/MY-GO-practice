package main

import "fmt"

func main() {
	var arr = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Println()
	}

	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		fmt.Printf("第%d班级的总分为%v，平均分为%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
}
