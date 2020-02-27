# rabbitMQ
rabbitMQ详情见：https://www.sojson.com/blog/48.html
###RabbitMQ 选型和对比
1.  **从社区活跃度**
按照目前网络上的资料，RabbitMQ 、activeM 、ZeroMQ 三者中，综合来看，RabbitMQ 是首选。
2.  **持久化消息比较**
ZeroMq 不支持，ActiveMq 和RabbitMq 都支持。持久化消息主要是指我们机器在不可抗力因素等情况下挂掉了，消息不会丢失的机制。
3.  **综合技术实现**
可靠性、灵活的路由、集群、事务、高可用的队列、消息排序、问题追踪、可视化管理工具、插件系统等等。
RabbitMq / Kafka 最好，ActiveMq 次之，ZeroMq 最差。当然ZeroMq 也可以做到，不过自己必须手动写代码实现，代码量不小。尤其是可靠性中的：持久性、投递确认、发布者证实和高可用性。
4.  **高并发**
毋庸置疑，RabbitMQ 最高，原因是它的实现语言是天生具备高并发高可用的erlang 语言。
5.  **比较关注的比较** RabbitMQ 和 Kafka
RabbitMq 比Kafka 成熟，在可用性上，稳定性上，可靠性上，  RabbitMq  胜于  Kafka  （理论上）。
6.  Kafka 的定位主要在日志等方面， 因为Kafka 设计的初衷就是处理日志的，可以看做是一个日志（消息）系统一个重要组件，针对性很强，所以 如果业务方面还是建议选择 RabbitMq 。
7.  Kafka 的性能（吞吐量、TPS ）比RabbitMq 要高出来很多。

