# RPC Learning

[TOC]

## RPC

RPC 代指远程过程调用（Remote Procedure Call），它的调用包含了传输协议和编码（对象序列号）协议等等。允许运行于一台计算机的程序调用另一台计算机的子程序，而**开发人员无需额外地为这个交互作用编程**。RPC 是在出于考虑高并发高可用高性能的分布式服务的基础上促成的。原本各个服务都在同一个地址空间，现在将服务拆分开，将总的服务拆解为一个一个粒度更高的微服务，例如订单服务系统，管理服务系统，用户服务系统等，在这样的情况下，我们选择**使用 RPC 来进行服务与服务之间的调用**。

RPC要解决的两个问题：

1. 解决分布式系统中，服务之间的调用问题。
2. 远程调用时，要能够像本地调用一样方便，让调用者感知不到远程调用的逻辑。

> [如何解释什么是RPC - 简书 (jianshu.com)](https://www.jianshu.com/p/2accc2840a1b)

### Protobuf

在这个 RPC 调用服务的过程中，是使用 socket 进行服务器和服务器之间的通信的，而这样的通信避不开的是对象的序列化和反序列化，我选择 Protobuf 来进行这个操作。Protobuf 支持 RPC，相比 JSON 占用内存更小，速度更快，并且对代码没有侵入，可以在不变更代码的情况下修改数据结构。

#### IDL

IDL 是Interface description language的缩写，指接口描述语言，是[CORBA](https://baike.baidu.com/item/CORBA/2776997)规范的一部分，是跨平台开发的基础。Protobuf 便是属于这个范围。

## CAP & BASE理论

- CAP 理论
  CAP原则又称CAP定理，指的是在一个分布式系统中， Consistency（一致性）、 Availability（可用性）、Partition tolerance（分区容错性），三者不可得兼。

- BASE
  BASE是Basically Available (基本可用) 、Soft state (软状态) 和Eventually consistent (最终一致性) 三个短语的简写，BASE是对CAP中一致性和可用性权衡的结果，其来源于对大规模互联网系统分布式实践的结论，是基于CAP定理逐步演化而来的，其核心思想是即使无法做到强一致性 (Strong consistency)，但每个应用都可以根据自身的业务特点，采用适当的方式来使系统达到最终一致性 (Eventual consistency)。

二者是描述分布式系统的基石。