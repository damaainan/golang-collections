# Summary

* [前言](preface.md)
  * [Go语言起源](ch0/ch0-01.md)
  * [Go语言项目](ch0/ch0-02.md)
  * [本书的组织](ch0/ch0-03.md)
  * [更多的信息](ch0/ch0-04.md)
  * [致谢](ch0/ch0-05.md)
* [入门](ch1入门/ch1-入门.md)
  * [Hello, World](ch1入门/ch1-01-Hello, World.md)
  * [命令行参数](ch1入门/ch1-02-命令行参数.md)
  * [查找重复的行](ch1入门/ch1-03-查找重复的行.md)
  * [GIF动画](ch1入门/ch1-04-GIF动画.md)
  * [获取URL](ch1入门/ch1-05-获取URL.md)
  * [并发获取多个URL](ch1入门/ch1-06-并发获取多个URL.md)
  * [Web服务](ch1入门/ch1-07-Web服务.md)
  * [本章要点](ch1入门/ch1-08-本章要点.md)
* [程序结构](ch2程序结构/ch2-程序结构.md)
  * [命名](ch2程序结构/ch2-01-命名.md)
  * [声明](ch2程序结构/ch2-02-声明.md)
  * [变量](ch2程序结构/ch2-03-变量.md)
  * [赋值](ch2程序结构/ch2-04-赋值.md)
  * [类型](ch2程序结构/ch2-05-类型.md)
  * [包和文件](ch2程序结构/ch2-06-包和文件.md)
  * [作用域](ch2程序结构/ch2-07-作用域.md)
* [基础数据类型](ch3基础数据类型/ch3-基础数据类型.md)
  * [整型](ch3基础数据类型/ch3-01-整型.md)
  * [浮点数](ch3基础数据类型/ch3-02-浮点数.md)
  * [复数](ch3基础数据类型/ch3-03-复数.md)
  * [布尔型](ch3基础数据类型/ch3-04-布尔型.md)
  * [字符串](ch3基础数据类型/ch3-05-字符串.md)
  * [常量](ch3基础数据类型/ch3-06-常量.md)
* [复合数据类型](ch4复合数据类型/ch4-复合数据类型.md)
  * [数组](ch4复合数据类型/ch4-01-数组.md)
  * [Slice](ch4复合数据类型/ch4-02-Slice.md)
  * [Map](ch4复合数据类型/ch4-03-Map.md)
  * [结构体](ch4复合数据类型/ch4-04-结构体.md)
  * [JSON](ch4复合数据类型/ch4-05-JSON.md)
  * [文本和HTML模板](ch4复合数据类型/ch4-06-文本和HTML模板.md)
* [函数](ch5函数/ch5-函数.md)
  * [函数声明](ch5函数/ch5-01-函数声明.md)
  * [递归](ch5函数/ch5-02-递归.md)
  * [多返回值](ch5函数/ch5-03-多返回值.md)
  * [错误](ch5函数/ch5-04-错误.md)
  * [函数值](ch5函数/ch5-05-函数值.md)
  * [匿名函数](ch5函数/ch5-06-匿名函数.md)
  * [可变参数](ch5函数/ch5-07-可变参数.md)
  * [Deferred函数](ch5函数/ch5-08-Deferred函数.md)
  * [Panic异常](ch5函数/ch5-09-Panic异常.md)
  * [Recover捕获异常](ch5函数/ch5-10-Recover捕获异常.md)
* [方法](ch6方法/ch6-方法.md)
  * [方法声明](ch6方法/ch6-01-方法声明.md)
  * [基于指针对象的方法](ch6方法/ch6-02-基于指针对象的方法.md)
  * [通过嵌入结构体来扩展类型](ch6方法/ch6-03-通过嵌入结构体来扩展类型.md)
  * [方法值和方法表达式](ch6方法/ch6-04-方法值和方法表达式.md)
  * [示例: Bit数组](ch6方法/ch6-05-示例: Bit数组.md)
  * [封装](ch6方法/ch6-06-封装.md)
* [接口](ch7接口/ch7-接口.md)
  * [接口是合约](ch7接口/ch7-01-接口是合约.md)
  * [接口类型](ch7接口/ch7-02-接口类型.md)
  * [实现接口的条件](ch7接口/ch7-03-实现接口的条件.md)
  * [flag.Value接口](ch7接口/ch7-04-flag.Value接口.md)
  * [接口值](ch7接口/ch7-05-接口值.md)
  * [sort.Interface接口](ch7接口/ch7-06-sort.Interface接口.md)
  * [http.Handler接口](ch7接口/ch7-07-http.Handler接口.md)
  * [error接口](ch7接口/ch7-08-error接口.md)
  * [示例: 表达式求值](ch7接口/ch7-09-示例: 表达式求值.md)
  * [类型断言](ch7接口/ch7-10-类型断言.md)
  * [基于类型断言识别错误类型](ch7接口/ch7-11-基于类型断言识别错误类型.md)
  * [通过类型断言查询接口](ch7接口/ch7-12-通过类型断言查询接口.md)
  * [类型分支](ch7接口/ch7-13-类型分支.md)
  * [示例: 基于标记的XML解码](ch7接口/ch7-14-示例: 基于标记的XML解码.md)
  * [补充几点](ch7接口/ch7-15-补充几点.md)
