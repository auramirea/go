package bar

import "fmt"

func Bar(day, name string) string {
	return fmt.Sprintf("Happy strawberry %s, %s!", day, name)
}

func SumDiff(x, y int) (int, int) {
	return x + y, x - y
}

func SumDiff2(x, y int) (sum, diff int) {
	sum = x + y
	diff = x - y
	return
}
func SumLoop(count int) (sum int) {
	for i := 0; i < count; i++ {
		sum += i
	}
	return
}

func DeferExample() {
	defer fmt.Println("This will be executed at the end of the method")
	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	defer fmt.Println("bla bla")
	defer func() {
		x := recover()
		fmt.Println(x)
		panic("Hallo2")
	}()
	panic("Hallo")
	fmt.Println("I'm the first statement")
}

func Slices() {
	x := make([]int, 4)
	x[0] = 1
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(fmt.Sprintf("Capacity is %d", cap(x)))
	x = x[0:2:3]
	fmt.Println(fmt.Sprintf("Capacity is %d", cap(x)))
	for _, v := range x {
		fmt.Println(v)
	}
	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i)
	}
	for _, v := range y {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()

	var slice1, slice2 = x[0:1], x[0:1]
	slice1[0] = 10
	for _, v := range slice2 {
		fmt.Println("Slices are just views on arrays as long as they don't get reallocated:", v)
	}
}

func Maps() {
	x := map[int]string{
		2: "two",
		3: "three",
	}
	y := make(map[int]string)
	y[1] = "one"
	fmt.Println(x)
	fmt.Println(y)
	delete(y, 1)
	fmt.Println(y)
}
