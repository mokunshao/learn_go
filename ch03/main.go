package main

import "fmt"

func main() {
	if i := 6; i > 10 {
		fmt.Println(">10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5~10")
	} else {
		fmt.Println("else...")
	}

	switch i := 6; {
	case i > 10:
		fmt.Println(">10")
	case i > 5 && i <= 10:
		fmt.Println("5~10")
	default:
		fmt.Println("else...")
	}

	switch i := 1; i {
	case 1:
		fmt.Println("1")
		fallthrough
	case 6:
		fmt.Println("6")
		fallthrough
	case 7:
		fmt.Println("7")
	case 8:
		fmt.Println("8")
	default:
		fmt.Println("else...")
	}

	switch 2 > 10 {
	case true:
		fmt.Println(true)
	case false:
		fmt.Println(false)
	}

	i := 1
	for i < 100 {
		i++
	}
	fmt.Println(i)
}
