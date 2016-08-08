// Exfmt03.go
// 测试和说明：
// 整数部分
// %b 二进制表示
// %c 数值对应的Unicode编码字符
// %d 十进制表示
// %o 八进制表示
// %x 十六进制表示，使用a-f
// %X 十六进制表示，使用A-F
// %U Unicode格式： U+1234，等价于"U+%04X"
// %q 单引号
package main

import "fmt"

func main() {
        fmt.Printf("%b\n", 888)
        fmt.Printf("%c\n", '太')
        fmt.Printf("%d\n", 888)
        fmt.Printf("%o\n", 888)
        fmt.Printf("%x\n", 888)
        fmt.Printf("%X\n", 888)
        fmt.Printf("%U\n", 888)
        fmt.Printf("%q\n", 88)
}
// 输出：
//1101111000
//太
//888
//1570
//378
//378
//U+0378
//'X'
