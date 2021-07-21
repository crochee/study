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


参考：https://segmentfault.com/a/1190000011338151
## 通用参考
*   不同的云服务API的通用语义的参数定义规范参考
*   原则上查询类、删除类的API的请求参数，需要进行检查，参数非法则报错。备注：涉及开源社区的，可以以社区为准

|公共参数 |type|描述 |
|:---:|:---:|:---:|
|name | string|名称字段应包含相对应资源名称|
|create_time | string|实体的创建时间戳|
|update_time | string|实体的最后更新时间戳|
|delete_time | string|实体的删除时间戳|
|expire_time | string|实体的过期时间戳|
|start_time | string|时间戳记某个时间段的开始|
|end_time | string|标记某个时间段或某个操作结束的时间戳|
|read_time | string|应该读取或读取特定实体的时间戳|
|time_zone | string|时区名称。它应该是IANA TZ名称，如America/Los_Angeles|
|region_code | string|某个位置的Unicode国家/地区代码|
|order_by | string|指定列表请求的结果顺序|
|limit | integer|查询返回记录的数量限制|
|offset | integer|偏移量,表示查询该偏移量后面的记录|
|sort_key | string|返回结果按该关键字排序（支持id，status，size，created_at等关键字，默认为“created_at”）|
|sort_dir | string|降序或升序（分别对应desc和asc，默认为“desc”）|


### 请求头
|name | 描述|是否必选|
|:---:|:---:|:---:|
|Content-Type | 默认值：application/json;charset=UTF-8|可选，有Body体的情况下必选，没有Body体则无需填写和校验|
|X-sdk-date | 请求的发生时间，格式为(YYYYMMDD'T'HHMMSS'Z')。取值为当前系统的GMT时间|如果使用ak/sk做接口认证的时候，那么此字段必须设置；如果使用PKI token的时候，不必设置|
|Authorization | 该值来源于请求签名结果。类型：字符串 默认值：无|否，使用AK/SK进认证时该字段必选|
|Host | 请求的服务器信息，从服务API的URL中获取。值为hostname[:port]。端口缺省时使用默认的端口，https的默认端口为443|否，使用AK/SK认证时该字段必选。|
|X-Auth-Token | 如果使用Token认证的方式，此字段携带认证密钥 类型：字符串。默认值：无。|否 使用Token认证时该字段必选|

### 状态码定义
参考 https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
### 响应头部
|name | 描述|是否必选|
|:---:|:---:|:---:|
|Content-Type | 用于指明发送给接收者的实体正文的媒体类型类型：字符串。 默认值：application/json; charset=UTF-8|可选，有Body体的情况下必选，没有Body体则无需填写和校验|
|X-Request-Id | 此字段携带请求ID号，以便任务跟踪.类型：字符串 request_id-timestamp-hostname(request_id在服务器端生成UUID， timestamp为当前时间戳，hostname为处理当前接口的服务器名称)默认值：无。|建议服务加上，以便后续进行任务跟踪|
### API设计要求
#### 请求行必须按照统一资源标识符定义的格式进行设计
请求行主要是对于URI（统一资源标识符）的要求。URI是用户看到接口最直观的内容之一，针对URI的定义，必须要按照规定的格式进行设计（OpenStack或者社区接口以原生定义为准）。统一资源标识符定义格式为：           https://endpoint/{version}/[project_id/tenant_id]/{resources}

备注：这里的version表示接口的版本号，如v1，本文中project_id和tenant_id含义一样都表示项目ID。
####   URI只允许出现名词，并且使用小写，尽量少用连接单词
(1)  URI和参数经过URI编码之后的总长度不超过2048；

(2)  URI中不要出现无谓的缩写单词(除非缩写单词业界通用)，名称要保证“望文知义”，以免造成理解障碍，不能使用不常见的缩略语，单词使用要正确，跨API的同一含义单词应一致，便于开发者理解API；

(3)  对于URI中资源的命名方式要求用中划线连词符“-”分隔的脊柱命名法，且所有字母小写；

(4)  URI中的参数或body请求体部分要求以下划线“_”分隔的蛇形命名法，所有字母均小写

举例：POST /v2/{project_id}/os-vender-replications
```json
{
  "volume": {
    "backup_id": null,
    "availability_zone": "az1.dc1"
  }
}
```
####  Version明确标识出服务API的版本号
1.    选项：非必选

2.    说明：每个API都必须要带上版本号，由于历史原因当前API的版本号存vX.Y和vX的两种状况，为规范API的版本号定义，后续新服务的API 版本定义建议统一规范成vX，这里X是一个整数(OpenStack 原生接口保持与原生社区一致保持兼容)，如网络服务新增一套VPC接口统一定义为v3，如果后期再推出一套全新接口就在此基础上增加X的值。

