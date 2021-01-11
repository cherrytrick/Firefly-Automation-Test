package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

//命令行执行Python脚本示例
//执行python脚本
func CmdPythonSaveImageDpi(wg *sync.WaitGroup) (err error) {
	//args := []string{"main.py", filePath, newFilePath}
	//out, err := exec.Command("python", args...).Output()
	fmt.Println("开始执行")
	out, err := exec.Command("python", "test.py").Output()
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}
	result := string(out)
	if strings.Index(result, "success") != 0 {
		err = errors.New(fmt.Sprintf("main.py error：%s", result))
	}
	fmt.Println("执行结束")
	wg.Done()
	return
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go CmdPythonSaveImageDpi(&wg)

	wg.Wait()
}
