# GO学习笔记

## GO工具

### govendor包依赖管理工具

使用 govendor 进行项目依赖管理，该工具将 项目依赖的外部包拷贝到项目中的vendor目录下，并通过 vendor.json 文件来记录依赖包的版本,方便用户使用相对稳定的依赖。<font color=red>编译的时候，系统优先从vendor目录中寻找依赖包，如果vendor中没有，然后再去GOPATH中寻找</font>。

#### 命令集合

|  指令  |                          含义                          |
| :----: | :----------------------------------------------------: |
|  init  |       创建 `vendor` 文件夹和 `vendor.json` 文件        |
|  list  |                  列出已经存在的依赖包                  |
|  add   |     从 `$GOPATH` 中添加依赖包，会加到 vendor.json      |
| update |                从 `$GOPATH` 升级依赖包                 |
| remove |               从 `vendor` 文件夹删除依赖               |
| status |        列出本地丢失的、过期的和修改的`package`         |
| fetch  |      从远端库添加或者更新 `vendor` 文件中的依赖包      |
|  sync  | 本地存在`vendor.json` 时候拉取依赖包，匹配所记录的版本 |
|  get   |                     等同于`go get`                     |



#### vendor中依赖包的类型

对于 govendor 来说,依赖包主要有以下多种类型:

| 状态      | 缩写状态 | 含义                                           |
| --------- | -------- | ---------------------------------------------- |
| +local    | l        | 本地包,及项目自身的包组织                      |
| +external | e        | 外部包,即被 $GOPATH 管理,但不在 vendor 目录下  |
| +vendor   | v        | 已被 govendor 管理,即在vendor目录下            |
| +std      | s        | 标准库中的包                                   |
| +unused   | u        | 未使用的包,即包在vendor目录下,但项目并没有使用 |
| +missing  | m        | 代码引用了依赖包,但该报并没有找到              |
| +program  | p        | 主程序包,意味着可以编译为执行文件              |
| +outside  |          | 外部包和缺失的包                               |
| +all      |          | 所有的包                                       |

#### 简单的使用

- 安装，该命令会将govendor可执行文件下载到 $GOPATH 的 bin 目录下

  ```shell
  $ go get github.com/kardianos/govendor
  ```

  命令行执行`govendor`	,查看安装结果.

  > **注意**: 需要把 `$GOPATH/bin/`加到`PATH`中

- 使用如下：

  ```shell
  # 进到 GOPATH 中的一个项目中
  cd "my project in GOPATH"
  
  # 初始化 vendor 目录, project 下出现 vendor 目录
  govendor init
  
  # 将该项目依赖的GOPATH中的包，添加到vendor中
  govendor add +external
  
  # 查看vendor中的包依赖情况
  govendor list
  
  # 查看vendor中某一特定包的依赖情况
  govendor list -v fmt
  
  # 从远程获取 golang.org/x/net/context 包 指定的版本或修订
  govendor fetch golang.org/x/net/context@a4bbce9fcae005b22ae5443f6af064d80a6f5a55
  govendor fetch golang.org/x/net/context@v1   
  
  # 获取标签和分支名为 v1 的context包
  govendor fetch golang.org/x/net/context@=v1  
  
  # 更新一个包到最新
  govendor fetch golang.org/x/net/context
  
  # 格式化本地依赖库库
  govendor fmt +local
  
  # Build everything in your repository only
  govendor install +local
  
  # Test your repository only
  govendor test +local
  ```



### gofmt工具的使用

```shell
gofmt 文件名 		# - 输出格式化后的代码

gofmt -w 文件名 	# - 重新格式化代码并更新文件

gofmt -r'rule' 文件名 	# - 格式化代码前执行指定的规则

gofmt 包所在的路径 	# - 格式化整个包下的源文件
```



### 设置自动代理