* [Goroutines和Channels](ch8Goroutines和Channels/ch8-Goroutines和Channels.md)
  * [Goroutines](ch8Goroutines和Channels/ch8-01-Goroutines.md)
  * [示例: 并发的Clock服务](ch8Goroutines和Channels/ch8-02-示例: 并发的Clock服务.md)
  * [示例: 并发的Echo服务](ch8Goroutines和Channels/ch8-03-示例: 并发的Echo服务.md)
  * [Channels](ch8Goroutines和Channels/ch8-04-Channels.md)
  * [并发的循环](ch8Goroutines和Channels/ch8-05-并发的循环.md)
  * [示例: 并发的Web爬虫](ch8Goroutines和Channels/ch8-06-示例: 并发的Web爬虫.md)
  * [基于select的多路复用](ch8Goroutines和Channels/ch8-07-基于select的多路复用.md)
  * [示例: 并发的目录遍历](ch8Goroutines和Channels/ch8-08-[示例: 并发的目录遍历.md)
  * [并发的退出](ch8Goroutines和Channels/ch8-09-并发的退出.md)
  * [示例: 聊天服务](ch8Goroutines和Channels/ch8-10-示例: 聊天服务.md)
* [基于共享变量的并发](ch9基于共享变量的并发/ch9-基于共享变量的并发.md)
  * [竞争条件](ch9基于共享变量的并发/ch9-01-竞争条件.md)
  * [sync.Mutex互斥锁](ch9基于共享变量的并发/ch9-02-sync.Mutex互斥锁.md)
  * [sync.RWMutex读写锁](ch9基于共享变量的并发/ch9-03-sync.RWMutex读写锁.md)
  * [内存同步](ch9基于共享变量的并发/ch9-04-内存同步.md)
  * [sync.Once惰性初始化](ch9基于共享变量的并发/ch9-05-[sync.Once惰性初始化.md)
  * [竞争条件检测](ch9基于共享变量的并发/ch9-06-竞争条件检测.md)
  * [示例: 并发的非阻塞缓存](ch9基于共享变量的并发/ch9-07-示例: 并发的非阻塞缓存.md)
  * [Goroutines和线程](ch9基于共享变量的并发/ch9-08-Goroutines和线程.md)
* [包和工具](ch10包和工具/ch10-包和工具.md)
  * [包简介](ch10包和工具/ch10-01-包简介.md)
  * [导入路径](ch10包和工具/ch10-02-导入路径.md)
  * [包声明](ch10包和工具/ch10-03-包声明.md)
  * [导入声明](ch10包和工具/ch10-04-导入声明.md)
  * [包的匿名导入](ch10包和工具/ch10-05-包的匿名导入.md)
  * [包和命名](ch10包和工具/ch10-06-包和命名.md)
  * [工具](ch10包和工具/ch10-07-工具.md)
* [测试](ch11测试/ch11-测试.md)
  * [go test](ch11测试/ch11-01-go test.md)
  * [测试函数](ch11测试/ch11-02-测试函数.md)
  * [测试覆盖率](ch11测试/ch11-03-测试覆盖率.md)
  * [基准测试](ch11测试/ch11-04-基准测试.md)
  * [剖析](ch11测试/ch11-05-剖析.md)
  * [示例函数](ch11测试/ch11-06-示例函数.md)
* [反射](ch12反射/ch12-反射.md)
  * [为何需要反射?](ch12反射/ch12-01-为何需要反射?.md)
  * [reflect.Type和reflect.Value](ch12反射/ch12-02-reflect.Type和reflect.Value.md)
  * [Display递归打印](ch12反射/ch12-03-Display递归打印.md)
  * [示例: 编码S表达式](ch12反射/ch12-04-示例: 编码S表达式.md)
  * [通过reflect.Value修改值](ch12反射/ch12-05-通过reflect.Value修改值.md)
  * [示例: 解码S表达式](ch12反射/ch12-06-示例: 解码S表达式.md)
  * [获取结构体字段标签](ch12反射/ch12-07-[获取结构体字段标签.md)
  * [显示一个类型的方法集](ch12反射/ch12-08-显示一个类型的方法集.md)
  * [几点忠告](ch12反射/ch12-09-几点忠告.md)
* [底层编程](ch13底层编程/ch13-底层编程.md)
  * [unsafe.Sizeof, Alignof 和 Offsetof](ch13底层编程/ch13-01-unsafe.Sizeof, Alignof 和 Offsetof.md)
  * [unsafe.Pointer](ch13底层编程/ch13-02-unsafe.Pointer.md)
  * [示例: 深度相等判断](ch13底层编程/ch13-03-示例: 深度相等判断.md)
  * [通过cgo调用C代码](ch13底层编程/ch13-04-通过cgo调用C代码.md)
  * [几点忠告](ch13底层编程/ch13-05-几点忠告.md)
* [附录](appendix/appendix.md)
  * [附录A：原文勘误](appendix/appendix-a-errata.md)
  * [附录B：作者译者](appendix/appendix-b-author.md)
  * [附录C：译文授权](appendix/appendix-c-cpoyright.md)
  * [附录D：其它语言](appendix/appendix-d-translations.md)