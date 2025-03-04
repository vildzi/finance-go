package main

import (
	"fmt"
	"github.com/vildzi/finance-go/equity"
)

func main() {
	eq, err := equity.Get("AAPL")
	fmt.Println(eq)
	fmt.Println(err)
}
