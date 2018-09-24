package main

import "fmt"

type Saiyan struct {
	Name   string
	Age    int
	father *Saiyan
}

type yo struct {
	yoname string
	yoage  int
}

func main() {
	a := NewSaiyan("王花花", 20)
	fmt.Println(a);
}

//工厂模式的
//可以批量产生新的Saiyan
func NewSaiyan(name string, Age int) *Saiyan {
	return &Saiyan{
		Name: name,
		Age:  Age,
	}
}
