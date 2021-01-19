## 介绍
 Go 是一种全新的编程语言，尽管它借鉴了现有的其它语言，但是它与众不同的特性，使得Go程序与其它编程语言有根本的不同。
 要学好Go语言写出易于理解的代码，需要了解Go的特性和编程惯例，例如，命名、格式化、程序构造等。
 
## 示例

Go源码库（https://golang.org/src/）中的许多软件包都包含可运行的，自包含的可执行示例，您可以直接从golang.org网站上运行该示例。（例如：https://golang.org/pkg/strings/#example_Map）

## 格式化
代码的格式化总是充满了争议，团队中的成员使用不同的代码风格使得程序后期的维护工作变得困难。
使用Go，我们可以采用一种不寻常的方法，让机器处理大多数格式化问题。 gofmt程序（也可作为go fmt使用，它在软件包级别而不是源文件级别运行）读取Go程序，
并以缩进和垂直对齐的标准样式发出源代码，并保留注释，并在必要时重新格式化注释。 

例如，不需要花时间在结构的字段上排列注释。Gofmt会帮你完成这些工作
```go
type T struct {
    name string // name of the object
    value int // its value
}
```
gofmt将使各列对齐：
```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

以下是一些关于格式化的细节
* 缩进
gofmt默认会使用制表符tab进行缩进，仅在必要时使用空格。
* 行的长度
Go没有行长限制，如果感觉太长可以对行进行拆分。
* 括号
Go需要的括号比C和Java少：控制结构（if,for,switch）的语法中没有括号。 同样，运算符优先级层次更短更清晰，
示例，
```go
x<<8 + y<<16
```
## 注释
Go提供了/ * * /块注释和//行注释。通常在代码中我们会使用行注释，块注释主要用于为程序包注释，但在表达式中或者禁用大量代码时很有用。
每个包都应该进行注释，对于拥有多个文件的包，只需要一个包注释即可（任意文件），包注释应当对该包进行介绍，并提供与包装整体相关的信息。
 
## 命名
与其它编程语言相比Go的命名具有特殊用途，了解Go程序中的命名约定是很有必要的。
### 包名
当一个包被导入，包名即是包中所有内容的访问器。
```go
import "bytes"
```
导入bytes包后可以通过bytes.Buffer进行调用。包的名称应当简洁且易于理解。 按照惯例，软件包应当使用小写的单字名称，无需下划线和驼峰命名。  
另一个惯例是，程序包名称是其源码目录的名称。 例如，文件夹src/encoding/base64中的程序包被导入为“ encoding/base64”，但名称为base64，而不是encoding_base64，也不是encodingBase64。
利用程序包的导出的名称可以避免啰嗦的语法。 
例如，bufio包中的缓冲读取器名为Reader，而不是BufReader，因为用户将其视为bufio.Reader，这是一个简洁明了的名称。 
此外，由于导入的实体始终使用其包名称来寻址，因此，bufio.Reader不会与io.Reader冲突。 
类似地，用于创建ring的新实例的函数。Ring（这是Go中构造函数的定义）通常被命名为NewRing，
但是由于Ring是该包导出的唯一类型，并且由于该包被称为ring，因此，我们可将其命名为“New”,使用时可通过“ring.New”进行调用。
### Getters
Go没有提供自动化的getter和setter方法支持。它通常由开发者自己实现，但是将Get用作getter方法的前缀不符合惯例也没有必要。 
如果您有一个名为owner的字段，则getter方法应称为Owner，而不是GetOwner。而setter方法可以会被命名为SetOwner。
示例，
```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```
## 流程控制
Go没有do或while循环，只有一个通用的for，
switch更灵活； if和switch可以像for一样使用初始化语句； 
break和continue语句带有一个可选标签以标识要中断或继续的内容； 
选择了新的控制结构，包括type switch 和多路通信复用器select。 
语法也略有不同：没有括号，并且主体必须始终用大括号分隔。

### if
在Go中一个简单的if语句通常如下所示
```go
if x > 0 {
    return y
}
```
强制大括号鼓励在多行上编写简单的if语句。这是一种良好的代码风格，尤其是当主体包含控制语句（例如return或break）时。
由于if和switch可使用初始化语句，因此我们常常能见到以下格式的代码，
```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```
在Go中if语句不再向下执行时（即以break，continue，goto或return结尾），通常会省略不必要的else。
示例，
```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

## 方法
### 多返回值
Go中的方法和函数可以返回多个值，
在Go中，Write可以返回一个计数和一个错误, 例如，os包中的Write方法的签名如下
```go
func (file *File) Write(b []byte) (n int, err error)
```
如文档所述，当n！= len（b）时，它返回写入的字节数和非nil错误，这是一种常见的样式。
### 命名返回参数
我们可以像入参那样对结果参数进行命名，
命名后，函数开会将它们初始化为参数类型的零值。 如果函数执行不带参数的return语句，则将结果参数的当前值用作返回值。
结果参数的命名不是强制的，但是一个好的命名可以提高程序的可读性
### 延迟函数（Defer）
Go的defer语句可以使一个方法的调用，在当前函数返回之前立即运行。 使用defer可以便利的进行异常处理，例如，io资源的释放，解锁互斥锁
```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```
defer对Close这样的函数的调用有两个优点。
首先，它保证您永远不会忘记关闭文件；
其次，它意味着close位于open附近，这比将其放在函数的末尾可读性要更好。

