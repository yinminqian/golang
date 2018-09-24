package main

import "fmt"

//组合结构体
func main() {
	goku := &Saiyan{
		Person: &Person{"王花花"},
		power:  80,
	}
	goku.Introduce();
	fmt.Println(goku.name)
	fmt.Println(goku.Person.name)
}

type Person struct {
	name string
}

func (p *Person) Introduce() {
	fmt.Println(p.name)
}

type Saiyan struct {
	*Person
	power int
}
