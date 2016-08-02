package main

import (
	"fmt"

	"github.com/auramirea/bar"
)

type x int

func (k x) add(v int) x {
	return k + x(v)
}

func main() {
	fmt.Println("Banana!")
	fmt.Println(bar.Bar("monday", "Aura"))
	fmt.Println("--------------------------------------------------")
	fmt.Println(bar.SumDiff(8, 7))
	fmt.Println(bar.SumDiff2(8, 7))
	fmt.Println("--------------------------------------------------")
	fmt.Println(bar.SumLoop(10))
	fmt.Println("--------------------------------------------------")
	bar.DeferExample()
	fmt.Println("--------------------------------------------------")
	bar.Slices()
	fmt.Println("--------------------------------------------------")
	bar.Maps()
	fmt.Println("--------------------------------------------------")
	u := x(10)
	fmt.Println(u.add(10))
	fmt.Println("--------------------------------------------------")
	fmt.Println(bar.Scope()())
	fmt.Println("--------------------------------------------------")
	bar.Test()
	fmt.Println("--------------------------------------------------")
	bar.TestStruct()
	fmt.Println("--------------------------------------------------")

}
