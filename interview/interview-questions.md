# 高频面试题

## base

- ▲ map 是线程安全的吗？如何保证map的线程安全？sync.map理解？
- ▲ go 中垃圾回收机制中如何判断对象需要回收？常见的 GC 回收算法有哪些？
- ▲ 深拷贝与浅拷贝区别是什么？
- ▲ 实现单例设计模式（懒汉，饿汉）
- ▲ slice的扩容机制1024以内成倍扩容，超过1024之后以1.25倍扩容
- ▲ go中内存对齐的使用场景
- ▲ 手写生产者消费者模型
- ▲ new与make的区别
- ▲ 什么是内存泄漏，怎么确定内存泄漏？

### # 操作系统

#### # 进程线程

- ▲ 进程有多少种状态？
- ▲ 进程间有哪些通信方式？**
- ▲ 线程间有哪些通信方式？*
- ▲ 简单介绍进程调度的算法
- ▲ 进程和线程之间有什么区别？****
- ▲ 什么情况下，进程会进行切换？
- ▲ 多线程和多进程的区别是什么？*
- ▲ 两个线程交替打印一个共享变量。*
- ▲ 为什么进程切换慢，线程切换快？
- ▲ 简述自旋锁与互斥锁的使用场景。
- ▲ 进程空间从高位到低位都有些什么？
- ▲ 进程通信中的管道实现原理是什么？*
- ▲ 线程有多少种状态，状态之间如何转换

#### # socket

- ▲ I/O多路复用中 select, poll, epoll之间有什么区别，各自支持的最大描述符上限以及原因是什么？
- ▲ 简述 socket 中 select 与 epoll 的使用场景以及区别，epoll 中水平触发以及边缘触发有什么不同？**

#### # linux

- ▲ 简单介绍进程调度的算法。
- ▲ 简述操作系统中的缺页中断
- ▲ 简述操作系统中的缺页中断。*
- ▲ 简述 Linux 零拷贝的原理
- ▲ 简述操作系统如何进行内存管理
- ▲ 简述 traceroute 命令的原理
- ▲ 简述 mmap 的使用场景以及原理
- ▲ 操作系统如何申请以及管理内存的？**
- ▲ 简述操作系统中 malloc 的实现原理
- ▲ 简述 Linux 虚拟内存的页面置换算法。*
- ▲ 简述几个常用的 Linux 命令以及他们的功能。**
- ▲ 简述同步与异步的区别，阻塞与非阻塞的区别。***

- ▲ Linux 如何查看实时的滚动日志？
- ▲ Linux 下如何排查 CPU 以及 内存占用过多？**
- ▲ Linux 下如何查看 CPU 荷载，正在运行的进程，某个端口对应的进程？**
- ▲ Linux 进程调度中有哪些常见算法以及策略？
- ▲ Linux 中虚拟内存和物理内存有什么区别？有什么优点？

- ▲ 什么时候会由用户态陷入内核态？
- ▲ LVS 的 NAT、TUN、DR 原理及区别
- ▲ 操作系统中，虚拟地址与物理地址之间如何映射？
- ▲ 简述 Linux 系统态与用户态，什么时候会进入系统态？
- ▲ 系统调用的过程是怎样的？操作系统是通过什么机制触发系统调用的？
- ▲ BIO、NIO 有什么区别？怎么判断写文件时 Buffer 已经写满？简述 Linux 的 IO模型

### # 网络协议

#### # TCP/IP

- ▲ 简述tcp协议。
  - tcp协议是一种面向连接的、可靠的、基于字节流的传输层同学协议。使用tcp的双方(c/s)在交换数据前，需要通过三次握手来建立tcp连接，建立连接后就可以基于字节流的双工通讯。由tcp内部实现保证通讯的可靠性，完全通讯完成后，通过四次挥手断开连接。

- ▲ 什么是 ARP 协议？

- ▲ 简述 TCP 半连接发生场景。
  - client和server在完成三次握手的过程中，会有一个半连接池(syn  queue)和一个圈连接池(accept queue)。
  - 在第二次握手时，server端在发送ack报文的同时会将连接放在半连接池中，半连接池的大小默认为1024，在收到client端的ack报文(即三第三次握手)之后会将请求从半连接池中取出放到全连接吃中。
  - 如果这个过程中，一直没有收到client的ack请求，服务端会一直重试发送ack+syn请求给client端，linux默认为5次。
  - 从半连接池放到全连接池的过程有多种情况：
    - 1.全连接池未满：从半连接池中放到全连接池中，之后server端accept()请求。
    - 2.全连接池满：tcp_abort_on_overflow=0; server端丢弃client端发送的ack数据，定时重试第二次握手数据，如果client端一直排不上队，报超时异常。
    - 3.全连接池满：tcp_abort_on_overflow=1; server端发送一个reset包给客户端，表示废弃本次连接，client端报错。

- ▲ 简述 TCP 的报文头部结构。
- ▲ 简述 TCP 的 TIME_WAIT。
  - tcp4次挥手之后，连接双方都不在交换数据。单主动关闭的一方保持这个连接在一段时间内不可用。time_wait是为了保证全双工的tcp连接正常终止。

- ▲ 简述 TCP 滑动窗口。**
  - 滑动窗口是为了提升tcp的传输效率
  - 滑动窗口的基本原理就是在任意时刻，发送方都维持了一个连续的允许发送的帧的序号，称为发送窗口。同时，接收方也维持了一个连续允许接收的帧的序号，称为接收窗口。发送和接收窗口的大小火上下界可以不一样。

