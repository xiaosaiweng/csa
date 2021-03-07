package main

import "fmt"

type Author struct {
	Name string //名字
	VIP bool  //是否是高贵的带会员
	Icon string  //头像
	Signature string  //签名
	Focus int  //关注人数
}
type vodio struct {
	num int //点赞
	shoucangNum int //收藏
	toubiNum int //投币
}
func(v *vodio) dianzan(author Author) {
	fmt.Println(author.Name,"点赞")
	v.num++
}
func(v *vodio) shoucang(author Author) {
	fmt.Println(author.Name,"收藏")
	v.shoucangNum++
}
func(v *vodio) toubi(author Author) {
	fmt.Println(author.Name,"投币")
	v.toubiNum++
}
func(v *vodio) yijiansanlian(author Author) {
	fmt.Println(author.Name,"一键三连")
	v.num++
	v.shoucangNum++
	v.toubiNum++
}
func main(){
	a :=Author{
		Name: "zas",
		VIP: true,
		Icon: "zas",
		Signature: "zas",
		Focus: 999,
	}
	v :=&vodio{
		num: 999,
		shoucangNum: 999,
		toubiNum: 999,
	}
	v.dianzan(a)
	v.toubi(a)
	v.shoucang(a)
	v.yijiansanlian(a)
//aaa
}