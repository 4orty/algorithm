package interview

import (
	"fmt"
	"math/rand"
)

// selection sort는 sorting 중에 제일 안좋은 거.
// 시간복잡도는 전부 n2이다.
// min 값 찾아서 switch

func Bubble() {
	a := []int{5,3,13,1,2,3,64,3,1,8,6}

	// 뒤에서부터 정렬임. 무조건 다 비교, 무조건 다 swap 시간복잡도 최악 같은 시간복잡도여도 더 느림.
	//
	//for i:=0; i<len(a); i++ {
	//	for j:=0; j<len(a)-i-1; j++ {
	//		if a[j] > a[j+1] {
	//			 temp := a[j]
	//			 a[j] = a[j+1]팁


	// 1.
	var efficiency int
	for i:=0; i<10; i++ {
		efficiency = rand.Intn(10)
		fmt.Println(efficiency)
	}

	// 2.
	for i:=0; i<10; i++ {
		readability := rand.Intn(10)
		fmt.Println(readability)
	}
}
