// Exfmt05.go
// 测试和说明：
// "%.*d",10,123: 打印整数，并保证10为长度，关键是长度运行时动态传入
// "%6.*f",2,888.666: 打印浮点数，并保证两位小数，关键是小数位运行时动态传入

package main

import "fmt"

func main() {
        fmt.Printf("%.*d\n",10,123)
        fmt.Printf("%6.*f\n",2,888.666)
}
