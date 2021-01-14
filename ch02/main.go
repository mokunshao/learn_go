package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		j int = 12
		k     = 13
	)
	var i int = 20
	fmt.Println(i, j, k)

	var f32 float32 = 2.2
	var f64 float64 = 2.28999
	fmt.Println(f32, f64)

	pi := &i
	fmt.Println(pi, *pi)
	i = 88
	fmt.Println(pi, *pi)

	const name = "love"

	const (
		one = 1
		two = 2
	)

	const (
		three = iota + 3
		four
		five
	)

	fmt.Println(three, four, five)

	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)

	i2f := float64(i)
	f2i := int(f64)
	fmt.Println(i2f, f2i)

	s1 := "china"
	fmt.Println(strings.HasPrefix(s1, "ch"))
	fmt.Println(strings.Index(s1, "hi"))
	fmt.Println(strings.ToUpper(s1))
}