- tcp的超时重传机制。*
  - tcp是一种面向连接的可靠的传输层协议，保证了数据的可靠传输，对于一些出错，数据丢包等问题，。tcp设计了超时与重传机制。
  - 基本原理：在发送一个数据之后，开启一个定时器，若是在这个时间内没有收到发送数据的ack确认报文，则对改报文进行重传，在达到一定次数还没有成功时放弃并发送一个复位信号。

- ▲ 简述 TCP 协议的延迟 ACK 和累计应答。
  - 接收方在收到数据后，并不会立即回复ACK,而是延迟一定时间。一般ACK延迟发送的时间为200ms，但这个200ms并非收到数据后需要延迟的时间。系统有一个固定的定时器每隔200ms会来检查是否需要发送ACK包。这样做有两个目的。
    - 1、这样做的目的是ACK是可以合并的，也就是指如果连续收到两个TCP包，并不一定需要ACK两次，只要回复最终的ACK就可以了，可以降低网络流量。
    - 2、如果接收方有数据要发送，那么就会在发送数据的TCP数据包里，带上ACK信息。这样做，可以避免大量的ACK以一个单独的TCP包发送，减少了网络流量。
  - ack机制：接收方在接收到数据后，不是立即会给发送方发送ACK的。
    - 1、收到数据包的序号前面还有需要接收的数据包。因为发送方发送数据时，并不是需要等上次发送数据被Ack就可以继续发送TCP包，而这些TCP数据包达到的顺序是不保证的，这样接收方可能先接收到后发送的TCP包（注意提交给应用层时是保证顺序的）。
    - 2、为了降低网络流量，ACK有延迟确认机制。
    - 3、ACK的值到达最大值后，又会从0开始。

- ▲ 简述 TCP 中的拥塞控制。
  - 慢启动：是tcp的一个拥塞控制机制，慢启动算法的基本思想是当tcp开始在一个网络中传输数据或发现数据丢失并开始重发时，首先慢慢的对网路实际容量进行试探，避免由于发送了过量的数据而导致阻塞。
  - 拥塞避免算法：网络中的拥塞发生会导致数据分组丢失，需要尽量避免。实际中，拥塞算法与慢启动通常一起实现。
- ▲ 简述 TCP 三次握手以及四次挥手的流程。为什么需要三次握手以及四次挥手？****
  - 3次握手
    - 1.client向server端发送建立tcp连接的请求报文，其中包含seq序号，为client随机生成的数字x。并且将报文中的syn置为1，表示需要建立tcp连接。(syn=1,seq=x)
    - 2.server端回复报文，其中seq为随机生成的数字y，ack为x+1，syn为1.(syn=1,ack=x+1,seq=y)
    - 3.client端收到server端报文后，回复请求做ack验证，在服务端发送过来的seq+1.(syn=1,ack=y+1,seq=x+1)
    - 完成三次握手之后，server端进入established模式，这时候尝试将消息放入accept queue中，完全tcp连接。
  - 4次挥手
    - 1.client发送请求断开连接的报文，其中包含随机生成的seq数字x。(fin=1,seq=x)
    - 2.server收到报文后回复报文，seq为随机生成的数字y,ack为x+1,以便client知道请求已经得到验证。(ACK=1,seq=y,ack=x+1)
    - 3.server并不会立即断开连接，而是等待传送到client端的数据发送完毕，之后会生成随机的seq返回报文。(FIN=1,ACK=1,seq=z,ack=x+1)
    - 4.client收到断开连接的请求后，会回复server端的断开连接。(FIN=1,seq=x+1,ack=z+1)。
    - 完成tcp断开连接。

