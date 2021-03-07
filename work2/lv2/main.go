package main

import "fmt"

func Receiver(v interface{}){
	switch v.(type) {
	case string:
		fmt.Println("这是个string")
	case int:
		fmt.Println("这是个int")
	case bool:
		fmt.Println("这是个bool")
	}
}
func main()  {
	Receiver("你好吗")
	Receiver(32)
	Receiver(true)
}