## type InternalBenchmark

结构：

	type InternalExample struct {
	    Name   string
	    F      func()
	    Output string
	}

类型说明：

一个内部类型，只为了跨包使用而成为可导出的；它是“go test”命令实现的一部分。

代码实例：

	package main

	import (
		"fmt"
		"testing"
	)

	func main() {
		iExample := testing.InternalExample{Name: "Example", F: func() {}, Output: ""}
	    fmt.Printf("Example: %v\n", iExample)
	}
