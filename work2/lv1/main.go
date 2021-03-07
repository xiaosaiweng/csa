package main

import "fmt"
//鸽子
type dove interface{
	gugugu()
}

//复读机
type repeater interface {
	repeat()
}
//柠檬精
type lemon interface {
	suan()
}
type ZhenXiangGuai interface {
	xiang()
}
type person struct {
	name string
	repeatNum int
}
func(p *person)gugugu(){
	fmt.Println(p.name,"又鸽了")
}
func(p *person)repeat(){
	fmt.Println(p.name,"重复了")
	p.repeatNum++
}
func(p *person)suan(){
	fmt.Println(p.name,"人酸了")
}
func (p *person) xiang(){
	fmt.Println(p.name,"人香了")
}
func main()  {
	p :=&person{
		name: "zas",
		repeatNum: 999,
	}
	p.gugugu()
	p.repeat()
	p.suan()
	p.xiang()
}