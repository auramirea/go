package bar

import "fmt"

func Scope() func() int {
	outer_var := 2
	foo := func() int { return outer_var }
	return foo
}

func Test() {
	// create a slice from an array
	x := [3]string{"unu", "doi", "trei"}
	s := x[:] // a slice referencing the storage of x
	fmt.Println(s)
}

type Vertex struct {
	x int
	y int
}

func TestStruct() {
	v1 := Vertex{y: 1, x: 2}
	fmt.Println(v1.x, v1.y)
	// v2 and v1 will point to the same memory
	v2 := &v1
	v2.x = 1
	v2.y = 2
	fmt.Println(v1.x, v1.y)
}
