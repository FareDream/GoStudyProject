package HeavyBrush

import (
	"fmt"
	"os"
	"reflect"
)

func ChannelsTest() {
	a := make(chan int)           // 整数类型的无缓冲信道
	b := make(chan int, 0)        // 整数类型的无缓冲信道
	c := make(chan *os.File, 100) // 指向文件指针的带缓冲信道

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println("%T,%T,%T", a, b, c)
}
