# GO学习笔记

### 问题集合

1. 为什么要注册一个别名为‘default’的数据库？	TODO
	. 是否应该设置本地时间 或者 系统时间？如何设置本地时间？	TODO
	. 没有orm.RunSyncdb这步，数据库中有没有表格？	没有



## 容易遗忘的基础知识

### 字符串处理

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

### 文件操作

#### Openfile的使用

```go
//	路径，选项（os.O_RDWR），操作模式(FileMode 如：ModeDir 目录操作)
f, err := os.OpenFile("C:/itcast/testFile.xyz", os.O_RDWR, 6)
```



#### 在指定的位置写入

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



#### 按行读取

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



#### 目录的读取

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





## goroutin

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

​	当多个协程操作同一内存地址时，会产生资源竞争而发生错误。于是就设计出用来协程之间通信的channel。channel背后也是用了锁的机制。



#### channel知识点

1. channel的读和写可以不同步（异步），但是如果只写不读，会死锁；
2. 



### 项目开发流程

![1541573422747](C:\Users\Shalom\AppData\Roaming\Typora\typora-user-images\1541573422747.png)



## 数组模拟环形队列



## orm操作

#### orm需求分析
##### 痛点：
​	当你开发一个应用程序的时候(不使用O/R Mapping),你可能会写不少数据访问层的代码，用来从数据库保存，删除，读取对象信息，等等。你在DAL中写了很多的方法来读取对象数据，改变状态对象等等任务。**而这些代码写起来总是重复的**。 
##### 解决方案：
1. 提高了开发效率。由于ORM可以自动对 对象与数据库 中的Table进行字段与属性的映射，所以我们实际已经不需要一个专用的、庞大的数据访问层。 
2. ORM提供了对数据库的映射，不用sql直接编码，能够像操作对象一样从数据库获取数据。



#### orm原理

##### 定义：

> **对象关系映射**（英语：**Object Relational Mapping**，简称**ORM**，或**O/RM**，或**O/R mapping**），是一种[程序设计](https://zh.wikipedia.org/wiki/%E7%A8%8B%E5%BC%8F%E8%A8%AD%E8%A8%88)技术，用于实现[面向对象](https://zh.wikipedia.org/wiki/%E7%89%A9%E4%BB%B6%E5%B0%8E%E5%90%91)编程语言里不同[类型系统](https://zh.wikipedia.org/wiki/%E9%A1%9E%E5%9E%8B%E7%B3%BB%E7%B5%B1)的数据之间的转换。从效果上说，它其实是创建了一个可在编程语言里使用的“虚拟[对象数据库](https://zh.wikipedia.org/wiki/%E7%89%A9%E4%BB%B6%E8%B3%87%E6%96%99%E5%BA%AB)”。

##### orm如何运作？

> - 每个类对应数据库中的一张表，每个对象对应数据库表的一行，对象的每个属性对应表中的字段。orm是一个处于对象和数据库中的中间层，这个中间层提供了对象与数据库的映射。
> - 你用O/R Mapping保存，删除，读取对象，O/R Mapping负责生成[SQL](http://www.itisedu.com/phrase/200604022014515.html) ，你只需要关心对象就好。



#### orm对象包含的函数

![1542198033232](C:\Users\Shalom\AppData\Roaming\Typora\typora-user-images\1542198033232.png)

### orm使用步骤

#### 1.对数据库的前期操作

##### 导包：

> - orm 属于中间层，底层操作数据库还需要导入数据库驱动，我们这里导入MySQL数据库驱动
> - 。。。。。。

![1542198205309](C:\Users\Shalom\AppData\Roaming\Typora\typora-user-images\1542198205309.png)

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





## 知识点小集合

### context包的使用

#### context包的作用

> 主要用于控制goroutine，防止goroutine泄漏



#### func Background() Context && func TODO() Context

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