- ▲ TCP 怎么保证可靠传输？*** [参考](https://www.cnblogs.com/deliver/p/5471231.html)
  - 确认和重传：接收方收到报文就会确认，发送方发送一段时间之后没有收到确认就重传。
  - 数据校验
  - 数据合理的分片和排序
  - 流量控制：当接收方来不及处理发送方的数据，能提升发送方降低发送的速率，防止包丢失
  - 拥塞控制：当网络拥塞时，减少数据的发送

- ▲ TCP 中常见的拥塞控制算法有哪些？**
  - Reno、Cubic

- ▲ TCP 中 SYN 攻击是什么？如何防止？**
  - SYN flooding攻击：是指client端恶意发送syn连接给server端，并且不做ack回应，将server端的半连接池恶意占满。
  - linux实现了一种叫做 SYNcookie 的机制。简单说就是讲client端的连接信息保存在ISN中返回给客户端，在client端发送ack回应的时候，从ISN中解析出连接信息完成三次握手，避免半连接池占满。

- ▲ TCP长连接和短连接有那么不同的使用场景？
  - 长连接：多用于操作频繁，点对点的同学，而且连接数量不太多的场景。如微信的扫描登录，聊天室等
  - 短连接：用户无需频繁操作，不需要一直获取服务端反馈的情况下，web网站的http服务一般都是有短连接

- ▲ TCP 在什么情况下服务端会出现大量 CLOSE_WAIT ？
  - 客户端主动关闭了socket连接，发送了FIN报文，服务端也发送了ACK报文，此时客户端处于FIN_WAIT_2状态，服务端处于CLOSE_WAIT
  - 服务端没有发送第二次FIN报文，一般都是服务端已经收到了客户端请求关闭，但是服务端未为关闭连接释放资源导致的。

- ▲ TCP的拥塞控制具体是怎么实现的？UDP有拥塞控制吗？

- ▲ TCP 四次挥手的时候 CLOSE_WAIT 的话怎么处理？***
  - 对方关闭连接后，自身程序没有检测。(被动方角度)
  - 本身忘了需要关闭连接，导致整个资源一直被程序占用。(主动方角度)
  - 解决：一般是由程序bug引起的，修改bug，及时释放资源，然后及时测试上线。

- ▲ TCP四次挥手过程以及所处状态，为什么还需要有 time_wait？
  - 被动关闭放发送fin(第3次挥手)，并等待主动关闭端返回ACK(第4次挥手)
  - 若最终ACK丢失(第4次挥手失败)，被动关闭放将重新发送fin(第3次挥手)，主动关闭方必须维护time_wait,保证自己可以接收，然后再重新发送ack，不能让主动方发送完报文之后立即进入close状态。
  - time_wait带来的问题
    - 主动断开放处于time_wait状态时，源端口无法使用
    - 端口最大数是65535，因此如果频繁主动断开tcp连接，将很快耗尽端口号。一旦达到上限，新的请求将无法被处理，将出现大量`Too Many OPen File异常`，还可能导致nginx，apache等挂掉。
  - 解决time_wait问题
    - 核心思想就是打开系统的time_wait的重用和快速回收机制。
    - net.ipv4.tcp_tw_recycle = 1表示开启TCP连接中TIME_WAIT sockets的快速回收，默认为0,表示关闭
    - net.ipv4.tcp_tw_reuse = 1 表示开启重用。允许将TIME_WAIT sockets重新用于新的TCP连接，默认为0，表示关闭

- ▲ TCP 与 UDP 在网络协议中的哪一层，他们之间有什么区别？***
  - 在网络协议中的传输层。
  - 区别：tcp面向连接的，udp无连接、tcp传输可靠，使用流量控制和拥塞控制，udp传输不可靠、tcp面向字节流，up面向报文、tcp首部开销大，首部最小20字节，最大60字节。udp首部开销小，仅8字节。

- ▲ TCP 的 keepalive 了解吗？说一说它和 http 的 keepalive 的区别？**
  - TCP 的 keepalive
    - tcp在客户端和服务端建立连接后，双方长时间未通讯时，通过tcp的报货机制即keepalive来确定对方连接是否健康且具有通讯能力。keepalive默认是关闭的，tcp的任何一方都可打开此功能。开启后如果在一段时间(报货时间：tcp_keepalive_time)内此连接都不活跃，开启保活功能的一端发送一个保活探测报文探测对方状态。
    - keepalive三个核心参数：tcp_keepalive_time、探测时间间隔：tcp_keepalive_intvl、探测循环次数：tcp_keepalive_probes。这三个参数，在linux上可以在/proc/sys/net/ipv4/路径下找到，或者通过sysctl -a | grep keepalive命令查看当前内核运行参数
  - http keep-alive与tcpkeepalive区别
    - 意图不一样，http keep-alivesh为了让tcp连接活得更久，以便在同一个连接上传送多个http，提高socket的效率。而tcp keepalive是tcp的一种保活机制。

- ▲ 从系统层面上，UDP如何保证尽量可靠？
  - 增加数据可靠性，例如重传等功能，保障数据的完整性
- ▲ 简述 OSI 七层模型，TCP，IP 属于哪一层？*
  - 物数网传会表应。一般使用的话按照四层划分分别为链路层(物数)、网络层、传输层、应用层
    - 链路层：负责封装和解封装ip报文，发送和接受arp/rarp报文等。
    - 网络层：负责路由以及把分组报文发送给目标网络或主机。
    - 传输层：负责对报文进行分组和重组，并以tcp或udp协议格式封装报文。
    - 应用层：负责向用户提供应用程序，如http，ftp，Telnet，dns，smtp等。
- ▲ DNS 查询服务器的基本流程是什么？DNS 劫持是什么？**

#### # http/https/http1.1/http2

- ▲ HTTP 的方法有哪些？** [http参考](https://zhuanlan.zhihu.com/p/72616216)
  - GET：请求指定的页面信息，并返回实体主体。
  - POST：向指定资源提交数据进行处理请求(如表单提交或文件上传),数据包含在请求体中。post请求可能会导致新的资源建立或已有资源修改。
  - HEAD：类似于get请求，只不过返回的响应数据中没有具体类容，用于获取报头信息。
  - PUT：从客户端想服务器传送的数据取代指定的文档的内容。
  - DELETE：请求服务器删除指定的页面。

- ▲ HTTP 中 GET 和 POST 区别。**
  - 都包含请求头请求行，post多了请求body。
  - get多用来查询，请求参数放在url中，不会对服务器上的内容产生作用。post一般用于提交。如账号密码放入body中。
  - get是直接添加到URL后面的，直接就可以在url中看到内容，而post是放在报文内部，用户无法直接看到。
  - get提交的数据长度有限制。具体限制由浏览器决定。而post没有限制。

- ▲ HTTP 与 HTTPS 有哪些区别？****
  - https是http协议的安全版本，http协议的数据传输都是明文的，是不安全的。https使用了ssl/tls协议进行加密处理
  - http和https使用连接方式不同，默认端口不同。http80、https443。

- ▲ 一次 HTTP 的请求过程中发生了什么？**** [参考](https://blog.csdn.net/qq_40804005/article/details/82876209)
  - 域名解析，获取对应的ip地址
  - 发起tcp请求，3次握手建立连接
  - 建立连接后发起http请求
  - 服务端响应http请求，客户端收到响应
  - 客户端解析响应数据

- ▲ 简述对称与非对称加密的概念。**
- ▲ 简述 HTTPS 的加密与认证过程。***
  - https解决的问题：
    - 信息加密传输：防止三方窃听截取
    - 校验机制：数据完整性校验，若数据被篡改，双方立马知道
    - 身份证书：防止身份被冒充
  - https传输流程
    - 客户端请求https连接，服务器返回证书(公钥)
    - 客户端收到证书(公钥)，校验证书，生成随机(对称)秘钥
    - 使用公钥对 对称秘钥 加密, 然后发送加密后的秘钥给服务端
    - 服务端利用自己的私钥解密出对称秘钥。
    - 服务端使用解密的秘钥加密信息发送给客户端,与客户端通信。

- ▲ 简述 HTTP 1.0，1.1，2.0 的主要区别。*** [参考](https://www.cnblogs.com/heluan/p/8620312.html)
  - http1.0和http1.1的区别。
    - 1.缓存处理。在HTTP1.0中主要使用header里的If-Modified-Since,Expires来做为缓存判断的标准，HTTP1.1则引入了更多的缓存控制策略例如Entity tag，If-Unmodified-Since, If-Match, If-None-Match等更多可供选择的缓存头来控制缓存策略。
    - 2.带宽优化及网络连接的使用，HTTP1.0中，存在一些浪费带宽的现象，例如客户端只是需要某个对象的一部分，而服务器却将整个对象送过来了，并且不支持断点续传功能，HTTP1.1则在请求头引入了range头域，它允许只请求资源的某个部分，即返回码是206（Partial Content），这样就方便了开发者自由的选择以便于充分利用带宽和连接。
    - 3.错误通知的管理，在HTTP1.1中新增了24个错误状态响应码，如409（Conflict）表示请求的资源与资源的当前状态发生冲突；410（Gone）表示服务器上的某个资源被永久性的删除。
    - 4.Host头处理，在HTTP1.0中认为每台服务器都绑定一个唯一的IP地址，因此，请求消息中的URL并没有传递主机名（hostname）。但随着虚拟主机技术的发展，在一台物理服务器上可以存在多个虚拟主机（Multi-homed Web Servers），并且它们共享一个IP地址。HTTP1.1的请求消息和响应消息都应支持Host头域，且请求消息中如果没有Host头域会报告一个错误（400 Bad Request）。
    - 5.长连接，HTTP 1.1支持长连接（PersistentConnection）和请求的流水线（Pipelining）处理，在一个TCP连接上可以传送多个HTTP请求和响应，减少了建立和关闭连接的消耗和延迟，在HTTP1.1中默认开启Connection： keep-alive，一定程度上弥补了HTTP1.0每次请求都要创建连接的缺点。
  - HTTP2.0和http1.x的新特性
    - 新的二进制格式：http1.x解析基于文本。基于文本协议的格式解析存在天然缺陷，文本的表现形式有多样性，要做到健壮性考虑的场景必然很多，二进制则不同，只认0和1的组合。基于这种考虑HTTP2.0的协议解析决定采用二进制格式，实现方便且健壮。
    - 多路复用(MultiPlexing)：即连接共享。每一个request都是用做连接共享机制的。一个request对应一个id，这样一个连接上可以有多个request，每个连接的request可以随机的混杂在一起，接收方可以根据request的id将request再归属到各自不同的服务端请求里面。
    - header压缩(HPack压缩算法)：http1.x的header带有大量的信息，而且每次都啊哟重复发送，http2.0使用encoder来来减少需要传输的header大小，通讯双方各自缓存一份header fields表，即避免了重复header传输，有减小了需要传输的大小。
    - 服务端推送(server push)：服务端推送能把客户端所需要的资源伴随着index.html一起发送到客户端，省去了客户端重复请求的步骤。正因为没有发起请求，建立连接等操作，所以静态资源通过服务端推送的方式可以极大地提升速度。例如我的网页有一个sytle.css的请求，在客户端收到sytle.css数据的同时，服务端会将sytle.js的文件推送给客户端，当客户端再次尝试获取sytle.js时就可以直接从缓存中获取到，不用再发请求了。
  - http2.0的多路复用和http1.x的长连接复用区别？
    - http1.x：一次请求-响应，建立一个连接，用完关闭每一个请求都要建立一个连接。
    - http1.1：pipline解决方式为。 若干个请求排队串行化单线程处理，后面的请求等待前面的请求返回才能获得执行机会。一旦有请求超时，后续请求只能被阻塞，别无他法。也就是常说的线头阻塞。
    - http2多个请求可以同时在一个连接上并行执行。某个请求任务耗时严重，不会影响其他链接的正常执行。

- ▲ 简述常见的 HTTP 状态码的含义（301，304，401，403）
  - 状态码分类：
    - 1XX-信息型：服务器收到请求，需要请求者继续操作。
    - 2XX-成功型：请求成功收到，理解并处理。
    - 3XX-重定向：需要进一步操作以完成请求。
    - 4XX-客户端错误：客户端错误，请求包含语法错误或无法完成请求，权限等。
    - 5XX-服务器错误：服务器再处理请求过程中发生了错误。
  - 常见状态码：
    - 200 客户端请求成功
    - 301 资源(网页等)被永久转移到其他URL
    - 302 临时跳转
    - 400 Bad Request 客户端请求语法错误，不能被服务器所完全理解。
    - 401 Unauthorized 请求未经授权，必须和www-authenticate报头域一起使用
    - 403 Forbidden 资源不可用，客户端权限不足
    - 404 请求资源不存在，可能是输入了错误的url
    - 500 服务器内部发了不可预期的错误
    - 503 Server Unavailable 服务完全当前不能处理客户端的请求，一段时间后可能恢复。
- ▲ SSL握手流程为什么要使用对称秘钥？**
  - 非对称加密的加解密效率是非常低的，而http的应用场景中通常端与端之间存在大量的交互，使用非对称加密的效率无法接受。
  - https的场景中只有服务端保存了私钥，一对公私钥只能实现单向的加解密，所以https内容传输采用的是对称加密。

- URI和URL的区别
  - http使用统一资源标识符(URI)来传输和建立连接
    - URI：Uniform Resourse Identifier 统一资源标识符
    - URL：Uniform Resourse Location 统一资源定位符
  - URI是用来标识 一个具体的资源的，可以通过URI知道一个资源是什么
  - URL用来定位具体的资源的，标示了一个具体的资源位置。互联网上的每个文件都有唯一的一个URL。

#### # socket/websocket

- ▲ 简述 WebSocket 是如何进行传输的。**
  WebSocket 是 HTML5 开始提供的一种在单个 TCP 连接上进行全双工通讯的协议。
  在 WebSocket API 中，浏览器和服务器只需要做一个握手的动作，然后，浏览器和服务器之间就形成了一条快速通道。两者之间就直接可以数据互相传送。
  浏览器通过 JavaScript 向服务器发出建立 WebSocket 连接的请求，
  连接建立以后，客户端和服务器端就可以通过 TCP 连接直接交换数据。
  当获取到 Web Socket 连接后，就可以通过 send() 方法来向服务器发送数据，
  并通过 onmessage 事件来接收服务器返回的数据。

#### # rpc/grpc/restful

- 简述rpc是什么？
  - rpc(Remote Procedure Call) 远程过程调用，他是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。rpc协议假定某些传输协议的存在，如tcp或udp，为通信程序之间携带信息数据。在osi网络通信模型中，rpc跨越了传输层和应用层。rpc使得开发包括网络分布式多程序在内的应用程序更加容易。

- ▲ RestFul 是什么？RestFul 请求的 URL 有什么特点？** [参考](https://zhuanlan.zhihu.com/p/334809573)
  - rest是一种架构风格，指的是一组架构约束条件和原则。满足这些约束条件和原则的应用程序或设计就是restful。reste规范把所有内容都是为资源。网络上一切皆资源。
  - restful的url的path通常组成为 /{version}/{resources}/{resource_id}
    - version：api版本号，有些版本号放置在头信息中也可以，通过控制版本号有利于应用迭代
    - resource：资源，restflu api推荐通小写英文单词的复数形式。
    - resource_id:资源的id，访问或操作改资源
  - restful api的url具体设计规范如下：
    - 1.不同大写字母，所有单词使用英文且小写。
    - 2.连字符用中杆“-”，而不用下杠"_"。
    - 3.正确使用"/"表示层级关系，url的层级不要过深，且越靠前的层级应相对越稳定。
    - 4.结尾不要包含正斜杠分隔符"/"。
    - 5.url中不要出现动词，请求方式表示动作。
    - 6.资源表示用复数不要用单数。
    - 7.不要使用文件拓展名。

- ▲ RestFul 与 RPC 的区别是什么？RestFul 的优点在哪里？
  - rest优缺点
    - 优点：耦合性低，兼容性好，熟悉后能提高开发效率，不用关系接口实现细节，相对更加规范，更加标准通用。跨语言支持。
    - 缺点：性能不如rpc高。
  - rpc优缺点
    - 优点：
      - 调用简单，清晰、透明。不用像rest一样复杂。就像调用背调方法一样简单。
      - 高效低延迟，性能高
      - 自定义协议(可以让传输报文变更小)
      - 性能消耗低，高效的序列化协议可以支持高效的二进制传输
      - 自带负载均衡
    - 缺点：耦合性强
  - rest和rpc选择
    - rpc适用于内网服务调用，对外提供服务一般采用rest
    - io密集的服务调用用rpc，低频服务用rest
    - 服务调用过于密集与复杂的使用rpc

#### # web安全/cooke/session

- ▲ 简述 JWT 的原理和校验机制？**
- ▲ 如何设计 API 接口使其实现幂等性？
- ▲ TCP 中 SYN 攻击是什么？如何防止？
- ▲ Cookie和Session的关系和区别是什么？***
- ▲ 什么是 SYN flood，如何防止这类攻击？
- ▲ 简述什么是 XSS 攻击以及 CSRF 攻击？**
- ▲ 什么是跨域，什么情况下会发生跨域请求？***
- ▲ DNS 查询服务器的基本流程是什么？DNS 劫持是什么？**

### # 算法

#### 排序

- ▲ 实现快速排序。***
- ▲ 实现归并排序。*
- ▲ 使用递归及非递归两种方式实现快速排序。**
- ▲ 常用的排序方式有哪些，时间复杂度是多少？**
- ▲ 64 匹马，8 个赛道，找出前 4 匹马最少需要比几次？*
- ▲ 快速排序的空间复杂度是多少？时间复杂度的最好最坏的情况是多少，有哪些优化方案？**
- ▲ 给定一个包含 40亿 个无符号整数的大型文件，使用最多 1G 内存，对此文件进行排序。

#### 链表、树

- ▲ 二叉树的层序遍历? *
- ▲ 有序链表插入的时间复杂度是多少？
- ▲ AVL 树和红黑树有什么区别？*
- ▲ 红黑树是怎么实现平衡的？它的优点是什么？*
- ▲ Hash 表常见操作的时间复杂度是多少？遇到 Hash 冲突是如何解决的？

#### topK查找

- ▲ 如何实现大数运算？*
- ▲ 如何随机生成不重复的 10个100 以内的数字？
- ▲ 10亿个数中如何高效地找到最大的一个数以及最大的第 K 个数。*
- ▲ 10亿条数据包括 id，上线时间，下线时间，请绘制每一秒在线人数的曲线图
- ▲ 给定 100G 的 URL 磁盘数据，使用最多 1G 内存，统计出现频率最高的 Top K 个 URL。*
- ▲ 两个文件包含无序的数字，数字的大小范围是0-500w左右。如何求两个文件中的重复的数据？*
- ▲ 两个 10G 大小包含 URL 数据的文件，最多使用 1G 内存，将这两个文件合并，并找到相同的 URL。*
- ▲ 1000台 机器，每台机器 1000个 文件，每个文件存储了 10亿个 整数，如何找到其中最小的 1000个 值？

#### Leetcode

- ▲ 丑数 II (Leetcode 264)*
- ▲ 最小栈 (Leetcode)*
- ▲ 爬楼梯 (Leetcode)*
- ▲ 路径总和 (Leetcode) *
- ▲ 环形链表 (Leetcode)*
- ▲ 反转链表 (Leetcode)*

- ▲ 两数相加 II (Leetcode 445) **
- ▲ 旋转数组 (Leetcode)*
- ▲ 旋转图像 (Leetcode)
- ▲ 按序打印 (Leetcode) **
- ▲ 多数元素 (Leetcode)*

- ▲ 用栈实现队列 (Leetcode)**
- ▲ 最大子序和 (Leetcode)**
- ▲ 最长公共子序列 (Leetcode)
- ▲ 最长连续子序列 (Leetcode)

- ▲ 二叉树的前序遍历 (Leetcode)**
- ▲ 二叉树的最近公共祖先 (Leetcode 236) ***
- ▲ 二叉树的锯齿形层次遍历 (Leetcode)

- ▲ 数组中的逆序对 (Leetcode)
- ▲ 和为 K 的子数组 (Leetcode)
- ▲ 搜索旋转排序数组 (Leetcode)**
- ▲ 旋转数组的最小数字 (Leetcode)*
- ▲ K 个一组翻转链表 (Leetcode 25)*
- ▲ 搜索旋转排序数组 II (Leetcode)**
- ▲ 数组中的第 K 个最大元素 (Leetcode)*
- ▲ 寻找旋转排序数组中的最小值 (leetcode)

- ▲ 合并两个有序链表 (Leetcode)*
- ▲ 链表倒数第K个数 (Leetcode)*
- ▲ 判断有环链表的环长度 (Leetcode)*
- ▲ 链表倒数第K个数 (Leetcode)
- ▲ 合并两个有序链表 (Leetcode)
- ▲ 删除排序链表中的重复元素 (Leetcode)

- ▲ 第一个只出现一次的字符 (Leetcode)
- ▲ 剑指 Offer 10- ▲ II. 青蛙跳台阶问题。*
- ▲ 用 Rand7() 实现 Rand10() (Leetcode)*
- ▲ 10亿个数中如何高效地找到最大的一个数以及最大的第 K 个数。*

#### 负载均衡 lru

- ▲ 简述常见的负载均衡算法。
- ▲ 简述 LRU 算法及其实现方式。**
- ▲ 如何从一个数组输出随机数组。
- ▲ 如果通过一个不均匀的硬币得到公平的结果？
- ▲ 常用的限流算法有哪些？简述令牌桶算法原理。
- ▲ 如何随机生成不重复的 10个100 以内的数字？
- ▲ 实现 LRU 算法，实现带有过期时间的 LRU 算法。
- ▲ 给定一个 foo 函数，60%的概率返回0，40%的概率返回1，如何利用 foo 函数实现一个 50% 返回 0 的函数？

### # 数据库

#### # mysql

- ▲ 模糊查询是如何实现的？
  - mysql的模糊查询like,需要用到通配符%。

- ▲ 简述 MySQL 三种日志的使用场景。

- ▲ 数据库的读写分离的作用是什么？如何实现？

- ▲ MySQL 中 join 与 left join 的区别是什么？**
  - join等价于inner join内连接抄，是返回两个表中都有的符合条件的行。
  - left join左连接，是返回左袭表知中所有的行及右表中符合条件的行。
  - right join右连接，是返回右表中所有的行及左表中符合条件的行。
  - full join全连接，是返回左表中所有的行及右表中所有的行，并按条件连接。
  - 通常情况下，left join肯定比inner join返回的行数多道。
- ▲ 简述 MySQL 的主从同步机制，如果同步失败会怎么样？

##### 存储引擎

- ▲ MySQL 有哪些常见的存储引擎？**
- ▲ 简述脏读和幻读的发生场景，InnoDB 是如何解决幻读的？**

##### 索引

- ▲ 联合索引的存储结构是什么？
  - 对于复合索引（多列b+tree，使用多列值组合而成的b+tree索引）。遵循最左侧原则，从左到右的使用索引中的字段，一个查询可以只使用索引中的一部份，但只能是最左侧部分。
- ▲ 数据库索引的实现原理是什么？
- ▲ MySQL 的索引什么情况下会失效？
- ▲ 简述 MySQL 常见索引数据，介绍一下覆盖索引。*
- ▲ 数据库有哪些常见索引？数据库设计的范式是什么？*
- ▲ 聚簇索引和非聚簇索引有什么区别？什么情况用聚集索引？*
  - 聚簇索引：将数据存储于索引放到一起，找到索引页找到了数据
  - 非聚簇索引：将数据存储与索引分开，索引结构的叶子节点指向了数据对应的行。myisam通过key_buffer把索引先缓存到内存中，当需要访问数据时(通过索引访问数据)，在内存中直接搜索索引，然后通过索引找到磁盘相应数据，也就是为什么索引不在key_buffer命中是，速度慢的原因。
- ▲ 唯一索引与普通索引的区别是什么？使用索引会有哪些优缺点？*
- ▲ MySQL 为什么使用 B+ 树来作索引，对比 B 树它的优点和缺点是什么？***
- ▲ 假设建立联合索引 (a, b, c) 如果对字段 a 和 c 查询，会用到这个联合索引吗？
  - 不会，会使用全表查询，当建立联合索引之后，需要满足最左原则才会使用联合索引。所以在a不确定的情况下将不会使用联合索引。

##### 事务

- ▲ 简述事务的四大特性。*
  - 原子性：一个事务执行时不可分割的基本工作单位，事务中的操作包括存储等。要么一起做，要么一起不做。
  - 一致性：事务的执行结果必须是使得数据库的操作从一个一致性状态到另外一个一致性状态。
  - 隔离性：一个事务的执行不能干扰其他事务的执行，也不被其他事务的执行干扰。
  - 持久性：一个事务的提交，它对数据库中的数据的改变将是永久性的改变。接下来的其他操作或故障都不会影响当前做操结果。
- ▲ 简述数据库中的 ACID 分别是什么？**
  - 原子性(Atomicity)、一致性(Consistency)、隔离性(Isolation )、持久性(Durability)。
- ▲ 并发事务会引发哪些问题？如何解决？
- ▲ 数据库的事务隔离级别有哪些？各有哪些优缺点？***
- ▲ 什么是数据库事务，MySQL 为什么会使用 InnoDB 作为默认选项。*

##### 分表分库

- ▲ 简述什么是最左匹配原则 [参考](https://www.jianshu.com/p/fd781d6e1158)
  - 因为联合索引的B+Tree是按照第一个关键字进行索引排列，
- ▲ 简述一致性哈希算法的实现方式及原理。
- ▲ SQL优化的方案有哪些，如何定位问题并解决问题？
- ▲ 简述数据库中什么情况下进行分库，什么情况下进行分表？

##### 锁

- ▲ 简述乐观锁以及悲观锁的区别以及使用场景。*
  - 乐观锁(Optimistic locking)：对加锁持有一种乐观的态度。即先进行业务操作，不到最后一步不进行加锁操作。`乐观` 认为加锁一定会成功。在最后一步更新数据的时候才进行加锁。
  - 悲观锁(Pressimistic Locking)：悲观锁对数据加锁持有一种悲观态度。因此，在整个数据处理过程中，将数据处于锁定状态。悲观锁的实现，往往依靠数据库提供的锁机制(也只有数据库层提供的锁机制才能真正保证数据访问排他性。否则，即使在本系统中实现了锁机制，也无法保证外部系统不会修改数据)。
  - 乐观锁适用于读多写少，并发冲突较少的场景。悲观锁适用于写多读少，对强一致性要求较高的场景。

- ▲ 什么情况下会发生死锁，如何解决死锁？*
  - 多个线程竞争资源而造成的一种互相等待。
    - 1.资源互斥使用。
    - 2.多个进程保持一定的资源，但又请求新的资源。
    - 3.资源不可被剥夺。
    - 4.多个进程循环等待。
  - 数据库产生死锁：
    - 当表进行了分区并且ALTERTABLE的LOCK_ESCALATION设置设为AUTO时也会发生死锁。
  - 解决死锁：顺序加锁，及时释放锁

- ▲ Kafka 发送消息是如何保证可靠性的？

#### # redis

- ▲ Redis 序列化有哪些方式？
- ▲ 为什么 Redis 在单线程下能如此快？*
  - 纯内存操作，避免了大量访问数据库，减少了直接读取磁盘数据，读写数据的时候都不会受到磁盘io速度限制，所以速度快。
  - 单线程操作，避免了不必要的上下文切换和竞争条件，也不存在多进程或多线程的切换而消耗cpu资源，不需要考虑各种锁的问题。不存在加解锁操作，避免了死锁导致的阻塞等。
  - 采用非阻塞io多路复用机制

- ▲ 简述 Redis 中跳表的应用以及优缺点。*
- ▲ 简述 Redis 如何处理热点 key 访问？
  - 做二级缓存
  - 备份热key
- ▲ 简述 Redis 中常见类型的底层数据结构。
  - 动态字符串、压缩链表、跳跃表、哈希表
- ▲ 简述 Redis 的过期机制和缓存淘汰策略。*
- ▲ 简述 Redis 的线程模型以及底层架构设计。
- ▲ 简述 Redis 中如何防止缓存雪崩和缓存击穿。*
- ▲ Redis 有几种数据结构？Zset 是如何实现的？* [参考1](https://www.cnblogs.com/MouseDong/p/11134039.html) [参考2](https://zhuanlan.zhihu.com/p/92536201)
  - 对外提供的数据类型由5种
    - 字符串(string)：字符串对象底层数据结构实现为简单动态字符串（SDS）和直接存储，但其编码方式可以是int、raw或者embstr，区别在于内存结构的不同。
    - 列表(list)：列表对象的编码可以是ziplist和linkedlist之一。
      - ziplist编码的哈希随想底层实现是压缩列表，每个压缩里列表节点保存了一个列表元素
      - linkedlist编码底层采用双端链表实现，每个双端链表节点都保存了一个字符串对象，在每个字符串对象内保存了一个列表元素。
      - 列表对象编码转换：列表对象使用ziplist编码需要满足两个条件：一是所有字符串长度都小于64字节，二是元素数量小于512，不满足任意一个都会使用linkedlist编码。
    - 哈希(hash)：哈希对象的编码可以是ziplist和hashtable之一。
      - 哈希对象编码转换：哈希对象使用ziplist编码需要满足两个条件：一是所有键值对的键和值的字符串长度都小于64字节；二是键值对数量小于512个；不满足任意一个都使用hashtable编码。
    - 集合(set)：集合对象的编码可以是intset和hashtable之一
      - 集合对象编码转换：集合对象使用intset编码需要满足两个条件：一是所有元素都是整数值；二是元素个数小于等于512个；不满足任意一条都将使用hashtable编码
    - 有序集合(SortedSet),即ZSet：有序集合的编码可以是ziplist和skiplist之一
      - 有序集合对象使用ziplist编码需要满足两个条件：一是所有元素长度小于64字节；二是元素个数小于128个；不满足任意一条件将使用skiplist编码
  - ZSet：
    - zset的数据结构同时包含一个字典和一个跳跃表，跳跃表按score从小到大保存所有集合元素。
    - 字典保存着从member到score的映射，这两种数据结构通过指针共享相同元素的member和score，不会造成额外的内存浪费。
  - 跳跃表：跳表是一个概率型的数据结构，元素的插入层数是随机指定的。
  - 层数计算过程：
    - 1.指定节点最大层数MAXLevel，指定概率P，默认层数lvl 为1。
    - 2.生成一个0~1的随机数r，若r < p,且lvl < MAXLevel,则lvl++。
    - 3.重复第2步，直至生成r > p为止，这时候的lvl就是要插入的层数。

- ▲ 假设Redis 的 master 节点宕机了，你会怎么进行数据恢复？

##### 数据备份

- ▲ 简述 Redis 的哨兵机制 *
  - 哨兵是对Redis的系统的运行情况的监控，它是一个独立进程，功能有二个：
    - 监控主从数据库是否正常运行。
    - 主数据出现故障后自动将从数据库转化为主数据库。

- ▲ 简述 Redis 持久化中 rdb 以及 aof 方案的优缺点。*
  - rdb是根据指定的贵州定时将内存中的数据备份到硬盘上，所以rdb备份文件是一个二进制文件。适合做冷备，全量复制的场景，相比aof恢复更快
  - aof是在每次只需命令后命令本身记录下来，即备份文件是一个文本文件。备份以append-only的方式追加，写入性能高，可以更好的保护数据不被丢失，若备份文件过大，恢复慢。
  
- ▲ Redis 中，sentinel和 cluster 的区别和适用场景是什么？

##### 分布式锁

- ▲ Redis 如何实现分布式锁？*
- ▲ Redis 如何实现延时队列，分布式锁的实现原理。
  - 延时队列可以用zset实现，根据score来判断是否到执行时间。
- 分布式锁：
  - 分布式锁常见的三种实现方式：
    - 数据库乐观锁；
    - 基于Redis的分布式锁；
    - 基于ZooKeeper的分布式锁
  - 一个可靠的、高可用的分布式锁需要满足以下几点：
    - 互斥性：任意时刻只能有一个客户端拥有锁，不能被多个客户端获取
    - 安全性：锁只能被持有该锁的客户端删除，不能被其它客户端删除
    - 死锁：获取锁的客户端因为某些原因而宕机，而未能释放锁，其它客户端也就无法获取该锁，需要有机制来避免该类问题的发生
    - 高可用：当部分节点宕机，客户端仍能获取锁或者释放锁

### # 系统设计

- ▲ 简述 CAP理论 ？什么是最终一致性？什么是幂等操作？***
  cap理论最多最多满足其中两个特性，分布式系统要么满足ca，要么满足cp，要么ap。
  - 一致性(Consistency)：在分布式系统完成某写操作后任何操作，都应该读取到该写操作写入的那个最新的值，相当于要求分布式系统中各节点时刻保持数据的一致性。
  - 可用性(Availability)：一直可以保持正常的读写操作。
  - 分区容错性(Partition tolerance)：分布式系统的某个节点或者网络分区出现故障的时候，整个系统任然能对外提供满足一致性和可用性的服务。
  - 一般常见是使用是cp(zk,redis,etcd)，少量的ap(12306等)，ap几乎不存在。

- ▲ 电商系统中，如何实现秒杀功能？如何解决商品的超卖问题？*
- ▲ 假如明天是活动高峰？QPS 预计会翻10倍，你要怎么做？*
- ▲ 简述生产消费者模式的流程。 *
- ▲ 设计一个阻塞队列。
- ▲ 简述 MapReduce 的原理
- ▲ 停车场有有限个车位，有多个车来抢车位，设计一个系统需要根据车辆进入和离开停车场的时间进行计费

### # 非技术

- ▲ 对加班有什么看法？
- ▲ 你的优势和劣势是什么？
- ▲ 你的优势和劣势是什么？
- ▲ 为什么要离开现在的公司？
- ▲ 下一份工作希望学习到什么？
- ▲ 最近在看什么书以及技术文章？
- ▲ 团队合作沟通中遇到过什么问题？
- ▲ 团队合作沟通中遇到过什么问题？
- ▲ 简单描述一下自己是怎么样的人？
- ▲ 成长过程中影响你最深的事件和人
- ▲ 目前为止，坚持得最久一件事情是什么？
- ▲ 最近一年内遇到的最有挑战的事情是什么？
- ▲ 项目中最难的地方是哪里？你学习到了什么？
- ▲ 与同事沟通的时候，如果遇到冲突了如何解决？
- ▲ 与同事沟通的时候，如果遇到冲突了如何解决？
- ▲ 最近在看什么书吗，有没有接触过什么新技术？
- ▲ 最近阅读哪些技术书籍，遇到技术问题是怎么去解决？
- ▲ 你对xxx公司了解多少？你选择xxx公司的原因是什么？
