package common

import "fmt"

var ch chan string

func init() {
	ch = make(chan string, 100)
	go demo()
}

/*
	将消息放入队列
	filepath:文件路径
*/
func PutMessage(filepath string) {
	ch <- filepath
}

/*
	取出消息
*/
func GetMessage() string {
	return <-ch
}

func demo() {
	for x := range ch {
		fmt.Println(x)
	}
}
