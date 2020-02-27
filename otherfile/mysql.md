# mysql相关

```sql
CREATE TABLE `order_commander` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `total` decimal(12,2) NOT NULL COMMENT '订单金额',
  `ccy` varchar(4) NOT NULL DEFAULT '' COMMENT '订单币种',
  `expire` varchar(16) NOT NULL DEFAULT '' COMMENT '支付截止时间',
  `order_stat` tinyint(2) NOT NULL COMMENT '订单状态',
  `operator` varchar(32) NOT NULL DEFAULT '' COMMENT '下单人',
  `ptid` varchar(5) NOT NULL DEFAULT '' COMMENT '渠道id',
  `reason` text NOT NULL COMMENT '失败原因',
  `disable` tinyint(1) NOT NULL DEFAULT '0',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```
```sql
alter table `trip_order_session` add column `ptid` varchar(5) NOT NULL DEFAULT '' COMMENT '渠道id'
```
# 相关知识
## 索引 
https://www.jianshu.com/p/0d6c828d3c70
### 索引优缺点和使用原则
**优点**
1.  所有的MySql列类型(字段类型)都可以被索引，也就是可以给任意字段设置索引
2.  大大加快数据的查询速度

**缺点**
1.  创建索引和维护索引要耗费时间，并且随着数据量的增加所耗费的时间也会增加
2.  索引也需要占空间，我们知道数据表中的数据也会有最大上线设置的，如果我们有大量的索引，索引文件可能会比数据文件更快达到上线值
3.  当对表中的数据进行增加、删除、修改时，索引也需要动态的维护，降低了数据的维护速度

**使用原则**
1.  对经常更新的表就避免对其进行过多的索引，对经常用于查询的字段应该创建索引
2.  数据量小的表最好不要使用索引，因为由于数据较少，可能查询全部数据花费的时间比遍历索引的时间还要短，索引就可能不会产生优化效果
3.  在一同值少的列上(字段上)不要建立索引，比如在学生表的"性别"字段上只有男，女两个不同值。相反的，在一个字段上不同值较多可以建立索引
### 索引分类
索引是在存储引擎中实现的，也就是说不同的存储引擎，会使用不同的索引

MyISAM和InnoDB存储引擎：只支持BTREE索引， 也就是说默认使用BTREE，不能够更换；MEMORY/HEAP存储引擎：支持HASH和BTREE索引

