# 原理
1. 首先根据HTTP协议头部的规则，应该持续从socket读取数据，直到读到了\r\n\r\n，表示头部报文结束。

2. 如果报文中使用Content-Length指定传输实体的大小，接下来不论HTTP客户/服务器都因该根据读取到Content-Length指定的实体大小

3. 对于分块传输的实体，传输编码为chunked。即Transfer-Encoding: chunked。 分快传输的编码，一般只适用于HTTP内容响应(HTTP请求也可以指定传输编码为chunked，但不是所有HTTP服务器都支持)。 这时可以读取定量的数据(如4096字节) ，交给parser解析。然后重复此过程，直到chunk编码结束。
