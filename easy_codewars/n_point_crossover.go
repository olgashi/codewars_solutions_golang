package main

import (
	"fmt"
	"sort"
)


func Crossover(ns []int, xs []int, ys []int) (xsOutput []int, ysOutput[]int) {
	xsOutput, ysOutput = []int{}, []int{}
	sort.Slice(ns, func(i, j int) bool { return ns[i] < ns[j] })
	ns = uniqueElements(ns)

	alternating1 := xs
	alternating2 := ys
	startIndex := 0
	endIndex := 0
	for _, xSIndex := range ns {
		endIndex = xSIndex
		xsOutput = append(xsOutput, alternating1[startIndex:endIndex]...)
		ysOutput = append(ysOutput, alternating2[startIndex:endIndex]...)

		alternating1, alternating2 = alternating2, alternating1
		startIndex = endIndex
	}
	xsOutput = append(xsOutput, alternating1[startIndex:]...)
	ysOutput = append(ysOutput, alternating2[startIndex:]...)

return
}

func main() {
	fmt.Println(Crossover([]int {1,3, 5}, []int {1,1,1,1,1,1,1}, []int {2,2,2,2,2,2,2}))
	fmt.Println(Crossover([]int {3,5,1,1,3}, []int {1,1,1,1,1,1,1}, []int {2,2,2,2,2,2,2}))
}

func uniqueElements(nums []int) (uniqueNums []int) {
  numsMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numsMap[num]; !ok {
			numsMap[num] = true
			uniqueNums = append(uniqueNums, num)
		}
	}
	return
}
/*
x,x,x,x,x,x,x => y,y,y,x,x,y,y

y,y,y,y,y,y,y,y => x,x,x,y,y,x,x

0, 3, 5



*/