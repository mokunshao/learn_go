package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])
	array2 := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(array2)
	array1 := [5]string{1: "b", 3: "d"}
	fmt.Println(array1)
	for i, v := range array1 {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
	for _, v := range array {
		fmt.Printf("对应值:%s\n", v)
	}

	array3 := [5]string{"a", "b", "c", "d", "e"}

	slice := array3[2:5]
	fmt.Println("前", array3)
	fmt.Println(slice)
	slice[2] = "z"
	fmt.Println("后", array3)
	fmt.Println(slice)

	slice1 := make([]string, 4)    // 长度为4
	slice2 := make([]string, 4, 8) // 长度为4，容量为8
	fmt.Println(slice1, slice2)

	slice3 := []string{"a", "b", "c", "d", "e"}

	fmt.Println(len(slice3), cap(slice3))

	//追加一个元素
	fmt.Println(append(slice1, "f"))

	//多加多个元素
	fmt.Println(append(slice1, "f", "g"))

	//追加另一个切片
	fmt.Println(append(slice1, slice...))

	nameAgeMap := make(map[string]int)
	nameAgeMap["飞雪无情"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22

	fmt.Println(nameAgeMap)
	nameAgeMap2 := map[string]int{"飞雪无情": 30}
	fmt.Println(nameAgeMap2)
	v, ok := nameAgeMap2["66"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no 66")

	}
	delete(nameAgeMap2, "飞雪无情")

	for k, v := range nameAgeMap {
		fmt.Println(k, ":", v)
	}

	s := "Hello飞雪无情"

	bs := []byte(s)

	fmt.Println(bs)

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println(s[0], s[1], s[15])

	for i, r := range s {
		fmt.Println(i, r,string(r))
	}

}
