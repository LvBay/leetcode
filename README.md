## 算法
1. 快速排序（包括算法步骤、平均算法复杂度、最好和最坏的情形），有人说校招要把算法写出来，我是社招，所以描述一下算法过程即可。
``` 从中间挑一个数n,使左边的数都小于n,右边的数都大于n```
2. 写二分查找算法，这个尽管是社招，但是一般也不难，所以要求面试者写出来。但是很多公司，比如不会直接让你写算法，而是结合一个具体场景来提问，然后让你自己联想到二分查找，比如求一个数的平方根。
3. 链表，常见的面试题有写一个链表中删除一个节点的算法、单链表倒转、两个链表找相交的部分，这个一般必须得完全无误的情况下写出来；
4. 自己实现一些基础的函数，例如strcpy / memcpy / memmov / atoi，同样的道理，这些必须完全无误且高效地写出来，比如你的实现中有动态分配堆内存，那么这道题目就算答错。第3点和第4点的关键点一般在于考察你的代码风格、对边界条件的处理，比如判断指针是否为空，千万不要故意不考虑这种情形，即使你知道也不行，只要你不写，一般面试官就认为你的思路不周详，容错率低；再比如，单链表的倒转，最后的返回值肯定是倒转后的链表头结点，这样才能引用一个链表，这些都是面试官想考虑的重点。
5. 哈希表，对哈希表的细节要求很高，比如哈希表的冲突检测、哈希函数常用实现、算法复杂度；比如百度二面就让我写一个哈希表插入元素算法，元素类型是任意类型。
6. AVL树和B树的概念、细节，比如会问mysql数据库的索引的实现原理，基本上就等于问你B树了。
7. 红黑树，这个基本上必问的一个数据结构，包括红黑树的概念、平均算法复杂度、最好最坏情况下的算法复杂度、、左右旋转、颜色变换。面试官常见的算法套路有：你熟悉C++的stl吗？你说熟悉，ok，stl的map用过吧？用过，ok，那map是如何实现的？红黑树，ok，那什么是红黑树？这样提问，红黑树就开始了。Java的也类似。

## golang
1. goroutine原理,轻量的原因
2. channel原理
3. 逃逸分析,内联
4. gc原理
5. 优化经验

##  网络通信
1. ### 协议栈的层级关系，三次握手和四次挥手的【细节】
    A与B建立了正常连接后，从未相互发过数据，这个时候B突然机器重启，问A此时的tcp状态处于什么状态？如何消除服务器程序中的这个状态？
2. 通信协议如何设计避免粘包；
3. http协议的get和post方法的区别（问的比较深的会让你画出http协议的格式，参照这篇文章中关于http协议格式的讲解：http://blog.csdn.net/analogous_love/article/details/72540130）

## 操作系统原理性的东西
1. 栈,堆概念
2. 堆和栈的区别，栈的结构，栈的细节.
3. 如函数参数入栈顺序、函数局部变量在栈中的布局、栈帧指针和栈顶指针位置
4. linux系统下可能还会问什么是daemon进程，如何产生daemo进程，什么是僵尸进程，僵尸进程如何产生和消除

## 开源技术
1. nosql技术的的redis 熟练使用redis甚至研究过redis源码，现在一般是对做后台开发的技术硬性要求或者说不可缺少的部分
2. 基于redis的面试题既可以聊算法与数据结构，也可以聊网络框架等等一类东西

## 聊自己的项目经验
1. beego优缺点
2. nsq印象较深的代码
3. etcd raft
4. 树形结构数据表结构设计