package main

import "fmt"

//type dove struct {
//	person
//	doveNum int
//}
//type repeater struct{
//	person
//	repeaterNum int
//}
//type person struct{
//	name string
//	age int
//	gender string
//}
//func main(){
//	d :=dove{
//		person: person{
//			name: "lmq",
//			age: 18,
//			gender: "male",
//		},
//		doveNum: 999,
//	}
//	d.gugugu()
//	fmt.Println(d.doveNum)
//}
//func(d *dove) gugugu(){
//	fmt.Println(d.name,"李梦琦")
//	d.doveNum++
//}
//func(r repeater) repeat(word string){
//	fmt.Println(word)
//	r.repeaterNum++
//}


//person 人类
type person struct {
	name string // 姓名
	age int // 年龄
	 gender string // 性别
}
//dove 鸽子
type dove interface {
	gugugu()
}
//repeater 复读机
type repeater interface {
	repeat(string)
}
//repeat 复读
func(p *person) repeat(word string){
	fmt.Println(word)
}
func(p *person)gugugu(){
	fmt.Println(p.name,"又鸽了")
}

func main(){
	p := &person{
		name: "lmq",
		age: 18,
		gender: "male",
	}
	p.gugugu()

	var r repeater

	r = p
	r.repeat("helloword")


}