package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "Constant"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	fmt.Println("Hello world")

	for k := 0; k <= 3; k++ {
		fmt.Println(k)
		if k == 1 {
			fmt.Println("k es ", k, "!")
		} else {
			fmt.Println("---")
			fmt.Println(s)
		}
	}

	t := time.Now()
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's not weekend")
	}

	whatIAm := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatIAm(true)
	whatIAm(2)
	whatIAm("Hey")

	var twoD [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)

	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	fmt.Println("---Maps---")

	m := make(map[string]int)

	m["k1"] = 7

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	fmt.Println("---Pointers---")

	i := 1

	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)

	fmt.Println("---Struct---")

	fmt.Println(person{"Bob", 29})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "alice", age: 12})
	fmt.Println(newPerson("Jon"))

	d := person{name: "Sean", age: 50}
	fmt.Println(d.name)

	sp := &d
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println("---Methods---")

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	re := rect{width: 3, height: 4}
	circle := circle{radius: 5}

	measure(re)
	measure(circle)

	co := container{
		base: base{
			num: 1,
		},
		str: "name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num: ", co.base.num)

	fmt.Println("describe: ", co.describe())

	type describer interface {
		describe() string
	}

	var desc describer = co
	fmt.Println("describer: ", desc.describe())

	fmt.Println("---Goroutines---")

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//time.Sleep(time.Second)
	fmt.Println("done")

	fmt.Println("---Channel---")

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

}
