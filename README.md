# 部署本地go语言书籍
### go高级编程
地址：
[https://github.com/chai2010/advanced-go-programming-book](https://github.com/chai2010/advanced-go-programming-book)
### go语言圣经
地址：
[https://github.com/gopl-zh/gopl-zh.github.com](https://github.com/gopl-zh/gopl-zh.github.com)
## 教程
1.安装nodejs[https://nodejs.org/en/](https://nodejs.org/en/)

2.安装gitbook插件`npm install gitbook-cli -g`

3.`go get`命令获取书籍源码

4.在源码目录执行`gitbook install`

5.使用`make`命令 生成_book文件夹

6.修改本源码中的书籍_book路径

7.使用`go build -ldflags '-w -s'`编译程序

8.安装upx [https://upx.github.io/](https://upx.github.io/)

9.使用upx -9 gobook命令压缩程序

10.`mv` gobook到path

11.启动本源码中的gobook.sh

12.浏览器使用localhost:9000/9001打开本地书籍

# go入门步骤
### 30分钟熟悉go基本语法
http://www.runoob.com/go/go-tutorial.html
### 熟悉go基本教程
https://tour.golang.org/welcome/1
### 高效实战go编程
https://golang.org/doc/effective_go.html
### go代码评审
https://github.com/golang/go/wiki/CodeReviewComments
### go圣经

### go高级编程