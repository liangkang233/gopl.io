package main

import "fmt"

func main() {
	s := []string{"12", "ad", "df"}
	fmt.Println(k(s))
}

/*
有时候我们需要一个map或set的key是slice类型，但是map的key必须是可比较的类型，
但是slice并不满足这个条件。不过，我们可以通过两个步骤绕过这个限制。
第一步，定义一个辅助函数k，将slice转为map对应的string类型的key，
确保只有x和y相等时k(x) == k(y)才成立。然后创建一个key为string类型的map，
在每次对map操作时先用k辅助函数将slice转化为string类型。 */

// 转义不可直接比较的map的key
func k(list []string) string { return fmt.Sprintf("%v", list) }
