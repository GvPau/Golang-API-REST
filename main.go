package main

import (
	"fmt" 
	"time"
)

const s string = "Constant"

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
}

