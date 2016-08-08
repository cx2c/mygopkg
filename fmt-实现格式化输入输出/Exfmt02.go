// Exfmt02.go
// 测试和说明：
// %t 值的true和false
// %T 值的类型在Go语言中的表示
// %% 打印一个 %
// %q 给打印的字串自动加引号
package main

import "fmt"

func main() {
        var Yes bool //bool 变量自动初始化为 false
        fmt.Printf("%t\n", Yes)
        fmt.Printf("%T\n", Yes)
        fmt.Printf("%%\n")
        fmt.Printf("%q\n","自动加引号")
}
// 输出：
//false
//bool
//%
//"自动加引号"