3.    对外提供API的服务需要提供查询获取本服务API版本号的接口（需要提供get /接口），要明确能在接口中说明清楚哪个版本号是目前服务主推的版本，哪些版本是支持但已经不推荐的版本，方便API使用者通过该接口快速了解与跟进服务API的变化。

4.    API 版本号信息属于服务对外的global 信息，不涉及资源的任何操作，不需要对接口进行token认证。如果get /或者get/{version}继承自开源社区，认证策略与社区保持一致；其中get/{version}不做强制要求。

整体设计要求如下：

响应参数要求：

|参数名称 | 类 型|是否必选|说 明|
|:---:|:---:|:---:|:---:|
|versions | Array|是|描述version 相关对象的列表|
|id | String|是|版本ID（版本号），如v1|
|links | String|是|API的URL地址|
|version | String|是|若该版本API支持微版本，则填支持的最大微版本号，如果不支持微版本，则填空|
|status | String|是|版本状态，为如下3种： CURRENT：表示该版本为主推版本 SUPPORTED：表示为老版本，但是现在还继续支持 DEPRECATED：表示为废弃版本，存在后续删除的可能|
|updated | String|是|版本发布时间，要求用UTC时间表示。如v1发布的时间2014-06-28T12:20:21Z|
|min_version | String|是|若该版本API 支持微版本，则填支持的最小微版本号， 如果不支持微版本，则填空|

举例：

Get /

响应样例：
```json
{
  "versions": [
    {
      "id": "v1",
      "links": [
        {
          "href": "https://xxxx.otc.t-systems.com/v1/",
          "rel": "self"
        }
      ],
      "min_version": "",
      "status": "SUPPORTED",
      "updated": "2016-12-09T00:00:00Z",
      "version": ""
    },
    {
      "id": "v2",
      "links": [
        {
          "href": "https://xxxx.otc.t-systems.com/v2/",
          "rel": "self"
        }
      ],
      "min_version": "2.0",
      "status": "CURRENT",
      "updated": "2017-12-09T00:00:00Z",
      "version": "2.26"
    }
  ]
}
```
Get /v2

响应样例：
```json
{
  "version": {
    "id": "v2",
    "links": [
      {
        "href": "https://xxxx.otc.t-systems.com/v2/",
        "rel": "self"
      }
    ],
    "min_version": "2.0",
    "status": "CURRENT",
    "updated": "2017-12-09T00:00:00Z",
    "version": "2.26"
  }
}
```          
#### 对于创建类以及批量处理的接口参数，Resources必须使用复数
1.  选项：必选
2.  说明：在请求消息体中增加创建参数；对于修改、删除、查询单一指定资源，在resources后面增加具体resouce_id.
3.  举例：  
    创建虚拟机：POST /v2/{tenant_id}/servers  
    查询指定虚拟机：GET /v2/{tenant_id}/servers/{server_id}  
    删除指定虚拟机：DELETE /v2/{tenant_id}/servers/{server_id}
####  HTTP动词使用标准化，便于开发者理解和正确使用API
1.  选项：必选
2.  Restful API的单个资源操作,资源的标准CRUD操作对应的HTTP动词如下，符合如下规则：

GET：

【规则】GET操作用于获取资源的场景。

【规则】GET操作必须具备安全性和幂等性。

说明：安全性指经过操作后不改变服务器状态。幂等性指不允许对资源状态做相对的改动。

【规则】GET操作成功返回状态码200（OK）。

【建议】 （1）若客户端存有资源缓存，在发送GET请求时可在Header中添加If-Modified-Since等条件信息，服务端经过判断后若发现资源未变更，则返回304（Not Modified），通知客户端缓存的咨询信息还可继续使用。

（2）为统一规范返回体格式，使用get方法，查询单个资源信息，返回body里面不要用列表格式[]，查询资源列表，返回body里面需要用列表格式[]。

POST：

【规则】适用于新建资源场景，以及CRUD无法表达的操作场景（Non-CRUD）。

说明：对于CRUD无法表达的操作的使用场景，参考“Non-CRUD资源”章节。

【规则】POST创建资源成功，返回状态码201（Created）。Non-CRUD操作返回状态码200（OK）。

PUT：

【规则】若操作的URL为一个新资源，则创建该资源。若URL为一个已存在的资源，则替换该资源。

说明：例如“PUT /users/admin”表达的含义为若id为admin的用户不存在，则创建一个id为admin的用户，并为该用户设置属性信息。如果存在，则替换该用户的所有信息。

【规则】PUT操作必须具备幂等性。

【规则】若使用PUT创建资源成功，返回状态码201（Created）。若修改资源成功，返回状态码200（OK）。

【规则】PUT操作传入的消息体需包含被替换资源的完整信息。若传递的信息不完整，在服务实现端需提供对应信息的默认值。

PATCH：

【规则】PATCH操作用于部分更新资源的场景。