deferred函数的参数(如果函数是一个方法，则包含接收器)在deferred执行时计算，而不是在调用执行时计算。
除了避免担心变量在函数执行时改变值之外，这意味着一个延迟调用站点可以延迟多个函数的执行。
示例1，
```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```
延迟函数是按照后进先出的顺序执行，所以当函数返回时，上述代码将输出4 3 2 1 0。
示例2，
```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```
输出，
```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

## Data
### new
Go有两个分配原语，内置函数new和make。 它们执行不同的操作，并应用于不同的类型。 
new是一个分配内存的内置函数，但是与其他语言中的同名函数不同，它不会初始化内存而是返回一个待分配类型的零值的地址。 
也就是说，new（T）为类型T的分配零存储空间并返回其地址，即类型* T的值。 在Go术语中，它返回一个指针，该指针指向新分配的类型T的零值。

由于new返回的内存为零，我们可以使用new创建一个数据结构并使用。 
例如，bytes.Buffer的默认零值为“空缓冲区”。sync.Mutex的零值为未锁定的互斥量。
示例，
```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```
通过new进行使用
```go
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```
### 构造函数与复合文字
有时零值不能满足需求，此时需要初始化构造函数来进行初始化，如os包中的NewFile函数。
```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```
那里冗余较多，我们可以使用复合文字来简化它，该文字每次求值时都会创建一个新实例。
```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

注意，与C语言不同，返回局部变量的地址完全没问题;函数返回后，与该变量相关的存储空间仍然存在。实际上，获取复合字面值的地址会在每次求值时分配一个新的实例，因此我们可以合并最后两行。

```go
    return &File{fd, name, nil, 0}
```
复合文字的字段按顺序排列，并且必须全部存在。但是，通过显式地将元素标记为field:value对，初始化器可以以任何顺序出现，缺失的元素作为它们各自的零值。因此我们可以说

```go
    return &File{fd: fd, name: name}
```
极端情况下，如果一个复合文字不包含任何字段，它将为该类型创建一个零值。 表达式new（File）和＆File {}是等效的。
复合文字也可用于array，slice和map的创建，其中字段标签为索引或map的key。 在这些示例中，无论Enone，Eio和Einval的值是什么，初始化都可以工作，只要它们是不同的即可。
```go
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```
### make
函数make（T，args）不同于new（T）的是，它仅创建slice，map和channel，并返回类型T（不是* T）的初始化值（非零值）。这三种类型本质上为引用数据类型，它们在使用前必须初始化。 例如，切片是一个具有三项内容的描述符，包含一个指向（数组内部）数据的指针、长度以及容量， 在这三项被初始化之前，该切片为 nil。对于切片、映射和信道，make 用于初始化其内部的数据结构并准备好将要使用的值。例如，
```go
make([]int, 10, 100)
```
会分配一个具有 100 个 int 的数组空间，接着创建一个长度为 10， 容量为 100 并指向该数组中前 10 个元素的切片结构。（生成切片时，其容量可以省略，更多信息见切片一节。） 与此相反，new([]int) 会返回一个指向新分配的，已置零的切片结构， 即一个指向 nil 切片值的指针。

### array
* 数组不是引用类型，如果传递一个数组到方法，本质上传递的是数组的复制，方法内对数组的修改不会影响到原数组
* 数组的大小是数组类型的一部分，因此[10]int 和 [6]int类型不相同
需要时我们可以通过slice对数组进行包装然后传递给方法，而不是&array
### slice
切片对数组进行包装，为数据序列提供了更通用、更强大和更方便的接口。除了具有显式维度的项(如转换矩阵)，Go中的大多数数组编程都是用切片而不是简单的数组来完成的。

切片包含对基础数组的引用，如果将一个切片分配给另一个切片，则两个切片均引用同一数组。 如果函数采用slice参数，则对slice的元素所做的更改将对调用者可见，这类似于将指针传递给基础数组。 因此，Read函数可以接受切片参数，而不是指针和计数。 切片内的长度设置了要读取多少数据的上限。 这是包os中File类型的Read方法的签名：
### Map
映射是一种便捷而强大的内置数据结构，它将一种类型（键）的值与另一种类型（元素或值）的值相关联。 键可以是定义了相等运算符的任何类型，例如整数，浮点数和复数，字符串，指针，接口（只要动态类型支持相等），结构和数组。 切片不能用作映射键，因为未在其上定义相等性。 像切片一样，映射保留对基础数据结构的引用。 如果将地图传递给更改地图内容的函数，则更改将在调用方中可见。

尝试使用映射中不存在的键来获取映射值时，将为映射中的条目类型返回零值。 例如，如果映射包含整数，则查找不存在的键将返回0。注意：无法作为key是否存在的依据，因为value可能保存的就是零值
使用map实现set，
```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
```

判断key是否真的存在
```go
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```
## Printing
Println版本还会在参数之间插入空格，并在输出中添加换行符，而Print版本仅在双方都不是字符串的操作数时才添加空格。

## Append
append所做的是将元素添加到片的末尾并返回结果。需要返回结果，因为与手写追加一样，底层数组可能会发生变化。这个简单的例子
使用...追加一个slice到另一个slice
```go
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
```

### 初始化
一个包中可以有多个 init 函数，但是它们的执行顺序并不确定，所以如果你定义了多个 init 函数的话，要确保它们是相互独立的，一定不要有顺序上的依赖。

初始化顺序：初始化导入的包->初始化变量->init函数
The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.

指针方法会修改接收者，当使用value类型作为接收者时，会复制value的copy导致修改不可见，与point语义不符因此不允许valur类型调用指针类型的方法

