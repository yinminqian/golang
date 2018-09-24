package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	name, age := power("王花花")
	if age == 12 {
		fmt.Println(name,age)
	}
}

func log(name string) {

}

func add(name string) string {
	return "123"
}

func power(people string) (string, int) {
	return people, 12
}
