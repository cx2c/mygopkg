## func NewEncoding(encoder string) *Encoding

参数列表：

- encoder 64字节长的字符串，用做转换表

返回值：

- Encoding 结构指针

功能说明：

计算 base64 解码输出需要的最大字节数

代码实例：

```go
package examples

import (
    "fmt"
    "encoding/base64"
)

// 我们的编码表，64 字节
const encodeTest = "----------------------------------------------------------------"

func ExampleNewEncoding1() {

    enc := base64.NewEncoding(encodeTest)

    src := "this is a test string."
    dst := enc.EncodeToString([]byte(src))

    // 最后不足4字节的会用"="补全
    fmt.Println(dst)
    fmt.Println(len(dst) % 4 == 0)

    // Output:
    // ------------------------------==
    // true

}
```
