package main;

import (
	"fmt"
)

func main() {
	//声明一个切片
	scores := []int{12, 123, 123}
	//长度为切片的长度,容量为底层数组的长度
	//mackcores是一个长度和容量都为10的切片
	mackcores := make([]int, 10)
	// testscore为一个长度为0,容量为10的切片
	testscores := make([]int, 0, 10)
	//如果切片的长度为0,而底层数组有容量的话,那么需要显示的拓展切片,才能进行赋值
	scores_ := testscores[0:8]

	fmt.Println(scores_);
	fmt.Println(testscores)
	fmt.Println(scores)
	fmt.Println(len(mackcores));
}
