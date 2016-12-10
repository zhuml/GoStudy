// GolangTest project main.go
package main

import (
	"fmt"
)

func main() {
	mapFun()
}

//第一个GO程序
func helloWorld() {
	fmt.Println("Hello World!")
}

//遍历字符串
func stringTest() {
	str := "你好 世界"
	ch := str[0]
	fmt.Printf("The first character of \"%s\" is \"%c\"\r\n", str, ch)
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}
}

//数组切片
func funSlice(baseArray bool) {
	if baseArray {
		//基于数组创建切片
		var myArray [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		var mySlice []int = myArray[:5]
		fmt.Println("Elements of myArray: ")
		for _, v := range myArray {
			fmt.Print(v, " ")
		}
		fmt.Println("\nElements of mySlice: ")
		for _, v := range mySlice {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		//创建一个初始元素个数为5的数组切片，元素初始值为0：
		mySlice1 := make([]int, 5)
		for i := 0; i < len(mySlice1); i++ {
			fmt.Println("mySlice1[", i, "] =", mySlice1[i])
		}
		fmt.Println()
		//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间：
		mySlice2 := make([]int, 5, 10)
		fmt.Println("len(mySlice2):", len(mySlice2)) //元素个数
		fmt.Println("cap(mySlice2):", cap(mySlice2)) //分配的空间大小
		fmt.Println()
		//追加元素1
		mySlice2 = append(mySlice2, 1, 2)
		fmt.Println("len(mySlice2):", len(mySlice2)) //元素个数
		fmt.Println("cap(mySlice2):", cap(mySlice2)) //分配的空间大小
		fmt.Println()
		//追加元素1
		tempSlice := []int{3, 4}
		mySlice2 = append(mySlice2, tempSlice...)
		fmt.Println("len(mySlice2):", len(mySlice2)) //元素个数
		fmt.Println("cap(mySlice2):", cap(mySlice2)) //分配的空间大小
		fmt.Println()
		for i := 0; i < len(mySlice2); i++ {
			fmt.Println("mySlice2[", i, "] =", mySlice2[i])
		}
		fmt.Println()
		//直接创建并初始化包含5个元素的数组切片：
		mySlice3 := []int{1, 2, 3, 4, 5}
		for i, v := range mySlice3 {
			fmt.Println("mySlice3[", i, "] =", v)
		}
		fmt.Println()
	}
}

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

//map
func mapFun() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	// 往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	// 从这个map查找键为"1234"的信息
	person, ok := personDB["12345"]
	// ok是一个返回的bool型，返回true表示找到了对应的数据
	if ok {
		fmt.Println("Found person", person.Name, "with ID ", person.ID, ".")
	} else {
		fmt.Println("Did not find person with ID 12345.")
	}
}