[GCTT | 【干货】go get 自动代理](https://mp.weixin.qq.com/s/N1tixHZuG6MLiWTd4vIQrQ)



## 字符串处理

```go
//	1. 字符串按 指定分割符拆分：	Split
ret := strings.Split(str, " I")

//	2. 字符串按 空格拆分： Fields
ret = strings.Fields(str)

//	3. 判断字符串结束标记 HasSuffix
flg := strings.HasSuffix("test.abc", ".mp3")

//	4. 判断字符串起始标记 HasPrefix
flg := strings.HasPrefix("test.abc", "tes.")
```

## 文件操作

### Openfile的使用

```go
//	路径，选项（os.O_RDWR），操作模式(FileMode 如：ModeDir 目录操作)
f, err := os.OpenFile("C:/itcast/testFile.xyz", os.O_RDWR, 6)
```



### 在指定的位置写入

```go
//	Seek(偏移量，起始偏移位置)，返回 相对开头的偏移量 和 错误信息
//	偏移量：+ 表示从头开始向左；- 表示从尾开始向右。
//	起始偏移位置：0,1,2分别是开头位置，当前位置和末尾位置，也可使用io.SeekStart，io.SeekCurrent 和 io.SeekEnd
off, _ := f.Seek(-5, io.SeekEnd)
fmt.Println("off:", off)

//	WriteAt()在指定位置写入数据，返回写入数据长度和错误信息
n, _ = f.WriteAt([]byte("1111"), off)
fmt.Println("WriteAt n :", n)
```



### 按行读取

使用 bufio 包

```go
// 创建一个带有缓冲区(用户缓冲)的 reader
reader := bufio.NewReader(f)
for {
    buf, err := reader.ReadBytes('\n')		// 读一行数据，遇到\n，将读取到的数据返回到切片中
    if err != nil && err == io.EOF {
        fmt.Println("文件读取完毕")
        return
    } else if err != nil {
        fmt.Println("ReadBytes err:", err)
    }
    fmt.Print(string(buf))
}
```



### 目录的读取

目录也是文件

>  Readdir() 返回一个接口切片，每个接口代表目录中的一个成员，接口中有该成员的一切信息

```go
// 打开目录
f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
if err != nil {
    fmt.Println("OpenFile err: ", err)
    return
}
defer f.Close()
// 读取目录项
info, err:= f.Readdir(-1)	// -1： 读取目录中所有目录项
if err != nil {
    fmt.Println("Readdir err: ", err)
    return
}
// 变量返回的切片
for _, fileInfo := range info {
    if fileInfo.IsDir() {			// 是目录
        fmt.Println(fileInfo.Name(), " 是一个目录")
    } else {
        fmt.Println(fileInfo.Name(), " 是一个文件")
    }
}
```





## orm操作

### orm需求分析
##### 痛点：
​	当你开发一个应用程序的时候(不使用O/R Mapping),你可能会写不少数据访问层的代码，用来从数据库保存，删除，读取对象信息，等等。你在DAL中写了很多的方法来读取对象数据，改变状态对象等等任务。**而这些代码写起来总是重复的**。 
##### 解决方案：
1. 提高了开发效率。由于ORM可以自动对 对象与数据库 中的Table进行字段与属性的映射，所以我们实际已经不需要一个专用的、庞大的数据访问层。 
2. ORM提供了对数据库的映射，不用sql直接编码，能够像操作对象一样从数据库获取数据。



### orm原理

##### 定义：

> **对象关系映射**（英语：**Object Relational Mapping**，简称**ORM**，或**O/RM**，或**O/R mapping**），是一种[程序设计](https://zh.wikipedia.org/wiki/%E7%A8%8B%E5%BC%8F%E8%A8%AD%E8%A8%88)技术，用于实现[面向对象](https://zh.wikipedia.org/wiki/%E7%89%A9%E4%BB%B6%E5%B0%8E%E5%90%91)编程语言里不同[类型系统](https://zh.wikipedia.org/wiki/%E9%A1%9E%E5%9E%8B%E7%B3%BB%E7%B5%B1)的数据之间的转换。从效果上说，它其实是创建了一个可在编程语言里使用的“虚拟[对象数据库](https://zh.wikipedia.org/wiki/%E7%89%A9%E4%BB%B6%E8%B3%87%E6%96%99%E5%BA%AB)”。

##### orm如何运作？

> - 每个类对应数据库中的一张表，每个对象对应数据库表的一行，对象的每个属性对应表中的字段。orm是一个处于对象和数据库中的中间层，这个中间层提供了对象与数据库的映射。
> - 你用O/R Mapping保存，删除，读取对象，O/R Mapping负责生成[SQL](http://www.itisedu.com/phrase/200604022014515.html) ，你只需要关心对象就好。



### orm对象包含的函数



### orm使用步骤

#### 1.对数据库的前期操作

##### 导包：

> - orm 属于中间层，底层操作数据库还需要导入数据库驱动，我们这里导入MySQL数据库驱动
> - 。。。。。。



##### 与数据库关联：

###### **？？为什么要注册一个别名为‘default’的数据库？	TODO**

```go
//	初次建立时，必须注册一个别名为‘default’的数据库；数据库驱动名；连接数据库的命令：“用户名：密码@tcp(IP:port)/数据库名称?编码格式”
orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
```

##### 时区设置：

###### **是否应该设置本地时间 或者 系统时间？如何设置本地时间？	TODO**

​	存取默认都为 UTC 时间，

```go
// 设置为 UTC 时间
orm.DefaultTimeLoc = time.UTC
```



##### 在内存中注册表格模型：

​	映射规则：

> - 第一个字母如果是大写，变为小写
> - 后边每遇到大写字母，变为'_小写'

```go
package main

import "github.com/astaxie/beego/orm"

type User struct {
    Id   int
    Name string
}
//	结构体是一个类，对应数据库中的一张表
func init(){
    orm.RegisterModel(new(User))
}
//	也可以同时注册多个 model
orm.RegisterModel(new(User), new(Profile), new(Post))
```



##### 创建表格：

​	**这一步才是真正将表格创建出来**

###### 没有orm.RunSyncdb这步，数据库中有没有表格？	没有

```go
// 数据库别名；更改数据前是否将表格清零；是否显示表格信息
orm.RunSyncdb("default", false, true)
```



#### 2.增删改查操作

```go
//	new一个orm对象
o := orm.NewOrm()
user := User{Name: "slene"}
// insert 增
id, err := o.Insert(&user)
// update 改
user.Name = "astaxie"
num, err := o.Update(&user)
// read one 查
u := User{Id: user.Id}
err = o.Read(&u)
// delete 删
num, err = o.Delete(&u)
```

- 查询表格

  ```go
  o := orm.NewOrm()
  //	获得名为 area 的表的查询对象
  qp := o.QueryTable("area")
  areaList := []models.Area{}
  //	将表 area 中的数据全部放在切片 areaList 中
  num, err := qp.All(&areaList)
  ```

- 



## go语言的基础数据结构

### 基础数据类型源码位置

> runtime 包里 和 buildin包里

### go语言中的error

#### 源码：

```go
// 路径：buildin/buildin.go
type error interface {
	Error() string
}
```

- 通过源码可以看出，error本质上是一个带方法的 interface；
- 带方法的interface，底层存储了两个字段：实现该interface的变量的原始类型 和 该变量的值；
- 只有上述两个字段都为 nil 的时候，判断 interface == nil 才为 true。

#### 可能会踩的坑

使用自定义的 错误类型 的时候，例码：

```go
// 这里定义了一个能够实现 error 的结构体，自定义的错误类型
type MyError struct {
	s string
}
func (e *MyError) Error() string {
	return e.s
}
func test() error {
	// 这里声明了该错误类型的变量，并且赋值为nil
	var a *MyError
	a = nil
	// 这里return的时候，a被自动转成 error类型，里面存储的类型不为nil
	return a
}
func main() {
    // 这里的err是一个接口，存储的类型为 MyError，值为空，所以这里返回的err永远都不为 nil
	err := test()
	fmt.Println(err == nil)
}
```

上述代码引用的是 errors.New() 方法的源代码：

```go
func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```



### Map





### recorver

recover 使用的三个要点：

1. recover 必须在 defer 中调用；
2. recover 必须在函数中调用，匿名不匿名都行；
3. recover 不能在多层函数中调用。

具体使用案例：

```go
//	1: 没有在defer中使用
func main() {
    if r := recover(); r != nil {
    	log.Fatal(r)
    }
    panic(123)
    if r := recover(); r != nil {
    	log.Fatal(r)
    }
}
//	2: 不可以在defer的多层函数中调用
func main() {
    defer func() {
        if r := MyRecover(); r != nil {
            fmt.Println(r)
        }
    }()
    panic(1)
}
func MyRecover() interface{} {
    log.Println("trace...")
    return recover()
}
//	3: 不可再defer的多层函数中调用
func main() {
    defer func() {
        defer func() {
            if r := recover(); r != nil {
            	fmt.Println(r)
        	}
    	}()
	}()
	panic(1)
}
//	4: 正确
func MyRecover() interface{} {
	return recover()
}
func main() {
    defer MyRecover()
    panic(1)
}
//	5: 必须要在defer的函数中调用
func main() {
    defer recover()
    panic(1)
}
//	6: 正确
func main() {
    defer func() {
        if r := recover(); r != nil { ... }
    }()
    panic(nil)
}
```





### goroutin

#### 并发和并行：

1. 并发：多个线程在一个CPU中运行，来回切换，微观上看，同一时刻只执行一个线程；
2. 并行：多个线程在多个CPU中运行，同时执行多个线程。

#### Go语言中的主线程和协程

1. 主线程其实是进程的另一种说法；
2. 协程是Go语言设计者在线程上优化得来的，比线程更轻巧（独立栈空间，共享堆空间）。



#### MPG模式

1. M : 相当于主线程；P：是协程运行的环境；G：在P中运行的协程
2. 多个M可运行在多个CPU上，叫做并行，运行在同一CPU上，叫做并发
3. M主线程和M1协程并发执行

##### 查看多线程竞争关系命令：

```go
go build -race main.go	//	然后再执行main.exe
```



### channel

#### channel的引出

当多个协程操作同一内存地址时，会产生资源竞争而发生错误。于是就设计出用来协程之间通信的channel。channel背后也是用了锁的机制。

#### 无缓冲chan中包含的GO并发内存模型

1. 发送数据前接收必须准备好，如果没有准备好会出现死锁； 
2. 接收完成之前发送必须已经结束，保证接收的数据完整； 

<font color=red>上述两点是并发模型的重要保证</font>

#### 导致死锁的情况

1. chan 关闭后，往该 chan 发送数据会导致 runtime panic； 
2. channel的读和写可以不同步（异步），但是如果只写不读，会死锁；

#### 使用小技巧

2. **判断chan是否关闭：**从该 chan 接收数据会立刻返回，同时可以加入第二个参数，判断是关闭了还是正常数据返回，即：`x, ok :=<-c` ，这时候 ok 是 false，因为此特性，close 一个 chan 可以用于广播（广播通道关闭的信号）； 
2. 往一个 nil chan 发送数据会永远阻塞





### Interface使用的细节

- 接口被指针类型变量实现 且 方法绑定在指针类型上，接口实例化后的对象和指针类型变量是同一个实例



#### 判断对象是否实现了某接口

例码：

```go
type MyWriter struct{}
func (m *MyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
// 声明一个匿名变量
var _ io.Writer = (*MyWriter)(nil)
```

**解释：**

1. 检查 *MyWriter 是否实现了 io.Writer 接口
2. (*MyWriter)(nil) 是将nil强转为 *MyWriter 类型
3. 若*MyWriter没有实现 io.Writer 接口，编译器会直接报错



### Slice源码分析



### 反射



### 闭包

```go
func Closure() func() int {
    var x int
    return func() int {
        x++
        return x
    }
}
```

调用这个函数会返回一个函数变量。`i := Closure()`：通过把这个函数变量赋值给 `i`，`i` 就成为了一个**闭包**。

**注意：** `i` 保存着对 `x` 的引用，可以理解 `i` 中有着一个指针指向 x 或 **i 中有 x 的地址**。由于 `i` 有着指向 `x` 的指针，所以可以修改 `x`。



## go的并发

### MPG模型

- 在单CPU的情况下，go的并发是非抢占的，后边的协程要执行，首先要有协程出现阻塞，延迟或者放弃执行

### 如何控制并发执行的 Goroutine 的最大数目？

例码：

```go
type pool struct {
	maxNum   int        // 最大Goroutine 数目
	taskChan chan *Task // 接收并传递任务的通道
}

func (pool) work() {
	for range taskChan {
		Task() // 这里执行任务
	}
}
func (pool) run() {
	for i := 0; i < pool.maxNum; i++ {
		go pool.work() // 这里只启动maxNum个go程
	}
}
```





## go程序中的规范

### 变量名声明规范

在 Go 编程中最好用短的变量名，尤其是那些作用域比较有限的局部变量

> 用 `c` 而不是 `lineCount`
>
> 用 `i` 而不是 `sliceIndex`

1. 基本规则：距离声明的地方越远，变量名需要越具可读性。

2. 作为一个函数接收者，1、2 个字母的变量比较高效。

3. 像循环指示变量和输入流变量，用一个单字母就可以。

4. 越不常用的变量和公共变量，需要用更具说明性的名字。



## 导包

### 点导包

测试的时候使用，用来将测试代码伪装成包内文件

### Go 包初始化流程：

![Fr2R83ovb9LYtta-DxOJQ1mUtZuq](assets/Fr2R83ovb9LYtta-DxOJQ1mUtZuq.png)

### go语言的执行顺序的规则

A依赖B，A虽然在B的前边，执行顺序依然是先B再A；

下面看一个初始化的例子，例码：

```go
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int { return c }
func g() int  { return a }
func sqr(x int) int { return x*x }
func u() int { return 1}
func v() int { return 2}
```

**分析：**

1. 首先执行的是给a赋值的语句，但是f() 依赖于c ，所以必须先执行 c 赋值语句；
2. 执行`sqr(u()) + v()`的时候依赖 u()，所以这里最先执行的是 u()；
3. 按照这种规则，推出函数的执行顺序是：u()、sqr()、v()、f()、v()、g()。





## 官方库解析

### context包的使用

#### context包的作用

> 主要用于控制goroutine，防止goroutine泄漏



#### func Background() Context 和 func TODO() Context

- 这是context包内部已经实现好了的两个空 context 对象；
- 可以通过调用小标题中的两个方法获得该 context 对象；
- 这两个函数获得的 context ，一般用来作为根，往下派生；



#### 四个派生函数

##### func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

- WithCancel() 函数只是单纯的生成一个parent的一个副本，相当于拷贝；



##### func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

- WithDeadline() 函数生成一个时间期限必须在 parent 之前的 context；
- 若时间期限在 parent 之后，则返回 parent 的副本；



##### func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

- WithTimeout() 函数是对 WithDeadline() 的一层封装，timeout不再是期限时间点，而是代表“经过多长时间”



##### func WithValue(parent Context, key, val interface{}) Context

- 源码：

  ```go
  func WithValue(parent Context, key, val interface{}) Context {
      if key == nil {
          panic("nil key")
      }
      if !reflect.TypeOf(key).Comparable() {
          panic("key is not comparable")
      }
      return &valueCtx{parent, key, val}
  }
  ```

- 官方应用案例：

  ```go
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
  ```

- Output:

  ```go
  found value: Go
  key not found: color
  ```

  暂时还看不出来这个函数有什么特殊作用...



##### Context 对象

- 源码：

  ```go
  type Context interface {
      // Done returns a channel that is closed when this Context is canceled
      // or times out.
      Done() <-chan struct{}
  
      // Err indicates why this context was canceled, after the Done channel
      // is closed.
      Err() error
  
      // Deadline returns the time when this Context will be canceled, if any.
      Deadline() (deadline time.Time, ok bool)
  
      // Value returns the value associated with key or nil if none.
      Value(key interface{}) interface{}
  }
  ```

- Done()，返回一个单向输出channel。当times out或者调用cancel方法时，将会close掉。

- Err()，返回一个错误。表明(indicate)该context为什么被取消掉。

- Deadline()，当goroutine快要被cancel的时候，返回截止时间。

- Value()，返回值。