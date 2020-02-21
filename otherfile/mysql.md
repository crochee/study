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