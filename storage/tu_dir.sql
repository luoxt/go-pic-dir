/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50562
Source Host           : localhost:3306
Source Database       : tu_dir

Target Server Type    : MYSQL
Target Server Version : 50562
File Encoding         : 65001

Date: 2019-01-20 20:42:49
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tu_path`
-- ----------------------------
DROP TABLE IF EXISTS `tu_path`;
CREATE TABLE `tu_path` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tu_name` char(100) DEFAULT NULL,
  `tu_path` varchar(250) DEFAULT NULL,
  `md5` varchar(200) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `md5` (`md5`(191)) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
