# http个版本的区别
详情：https://blog.csdn.net/qq_39207948/article/details/80969968

# 状态码
## 1 请求处理中
*   100 Continue 继续
*   101 Switching Protocols 交换协议
*   102 Processing 处理中
## 2 请求成功
*   200 OK 成功
*   201 Created 创建
*   202 Accepted 接受
*   203 Non-Authoritative Information 非权威信息
*   204 No Content 无内容
*   205 Reset Content 重置内容
*   206 Partial Content 部分内容
*   207 Multi-Status 多状态
## 3 重定向
*   300 Multiple Choices 多项选择
*   301 Moved Permanently 永久转移
*   302 Move Temporarily 临时转移 
*   303 See Other 见其他
*   304 Not Modified 未修改
*   305 Use Proxy 使用代理
*   306 Switch Proxy 关闭代理
*   307 Temporary Redirect 临时重定向
## 4 客户端错误
*   400 Bad Request 错误请求
*   401 Unauthorized 未经授权
*   402 Payment Required 需要付款
*   403 Forbidden 禁止
*   404 Not Found 找不到
*   405 Method Not Allowed 不允许的方法
*   406 Not Acceptable 不可接受
*   407 Proxy Authentication Required 需要代理身份验证
*   408 Request Timeout 请求超时
*   409 Conflict 冲突
*   410 Gone 不见了
*   411 Length Required 需要长度
*   412 Precondition Failed 前置条件失败
*   413 Request Entity Too Large 请求实体太大
*   414 Request-URI Too Long 请求URL太长
*   415 Unsupported Media Type 不支持的媒体类型
*   416 Requested Range Not Satisfiable 请求的范围不满足
*   417 Expectation Failed 期望失败
*   418 I'm a teapot
*   421 Too Many Connections 连接过多
*   422 Unprocessable Entity 不可处理的实体
*   423 Locked 锁定
*   424 Failed Dependency 失败的依赖项
*   425 Too Early 过早
*   426 Upgrade Required 需要升级
*   449 Retry With 重试
*   451 Unavailable For Legal Reasons 因法律原因不可用
*   499 客户端已关闭
## 5 服务器错误
*   500 Internal Server Error 内部服务器错误
*   501 Not Implemented 未实现
*   502 Bad Gateway 网关不可用
*   503 Service Unavailable 服务不可用
*   504 Gateway Timeout 网关超时
*   505 HTTP Version Not Supported 不支持HTTP版本
*   506 Variant Also Negotiates
*   507 Insufficient Storage 存储不足
*   509 Bandwidth Limit Exceeded 超过带宽限制
*   510 Not Extended 未拓展
*   600 Unparseable Response Headers 不可解析的响应头
## tcp3次握手和4次挥手
参考：https://www.jianshu.com/p/29868fb82890

# 注意
因为XMLHttpRequest规范的限制，浏览器中ajax发送的http请求，get，delete请求不能携带实体