package main

import "fmt"

//定义Saiyan结构体
type Saiyan struct {
	Name string
	Age  int
}

func main() {
	goku := &Saiyan{Name: "王花花", Age: 20}
	Super(goku)
	fmt.Println(goku)
	goku.addAge();
	fmt.Println(goku)
}

func Super(s *Saiyan) {
	fmt.Println(s)
	//s = &Saiyan{"李栓蛋", 80};
	s.Age = 80;
	fmt.Println(s)
}

//在结构体上面添加方法
func (s *Saiyan) addAge() {
	s.Age += 10000;
}