**单列索引**
1.  **普通索引**->没有什么限制，允许在定义索引的列中插入重复值和空值，纯粹为了查询数据更快一点
   创建表的时候带索引，key或index都可以
    ```sql
    CREATE TABLE `t1` (
      `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
      `total` decimal(12,2) NOT NULL COMMENT '订单金额',
      key(order_id)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
    ```
    插入一条数据
    ```sql
    INSERT INTO `t1`(`order_id`,`total`)VALUE ('test',12)
    ```
    查看是否用了索引
    ```sql
    EXPLAIN SELECT * FROM t1 WHERE order_id = 'test'
    ```
    结果一一对应
    "id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
    1,   "SIMPLE",      "t1",    NULL,       "ref","order_id",      "order_id",98,"const",1,     100.00,    NULL

2.  **唯一索引**->索引列中的值必须是唯一的，但是允许为空值
    创建
    ```sql
    CREATE TABLE `t2` (
      `id` int(11) NOT NULL,
      `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
      `total` decimal(12,2) NOT NULL COMMENT '订单金额',
      UNIQUE KEY `UniqIdx` (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    ```
    要查看其中查询时使用的索引，必须先往表中插入数据，然后在查询数据，不然查找一个没有的id值，是不会使用索引的
    当表中没有数据时，查询
    ```sql
    EXPLAIN SELECT * FROM t2 WHERE id = 1
    ```
    结果 "id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
       1,"SIMPLE",NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,"no matching row in const table"
    
    现在插入一条数据
    ```sql
    INSERT INTO t2 VALUES(1,'xxx',23)
    ```
    再用上面的查询
    "id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
    1,"SIMPLE","t2",NULL,"const","UniqIdx","UniqIdx",4,"const",1,100.00,NULL
3.  **主键索引**->是一种特殊的唯一索引，不允许有空值
    创建
    ```sql
    CREATE TABLE `t1` (
      `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
      `total` decimal(12,2) NOT NULL COMMENT '订单金额',
      PRIMARY KEY (`order_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    ```
    无数据直接
    ```sql
    EXPLAIN SELECT * FROM t1 WHERE order_id = 'test'
    ```
    结果
    "id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
    1,"SIMPLE","t1",NULL,"ALL","PRIMARY",NULL,NULL,NULL,1,100.00,"Using where"
    现在插入一条数据
    ```sql
    INSERT INTO `t1`(`order_id`,`total`)VALUE ('test',12)
    ```
    再查询，结果
    "id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
    1,"SIMPLE","t1",NULL,"const","PRIMARY","PRIMARY",98,"const",1,100.00,NULL

    
**组合索引**

在表中的多个字段组合上创建的索引，只有在查询条件中使用了这些字段的左边字段时，索引才会被使用，使用组合索引时遵循最左前缀集合

创建：
```sql
CREATE TABLE `t2` (
  `id` int(11) NOT NULL,
  `name` char(30) NOT NULL,
  `age` int(11) NOT NULL,
  `info` varchar(255) DEFAULT NULL,
  KEY `mutildx` (`id`,`name`,`age`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
组合索引就是遵从了**最左前缀**，利用索引中最左边的列集来匹配行，这样的列集称为最左前缀，不明白没关系，举几个例子就明白了，例如，这里由id、name和age3个字段构成的索引，索引行中就按id/name/age的顺序存放，索引可以索引下面字段组合(id，name，age)、(id，name)或者(id)。如果要查询的字段不构成索引最左面的前缀，那么就不会是用索引，比如，age或者（name，age）组合就不会使用索引查询

现在表里有数据，查询
```sql
EXPLAIN SELECT * FROM t2 WHERE id = 1 AND name = 'xxx'
```
结果
"id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
1,"SIMPLE","t2",NULL,"ref","mutildx","mutildx",94,"const,const",1,100.00,NULL

没有最左前缀的情况
```sql
EXPLAIN SELECT * FROM t2 WHERE age = 20 AND name = 'xxx'
```
结果
"id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
1,"SIMPLE","t2",NULL,"ALL",NULL,NULL,NULL,NULL,1,100.00,"Using where"
    
**全文索引**

全文索引，只有在MyISAM引擎上才能使用，只能在CHAR,VARCHAR,TEXT类型字段上使用全文索引，介绍了要求，说说什么是全文索引，就是在一堆文字中，通过其中的某个关键字等，就能找到该字段所属的记录行，比如有"你是个靓仔，靓女 ..." 通过靓仔，可能就可以找到该条记录

创建
```sql
CREATE TABLE `t1` (
  `id` int(11) NOT NULL,
  `name` char(30) NOT NULL,
  `age` int(11) NOT NULL,
  `info` varchar(255) DEFAULT NULL,
  FULLTEXT KEY `fulltextldx` (`info`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
```
插入数据
```sql
INSERT INTO t1 VALUES(8,'xxx',20,'text is so good，hei，my name is xxx'),(9,'BBB',4,'my name is gorlr')
```
查询
```sql
EXPLAIN SELECT * FROM t1 WHERE MATCH(info) AGAINST('gorlr')
```
结果
"id","select_type","table","partitions","type","possible_keys","key","key_len","ref","rows","filtered","Extra"
1,"SIMPLE","t1",NULL,"fulltext","fulltextldx","fulltextldx",0,"const",1,100.00,"Using where"

注意：在使用全文搜索时，需要借助MATCH函数，并且其全文搜索的限制比较多，比如只能通过MyISAM引擎，比如只能在CHAR,VARCHAR,TEXT上设置全文索引。比如搜索的关键字默认至少要4个字符，比如搜索的关键字太短就会被忽略掉。等等

**空间索引**

空间索引是对空间数据类型的字段建立的索引，MySQL中的空间数据类型有四种，GEOMETRY、POINT、LINESTRING、POLYGON。在创建空间索引时，使用SPATIAL关键字。要求，引擎为MyISAM，创建空间索引的列，必须将其声明为NOT NULL
```sql
CREATE TABLE `t1` (
  `g` geometry NOT NULL,
  SPATIAL KEY `spatldx` (`g`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
```

**关于索引的一些操作**

*   在已经存在的表上创建索引

ALTER TABLE 表名 ADD[UNIQUE|FULLTEXT|SPATIAL][INDEX|KEY] [索引名] (索引字段名)[ASC|DESC]
*   为表添加索引
```sql
ALTER TABLE book ADD INDEX BkNameIdx(bookname(30));
```
*   使用CREATE INDEX创建索引
```sql
CREATE INDEX BkBookNameIdx ON book(bookname)
```
* 查询索引

SHOW INDEX FROM表名\G；　　\G只是让输出的格式更好看
* 删除索引

ALTER TABLE 表名 DROP INDEX 索引名
DROP INDEX 索引名 ON 表名