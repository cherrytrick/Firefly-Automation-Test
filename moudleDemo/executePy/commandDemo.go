package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"golang.org/x/sync/errgroup"
)

// CmdPythonSaveImageDpi 命令行执行Python脚本示例，脚本必须返回"success"
//
// 不建议在业务方法的签名里填写WaitGroup等并发相关对象，那样就必须在并发场景使用
// 一个方法本身是无法决定自己被如何使用的。
func CmdPythonSaveImageDpi(args ...string) error {
	cmd := exec.Command("python", args...)
	var buffer bytes.Buffer
	// 单独提出来可以定制化
	// cmd.Stdin = nil
	cmd.Stdout = &buffer
	cmd.Stderr = &buffer
	// cmd.Dir = "."
	// cmd.Env = os.Environ()

	err := cmd.Run()
	if err != nil {
		return err
	}

	if bytes.HasPrefix(buffer.Bytes(), []byte("success")) {
		return nil
	}

	return fmt.Errorf("expect %q,actual %s", "success", buffer.Bytes())
}

func main() {
	// errgroup是WaitGroup的封装，具有高级功能
	eg, ctx := errgroup.WithContext(context.Background())

	// 可以自由地Go，不必处理Add和Done
	eg.Go(func() error {
		args := []string{"test.py"}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return CmdPythonSaveImageDpi(args...)
	})
	// 可以自由地Go，不必处理Add和Done
	eg.Go(func() error {
		args := []string{"test2.py"}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return CmdPythonSaveImageDpi(args...)
	})
	// 可以自由地Go，不必处理Add和Done
	eg.Go(func() error {
		args := []string{"test3.py"}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return CmdPythonSaveImageDpi(args...)
	})

	// 其中任何一个goroutine返回err，则其它goroutine不再执行
	// 这里的err就是最早出现的err，如果都没有出错则为nil
	err := eg.Wait()
	_ = err
}
