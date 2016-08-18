package somes

import "fmt"

//Say :some test
func Say() int {
	return 2
}

func Some() {
	sa := one{2, "my here"}
	fmt.Println("hello", sa.b)
}

type one struct {
	a int
	b string
}
