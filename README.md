# goblin-go

pcap实现的是网卡数据嗅探, 并非流量拦截, 因此只能实现单向转发.
考量到tcp/udp/transfer的feature实现, 需要引入内核的拦截模块, 提供api对数据报进行过滤、地址转换、具体业务处理等.

在Linux2.4.x之后新一代的Linux防火墙机制采用的是iptables#netfilter

mainFile: net/core/netfilter.c

mainHeadFile: include/linux/netfilter.h

ipv4: net/ipv4/netfilter/*.c

ipv4Head: 
  - include/linux/netfilter_ipv4.h
  - include/linux/netfilter_ipv4/*.h

api:
  - 连接跟踪模块（Conntrack）
  - 网络地址转换模块（NAT）
  - 据报修改模块（mangle）
  - etc.
