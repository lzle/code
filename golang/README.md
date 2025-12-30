## 目录

* [环境](#环境)
    * [配置](#配置) 
    * [Hello World](#hello-world) 
    * [编译](#编译) 
* [命名规范](#命名规范)
* [注释](https://go.dev/doc/comment)
* [基础类型](#基础类型)
    * [整数](#整数)
    * [浮点数](#浮点数)
    * [布尔值](#布尔值)
    * [字符串](#字符串)
* [复合类型](#复合类型)
    * [数组](#数组)
    * [切片](#切片)
    * [字典](#字典)
* [常量和枚举](#常量和枚举)
* [变量定义](#变量定义)
    * [全局变量](#全局变量)
    * [局部变量](#局部变量)
    * [零值](#零值)
* [流程控制](#流程控制)
    * [for](#for)
    * [select](#select)
    * [switch](#switch)
    * [for select](#for-select)
    * [for switch](#for-switch)
* [通道](#通道)
* [函数](#函数)
* [接口](#接口)
	* [接口类型](#接口类型)
	* [接口实现](#接口实现)
	* [接口值](#接口值)
	* [类型断言](#类型断言)
	* [类型分支](#类型分支)
* [比较操作](doc/comparison/README.md)
* [断言、类型转换](doc/conversion/README.md)
* [并发安全](#并发安全)
    * [竞态](#竞态)
    * [sync.Mutex](#互斥锁syncmutex)
    * [sync.RWMutex](#读写互斥锁syncrwmutex)
    * [sync.Once](#延迟初始化synconce)
    * [sync.Cond](#条件变量synccond)
    * [sync.Pool](#对象池syncpool)
    * [race detector](#静态检测race-detector)
    * [sync/atomic](#原子操作syncatomic)
* [协程](#协程)
    * [context](#context)
    * [gopkg.in/tomb.v2](#gopkg.in/tomb.v2)
* [gRPC](doc/grpc/README.md)
* [标准库](#标准库)
    * [time](#time)
    * [fmt](#fmt)
    * [flag](#flag)
    * [log](#log)
* [Tools](#Tools)
    * [gofmt](#gofmt)
    * [workerpool](#workerpool)
    * [bytebufferpool](#bytebufferpool)
    * [delve](#delve)
* [Test](#test)
    * [功能测试](#功能测试)
    * [黑盒、白盒测试](#黑盒白盒测试)
    * [覆盖率](#覆盖率)
    * [基准测试](#基准测试)
    * [示例函数](#示例函数)
* [Benchmark](#Benchmark)
* [PProf](#doc/pprof/README.md)
* [其他](#其他)
    * [gopls&staticcheck](#gopls--staticcheck)
    * [error规范](#error-规范)
    * [正则规范](#正则规范)
    * [text/template](#texttemplate)
    * [默认值slice&map&chan for循环](#默认值-slicemapchan-for循环)
    * [httputil](#httputil)
    * [优雅的recover](#优雅的-recover)
* [相关链接](#相关链接)


## 环境

`golang` 环境安装参考官网 [https://go.dev/doc/install](https://go.dev/doc/install)

### 配置

获取环境变量。

```shell
$ go env
AR='ar'
CC='cc'
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_ENABLED='1'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
CXX='c++'
GCCGO='gccgo'
GO111MODULE=''
GOAMD64='v1'
GOARCH='amd64'
GOAUTH='netrc'
GOBIN=''
GOCACHE='/Users/连长/Library/Caches/go-build'
GOCACHEPROG=''
GODEBUG=''
GOENV='/Users/连长/Library/Application Support/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFIPS140='off'
GOFLAGS=''
GOGCCFLAGS='-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -ffile-prefix-map=/var/folders/94/1mhpmp9j2fnbbgtt4p308t440000gn/T/go-build1668721803=/tmp/go-build -gno-record-gcc-switches -fno-common'
GOHOSTARCH='amd64'
GOHOSTOS='darwin'
GOINSECURE=''
GOMOD='/Users/连长/Code/Go/src/grpc-go/examples/go.mod'
GOMODCACHE='/Users/连长/Code/Go/pkg/mod'
GONOPROXY=''
GONOSUMDB=''
GOOS='darwin'
GOPATH='/Users/连长/Code/Go'
GOPRIVATE=''
GOPROXY='https://goproxy.cn,direct'
GOROOT='/usr/local/Cellar/go/1.24.1/libexec'
GOSUMDB='sum.golang.org'
GOTELEMETRY='local'
GOTELEMETRYDIR='/Users/连长/Library/Application Support/go/telemetry'
GOTMPDIR=''
GOTOOLCHAIN='auto'
GOTOOLDIR='/usr/local/Cellar/go/1.24.1/libexec/pkg/tool/darwin_amd64'
GOVCS=''
GOVERSION='go1.24.1'
GOWORK=''
PKG_CONFIG='pkg-config'
```

查看指定变量。

```shell
$ go env GOPROXY
https://proxy.golang.org,direct
```

配置 `GOPROXY` 为国内镜像（推荐）

```shell
$ go env -w GOPROXY=https://goproxy.cn,direct

$ go env GOPROXY
https://goproxy.cn,direct
```

### Hello World

执行下面代码，输出了"Hello World!"。
```go
package main

import "fmt"

func main() {
    /* 这是我的第一个简单的程序 */
    fmt.Println("Hello, World!")
}
```

执行

```shell
go run hello.go

Hello, World!
```


## 命名规范

### 包名

包名应该尽量简短且有意义，使用小写字母，不使用下划线或驼峰命名。包名通常与包所在的目录名称相同。

```go
package http
```

### 文件名

文件名应该使用小写字母，单词之间可以使用下划线分隔。文件名通常以 `.go` 结尾。

```go
net/http/
├── client.go
├── client_test.go
├── tranfer.go     
├── server.go 
├── request.go     
├── response.go     
├── chunked_writer.go 
```

### 函数名

函数名应该使用驼峰命名法，首字母大写表示导出函数，首字母小写表示非导出函数。

```go
// net/http/request.go
func NewRequest(method, url string, body io.Reader) (*Request, error) {
	// 导出函数
}

func removeZone(host string) string {
    // 非导出函数
}
```

导出和非导出是指函数是否可以被其他**包**访问。例如在 `net/http/request.go` 中调用 `net/http/transfer.go` 文件 `readTransfer` 函数，
函数首字母小写。

```go
// net/http/request.go
func readRequest(b *bufio.Reader) (req *Request, err error) {
    // 调用非导出函数
    if err := readTransfer(b, req); err != nil {
        return nil, err
    }
    // ...
}

// net/http/transfer.go
func readTransfer(b *bufio.Reader, req *Request) error {
    // 读取传输数据
    // ...
    return nil
}
```

Another short example is once.Do; once.Do(setup) reads well and would not be improved by writing once.DoOrWaitUntilDone(setup). Long names don't automatically make things more readable. A helpful doc comment can often be more valuable than an extra long name.

### 变量名、常量名

变量名和常量名应该使用驼峰命名法，根据是否需要导出决定首字母大小写。

```go
// fmt
const (
	ldigits = "0123456789abcdefx"
	udigits = "0123456789ABCDEFX"
)

// net/http
const (
    StatusOK                   = 200
    StatusCreated              = 201
    StatusAccepted             = 202
    StatusNonAuthoritativeInfo = 203
    StatusNoContent            = 204
    StatusResetContent         = 205
    StatusPartialContent       = 206
)
```

用变量还是常量取决于值在运行时是否可以变更，全局常量和变量不用大写。

不建议用类似 MAX_COUNT、HTTP_STATUS_OK 这类全大写加下划线的写法，Go 社区风格不喜欢这种风格。习惯采用驼峰式，并在需要时保持缩写时全大写

### 结构体

结构体名应该使用驼峰命名法，和变量名命名规则一致，根据是否需要导出决定首字母大小写，包括结构体内的变量。

```go
// net/http/request.go
type Request struct {
    Method string 
    URL    string 
    Header map[string]string
    Body   io.Reader
}
```

### 接口名

接口名通常是一个描述动作或能力的动词 + “-er” 结尾，如 Reader 表示“能读取”，Writer 表示“能写入”。

```go
type requestTooLarger interface {
    requestTooLarge()
}
```

## 基础类型

### 整数

Go 语言提供了以下整数类型：

- int8
- int16
- int32
- int64
- uint8
- uint16
- uint32
- uint64
- int
- uint
- uintptr

其中 uint 和 int 的大小是取决于操作系统的位数，32 位操作系统上是 32 位，64 位操作系统上是 64 位。
uintptr 是一个无符号整数类型，用于存放一个指针，大小也是取决于操作系统的位数。

int 是有符号整数类型，取值范围是 -2^31 到 2^31-1，即 -2147483648 到 2147483647。

除了上面基础数据类型外，还有特殊的整数类型：

- byte：type byte = uint8
- rune：type rune = int32

byte 和 rune 类型是为了更好的表示字符类型，byte 类型一般用于存储 ASCII 码，rune 类型一般用于存储 Unicode 码。

```go
package main

import (
    "fmt"
)

func main() {
    // 使用 byte 存储 ASCII 码
    var b byte = 'A'
    fmt.Printf("byte: %c, ASCII: %d\n", b, b)

    // 使用 rune 存储 Unicode 码
    var r rune = '你'
    fmt.Printf("rune: %c, Unicode: %U\n", r, r)
}

// byte: A, ASCII: 65
// rune: 你, Unicode: U+4F60
```

### 浮点数

Go 语言提供了以下浮点数类型：

- float32：32 位浮点数，IEEE-754 标准，取值范围约为 1.4E-45 到 3.4E+38，精度约为 6-9 位十进制数。
- float64：64 位浮点数，IEEE-754 标准，取值范围约为 4.9E-324 到 1.8E+308，精度约为 15-17 位十进制数

### 布尔值

Go 语言提供了布尔类型 bool，取值为 true 或 false。

### 字符串

字符串是⼀个不可改变的字节序列，和数组不同的是，字符串的元素不可修改，是⼀个只读的字节数组。

```go
package main

import (
    "fmt"
)

func main() {
    str := "abc"
    fmt.Println(str[0])
    str[1] = 'd' // panic: cannot assign to str[1] (neither addressable nor a map index expression)
    fmt.Println(str)
}
```

Go 语⾔字符串的底层结构在 `reflect.StringHeader` 中定义： 

```go
type StringHeader struct {
    Data uintptr
    Len  int
}
```

字符串结构由两个信息组成：第⼀个是字符串指向的底层字节数组，第⼆个是字符串的字节的⻓度。字符串其实是⼀个结构体，
因此字符串的赋值操作也就是 `reflect.StringHeader` 结构体的复制过程，并不会涉及底层字节数组的复制。

```go
import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    str := "Hello, World!"
    strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
    fmt.Printf("String: %s\n", str)
    fmt.Printf("Data pointer: %x\n", strHeader.Data)
    fmt.Printf("Length: %d\n", strHeader.Len)

    newStr := str
    newStrHeader := (*reflect.StringHeader)(unsafe.Pointer(&newStr))
    fmt.Printf("String: %s\n", newStr)
    fmt.Printf("Data pointer: %x\n", newStrHeader.Data)
    fmt.Printf("Length: %d\n", newStrHeader.Len)

    // String: Hello, World!
    // Data pointer: de0dae3
    // Length: 13
    // String: Hello, World!
    // Data pointer: de0dae3
    // Length: 13
}
```

字符串和其子串（[:4]）可以安全的共用数据，子串生成和字符串复制操作开销都很低，不会产生额外的内存分配。

```go
import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    str := "Hello, World!"

    newStr := str[:]
    fmt.Println(reflect.TypeOf(newStr))

    strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
    fmt.Printf("Data pointer: %x\n", strHeader.Data)
    newStrHeader := (*reflect.StringHeader)(unsafe.Pointer(&newStr))
    fmt.Printf("Data pointer: %x\n", newStrHeader.Data)

    // string
    // Data pointer: d6e3c4d
    // Data pointer: d6e3c4d
}
```

字符串的长度可以通过内置函数 `len()` 获取，字符串的索引从 0 开始，字符串的值是 Unicode 码点，而不是字节。

可以通过下面方式对字符串进行遍历。

```go
package main

import (
    "fmt"
)

func main() {
    str := "Hello, 世界"

    fmt.Println(len(str))  
    // 13

    // 使用索引遍历字符串
    for i := 0; i < len(str); i++ {
        fmt.Printf("字符: %c, Unicode: %U\n", str[i], str[i])
    }
    // 字符: H, Unicode: U+0048
    // 字符: e, Unicode: U+0065
    // 字符: l, Unicode: U+006C
    // 字符: l, Unicode: U+006C
    // 字符: o, Unicode: U+006F
    // 字符: ,, Unicode: U+002C
    // 字符:  , Unicode: U+0020
    // 字符: ä, Unicode: U+00E4
    // 字符: ¸, Unicode: U+00B8
    // 字符: , Unicode: U+0096
    // 字符: ç, Unicode: U+00E7
    // 字符: , Unicode: U+0095
    // 字符: , Unicode: U+008C
    
    // 使用 range 遍历字符串
    for _, r := range str {
        fmt.Printf("字符: %c, Unicode: %U\n", r, r)
    }
    // 字符: H, Unicode: U+0048
    // 字符: e, Unicode: U+0065
    // 字符: l, Unicode: U+006C
    // 字符: l, Unicode: U+006C
    // 字符: o, Unicode: U+006F
    // 字符: ,, Unicode: U+002C
    // 字符:  , Unicode: U+0020
    // 字符: 世, Unicode: U+4E16
    // 字符: 界, Unicode: U+754C
}
```

字符串可以和字节 `slice` 相互转换，概念上，转换之间会分配新的内存，进行数据拷贝，维持了字符串的不可变性：

```go
s := "hello"
b := []byte(s)
s2 := string(b)
```

另外 `bytes.Buffer` 类型是一个可变大小的字节缓冲区，可以用来缓存字符串，可以节省字符串拼接造成的内存开销。

```go
func intsToString(values []int) string { 
    var buf bytes.Buffer 
    buf.WriteByte('[') 
    for i, v := range values {
         if i > 0 { buf.WriteString(", ") 
    }
    fmt.Fprintf(&buf, "%d", v) }
    buf.WriteByte(']') 
    return buf.String() 
}
    
func main() { 
    fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]" 
}
```

注：4个标准包对字符串的操作特别重要：`bytes`、`strings`、`strconv` 和 `unicode` 。


## 复合类型

### 数组

数组是具有固定长度且拥有零个或者多个相同数据类型元素的序列，数组的长度固定。

数组索引从 0 开始，新数组中的元素初始值为元素类型的零值。

```go
package main

import (
    "fmt"
)

func main() {
    var a [3]int
    for i, v := range a {
        fmt.Println(i, v)
    }
    // 0 0
    // 1 0
    // 2 0

    var b [3]int = [3]int{1, 2}
    fmt.Println(b)
    // [1 2 0]

    var c [2]int 
    fmt.Println(c) // 不是 nil
    // [0 0]
}
```

数组的元素可以通过索引访问，进行修改。

```go
a := [3]int{1, 2, 3}
a[0] = 4
fmt.Println(a)
// [4 2 3]
```

如果省略号 "..." 出现在数组长度的位置，那么数组的长度由初始化数组的元素个数决定。

```go
q := [...]int{1, 2, 3}
```

数组的长度是数组类型的一部分，不同长度的数组是两种不同的数据类型。

```go
q := [3]int{1, 2, 3}
q = [4]int{1, 2, 3, 4} // 编译错误：不可以将[4]int 赋值给[3]int
```

数组的长度必须是常量表达式（值在程序编译时就可以确定）。

```go
package main

import (
    "fmt"
)

func main() {
    var length = 10
    var a [length]int // 编译错误：non-constant array bound length
    fmt.Println(a)

    var arr2 [getLength()]int // 编译错误：non-constant array bound getLength()
    fmt.Println(arr2)
}

func getLength() int {
    return 5
}
```

如果一个数组的元素类型是可比较的，那么这个数组也是可以比较的。需要注意的是如果数组的长度不同，数组比较会引发编译错误。

```go
package main

import (
    "fmt"
)

func main() {
    var a [2]int = [2]int{1, 2}
    var b [2]int = [2]int{1, 2}
    fmt.Println(a == b) // true

    var c [3]int = [3]int{1, 3}
    fmt.Println(a == c) // invalid operation: a == c (mismatched types [2]int and [3]int)
}
```

需要注意的是，**数组作为函数参数传递时，进行值传递**，会创建一个副本，传递大的数组会变得很低效。可以显式的传递指针。同理，数组赋值也是值拷贝。

```go
package main

import (
    "fmt"
)

func copyArray(arr [5]int) {
    fmt.Println(arr)
}

func copyArrayPointer(arr *[5]int) {
    fmt.Println(*arr)
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("原始数组:", arr)

    // 值传递
    copyArray(arr)

    // 指针传递
    copyArrayPointer(&arr)

    arr2 := arr
    arr2[0] = 100 // 修改 arr2 不会影响 arr
    fmt.Println("原始数组:", arr)
}

// 原始数组: [1 2 3 4 5]
// 原始数组: [1 2 3 4 5]
```

可以⽤ for 循环来迭代数组。下⾯常⻅的⼏种⽅式都可以⽤来遍历数组：

````go
for i := range a { 
    fmt.Printf("a[%d]: %d\n", i, a[i]) 
}

for i, v := range b { 
    fmt.Printf("b[%d]: %d\n", i, v) 
}

for i := 0; i < len(c); i++ { 
    fmt.Printf("c[%d]: %d\n", i, c[i]) 
}
````

### 切片

切⽚是动态数组，结构定义 `reflect.SliceHeader`：

```go
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```

看看切⽚的定义⽅式：

```go
package main

import (
    "fmt"
)

func main() {
    var (
        a []int             // nil切⽚, 和nil相等
        b = []int{}         // 空切⽚, 和nil不相等
        c = []int{1, 2, 3}  // 有3个元素的切⽚
        d = c[:2]           // 有2个元素的切⽚
        g = make([]int, 3)      // 有3个元素的切⽚, len和cap都为3
        i = make([]int, 0, 3)   // 有0个元素的切⽚, len为0, cap为3
    )
    fmt.Println(a, b, c, d, g, i)
}
```

遍历切⽚的⽅式和遍历数组的⽅式类似：

````go
for i := range a { 
    fmt.Printf("a[%d]: %d\n", i, a[i]) 
}

for i, v := range b { 
    fmt.Printf("b[%d]: %d\n", i, v) 
}

for i := 0; i < len(c); i++ { 
    fmt.Printf("c[%d]: %d\n", i, c[i]) 
}
````

切片可以通过 `len()` 和 `cap()` 函数获取切片的长度和容量。

```go
package main

import (
    "fmt"
)

func main() {
    // 创建一个切片
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println("切片:", slice)

    // 获取切片的长度
    length := len(slice)
    fmt.Println("切片的长度:", length)

    // 获取切片的容量
    capacity := cap(slice)
    fmt.Println("切片的容量:", capacity)
}

// 切片: [1 2 3 4 5]
// 切片的长度: 5
// 切片的容量: 5
```

还可以通过 `append()` 函数向切片中追加元素。

```go
package main

import (
    "fmt"
)

func main() {
    // 创建一个切片
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println("原始切片:", slice)

    // 向切片中追加一个元素
    slice = append(slice, 6)
    fmt.Println("追加一个元素后的切片:", slice)
}

// 原始切片: [1 2 3 4 5]
// 追加一个元素后的切片: [1 2 3 4 5 6]
```

需要注意的是切片在并发添加元素时，可能会导致**数据竞争**，需要用锁机制保障。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    slice := make([]int, 0)

    for i := 0; i < 10000; i++ {
        go func(i int) {
            slice = append(slice, i)
        }(i)
    }

    time.Sleep(time.Second)

    fmt.Println(len(slice)) 

    // Output: 8812
}
```

切片是不可比较的，不能使用 `==` 操作符来判断两个切片是否相等，唯一的例外是与 nil 进行比较，你可以检查一个切片是否为 nil。

赋值和函数传递时，切片是引用传递，传递的是切片的引用，修改切片的元素会影响到原切片。

如果想检查一个 slice 是否为空，可以使用 `len(s) == 0` 而不是 `s == nil`。

```go
var s []int 	// len(s) == 0, s == nil
s = nil         // len(s) == 0, s == nil
s = []int(nil)  // len(s) == 0, s == nil
s = []int{} 	// len(s) == 0, s != nil
```

注意 `slice` 为 `nil` 时，也可以进行切片，`len()` 和 `cap()`、`append()` 操作。

```go
package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3, 4, 5}
    a  = nil
    // var a []int

    b := a[:0]
    fmt.Println(b)
    fmt.Println(len(a))

    // nil slice append
    a = append(a, 1)
    fmt.Println(a)
}
```


### 字典

字典是一种无序的集合，是由键值对组成的。键的类型必须是可以通过操作符 `==` 来进行比较的数据类型，值的类型没有限制，
虽然浮点数类型也可以作为键，但是比较浮点数的相等性不是一个好主意。

```go
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}

ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```

上面两种创建字典的方式是等价的。

如果获取的键不存在，返回值是该类型的零值；删除不存在的值也是安全的；快捷赋值方式（x+=1，y++）对字典中的元素也是使用的。

```go
import (
    "fmt"
)

func main() {
    ages := map[string]int{
        "alice":   31,
        "charlie": 34,
    }

    // 不存在的 key 返回 value 类型零值
    fmt.Println(ages["bob"]) // "0"

    // 删除不存在的值也是安全的
    delete(ages, "bob") // 删除元素

    // 语法糖，检查是否存在并赋值
    ages["bob"]++
    fmt.Println(ages["bob"]) // "1"
}
```

`map` 类型的零值是 `nil`，对于 `nil` 的 `map` 不能进行赋值操作，但是可以进行读取操作。

```go
var ages map[string]int
fmt.Println(ages == nil) // "true"
fmt.Println(len(ages) == 0) // "true"

ages["bob"] = 1 // panic: assignment to entry in nil map
```

判断元素是否存在的更优雅的方式如下：

```go
age, ok := ages["bob"]
if !ok {
    fmt.Println("bob not found")
}
```

和切片一样，字典也是引用传递，函数传参和赋值操作，修改字典的元素会影响到原字典。

字典是不可比较的，唯一合法的就是和 nil 进行比较。为了判断两个字典是否拥有相同的键和值，需要遍历字典。

```go
func equal(a, b map[string]int) bool {
    if len(a) != len(b) {
        return false
    }

    for k, av := range a {
        if bv, ok := b[k]; !ok || bv != av {
            return false
        }
    }

    return true
}
```

字典多线程操作是不安全的，下面操作会导致 panic。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	// 启动 10 个 goroutine 并发写 map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			m[key] = key // 并发写操作
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}

// fatal error: concurrent map writes
```

## 常量和枚举

常量是在编译时就确定的值，不可更改，Go 语言中的常量可以是字符、字符串、布尔值或数值。

```go
const (
    class1 = 0
    class2 // class2 = 0
    class3 = iota  //iota is 2, so class3 = 2
    class4 // class4 = 3
    class5 = "abc" 
    class6 // class6 = "abc"
    class7 = iota // class7 is 6
)
```

## 变量定义

### 零值

数值：所有数值类型的零值都是 0

布尔值：零值是 false

字符串：零值是空字符串 ""

指针：var a *int 零值是 nil

切片：var a []int 零值是 nil

map：var a map[string] int 零值是 nil

函数：var a func(string) int 零值是 nil

channel：var a chan int 零值是 nil

接口：var a interface_type 零值是 nil

结构体: var instance StructName 结构体里每个 field 的零值是对应 field 的类型的零值


## 流程控制

### for

`for` 是 Go 语言中唯一的循环语句，功能强大。

经典的三段式：

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

只带条件的循环：

```go
for i < 10 {
    fmt.Println(i)
    i++
}
```

无限循环：

```go
for {
    fmt.Println("Infinite Loop")
}
```

遍历切片、数组、字符串等：

```go
for i := range a { 
    fmt.Printf("a[%d]: %d\n", i, a[i]) 
}

for i, v := range b { 
    fmt.Printf("b[%d]: %d\n", i, v) 
}
```

遍历 map:

```go
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    fmt.Println(k, v)
}

for k := range m {
	fmt.Println(k)
}
```

### select 

`select` 支持关键字 `break`、`goto` 和 `return`，无法使用 `continue`。

`break` 用于退出 `select` 语句，`return` 用于退出函数。

```go
import (
	"fmt"
	"time"
)

func main() {
    timer := time.NewTicker(1 * time.Second)
    complete := make(chan int, 1)

    select {
    case <- timer.C:
        fmt.Println("Timeout")
        if a := 1; a == 1 {
            fmt.Println("Break Case")
            break
        }
    case <- complete:
        fmt.Println("Complete")
        return
    }

    fmt.Println("Done")
}

// Timeout
// Break Case
// Done
```

`goto` 用于跳转到 `select` 语句外部标签，不能跳转到 `select` 内部。

```go
import (
	"fmt"
	"time"
)

func main() {
    timer := time.NewTicker(1 * time.Second)
    complete := make(chan int, 1)

    select {
    case <- timer.C:
        fmt.Println("Timeout")
        if a := 1; a == 1 {
            fmt.Println("Break Case")
            goto Break
        }
    case <- complete:
        fmt.Println("Complete")
        return
    }

    Break:
        fmt.Println("Break Label")
    
    fmt.Println("Done")
}

// Timeout
// Break Case
// Break Label
// Done
```

### switch 

`switch` 支持关键字 `break`、`goto`、`fallthrough` 和 `return`，无法使用 `continue`。

`break` 用于退出 `switch` 语句，`return` 用于退出函数。

```go
import (
	"fmt"
)

func main() {
    x := 1
    switch x {
    case 1:
        fmt.Println("x is 1")
        if x == 1 {
            fmt.Println("Break out of switch")
            break
        }
        fmt.Println("Do something")
    case 2:
        fmt.Println("x is 2")
        return
    case 3:
        fmt.Println("x is 3")
    }
    fmt.Println("Done")
}

// x is 1
// Break out of switch
// Done
```

`goto` 用于跳转到 `switch` 语句外部标签，不能跳转到 `switch` 内部。

```go
import (
    "fmt"
)

func main() {
    x := 1
    switch x {
    case 1:
        fmt.Println("x is 1")
        if x == 1 {
            fmt.Println("Break out of switch")
            goto Break
        }
        fmt.Println("Do something")
    case 2:
        fmt.Println("x is 2")
        return
    case 3:
        fmt.Println("x is 3")
    }

    Break:
        fmt.Println("Break Label")

    fmt.Println("Done")

// x is 1
// Break out of switch
// Break Label
// Done
}
```

`fallthrough` 用于执行下一个 `case` 语句，不会判断下一个 `case` 的条件。

```go
import (
    "fmt"
)

func main() {
    x := 1
    switch x {
    case 1:
        fmt.Println("x is 1")
        fallthrough
    case 2:
        fmt.Println("x is 2")
    case 3:
        fmt.Println("x is 3")
    }

    fmt.Println("Done")
}

// x is 1
// x is 2
// Done
```

### for select 

`continue` 用于跳过当前 `for` 循环，执行下一次循环。

```go
import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTicker(10 * time.Second)
    complete := make(chan int, 1)

    for {
        fmt.Println("Start")
        select {
        case <- timer.C:
            fmt.Println("Timeout")
            continue
        case <- complete:
            fmt.Println("Complete")
        }
        fmt.Println("Looping")
    }

    fmt.Println("Done")
}

// Start
// Timeout
// Start
// Timeout
```

`continue lable` 用于跳过 `lable` 指定的 `for` 循环，执行下一次循环。

```go
import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTicker(1 * time.Second)
    complete := make(chan int, 1)

    loop:
    for {
        fmt.Println("Start loop")
        for i:=0; i<3; i++ {
            fmt.Println(i)
            select {
            case <- timer.C:
                fmt.Println("Timeout")
                continue loop
            case <- complete:
                fmt.Println("Complete")
            }
        }
        fmt.Println("Looping")
    }

    fmt.Println("Done")
}

// Start loop
// 0
// Timeout
// Start loop
// 0
// Timeout
// Start loop
// 0
```

`break` 用于跳出 `select` 语句，`for` 循环继续执行。

```go
import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTicker(1 * time.Second)
    complete := make(chan int, 1)

    for {
        fmt.Println("Start loop")
        select {
        case <- timer.C:
            fmt.Println("Timeout")
            if a := 1; a == 1 {
                fmt.Println("Break Case")
                break
            }
        case <- complete:
            fmt.Println("Complete")
            return
        }
        fmt.Println("Looping")
    }
    fmt.Println("Done")
}
```

`break lable` 用于跳出 `lable` 指定的 `for` 循环。

```go
import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTicker(1 * time.Second)
    complete := make(chan int, 1)

    loop:
    for {
        fmt.Println("Start loop")
        for i:=0; i<3; i++ {
            fmt.Println(i)
            select {
            case <- timer.C:
                fmt.Println("Timeout")
                break loop
            case <- complete:
                fmt.Println("Complete")
            }
        }
        fmt.Println("Looping")
    }

    fmt.Println("Done")
}

// Start loop
// 0
// Timeout
// Done
```

### for switch 

使用方式和 `for select` 类似。


## 通道

### 非阻塞通道

`select` 语句可以实现非阻塞的通道操作，通过 `default` 分支实现。

```go

import (
    "fmt"
)

func main() {
    ch := make(chan int)

    select {
    case <- ch:
        fmt.Println("Read from ch")
    default:
        fmt.Println("Default")
    }

    select {
    case ch <- 1:
        fmt.Println("Write to ch")
    default:
        fmt.Println("Default")
    }
}
```

标准库中的实现：

```go
// sendTime does a non-blocking send of the current time on c.
func sendTime(c any, seq uintptr, delta int64) {
	// delta is how long ago the channel send was supposed to happen.
	// The current time can be arbitrarily far into the future, because the runtime
	// can delay a sendTime call until a goroutine tries to receive from
	// the channel. Subtract delta to go back to the old time that we
	// used to send.
	select {
	case c.(chan Time) <- Now().Add(Duration(-delta)):
	default:
	}
}
```

## 函数

### 普通函数

函数是 Go 语言的基本构建块，函数可以有参数和返回值。

```go
package main

import (
    "fmt"
)

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(1, 2)
    fmt.Println("Result:", result) // Result: 3
}
```

### 匿名函数

匿名函数是没有名字的函数，可以在定义时直接调用，也可以赋值给变量。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    rot13 := func(r rune) rune {
        switch {
        case r >= 'A' && r <= 'Z':
            return 'A' + (r-'A'+13)%26
        case r >= 'a' && r <= 'z':
            return 'a' + (r-'a'+13)%26
        }
        return r
    }
    fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}
```

## 接口

接口是一种抽象类型，没有暴露所含数据的布局或者内部结构，所提供的仅仅是一些方法。

### 接口类型

一个接口类型定义了一套方法，如果一个具体类型要实现该接口，那么必须实现接口类型中定义的方法。

与嵌入式函数类似，也可以实现嵌入式接口。

```go
package io 

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type ReadCloser interface {
    Reader
    Closer
}
```

### 接口实现

接口封装了所对应的类型和数据，只有通过接口暴露的方法才可以调用，类型的其他方法无法通过接口来调用。

```go
os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method 
os.Stdout.Close() // OK: *os.File has Close method 

var w io.Writer 
w = os.Stdout 
w.Write([]byte("hello")) // OK: io.Writer has Write method 
w.Close() 	// compile error: io.Writer lacks Close method
```

空接口类型 `interface{}` 可以保存任何类型的值，对实现类型没有任何要求。

```go
var any interface{} 
any = true 
any = 12.34 
any = "hello" 
any = map[string]int{"one": 1} 
any = new(bytes.Buffer)
```

对于创建的一个 `interface{}` 值持有一个boolean，float，string，map，pointer，或者任意其它的类型；
我们不能直接对它持有的值做操作，因为 `interface{}` 没有任何方法。需要使用类型断言从空接口还原出实际值。

### 接口值

从概念上来讲，一个接口类型的值（简称接口值）其实有两个部分：一个具体类型和该类型的值。二者称为接口的动态类型和动态值。

获取接口的动态类型：

```go
var w io.Writer 
fmt.Printf("%T\n", w) // "<nil>" 

w = os.Stdout 
fmt.Printf("%T\n", w) // "*os.File" 

w = new(bytes.Buffer) 
fmt.Printf("%T\n", w) // "*bytes.Buffer"
```

注意：含有空指针的非空接口

```go
const debug = false 

func main() { 
    var buf *bytes.Buffer 
    if debug { 
        buf = new(bytes.Buffer) // enable collection of output 
    }
    f(buf) // NOTE: subtly incorrect! 
    if debug { 
        // ...use buf... 
    } 
}

// If out is non‐nil, output will be written to it. 
func f(out io.Writer) { 
    // ...do something... 
    if out != nil { 
        out.Write([]byte("done!\n")) 
    } 
}
```

当 `main` 函数调用函数 `f` 时，把一个类型为 `*bytes.Buffer` 的空指针赋值给了 `out` 参数，所以 `out` 的动态值确实为空，
但他的动态类型是 `*bytes.Buffer`，表示 `out` 是一个包含空指针的非空接口，所以防御性检查 out != nil 会返回 true。

对于 `*bytes.Buffer`，空指针调用 `Write` 方法会导致 panic。

解决方案是把 `main` 函数中的 `buf` 类型修改为 `io.Writer`，函数 `f` 的 `out` 值为nil。

```go
var buf io.Writer
if debug { 
    buf = new(bytes.Buffer) 
}
f(buf)
```

### 类型断言

类型断言是一个作用在接口值上的操作，写出来类似于 x.(T)，其中 x 是一个接口类型的表达式，T 是一个类型（称为断言类型）。
类型断言检查接口值的动态类型是否和断言类型匹配。

有两种可能，如果断言类型 T 是一个具体类型，那么类型断言会检查 x 的动态类型是否就是 T。如果检查成功，类型断言的结果就是 x 的动态值，类型是 T。
类型断言是用来从操作数中把具体的类型的值提取出来的操作。如果检查失败，那么操作崩溃。

```go
var w io.Writer
w = os.Stdout
f, ok := w.(*os.File) 		// success: f == os.Stdout, ok == true
b, ok := w.(*bytes.Buffer)  // 崩溃: w 的动态类型是 *os.File, 不是 *bytes.Buffer
```

其次，如果相反断言的类型 T 是一个接口类型，然后类型断言检查是否 x 的动态类型满足 T。如果这个检查成功了，动态值没有获取到；
这个结果仍然是一个有相同类型和值部分的接口值，但是结果是类型 T。换句话说，对一个接口类型的类型断言改变了类型的表述方式，
改变了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型和值的部分。

如下类型断言代码中，w 和rw 都持有` os.Stdout`，于是所有对应的动态类型都是 `*os.File` ，
但变量w是一个 `io.Writer` 类型只对外公开出文件的 Write 方法，而 rw 变量还暴露它的 Read 方法。

```go
var w io.Writer
w = os.Stdout
rw := w.(io.ReadWriter) // success: *os.File has both Read and Write

w = new(ByteCounter)
rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
```

类型断言的第二种形式是带有两个结果的类型断言，如果断言失败，第二个结果是 false，否则是 true。

```go
var w io.Writer = os.Stdout 
f, ok := w.(*os.File)       // success: ok, f == os.Stdout
b, ok := w.(*bytes.Buffer)  // failure: !ok, b == nil
```

### 类型分支

类型分支是一种多分支的类型断言，可以用来判断一个接口值的动态类型是哪个。

```go
switch x := x.(type) {
case nil:
    // x is nil
case int:
    // x is int
case float64:
    // x is float64
}
```

在每个单一类型得分分支块内，变量 x 的类型都与该分支的类型一致，可以直接使用。

## 并发安全

### 竞态

数据竞争的定义：数据竞争会在两个以上的 `goroutine` 并发访问相同的变量且至少其中一个为写操作时发生。

`Go` 箴言：不要通过共享内存来通信，而应该通过通信来共享内存。

#### Map

数据类型 `map` 的操作是并发不安全的，当多个 `goroutine` 同时对同一个 `map` 进行读写操作时，会直接触发 `panic`。

```go
package main

import (
    "sync"
)

func main() {
    m := make(map[string]int)
    m["foo"] = 1

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for i := 0; i < 1000; i++ {
            m["foo"]++
        }
        wg.Done()
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            m["foo"]++
        }
        wg.Done()
    }()

    wg.Wait()
}
```

在这个例子中，我们可以看到，两个 `goroutine` 将尝试同时对 `map` 进行写入。运行这个程序时，有可能会程序崩溃：

```bash
$ go run main.go
fatal error: concurrent map writes

goroutine 17 [running]:
internal/runtime/maps.fatal({0xf55d720?, 0x0?})
        /usr/local/go/src/runtime/panic.go:1058 +0x18
main.main.func1()
        /Users/lzle/Code/Go/src/learn/main.go:16 +0x45
created by main.main in goroutine 1
        /Users/lzle/Code/Go/src/learn/main.go:14 +0xa7

goroutine 1 [sync.WaitGroup.Wait]:
sync.runtime_SemacquireWaitGroup(0xc000116018?)
        /usr/local/go/src/runtime/sema.go:110 +0x25
sync.(*WaitGroup).Wait(0xf582ac0?)
        /usr/local/go/src/sync/waitgroup.go:118 +0x48
main.main()
        /Users/lzle/Code/Go/src/learn/main.go:28 +0xff
exit status 2
```

#### Slice

数据类型 `slice` 也是并发不安全的，原因在于 `slice` 底层是一个结构体，会有数据竞争和内存重分配的问题，当发生数据竞争时可能会导致程序崩溃。

```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

* 数据竞争，一个协程在修改 `slice` 的长度，而另一个协程同时在读取或修改 `slice` 的内容。

* 内存重分配，在向 `slice` 中追加元素时，可能会触发 `slice` 的扩容操作，如果有其他协程访问了 `slice`，
就会导致指向底层数组的指针出现异常。

下面示例两个 `goroutine` 同时对 `slice` 进行读写操作：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	slice := []int{1, 2, 3}

	// 写协程
	go func() {
		for i := 0; i < 10000; i++ {
			slice = append(slice, i)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	// 读协程
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(slice[i])
			time.Sleep(1 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
}
```

执行代码，有一定概率触发 `panic` 错误。

```bash
$ go run main.go

2563
2564
2565
2566
2567
panic: runtime error: index out of range [2571] with length 2571

goroutine 19 [running]:
main.main.func2()
        /Users/lzle/Code/Go/src/learn/main.go:22 +0xb6
created by main.main in goroutine 1
        /Users/lzle/Code/Go/src/learn/main.go:20 +0xf6
exit status 2
```

下面是内存重分配造成数据竞争的示例：

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    slice := make([]int, 0)

    for i := 0; i < 10000; i++ {
        go func(i int) {
            slice = append(slice, i)
        }(i)
    }

    time.Sleep(time.Second)
    fmt.Println(len(slice)) 
}
```

执行程序，可以看到结果并不是 10000.

```go
7169
```

程序执行时并没有 `panic` 错误，但是结果不符合预期。

```bash
$ go run -race main.go
==================
WARNING: DATA RACE
Write at 0x00c000224000 by goroutine 11:
  main.main.func1()
      /Users/lzle/Code/Go/src/learn/main.go:13 +0xa8
  main.main.gowrap1()
      /Users/lzle/Code/Go/src/learn/main.go:14 +0x41

Previous read at 0x00c000224000 by goroutine 7:
  main.main.func1()
      /Users/lzle/Code/Go/src/learn/main.go:13 +0x33
  main.main.gowrap1()
      /Users/lzle/Code/Go/src/learn/main.go:14 +0x41

Goroutine 11 (running) created at:
  main.main()
      /Users/lzle/Code/Go/src/learn/main.go:12 +0x7d

Goroutine 7 (running) created at:
  main.main()
      /Users/lzle/Code/Go/src/learn/main.go:12 +0x7d
==================
```

#### Struct

结构体中的普通字段被多个 `goroutine` 同时访问时，也会出现数据竞争的问题，但这种数据竞争不会造成程序崩溃，但可能会导致数据不一致。

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Counter struct {
    n int
}

func main() {
    c := Counter{}
    var wg sync.WaitGroup
    wg.Add(2)

    // goroutine 写
    go func() {
        for i := 0; i < 10000; i++ {
            c.n++ 
            time.Sleep(time.Millisecond)
        }
        wg.Done()
    }()

    // goroutine 读
    go func() {
        for i := 0; i < 10000; i++ {
            fmt.Println(c.n)
            time.Sleep(time.Millisecond)
        }
        wg.Done()
    }()

    wg.Wait()
}
```

执行 `race` 检查，有概率检查到数据竞争。

```bash
$ go run -race main.go

10000
10000
Found 1 data race(s)
exit status 66
```

### 互斥锁：sync.Mutex

`sync.Mutex` 是一个互斥锁，是 `Go` 语言中使用最广泛的同步原语，也称为并发原语，解决数据竞争问题。`sync.Mutex` 不支持可重入锁。

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    countNum := 0

    var wg sync.WaitGroup

    wg.Add(100)
    for i := 0; i < 100; i++ {
        go func() {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                mu.Lock()
                countNum++
                mu.Unlock()
            }
        }()
    }
    wg.Wait()
    fmt.Printf("countNum: %d", countNum)
}
```

程序执行结果也是符合预期的。

```go
countNum: 10000
```

`defer Unlock()` 会在函数 `panic` 时依然会执行，`defer` 的执行成本比显式调用 `Unlock()` 略大一些，但不足以成为
代码不清晰的理由，在处理并发程序时，应当优先考虑清晰度。

```go
func main() {
    mu.Lock()
    defer mu.Unlock()
    ......
}
```

### 读写互斥锁：sync.RWMutex

`sync.RWMutex` 是一个读写锁，读锁占用的情况下会阻止写，但不会阻止读。仅在绝大部分 `goroutine` 都在获取读锁并且
锁竞争比较激烈时，`sync.RWMutex` 才有优势，因为 `sync.RWMutex`需要更复杂的内部簿记工作，所以在竞争不激烈时它比普通的互斥锁慢。

### 内存同步

现在计算机一般会有多个处理器，每个处理器都有内存的本地缓存，为了提高效率，对内存的写入是缓存在每个处理器中的，只有在必要时才会刷回内存。
刷回内存并提交的数据才会在其他处理器的 `goroutine` 中可见。

考虑如下代码片段的可能输出：

```go
var x, y int
go func() {
    x = 1
    fmt.Print("y:", y, " ")
}()
go func() {
    y = 1
    fmt.Print("x:", x, " ")
}()
```

可能会输出下面的结果：

```
y:0 x:0
x:0 y:0
```

单 `goroutine` 中的代码可以保证是顺序执行的，但在缺乏使用通道或者互斥量来显式同步的情况下，并不能保证所有的 `goroutine` 看到的事件顺序是一致的。

1、编译器可能会认为语句的执行顺序不会影响结果，然后就交换了这两个语句的执行顺序。

2、如果两个 `goroutine` 在不同的处理器上运行，每个处理器都有自己的的缓存，
那么一个 `goroutine` 的写入操作在同步到内存之前对另外一个 `goroutine` 是不可见的。

### 延迟初始化：sync.Once

sync.Once 是惰性初始化，将初始化延迟到需要的时候再去做。用于确保某个函数在整个程序运行期间只会执行一次。

```go
package main

import (
    "fmt"
    "sync"
)

func initialize() {
    fmt.Println("Hello, World")
}

func main() {
   var once sync.Once

    for i := 0; i < 10; i++ {
        // 保证只会执行一次
        once.Do(initialize)
    }
}
```

### 条件变量：sync.Cond

sync.Cond 是条件变量，用于等待或者通知，可以用于多个 goroutine 之间的通信。

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    /**/
    var mutex sync.Mutex
    cond := sync.NewCond(&mutex)
    size := 10
    wg.Add(size+1)
    
    for i:=0; i<size; i++ {
        i := i
        go func() {
            defer wg.Done()
            /*调用Wait方法时，要对L加锁*/
            cond.L.Lock()
            fmt.Printf("%d ready\n", i)
            /*Wait实际上是会先解锁cond.L，再阻塞当前goroutine
            这样其它goroutine调用上面的cond.L.Lock()才能加锁成功，才能进一步执行到Wait方法，
            等待被Broadcast或者signal唤醒。
            Wait被Broadcast或者Signal唤醒的时候，会再次对cond.L加锁，加锁后Wait才会return
            */
            cond.Wait()
            fmt.Printf("%d done\n", i)
            cond.L.Unlock()
        }()
    }
    
    /*这里sleep 2秒，确保目标goroutine都处于Wait阻塞状态
    如果调用Broadcast之前，目标goroutine不是处于Wait状态，会死锁
    */
    time.Sleep(2*time.Second)
    go func() {
        defer wg.Done()
        cond.Broadcast()
    }()
    wg.Wait()
}
```

### 对象池：sync.Pool

`sync.Pool` 是一组临时对象，这些对象可以被单独保存和获取。主要目的缓存和复用临时对象，减少内存分配和垃圾回收的开销。操作是线程安全的，适合高并发的场景。
存储在 `sync.Pool` 中的对象可能会在任何时间被垃圾回收，且不会有通知。

最佳实践的例子是 `fmt` 包，用于打印格式化的数据，使用 sync.Pool 来管理临时输出缓冲区。在负载增加时（例如多个 goroutine 同时打印），缓冲区会动态扩展；
在空闲时，缓冲区会收缩。

```go
// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.
type pp struct {
    buf buffer

    // arg holds the current item, as an interface{}.
    arg any

    // value is used instead of arg for reflect values.
    value reflect.Value

    // fmt is used to format basic items such as integers or strings.
    fmt fmt

    // reordered records whether the format string used argument reordering.
    reordered bool
    // goodArgNum records whether the most recent reordering directive was valid.
    goodArgNum bool
    // panicking is set by catchPanic to avoid infinite panic, recover, panic, ... recursion.
    panicking bool
    // erroring is set when printing an error string to guard against calling handleMethods.
    erroring bool
    // wrapErrs is set when the format string may contain a %w verb.
    wrapErrs bool
    // wrappedErrs records the targets of the %w verb.
    wrappedErrs []int
}

var ppFree = sync.Pool{
	New: func() any { return new(pp) },
}

// newPrinter allocates a new pp struct or grabs a cached one.
func newPrinter() *pp {
    p := ppFree.Get().(*pp)
    p.panicking = false
    p.erroring = false
    p.wrapErrs = false
    p.fmt.init(&p.buf)
    return p
}

// free saves used pp structs in ppFree; avoids an allocation per invocation.
func (p *pp) free() {
    // Proper usage of a sync.Pool requires each entry to have approximately
    // the same memory cost. To obtain this property when the stored type
    // contains a variably-sized buffer, we add a hard limit on the maximum
    // buffer to place back in the pool. If the buffer is larger than the
    // limit, we drop the buffer and recycle just the printer.
    //
    // See https://golang.org/issue/23199.
    if cap(p.buf) > 64*1024 {
        p.buf = nil
    } else {
        p.buf = p.buf[:0]
    }
    if cap(p.wrappedErrs) > 8 {
        p.wrappedErrs = nil
    }

    p.arg = nil
    p.value = reflect.Value{}
    p.wrappedErrs = p.wrappedErrs[:0]
    ppFree.Put(p)
}
```

静态网站生成工具 `Hugo` 在模块 `bufferpool` 中使用 `sync.Pool` 来管理缓冲区。

```go
package bufferpool

import (
    "bytes"
    "sync"
)

var bufferPool = &sync.Pool{
    New: func() any {
        return &bytes.Buffer{}
    },
}

// GetBuffer returns a buffer from the pool.
func GetBuffer() (buf *bytes.Buffer) {
    return bufferPool.Get().(*bytes.Buffer)
}

// PutBuffer returns a buffer to the pool.
// The buffer is reset before it is put back into circulation.
func PutBuffer(buf *bytes.Buffer) {
    buf.Reset()
    bufferPool.Put(buf)
}
```

`sync.Pool` 的设计目标是为跨多个 `goroutine` 的重用对象提供高效的内存复用机制，尤其适合那些生命周期较长或频繁创建销毁的对象。然而，当对象的生命周期非常短暂且重用范围局限时，使用 `sync.Pool` 可能会引入不必要的开销，此时更适合让对象自行管理复用逻辑。

线程竞争激烈，频繁创建和销毁对象，其他情况如单线程，执行频率低不用考虑使用。

#### 内存泄漏

回到 `Hugo` 刚才的例子，使用 `bufferpool` 可能会导致内存泄漏。取出来的 `bytes.Buffer` 在使用的时候，我们可以往这个元素中增加大量的 `byte` 数据，这会导致底层的 `byte slice` 的容量可能会变得很大。这个时候，即使 `Reset` 再放回到池子中，这些 `byte slice` 的容量不会改变，所占的空间依然很大。而且，因为 `Pool` 回收的机制，这些大的 `Buffer` 可能不被回收，而是会一直占用很大的空间，这属于内存泄漏的问题。

`golang` 官网 [issue 23199](https://github.com/golang/go/issues/23199)提供了一个简单的可重现的例子，演示了内存泄漏的问题。

解决方法也比较简单，就是在放回池子之前，对 `buf` 做下容量判断，超过一定的值不放回池中。

#### 内存浪费

除了内存泄漏以外，另一个常见的问题是内存浪费，就是池子中的 `buffer` 都比较大，但在实际使用的时候，很多时候只需要一个小的 `buffer`。

解决方案是做到物尽其用，可以将 `buffer` 池分成几层。标准包 [net/http](https://github.com/golang/go/blob/7e394a2/src/net/http/h2_bundle.go#L1033)有多层缓存池的实现。

```go
// Buffer chunks are allocated from a pool to reduce pressure on GC.
// The maximum wasted space per dataBuffer is 2x the largest size class,
// which happens when the dataBuffer has multiple chunks and there is
// one unread byte in both the first and last chunks. We use a few size
// classes to minimize overheads for servers that typically receive very
// small request bodies.
//
// TODO: Benchmark to determine if the pools are necessary. The GC may have
// improved enough that we can instead allocate chunks like this:
// make([]byte, max(16<<10, expectedBytesRemaining))
var (
    http2dataChunkSizeClasses = []int{
        1 << 10,
        2 << 10,
        4 << 10,
        8 << 10,
        16 << 10,
    }
    http2dataChunkPools = [...]sync.Pool{
        {New: func() interface{} { return make([]byte, 1<<10) }},
        {New: func() interface{} { return make([]byte, 2<<10) }},
        {New: func() interface{} { return make([]byte, 4<<10) }},
        {New: func() interface{} { return make([]byte, 8<<10) }},
        {New: func() interface{} { return make([]byte, 16<<10) }},
    }
)

func http2getDataBufferChunk(size int64) []byte {
    i := 0
    for ; i < len(http2dataChunkSizeClasses)-1; i++ {
        if size <= int64(http2dataChunkSizeClasses[i]) {
            break
        }
    }
    return http2dataChunkPools[i].Get().([]byte)
}

func http2putDataBufferChunk(p []byte) {
    for i, n := range http2dataChunkSizeClasses {
        if len(p) == n {
            http2dataChunkPools[i].Put(p)
            return
        }
    }
    panic(fmt.Sprintf("unexpected buffer len=%v", len(p)))
}
```

YouTube 开源的知名项目 `vitess` 中提供了[bucketpool](https://github.com/vitessio/vitess/blob/main/go/bucketpool/bucketpool.go)的实现，它提供了更加通用的多层 `buffer` 池。

#### 伪共享：False Sharing

`sync.Pool` 的实现中，涉及到了避免伪共享的设计，为了避免伪共享，使用了 `pad` 字段，可以展开理解下什么是伪共享。

```go
type poolLocal struct {
    poolLocalInternal

    // Prevents false sharing on widespread platforms with
    // 128 mod (cache line size) = 0 .
    pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte
}
```

伪共享（False Sharing）是多线程编程中的一种性能问题，它发生在多个线程同时访问不同的变量，但这些变量却共享同一缓存行（cache line）时。尽管这些变量并不相互依赖，但由于它们的存储位置在缓存中靠得很近，导致处理器频繁地无效化（invalidate）缓存行，从而影响性能。

避免伪共享的方法：

- 内存对齐

- 缓存行填充

- 分离数据结构

- 使用原子操作

参考 [什么是伪共享?](https://juejin.cn/post/7427468776982085651)

#### 第三方库

[bytebufferpool](https://github.com/valyala/bytebufferpool)

[oxtoacart/bpool](https://github.com/oxtoacart/bpool)

[alitto/pond](https://github.com/alitto/pond)

参考 [Pool：性能提升大杀器](https://time.geekbang.org/column/article/301716?code=L6RL-eocu27wznXuQuV7XXvNA01tPBYxsdUgLU6wRLI%3D)


### 静态检测：race detector

即使我们小心到不能再小心，但在并发程序中犯错还是太容易了。幸运的是， Go 的 runtime 和工具链为我们装备了一个复杂但好用的动态分析工具， 竞争检查器（the race detector）。

只要在 go build、go run 或者 go test 命令后面加上 -race 的 flag，就会使编译器创建一个你的应用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的 test，
并且会记录下每一个读或者写共享变量的 goroutine 的身份信息。另外，修改版的程序会记录下所有的同步事件，比如 go 语句，channel 操作，以及对 (*sync.Mutex).Lock，(*sync.WaitGroup)
.Wait 等等的调用。

Go race detector 是基于 Google 的 C/C++ sanitizers 技术实现的，编译器通过探测所有的内存访问，加入代码能监视对这些内存地址的访问（读还是写）。在代码运行的时候，race detector 就能监控到对共享变量的非同步访问，出现 race 的时候，就会打印出警告信息。

通过在编译的时候插入一些指令，在运行时通过这些插入的指令检测并发读写从而发现 data race 问题，就是这个工具的实现机制。而且，在运行的时候，只有在触发了 data race 之后，才能检测到，如果碰巧没有触发，是检查不出来的。

```shell
$ go test -race mypkg    // test the package
$ go run -race mysrc.go  // compile and run the program
$ go build -race mycmd   // build the command
$ go install -race mypkg // install the package
```

[Introducing the Go Race Detector](https://go.dev/blog/race-detector)

### 原子操作：sync/atomic

原子操作是指这些操作要么全部执行完毕，要么完全不执行，不会出现中间状态。

如果一个操作是由一个 CPU 指令来实现的，那么它就是原子操作；如果操作是基于多条指令来实现的，那么，执行的过程中可能会被中断，并执行上下文切换，这样的话，原子性的保证就被打破了，因为这个时候，操作可能只执行了一半。

有的代码也会因为架构的不同而不同。有时看起来貌似一个操作是原子操作，但实际上，对于不同的架构来说，情况是不一样的。

```go
package main

const x int64 = 1 + 1<<33

func main() {
    var i = x
    _ = i
}
```

如果你使用 GOARCH=386 的架构去编译这段代码，那么，第 6 行其实是被拆成了两个指令，分别操作低 32 位和高 32 位。

```shell
$ GOARCH=386 go tool compile -N -l main.go; GOARCH=386 go tool objdump -gnu main.o
TEXT main.main(SB) /Users/连长/Code/Go/src/learn/orgs/main.go
  main.go:5             0x4cc                   658b0d00000000          MOVL GS:0, CX                        // mov %gs:,%ecx           [3:7]R_TLS_LE
  main.go:5             0x4d3                   3b6108                  CMPL SP, 0x8(CX)                     // cmp 0x8(%ecx),%esp
  main.go:5             0x4d6                   7616                    JBE 0x4ee                            // jbe 0x4ee
  main.go:5             0x4d8                   83ec08                  SUBL $0x8, SP                        // sub $0x8,%esp
  main.go:6             0x4db                   c7042401000000          MOVL $0x1, 0(SP)                     // movl $0x1,(%esp)
  main.go:6             0x4e2                   c744240402000000        MOVL $0x2, 0x4(SP)                   // movl $0x2,0x4(%esp)
  main.go:8             0x4ea                   83c408                  ADDL $0x8, SP                        // add $0x8,%esp
  main.go:8             0x4ed                   c3                      RET                                  // ret
  main.go:5             0x4ee                   e800000000              CALL 0x4f3                           // call 0x4f3              [1:5]R_CALL:runtime.morestack_noctxt
  main.go:5             0x4f3                   ebd7                    JMP main.main(SB)                    // jmp 0x4cc
```

如果 GOARCH=amd64 的架构去编译这段代码，那么，第 6 行其中的赋值操作其实是一条指令：

```shell
$ GOARCH=amd64 go tool compile -N -l main.go; GOARCH=amd64 go tool objdump -gnu main.o
TEXT main.main(SB) /Users/连长/Code/Go/src/learn/orgs/main.go
  main.go:5             0x464                   55                      PUSHQ BP                             // push %rbp
  main.go:5             0x465                   4889e5                  MOVQ SP, BP                          // mov %rsp,%rbp
  main.go:5             0x468                   4883ec08                SUBQ $0x8, SP                        // sub $0x8,%rsp
  main.go:6             0x46c                   48b80100000002000000    MOVQ $0x200000001, AX                // mov $0x200000001,%rax
  main.go:6             0x476                   48890424                MOVQ AX, 0(SP)                       // mov %rax,(%rsp)
  main.go:8             0x47a                   4883c408                ADDQ $0x8, SP                        // add $0x8,%rsp
  main.go:8             0x47e                   5d                      POPQ BP                              // pop %rbp
  main.go:8             0x47f                   c3                      RET                                  // retq
```

Go 语言标准库中的 `sync/atomic` 包提供了偏底层的原子内存原语(atomic memory primitives)，用于实现同步算法，其本质是将底层 CPU 提供的原子操作指令封装成了 Go 函数。

`atomic` 包提供内存屏障的功能，不仅仅可以保证赋值的数据完整性，还能保证数据的可见性，
一旦一个核更新了该地址的值，其它处理器总是能读取到它的最新值。但是，需要注意的是，
因为需要处理器之间保证数据的一致性，`atomic` 的操作也是会降低性能的。

[官方文档地址](https://pkg.go.dev/sync/atomic)

#### Add

The add operation, implemented by the AddT functions, is the atomic equivalent of:

```
*addr += delta
return *addr
```

支持 `int32`, `int64`, `uint32`, `uint64`, `uintptr` 5 种基础数据类型的操作，下面 AddInt32 示例：

```go
// add.go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

var wg sync.WaitGroup

func test1() {
    var sum int32 = 0
    N := 100
    wg.Add(N)
    for i := 0; i < N; i++ {
        go func(i int32) {
            sum += i
            wg.Done()
        }(int32(i))
    }
    wg.Wait()
    fmt.Println("func test1, sum=", sum)
}

func test2() {
    var sum int32 = 0
    N := 100
    wg.Add(N)
    for i := 0; i < N; i++ {
        go func(i int32) {
            atomic.AddInt32(&sum, i)
            wg.Done()
        }(int32(i))
    }
    wg.Wait()
    fmt.Println("func test2, sum=", sum)
}

func main() {
    test1()
    test2()
}

// func test1, sum= 4815
// func test2, sum= 4950
```

#### CAS （CompareAndSwap）

The compare-and-swap operation, implemented by the CompareAndSwapT functions, is the atomic equivalent of:

```
if *addr == old {
    *addr = new
    return true
}
return false
```

操作实现的功能是先比较 addr 指针指向的内存里的值是否为旧值 old 相等。

* 如果相等，就把 addr 指针指向的内存里的值替换为新值 new，并返回 true，表示操作成功。
* 如果不相等，直接返回 false，表示操作失败。

支持 `int32`, `int64`, `uint32`, `uint64`, `uintptr`, `unsafe.Pointer` 这 6 种基本数据类型，用 CompareAndSwapInt32 举个例子：

```go
// compare-and-swap.go
package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var dst int32 = 100
    oldValue := atomic.LoadInt32(&dst)
    var newValue int32 = 200

    swapped := atomic.CompareAndSwapInt32(&dst, oldValue, newValue)

    fmt.Printf("old value: %d, swapped value: %d, swapped success: %v\n", oldValue, dst, swapped)
}
```

#### Swap

The swap operation, implemented by the SwapT functions, is the atomic equivalent of:

```
old = *addr
*addr = new
return old
```

如果不需要比较旧值，只是比较粗暴地替换的话，就可以使用 Swap 方法，它替换后还可以返回旧值。


#### Load & Store

The load and store operations, implemented by the LoadT and StoreT functions, are the atomic equivalents of "return *addr" and "*addr = val".

Load & Store 操作保证在多处理器、多核、有 CPU cache 的情况下是原子操作，这里主要是解决线程间内存可见性问题。
下面示例可以很好的阐述 Load & Store 的作用：

执行有问题的代码：

```go
package main

import (
    "fmt"
    "time"
)

var x int32 = 1

func storeFunc() {
    for i := 0; ; i++ {
        if i % 2 == 0 {
            x = 2
        } else {
            x = 3
        }
    }
}

func main() {
    go storeFunc()
    for {
        fmt.Println(x)
        time.Sleep(100 * time.Millisecond)
    }
}
```

输出结果

```shell
1
1
1
1
1
```

可以看到结果一直是 1，实际上 x 的值是被修改了的。

正确的示例：

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

var x int32 = 1

func storeFunc() {
    for i := 0; ; i++ {
        if i % 2 == 0 {
            atomic.StoreInt32(&x, 2)
        } else {
            atomic.StoreInt32(&x, 3)
        }
    }
}

func main() {
    go storeFunc()
    for {
        fmt.Println(atomic.LoadInt32(&x))
        time.Sleep(100 * time.Millisecond)
    }
}
```

输出结果

```shell
1
3
2
3
3
```

#### Value

`atomic` 还提供了一个特殊的类型：Value。它可以原子地存取对象类型，
但也只能存取，不能 CAS 和 Swap，常常用在配置变更等场景中。

官网示例：

```
package main

import (
    "sync/atomic"
    "time"
)

func loadConfig() map[string]string {
    return make(map[string]string)
}

func requests() chan int {
    return make(chan int)
}

func main() {
    var config atomic.Value // holds current server configuration
    // Create initial config value and store into config.
    config.Store(loadConfig())
    go func() {
        // Reload config every 10 seconds
        // and update config value with the new version.
        for {
            time.Sleep(10 * time.Second)
            config.Store(loadConfig())
        }
    }()
    // Create worker goroutines that handle incoming requests
    // using the latest config value.
    for i := 0; i < 10; i++ {
        go func() {
            for r := range requests() {
                c := config.Load()
                // Handle request r using config c.
                _, _ = r, c
            }
        }()
    }
}
```


## 协程

### context

`context` 包提供了一个用于跟踪请求的上下文，可以在多个 `goroutine` 之间传递请求特定的数据、取消信号、截止时间等。

```go
package main
import (
    "context"
    "fmt"
    "time"
)
func main() {
    ctx := context.Background()

    ctx, cancel := context.WithCancel(ctx)

    go func() {
        time.Sleep(1 * time.Second)
        cancel()
    }()

    select {
    case <-ctx.Done():
        fmt.Println("Context cancelled:", ctx.Err())
    case <-time.After(2 * time.Second):
        fmt.Println("Context still active")
    }

	fmt.Println(ctx.Deadline())
}

// 2025-02-08 17:29:05.53889 +0800 CST m=+1.000134132 true
// 2025-02-08 17:29:05.53889 +0800 CST m=+1.000134132 true
```

2、`Done()` 方法返回 `channel` 对象，多次调用返回的结果是一样的。

WithCancel arranges for Done to be closed when cancel is called; WithDeadline arranges for Done to be closed when the deadline expires; WithTimeout arranges for Done to be closed when the timeout elapses.

```go
package  main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	}()

	cancel()
	time.Sleep(1 * time.Second)
}

// context canceled
```

3、`Err()` 方法返回 `context` 被取消的原因，如果 `context` 没有被取消，`Err` 返回 nil。多次调用返回的结果是一样的。

4、`Value()` 方法返回 `context` 中与 key 关联的值，如果没有与 key 关联的值，返回 nil`。
`context.Value()` 实现了链式查找。如果不存在，会向 parent `context` 去查找。

```go
// Package user defines a User type that's stored in Contexts.
package user

import "context"

// User is the type of value stored in the Contexts.
type User struct {...}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userKey is the key for user.User values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var userKey key

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userKey).(*User)
	return u, ok
}
```

#### context.Background() & context.TODO()

`context.Background()` 返回一个非 nil 的、空的 `context`，没有任何值，不会被 cancel，不会超时，没有截止日期。一般用在主函数、初始化、测试以及创建根 `context` 的时候。

`context.TODO()` 底层和 `context.Background()` 一样，一般不用。

```go
type backgroundCtx struct{ emptyCtx }

func (backgroundCtx) String() string {
	return "context.Background"
}

type todoCtx struct{ emptyCtx }

func (todoCtx) String() string {
	return "context.TODO"
}

// Background returns a non-nil, empty [Context]. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
	return backgroundCtx{}
}

// TODO returns a non-nil, empty [Context]. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context {
	return todoCtx{}
}
```

#### context.WithValue()

`context.WithValue()` 方法返回一个父 `context` 的副本，并将 `key` 和 `val` 存储在其中，用于在 `context` 之间传递值。

```go
// WithValue returns a copy of parent in which the value associated with key is
// val.
func WithValue(parent Context, key, val any) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if key == nil {
		panic("nil key")
	}
	if !reflectlite.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}

// A valueCtx carries a key-value pair. It implements Value for that key and
// delegates all other calls to the embedded Context.
type valueCtx struct {
	Context
	key, val any
}
```

其中 `key` 必须是可比较的，`val` 可以是任意类型。`key` 一般是一个自定义的类型，避免 `packages` 使用冲突（The provided key must be comparable and should not be of type string or any other built-in type to avoid collisions between packages using context. Users of WithValue should define their own types for keys）。

```go
import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}

// found value: Go
// key not found: color
```

#### context.WithCancel()

`context.WithCancel()` 方法返回一个父 `context` 的副本，当调用返回的 `cancel` 函数或者父 `context` 的 `Done channel` 被关闭时，返回的 `context` 的 `Done channel` 会被关闭。

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// 1
// 2
// 3
// 4
// 5
```

#### context.WithDeadline()

WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d. If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent to parent. The returned [Context.Done] channel is closed when the deadline expires, when the returned cancel function is called, or when the parent context's Done channel is closed, whichever happens first.

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
    // 如果parent的截止时间更早，直接返回一个cancelCtx即可
    if cur, ok := parent.Deadline(); ok && cur.Before(d) {
        return WithCancel(parent)
    }
    c := &timerCtx{
        cancelCtx: newCancelCtx(parent),
        deadline:  d,
    }
    propagateCancel(parent, c) // 同cancelCtx的处理逻辑
    dur := time.Until(d)
    if dur <= 0 { //当前时间已经超过了截止时间，直接cancel
        c.cancel(true, DeadlineExceeded)
        return c, func() { c.cancel(false, Canceled) }
    }
    c.mu.Lock()
    defer c.mu.Unlock()
    if c.err == nil {
        // 设置一个定时器，到截止时间后取消
        c.timer = time.AfterFunc(dur, func() {
            c.cancel(true, DeadlineExceeded)
        })
    }
    return c, func() { c.cancel(true, Canceled) }
}
```

exmaple:

```go
d := time.Now().Add(shortDuration)
ctx, cancel := context.WithDeadline(context.Background(), d)

// Even though ctx will be expired, it is good practice to call its
// cancellation function in any case. Failure to do so may keep the
// context and its parent alive longer than necessary.
defer cancel()

select {
case <-neverReady:
	fmt.Println("ready")
case <-ctx.Done():
	fmt.Println(ctx.Err())
}
```

#### context.WithTimeout()

`WithTimeout()` 调用 `WithDeadline()` 函数。

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}
```

#### context.AfterFunc()

功能描述可以查看官网文档，这里给出一个使用例子：

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 在 2 秒后执行指定的函数
    timer := time.AfterFunc(2*time.Second, func() {
        fmt.Println("2 秒后执行的函数")
    })

    // 等待 3 秒以确保定时器触发
    time.Sleep(3 * time.Second)

    // 停止定时器（如果它还没有触发）
    stopped := timer.Stop()
    if stopped {
        fmt.Println("定时器被停止")
    } else {
        fmt.Println("定时器已经触发")
    }
}

// 2 秒后执行的函数
// 定时器已经触发
```

### gopkg.in/tomb.v2

`tomb` 是一个用于管理 `goroutine` 的库，可以用于优雅的关闭 `goroutine`。下面是一个简单的例子：

```go
package main

import (
	"fmt"
	"time"

	"gopkg.in/tomb.v2"
)

func worker(t *tomb.Tomb, id int) error {
	for {
		select {
		case <-t.Dying(): // 监听退出信号
			fmt.Printf("Worker %d is dying...\n", id)
			return nil
		default:
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var t tomb.Tomb

	// 启动多个 goroutine
	for i := 1; i <= 3; i++ {
		id := i
		t.Go(func() error {
			return worker(&t, id)
		})
	}

	// 让 worker 运行一段时间
	time.Sleep(2 * time.Second)

	// 停止所有 goroutine
	t.Kill(nil) // 发送退出信号

	// 等待所有 goroutine 退出
	err := t.Wait()
	if err != nil {
		fmt.Println("Some workers exited with error:", err)
	} else {
		fmt.Println("All workers exited gracefully")
	}
}

// Worker 1 is working...
// Worker 3 is working...
// Worker 2 is working...
// Worker 1 is working...
// Worker 3 is working...
// Worker 2 is working...
// Worker 2 is working...
// Worker 3 is working...
// Worker 1 is working...
// Worker 3 is working...
// Worker 1 is working...
// Worker 2 is working...
// Worker 2 is dying...
// Worker 3 is dying...
// Worker 1 is dying...
// All workers exited gracefully
```

## 标准库

### time

#### NewTicker

`time.NewTicker()` 返回一个新的 `Ticker`，该 `Ticker` 包含一个通道字段 `C`，每隔一段时间就会向该通道发送一个时间值。

```go
package main

import (
	"fmt"
	"time"
)

// Before Go 1.23, the garbage collector did not recover
// tickers that had not yet expired or been stopped, so code often
// immediately deferred t.Stop after calling NewTicker, to make
// the ticker recoverable when it was no longer needed.
// As of Go 1.23, the garbage collector can recover unreferenced
// tickers, even if they haven't been stopped.
// The Stop method is no longer necessary to help the garbage collector.
// (Code may of course still want to call Stop to stop the ticker for other reasons.)
func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
```

### fmt

格式化动词说明：

| 动词  | 说明                                      |
|-------|-------------------------------------------|
| `%v`  | 默认格式的值（自动类型推断）              |
| `%+v` | 结构体字段名 + 值（对调试更友好）          |
| `%#v` | Go 语法表示的值（如结构体的完整定义）      |
| `%T`  | 输出值的类型                              |
| `%d`  | 十进制整数                                |
| `%f`  | 浮点数（默认 6 位小数）                   |
| `%s`  | 字符串                                    |
| `%q`  | 带引号的字符串（安全转义）                |
| `%p`  | 指针地址                                  |


完整示例：

```go
package main

import (
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	// 标准输出
	fmt.Printf("整数: %d, 字符串: %s\n", 123, "hello")

	// 字符串构建 return string
	s := fmt.Sprintf("温度: %.1f℃", 23.5)
	fmt.Println(s) // 输出: 温度: 23.5℃

	// 错误生成  return error
	err := fmt.Errorf("文件 %q 不存在", "data.txt")
	fmt.Println(err) // 输出: 文件 "data.txt" 不存在

	// 结构体格式化
	p := Point{1, 2}
	fmt.Printf("%v\n", p)   // 输出: {1 2}
	fmt.Printf("%+v\n", p)  // 输出: {X:1 Y:2}
	fmt.Printf("%#v\n", p)  // 输出: main.Point{X:1, Y:2}

	// 类型信息
	fmt.Printf("类型: %T\n", p) // 输出: 类型: main.Point

	// 写入文件
	file, _ := os.Create("output.txt")
	defer file.Close()
    // 直接输出参数到 io.Writer（自动格式化，无显式格式字符串）
	fmt.Fprintf(file, "写入文件内容: %d", 42, "\n") 
    // 类似 Fprint，但自动添加换行符和空格分隔
	fmt.Fprintln(file, "写入文件内容: ", 42) // 输出: 写入文件内容:  42
}
```

### flag

`flag` 包提供了命令行参数解析的功能，可以定义命令行参数的名称、类型和默认值，并解析命令行参数。

```go

var (
	configFile       = flag.String("c", "", "config file path")
	configVersion    = flag.Bool("v", false, "show version")
	configForeground = flag.Bool("f", false, "run foreground")
	redirectSTD      = flag.Bool("redirect-std", true, "redirect standard output to file")
)

func main() {
	flag.Parse()

	Version := proto.DumpVersion("Server")
	if *configVersion {
		fmt.Printf("%v", Version)
		os.Exit(0)
	}

	/*
	 * LoadConfigFile should be checked before start daemon, since it will
	 * call os.Exit() w/o notifying the parent process.
	 */
	cfg, err := config.LoadConfigFile(*configFile)
	if err != nil {
		daemonize.SignalOutcome(err)
		os.Exit(1)
	}
}
```

代码来源于 `cubefs` 中的 `cmd/cmd.go` 文件。

### log

`log` 包提供了简单的日志记录功能，默认日志输出到标准输出中。

```
package main

import (
	"log"
)

func main() {
	a := 1
	b := 2
	c := a + b
	log.Printf("Result: %d", c)
}


// 2023/10/08 12:00:00 Result: 3
```

## Tools

### gofmt

`gofmt` 是 Go 语言自带的代码格式化工具，可以格式化 Go 代码，使其符合 Go 语言的编码规范。

```bash
$ gofmt -v
flag provided but not defined: -v
usage: gofmt [flags] [path ...]
  -cpuprofile string
    	write cpu profile to this file
  -d	display diffs instead of rewriting files
  -e	report all errors (not just the first 10 on different lines)
  -l	list files whose formatting differs from gofmt's
  -r string
    	rewrite rule (e.g., 'a[b:len(a)] -> a[b:]')
  -s	simplify code
  -w	write result to (source) file instead of stdout
```

`-d` 参数可以显示格式化前后的差异

`-w` 参数可以直接修改文件

`-l` 参数可以列出格式化不一致的文件

`-s` 参数可以简化代码

`path` 可以是文件或者目录，如果是目录，会递归处理目录下的所有文件。

```
$ gofmt -d main.go

$ gofmt -w directory
```

### workerpool

`workerpool` 是一个 `fasthttp` 中工作池实现，可以用来限制并发执行的 goroutine 数量。

可以创建一个固定数量的 `goroutine`（Worker），由这一组 Worker 去处理任务队列中的任务。

[fasthttp workerpool](https://github.com/valyala/fasthttp/blob/9f11af296864153ee45341d3f2fe0f5178fd6210/workerpool.go)


### bytebufferpool

`fasthttp` 作者 valyala 提供的一个 `buffer` 池，基本功能和 `sync.Pool` 相同。它的底层也是使用 `sync.Pool` 实现的，
包括会检测最大的 `buffer`，超过最大尺寸的 `buffer`，就会被丢弃。

[bytebufferpool](https://github.com/valyala/bytebufferpool)


### delve

`delve` 是 Go 语言的调试器，可以用来调试 Go 代码，支持断点设置、变量查看、堆栈跟踪等功能。

调试正在运行的 Go 程序，查死锁。

```bash
dlv attach $(pidof your_app)
```

进入后让程序先 停下来：

```bash
halt
```

列出所有 `goroutine`:

```bash
goroutines -t

Goroutine 34 - User: /usr/local/go/src/sync/mutex.go:123 sync.Mutex.Lock (0x4627e3)
Goroutine 56 - User: myapp/worker.go:88 (*Worker).Run (0x556abc) (blocked)
Goroutine 112 - chan receive
Goroutine 188 - IO wait
```

查看某个 `goroutine` 的堆栈。

```
goroutine 56
stack

goroutine 56 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000140c4, 0x0)
    /usr/local/go/src/runtime/sema.go:71 +0x3d
sync.(*Mutex).Lock(0xc0000140c0)
    /usr/local/go/src/sync/mutex.go:134 +0x109
myapp/worker.go:88 +0x45
```

说明 `goroutine` 56 卡在锁上。

[go-delve/delve](https://github.com/go-delve/delve)


## Test

保证复杂程序的正确性有两个关键的技术尤为有效，第一是软件发布前的例行同行评审，另一个
就是测试。

### 功能测试

`go test` 在不指定参数的情况下，会运行当前目录下所有的测试文件（*_test.go）。

`go test -v` 可以输出包中每个测试用例的名称和执行时间。

```go
$ go test ‐v 
=== RUN TestPalindrome 
‐‐‐ PASS: TestPalindrome (0.00s) 
=== RUN TestNonPalindrome 
‐‐‐ PASS: TestNonPalindrome (0.00s) 
=== RUN TestFrenchPalindrome 
‐‐‐ FAIL: TestFrenchPalindrome (0.00s) 
    word_test.go:28: IsPalindrome("été") = false 
=== RUN TestCanalPalindrome 
‐‐‐ FAIL: TestCanalPalindrome (0.00s) 
    word_test.go:35: IsPalindrome("A man, a plan, a canal: Panama") = false 
FAIL 
exit status 1 
FAIL gopl.io/ch11/word1 0.017s
```

参数 `‐run` 对应一个正则表达式，只有测试函数名被它正确匹配的测试函数才会被`go test`测试命令 运行：

```go
 go test ‐v ‐run="French|Canal" 
 === RUN TestFrenchPalindrome 
 ‐‐‐ FAIL: TestFrenchPalindrome (0.00s) 
    word_test.go:28: IsPalindrome("été") = false 
=== RUN TestCanalPalindrome 
‐‐‐ FAIL: TestCanalPalindrome (0.00s) 
    word_test.go:35: IsPalindrome("A man, a plan, a canal: Panama") = false 
FAIL 
exit status 1 
FAIL gopl.io/ch11/word1 0.014s
 ```

一般使用 `t.Errorf()` 将错误信息记录到测试日志中，不会终止测试函数的执行；`t.Fatal` 或 `t.Fatalf` 会终止测试函数的执行。

测试错误消息一般格式是 "f(x)=y, want z"。

```go
package main

import (
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    var tests = []struct {
        input string
        want  bool
    }{
        {"", true},
        {"a", true},
        {"aa", true},
        {"ab", false},
        {"kayak", true},
        {"detartrated", true},
        {"A man, a plan, a canal: Panama", true},
        {"Evil I did dwell; lewd did I live.", true},
        {"Able was I ere I saw Elba", true},
        {"été", true},
        {"Et se resservir, ivresse reste.", true},
        {"palindrome", false}, // non-palindrome
        {"desserts", false},   // semi-palindrome
    }

    for _, test := range tests {
        if got := IsPalindrome(test.input); got != test.want {
            t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, got, test.want)
        }
    }
}
```

可以再测试函数中生成随机字符串，来进行**随机测试**，在测试的整个生命周期中，每次运行测试函数都会生成不同的输入。

### 黑盒、白盒测试

黑盒测试是指测试者只关心程序的输入和输出，而不关心程序内部的实现细节。

白盒测试是指测试者关心程序的内部实现细节，比如测试函数的覆盖率。白盒测试可以通过修改全局变量、调用私有函数等方式来测试。

### 覆盖率

著名计算机科学家 Edsger W. Dijkstra 曾说过：“测试旨在发现bug，而不是证明其不存在”。

查看测试覆盖率：

```go
$ go test -v -run=Coverage -coverprofile=c.out gopl.io/ch7/eval
=== RUN   TestCoverage
--- PASS: TestCoverage (0.00s)
PASS
coverage: 63.8% of statements
ok  	gopl.io/ch7/eval	1.128s	coverage: 63.8% of statements
```

通过运行 `go tool cover` 命令，生成一个 HTML 格式的覆盖率报告：

```go
$ go tool cover -html=c.out
```

### 基准测试

基准测试是一种测试方法，用于评估程序的性能。基准测试的目的是测量程序的性能，以便在重构后确保程序的性能没有退化。

```go
import "testing" 

func BenchmarkIsPalindrome(b *testing.B) { 
    for i := 0; i < b.N; i++ { 
        IsPalindrome("A man, a plan, a canal: Panama") 
    } 
}
```

标记 `-bench` 的参数指定了要运行的基准测试。

```go
$ go test -bench=. gopl.io/ch11/word2
goos: darwin
goarch: amd64
pkg: gopl.io/ch11/word2
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkIsPalindrome-16    	 4660534	       234.9 ns/op
PASS
ok  	gopl.io/ch11/word2	1.751s
```

命令行标记 ``-benchmem` 可以显示内存分配的次数和字节数。

```go
$ go test -bench=. gopl.io/ch11/word2 -benchmem
goos: darwin
goarch: amd64
pkg: gopl.io/ch11/word2
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkIsPalindrome-16    	 4806658	       240.4 ns/op	     248 B/op	       5 allocs/op
PASS
ok  	gopl.io/ch11/word2	1.765s
```

`-cpuprofile` 参数可以生成一个 CPU profile 文件，可以使用 `go tool pprof` 命令查看。

```go
$ go test -bench=. gopl.io/ch11/word2 -cpuprofile=cpu.log
$ go tool pprof -text -nodecount=10 cpu.log
File: word2.test
Type: cpu
Time: 2025-03-29 22:02:44 CST
Duration: 1.62s, Total samples = 1.46s (89.92%)
Showing nodes accounting for 1.37s, 93.84% of 1.46s total
Showing top 10 nodes out of 79
      flat  flat%   sum%        cum   cum%
        1s 68.49% 68.49%         1s 68.49%  runtime.kevent
     0.18s 12.33% 80.82%      0.18s 12.33%  runtime.madvise
     0.05s  3.42% 84.25%      0.18s 12.33%  gopl.io/ch11/word2.IsPalindrome
     0.04s  2.74% 86.99%      0.04s  2.74%  unicode.ToLower
     0.03s  2.05% 89.04%      0.03s  2.05%  runtime.pthread_kill
     0.02s  1.37% 90.41%      0.02s  1.37%  runtime.(*mspan).init
     0.02s  1.37% 91.78%      0.05s  3.42%  runtime.mallocgcSmallNoscan
     0.01s  0.68% 92.47%      0.19s 13.01%  gopl.io/ch11/word2.BenchmarkIsPalindrome
     0.01s  0.68% 93.15%      0.01s  0.68%  runtime.(*lfstack).push
     0.01s  0.68% 93.84%      0.01s  0.68%  runtime.(*mcentral).partialSwept (inline)
```

### 示例函数

示例函数是一种特殊的测试函数，它们以 `Example` 为前缀，没有参数，没有返回值，不需要导入 `testing` 包。

示例函数的代码会被提取到文档中，可以通过 `go doc` 命令查看。

```go
func ExampleIsPalindrome() { 
    fmt.Println(IsPalindrome("A man, a plan, a canal: Panama")) 
    fmt.Println(IsPalindrome("palindrome")) 
    // Output: 
    // true 
    // false 
}
```


## Benchmark

`benchmark` 是 Go 语言自带的性能测试工具，可以用来测试 Go 代码的性能。

对比测试使用 `sync.Pool` 和不使用两种情况在高并发使用 `map` 的性能对比：

```go
package main_test

import (
	"sync"
	"testing"
)

func BenchmarkNormalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for j := 0; j < 100000; j++ {
			m[j] = j
		}

		for j := 0; j < 100000; j++ {
			_ = m[j]
		}
	}
}

func BenchmarkPreallocMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, 100000)
		for j := 0; j < 100000; j++ {
			m[j] = j
		}

		for j := 0; j < 100000; j++ {
			_ = m[j]
		}
	}
}

var pool = &sync.Pool{
	New: func() interface{} {
		return make(map[int]int, 100000)
	},
}

func BenchmarkSyncPoolMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := pool.Get().(map[int]int)
		for j := 0; j < 100000; j++ {
			m[j] = j
		}
		clear(m)
		pool.Put(m)
	}
}
```

运行基准测试，`-bench=Benchmark` 运行名称匹配 `Benchmark` 的基准测试函数：

```bash
$ go test -bench=Benchmark -benchmem
goos: darwin
goarch: amd64
pkg: learn/bench
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkNormalMap-16                175           6509358 ns/op         4729596 B/op        530 allocs/op
BenchmarkPreallocMap-16              312           3840862 ns/op         2364800 B/op        257 allocs/op
BenchmarkSyncPoolMap-16              585           1958362 ns/op              71 B/op          0 allocs/op
PASS
ok      learn/bench     5.242s
```

`go help testflag` 可以查看所有的测试标志。

```bash
The following flags are also recognized by 'go test' and can be used to
profile the tests during execution:

	-benchmem
	    Print memory allocation statistics for benchmarks.
	    Allocations made in C or using C.malloc are not counted.

	-blockprofile block.out
	    Write a goroutine blocking profile to the specified file
	    when all tests are complete.
	    Writes test binary as -c would.

	-blockprofilerate n
	    Control the detail provided in goroutine blocking profiles by
	    calling runtime.SetBlockProfileRate with n.
	    See 'go doc runtime.SetBlockProfileRate'.
	    The profiler aims to sample, on average, one blocking event every
	    n nanoseconds the program spends blocked. By default,
	    if -test.blockprofile is set without this flag, all blocking events
	    are recorded, equivalent to -test.blockprofilerate=1.

	-coverprofile cover.out
	    Write a coverage profile to the file after all tests have passed.
	    Sets -cover.

	-cpuprofile cpu.out
	    Write a CPU profile to the specified file before exiting.
	    Writes test binary as -c would.

	-memprofile mem.out
	    Write an allocation profile to the file after all tests have passed.
	    Writes test binary as -c would.

	-memprofilerate n
	    Enable more precise (and expensive) memory allocation profiles by
	    setting runtime.MemProfileRate. See 'go doc runtime.MemProfileRate'.
	    To profile all memory allocations, use -test.memprofilerate=1.

	-mutexprofile mutex.out
	    Write a mutex contention profile to the specified file
	    when all tests are complete.
	    Writes test binary as -c would.

	-mutexprofilefraction n
	    Sample 1 in n stack traces of goroutines holding a
	    contended mutex.

	-outputdir directory
	    Place output files from profiling in the specified directory,
	    by default the directory in which "go test" is running.

	-trace trace.out
	    Write an execution trace to the specified file before exiting.
```

## 其他

### gopls & staticcheck

`gopls` 是 Go 语言的官方语言服务器（Go Language Server Protocol (LSP) ），提供了代码补全、跳转、重构等功能。它可以与编辑器集成，提供更好的开发体验。

```bash
$ go install golang.org/x/tools/gopls@latest

$ gopls check Source/hdfs/hdfs.go
/Users/连长/Code/Go/src/azshara/Source/hdfs/hdfs.go:50:36-66: should use raw string (`...`) with regexp.Compile to avoid having to escape twice
```

`staticcheck` 是一个 Go 语言的静态分析工具，可以检查代码中的潜在问题、性能问题和不符合最佳实践的代码。它可以与 `gopls` 集成，提供更好的代码质量检查。

```bash
$ go install honnef.co/go/tools/cmd/staticcheck@latest

$ staticcheck Source/hdfs/hdfs.go
Source/hdfs/hdfs.go:22:2: should not use dot imports (ST1001)
Source/hdfs/hdfs.go:50:21: should use raw string (`...`) with regexp.Compile to avoid having to escape twice (S1007)
Source/hdfs/hdfs.go:52:2: error var FileCorruptedError should have name of the form ErrFoo (ST1012)
Source/hdfs/hdfs.go:290:5: error var TmpError should have name of the form ErrFoo (ST1012)
Source/hdfs/hdfs.go:602:8: unnecessary assignment to the blank identifier (S1005)
```

### error 规范

**ST1012**: Poorly chosen name for error variable
Error variables that are part of an API should be called errFoo or ErrFoo.

**ST1005**: Incorrectly formatted error string
use fmt.Errorf(“something bad”) not fmt.Errorf(“Something bad”)

### 正则规范

**S1007**: should use raw string (`...`) with regexp.Compile to avoid having to escape twice
When using regexp.Compile, use a raw string literal (`` `...` ``) instead of a double-quoted string literal (`"..."`) to avoid having to escape backslashes twice.

```go
// Before:

    regexp.Compile("\\A(\\w+) profile: total \\d+\\n\\z")
// After:

    regexp.Compile(`\A(\w+) profile: total \d+\n\z`)
```

### text/template

标准库 `text/template` 包提供了一个简单的模板引擎，可以用来生成文本输出。它支持变量替换、条件判断、循环等功能。

```go
import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}

// Dear Aunt Mildred,

// It was a pleasure to see you at the wedding.
// Thank you for the lovely bone china tea set.

// Best wishes,
// Josie

// Dear Uncle John,

// It is a shame you couldn't make it to the wedding.
// Thank you for the lovely moleskin pants.

// Best wishes,
// Josie

// Dear Cousin Rodney,

// It is a shame you couldn't make it to the wedding.

// Best wishes,
// Josie
```

### 默认值 slice&map&chan for循环

在 Go 语言中，`nil` 切片、映射和通道在使用 `for` 循环时不会引发运行时错误。相反，它们会被视为空集合，因此循环体不会执行。

```go
package main

import "fmt"

func main() {
	var s []int
	var m map[string]int
	var c chan int

	for i := range s {
		fmt.Println(i)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for v := range c {
		fmt.Println(v)
	}
}
```

如果 `chan` 是 `nil`，循环会永久阻塞，因为 `nil` `channel` 无法接收数据。

### httputil

`httputil` 可以输出 `HTTP` 请求中协议发送的全部内容。

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	// 构造一个简单的 POST 请求
	body := strings.NewReader(`{"msg":"hello"}`)
	req, err := http.NewRequest("POST", "http://example.com/index.html?debug=true", body)
	if err != nil {
		panic(err)
	}

	// 添加一些头部
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-Demo-Client/1.0")

	// DumpRequestOut 会打印出即将发送到服务器的完整请求
	dump, err := httputil.DumpRequestOut(req, true) // 第二个参数 true 表示包含 body
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))

	// 如果你真的要发出去，可以用 http.DefaultClient.Do(req)
	// resp, _ := http.DefaultClient.Do(req)
	// defer resp.Body.Close()
}

// POST /index.html?debug=true HTTP/1.1
// Host: example.com
// User-Agent: Go-Demo-Client/1.0
// Content-Length: 15
// Content-Type: application/json
// Accept-Encoding: gzip

// {"msg":"hello"}
```

### 优雅 recover

参考 `cubefs` 中 raft 模块中实现。

```go
package util

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/cubefs/cubefs/depends/tiglabs/raft/logger"
)

func HandleCrash(handlers ...func(interface{})) {
	if r := recover(); r != nil {
		debug.PrintStack()
		logPanic(r)
		for _, fn := range handlers {
			fn(r)
		}
	}
}

func logPanic(r interface{}) {
	callers := ""
	for i := 0; true; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		callers = callers + fmt.Sprintf("%v:%v\n", file, line)
	}
	logger.Error("Recovered from panic: %#v (%v)\n%v", r, r, callers)
}

func RunWorker(f func(), handlers ...func(interface{})) {
	go func() {
		defer HandleCrash(handlers...)

		f()
	}()
}

func RunWorkerUtilStop(f func(), stopCh <-chan struct{}, handlers ...func(interface{})) {
	go func() {
		for {
			select {
			case <-stopCh:
				return

			default:
				func() {
					defer HandleCrash(handlers...)
					f()
				}()
			}
		}
	}()
}
```


## 相关链接

[Standard Go Project Layout](https://github.com/golang-standards/project-layout)

[Go package for reading from continously updated files](https://github.com/hpcloud/tail)

[The Go Blog](https://go.dev/blog/)

[Effective Go](https://go.dev/doc/effective_go)

[Gopls: Analyzers](https://tip.golang.org/gopls/analyzers#)

[Staticcheck](https://staticcheck.dev/)
