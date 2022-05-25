CREATE DATABASE IF NOT EXISTS zsxq default charset utf8 COLLATE utf8_general_ci;
use zsxq;

DROP TABLE IF EXISTS `knowledge`;
CREATE TABLE `knowledge` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `message` varchar(100) DEFAULT NULL COMMENT '内容',
  `name` varchar(50) DEFAULT NULL COMMENT '作者',
  `create_time` varchar(30) DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
