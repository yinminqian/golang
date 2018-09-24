package main

import (
	"fmt"
	"os"
)

func main() {
	for a, arg := range os.Args {
		fmt.Println(a);
		fmt.Println(arg);
	}
	fmt.Println(len(os.Args));
}