【规则】部分更新操作成功，返回状态200（OK）。

【建议】若使用PUT操作所需输入的整体资源信息内容大小与PATCH操作无太大差异，优先使用PUT操作，不推荐使用PATCH操作。

说明：RFC 5789（PATCH Method for HTTP）文档明确说明：“ PATCH is neither safe nor idempotent as defined by [RFC2616], Section9.1.” ，即 PATCH 方法非幂等。此外，由于业界对PATCH方法的使用风格不统一，以及多数标准开发框架对PATCH的支持不完善，所以在能够使用PUT请求定义操作接口的情况下，不建议业务使用PATCH操作。若某些业务必须使用，需定义详细的请求消息体格式和使用场景。

DELETE ：

【规则】DELETE操作用于删除资源的场景。

【规则】DELETE操作必须具备幂等性。

【规则】若删除资源资源成功，返回状态码200（OK）或204 (No Content)。若服务接收请求，但操作未立即执行，返回状态码202（Accepted）。
####  批量资源操作设计
1.  选项：可选
2.  批量创建举例

POST      http://{endpoint}/v3/{project_id}/{resource}/batch-create 

Body体中包含所有需要创建的对象
```json
{

            "resources":[

               {.......

              },

              {......

               }

           ]

}
```
3.批量删除举例

POST      http://{endpoint}/v3/{project_id}/{resource}/batch-delete 

Body体中包含所有需要删除的对象
```json
{
  "resources": [
    {
      "id": "id1"
    },
    {
      "id": "id2"
    },
    {
      "id": "id3"
    },
    {
      "id": "id4"
    }
  ]
}
```
4.批量变更举例

POST      http://{endpoint}/v3/{project_id}/{resource}/batch-update

Body体中包含所有需要更新的对象
```json
{
  "resources": [
    {
      "id": "id1"
    },
    {
      "id": "id2"
    },
    {
      "id": "id3"
    },
    {
      "id": "id4"
    }
  ]
}
```
5.返回值  
返回分为几种场景:全部成功或者全部失败，部分成功部分失败，异常，从接口行为又分同步操作和异步操作接口  

|操作方式 | 返回情况|返回样式|场景说明|
|:---:|:---:|:---:|:---:|
|同步操作 | 全部成功 全部失败|{"resources":[{"id":"xxx","xxx":"xxx"},{ "id":"xxx","xxx":"yyyy"}]}|若API仅支持全部成功或全部失败，返回值无需标记单个资源状态，全部返回成功信息即可，否则全部失败，报错。|
|同步操作 | 部分成功部分失败|{"resources": [{"id": "xxx", "ret_status": "successful"}, {"id": "xxx", "ret_status": "error"}]}|若需要支持部分成功部分失败场景返回的资源清单中包含每个资源的返回状态，用于标识每个资源的创建状态，例如返回样式中的ret_status。|
|同步操作 | 失败报错|标准的错误返回 {"error_msg": "The format of message is error", "error_code": "服务名.00000001"}|配额不足，全量失败，权限不足，依赖资源冲突，参数错误，内部服务故障等场景。|
|异步操作 | 全部成功 部分成功部分失败|{"resources": [{"id": "xxx"}, {"id": "xxx"}]}|异步操作下，不管是支持全部成功全部失败的api还是支持部分成功部分失败的api均返回的是资源id清单，最终成功失败状态需要使用查询接口进行轮询检查。|
|异步操作 | 失败报错|标准的错误返回 {"error_msg": "The format of message is error", "error_code": "服务名.00000001"}|异步场景失败报错特指立即能返回的错误，配额不足，权限不足，参数错误，内部服务故障等，此类错误无需等待任务运行即可检查出来。|
#### 统一的错误码格式和错误响应体格式规范
请求失败时，必须在响应消息体中返回应用级的错误码（error_code）和错误描述(error_msg)。  
请求成功时，不在响应消息体中包含错误码（error_code）和错误描述（error_msg）。  
错误码要以服务简写开头+”.”+数字编码方式的错误码  
![Image text](错误码定义.jpg)
#### 时间格式 
1.  选项：必选
2.  时间格式：遵循RFC3339规范，非特殊场景默认使用UTC时间。
3.  示例：
精确到秒：2020-09-01T18:50:20Z  
精确到毫秒：2020-09-01T18:50:20.200Z  
带时区：：2020-09-01T18:50:20+08:00
####  幂等性：API设计要遵循幂等性的几个场景
从定义上看，HTTP方法的幂等性是指一次和多次请求某一个资源应该具有同样的副作用。
(1) 查询类接口(Get)不涉及；

(2) 资源修改类接口(PUT、Patch)，建议要支持幂等；

(3) 创建类接口(POST)，要实现幂等比较困难，不作要求；

(4) 删除接口(Delete),建议支持幂等