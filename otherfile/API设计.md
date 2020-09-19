## RESTFUL API 设计总结
### 协议&域名&版本
* 协议：API与用户的通信协议，总是使用HTTPS
* 域名: 可以部署专有域名或者是主域名下加入URL里面
```text
https://api.github.com
https://github.com/api
```
* 版本: 应该将API的版本号放入URL
```text
https://github.com/v1
```
