# FlameGraph和pprof的使用

## 安装

### Ubuntu

1. Go tool pprof 辅助工具安装(图形工具graphviz为例)，使用apt-get安装graphviz 

   ```shell
   sudo apt-get install graphviz
   ```


2. 图形工具，火焰图安装： 

   ```shell
   # 从GitHub上将FlameGraph的源码clone下来
   git clone https://github.com/brendangregg/FlameGraph.git
   # 将里边的主要可执行文件flamegraph.pl拷贝到全局目录下
   cd FlameGraph-master
   cp flamegraph.pl /usr/local/bin
   # 执行测试命令，安装成功会显示一系列命令帮助信息
   flamegraph.pl -h
   ```

3. 最后安装 go-torch 

   - 将代码从GitHub上拉下来，然后编译成可执行文件。GitHub链接：

     ```shell
     https://github.com/uber-archive/go-torch.git
     ```

   - 它会依赖很多本地没有的依赖包，不知道 go mod 能否解决自动拉取依赖包的问题

### Windows

1. graphviz 下载
   - 官方下载安装包，[下载Stable稳定版本(.msi) ](<http://www.graphviz.org/download/>)
   - 安装好后，配置环境变量
2. 火焰图的安装和Ubuntu一样，同样需要配置flamegraph.pl的路径到 PATH 中，只是Windows不能够执行 flamegraph.pl 文件
   - 安装[Activeperl](<https://www.activestate.com/products/activeperl/downloads/>)；
   - 执行命令 `perl flamegraph.pl -h` 看看是否安装成功
3.  编译安装 go-torch，和Ubuntu一样，生成的可执行文件用来生成火焰图
   - 这个可执行文件使用会调用全局的 flamegraph.pl ，所以如果系统不能够执行 .pl 文件，那么将无法正常使用。

## 使用

1. 首先，先获得程序执行时的CPU和内存使用的文件

2. 生成流程图命令：

   ```shell
   go tool pprof *filepath # 
   ```

3. 生成火焰图命令：

   ```shell
   ./go-torch *filepath 	# 火焰图,go-touch是编译好的可执行文件
   		-f 				# 指定输出路径 
```
   
   



