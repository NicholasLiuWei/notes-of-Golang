# Linux下安装Goland

1. 下载安装包：<https://www.jetbrains.com/go/> 

2. 解压后，执行bin目录下的 goland.sh 文件，进行第一次打开

3. 配置 goland 在 Ubuntu 中的环境变量

   1. ```$ vi ~/.bashrc``` 
   2. 在最后补充： ```export PATH=$PATH:goland.sh所在目录的绝对路径 ```
   3. 完成后，在终端中就可以输入goland.sh命令打开Goland，不带./

4. 创建快捷方式

   1. 在 /usr/share/applications 目录中使用 vim 编辑器创建文件 goland.desktop 
   2. 将如下代码写入该文件：

   ```shell
   [Desktop Entry]
   
   Name=GoLand
   
   Comment=GoLand
   # Exec指定启动的文件(根据自己电脑指定)
   Exec=/home/ubuntu/GoLand-2017.3/bin/goland.sh
   # Icon指定启动软件的图标(根据自己电脑指定)
   Icon=/home/ubuntu/GoLand-2017.3/bin/goland.png
   
   Terminal=false
   
   Type=Application
   
   Categories=Developer;
   ```

   3. 将生成的文件 goland.desktop 拖入启动器或者复制到桌面