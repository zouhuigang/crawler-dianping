/*
Navicat MySQL Data Transfer

Source Server         : qa
Source Server Version : 50713
Source Host           : 139.196.16.67:3306
Source Database       : anooc3

Target Server Type    : MYSQL
Target Server Version : 50713
File Encoding         : 65001

Date: 2017-08-13 10:51:16
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for crawl_crontab
-- ----------------------------
DROP TABLE IF EXISTS `crawl_crontab`;
CREATE TABLE `crawl_crontab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `c_name` varchar(50) NOT NULL COMMENT '爬虫名称',
  `start_time` int(6) NOT NULL DEFAULT '0' COMMENT '爬虫生效时间',
  `end_time` int(6) NOT NULL DEFAULT '0' COMMENT '爬虫生效时间',
  `seconds` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，秒',
  `minutes` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，分',
  `hours` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，小时',
  `day` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，天，1-31号',
  `month` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，月，1-12',
  `week` varchar(50) NOT NULL DEFAULT '*' COMMENT '定时任务，星期，0-6',
  `is_open` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1开启状态，0关闭状态',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0当前爬取行为未完成，不允许开始新的任务',
  `entryid` int(11) NOT NULL DEFAULT '0' COMMENT '定时任务运行的线程id，用与移除任务或暂停等',
  `is_showlog` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0不显示日志，1显示日志',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ctime` (`ctime`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='网站抓取规则表';

-- ----------------------------
-- Records of crawl_crontab
-- ----------------------------
INSERT INTO `crawl_crontab` VALUES ('1', '测试5秒钟刷新一次aaa', '0', '0', '*/5', '*', '*', '*', '*', '*', '1', '1', '2', '0', '2017-08-13 10:50:13');
INSERT INTO `crawl_crontab` VALUES ('3', '测试3秒钟刷新一次', '0', '0', '*/3', '*', '*', '*', '*', '*', '1', '1', '4', '1', '2017-08-13 10:50:14');
INSERT INTO `crawl_crontab` VALUES ('4', '测试1分钟刷新一次', '0', '0', '*', '*/1', '*', '*', '*', '*', '0', '1', '0', '0', '2017-08-11 15:11:13');
SET FOREIGN_KEY_CHECKS=1;
