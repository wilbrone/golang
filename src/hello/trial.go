package main

import "fmt"

func main() {
	x := []int{8, 5, 4, 3, 9}
	fmt.Println(x)

	y := make([]int, 0, 100)
	y = append(y, 23)
	y = append(y, 44)
	y = append(y, 5)
	y = append(y, 10)
	fmt.Println(y)

	m := []int{45, 65, 30, 40}
	for i, _ := range m {
		fmt.Println(i, "-", m[i])
	}

	n := make([]int, 0, 10)
	n = append(n, 777)
	n = append(n, 555)
	for i, v := range n {
		fmt.Println(i, "-", v)
	}

	k := map[string]int{
		"Todd": 50,
		"Rick": 60,
		"Dev":  30,
	}
	for i, v := range k {
		fmt.Println(i, "-", v)
	}

	l := map[string]int{
		"Jack":  32,
		"Will":  90,
		"Alpha": 65,
	}
	for i, r := range l {
		fmt.Println(i, "-", r)
	}

	w := make(map[string]string)
	w["James"] = "777"
	fmt.Println(w["James"])
}
