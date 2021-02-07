package common

import "fmt"

var ch = make(chan string, 100)

func init() {
	for x := range ch {
		fmt.Println(x)
	}
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
func getMessage() string {
	return <-ch
}
