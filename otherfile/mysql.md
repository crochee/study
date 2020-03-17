# mysql相关

* 完成业务功能，懂基本的SQL语句
* 做性能优化，懂索引，懂引擎
* 做分库分表，懂主从，懂读写分离
* 做安全，懂权限，懂备份，懂日志
* 做云数据库，懂源码，懂瓶颈

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
3.  在不同值少的列上(字段上)不要建立索引，比如在学生表的"性别"字段上只有男，女两个不同值。相反的，在一个字段上不同值较多可以建立索引
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
## 事务
MySQL 事务主要用于处理操作量大，复杂度高的数据。比如说，在人员管理系统中，你删除一个人员，你既需要删除人员的基本资料，也要删除和该人员相关的信息，如信箱，文章等等，这样，这些数据库操作语句就构成一个事务！

**注意**：在 MySQL 中只有使用了 Innodb 数据库引擎的数据库或表才支持事务。
事务处理可以用来维护数据库的完整性，保证成批的 SQL 语句要么全部执行，要么全部不执行。
事务用来管理 insert,update,delete 语句

一般来说，事务是必须满足4个条件（ACID）：：原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）。

$color{red}{**原子性**}$：一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。

**一致性**：在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。

**隔离性**：数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。

**持久性**：事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。

在 MySQL 命令行的默认设置下，事务都是自动提交的，即执行 SQL 语句后就会马上执行 COMMIT 操作。因此要显式地开启一个事务务须使用命令 BEGIN 或 START TRANSACTION，或者执行命令 SET AUTOCOMMIT=0，用来禁止使用当前会话的自动提交。

**事务控制语句**：
*   BEGIN 或 START TRANSACTION 显式地开启一个事务；
*   COMMIT 也可以使用 COMMIT WORK，不过二者是等价的。COMMIT 会提交事务，并使已对数据库进行的所有修改成为永久性的；
*   ROLLBACK 也可以使用 ROLLBACK WORK，不过二者是等价的。回滚会结束用户的事务，并撤销正在进行的所有未提交的修改；
*   SAVEPOINT identifier，SAVEPOINT 允许在事务中创建一个保存点，一个事务中可以有多个 SAVEPOINT；
*   RELEASE SAVEPOINT identifier 删除一个事务的保存点，当没有指定的保存点时，执行该语句会抛出一个异常；
*   ROLLBACK TO identifier 把事务回滚到标记点；
*   SET TRANSACTION 用来设置事务的隔离级别。InnoDB 存储引擎提供事务的隔离级别有READ UNCOMMITTED、READ COMMITTED、REPEATABLE READ 和 SERIALIZABLE。

**MYSQL 事务处理主要有两种方法**：
1.  用 BEGIN, ROLLBACK, COMMIT来实现

BEGIN 开始一个事务
ROLLBACK 事务回滚
COMMIT 事务确认

2.  直接用 SET 来改变 MySQL 的自动提交模式:

SET AUTOCOMMIT=0 禁止自动提交
SET AUTOCOMMIT=1 开启自动提交

## 其他
1.  应尽量避免在 where 子句中使用 or 来连接条件，否则将导致引擎放弃使用索引而进行全表扫描

    select id from t where num=10 or num=20
    可以这样查询：
    select id from t where num=10 union all select id from t where num=20
2.  in 和 not in 也要慎用，否则会导致全表扫描，如：

    select id from t where num in(1,2,3)
    对于连续的数值，能用 between 就不要用 in 了：
    select id from t where num between 1 and 3
3.  很多时候用 exists 代替 in 是一个好的选择：
    
    select num from a where num in(select num from b)
    用下面的语句替换：
    select num from a where exists(select 1 from b where num=a.num)
    
    
    