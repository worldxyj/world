/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50723
Source Host           : localhost:3306
Source Database       : world

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2020-08-19 14:54:46
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for w_admin
-- ----------------------------
DROP TABLE IF EXISTS `w_admin`;
CREATE TABLE `w_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL DEFAULT '0',
  `name` varchar(20) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `tel` char(11) NOT NULL DEFAULT '',
  `ip` varchar(20) NOT NULL DEFAULT '',
  `login_at` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_password` (`name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_admin
-- ----------------------------
INSERT INTO `w_admin` VALUES ('1', '10', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '', '127.0.0.1', '1597820008', '1562057575');
INSERT INTO `w_admin` VALUES ('13', '11', '测试', 'e10adc3949ba59abbe56e057f20f883e', '', '127.0.0.1', '1597816947', '1596621542');

-- ----------------------------
-- Table structure for w_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `w_admin_role`;
CREATE TABLE `w_admin_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int(10) unsigned NOT NULL DEFAULT '0',
  `role_id` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_admin_role
-- ----------------------------

-- ----------------------------
-- Table structure for w_menu
-- ----------------------------
DROP TABLE IF EXISTS `w_menu`;
CREATE TABLE `w_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(16) NOT NULL DEFAULT '',
  `css` varchar(16) NOT NULL DEFAULT '',
  `url` varchar(32) NOT NULL DEFAULT '',
  `pid` int(10) unsigned NOT NULL DEFAULT '0',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '100',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1显示',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `url` (`url`) USING BTREE,
  KEY `sort` (`sort`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_menu
-- ----------------------------
INSERT INTO `w_menu` VALUES ('1', '系统设置', 'fa-cogs', 'sys', '0', '101', '1', '1562304671', '1597211968');
INSERT INTO `w_menu` VALUES ('2', '菜单管理', '', 'sys/menu', '1', '100', '1', '1562315306', '1562315306');
INSERT INTO `w_menu` VALUES ('3', '管理员列表', '', 'sys/admin', '1', '100', '1', '1562565735', '1562565735');
INSERT INTO `w_menu` VALUES ('4', '角色管理', '', 'sys/role', '1', '100', '1', '1562567947', '1597298764');
INSERT INTO `w_menu` VALUES ('5', '添加', '', 'sys/admin/add', '3', '100', '1', '1597289625', '1597289625');
INSERT INTO `w_menu` VALUES ('55', '添加', '', 'sys/role/add', '4', '100', '1', '1597305582', '1597305582');

-- ----------------------------
-- Table structure for w_permission
-- ----------------------------
DROP TABLE IF EXISTS `w_permission`;
CREATE TABLE `w_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `http_method` char(10) NOT NULL DEFAULT '',
  `http_route` varchar(32) NOT NULL,
  `create_at` int(10) unsigned NOT NULL DEFAULT '0',
  `update_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_permission
-- ----------------------------

-- ----------------------------
-- Table structure for w_role
-- ----------------------------
DROP TABLE IF EXISTS `w_role`;
CREATE TABLE `w_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_role
-- ----------------------------
INSERT INTO `w_role` VALUES ('10', '超级管理员', '1597311836');
INSERT INTO `w_role` VALUES ('11', '网站管理员', '1597717035');

-- ----------------------------
-- Table structure for w_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `w_role_menu`;
CREATE TABLE `w_role_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL DEFAULT '0',
  `menu_id` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=175 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_role_menu
-- ----------------------------
INSERT INTO `w_role_menu` VALUES ('162', '10', '56');
INSERT INTO `w_role_menu` VALUES ('163', '10', '57');
INSERT INTO `w_role_menu` VALUES ('164', '10', '1');
INSERT INTO `w_role_menu` VALUES ('165', '10', '2');
INSERT INTO `w_role_menu` VALUES ('166', '10', '3');
INSERT INTO `w_role_menu` VALUES ('167', '10', '5');
INSERT INTO `w_role_menu` VALUES ('168', '10', '4');
INSERT INTO `w_role_menu` VALUES ('169', '10', '55');
INSERT INTO `w_role_menu` VALUES ('170', '11', '56');
INSERT INTO `w_role_menu` VALUES ('171', '11', '57');
INSERT INTO `w_role_menu` VALUES ('172', '11', '1');
INSERT INTO `w_role_menu` VALUES ('173', '11', '3');
INSERT INTO `w_role_menu` VALUES ('174', '11', '5');

-- ----------------------------
-- Table structure for w_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `w_role_permission`;
CREATE TABLE `w_role_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL DEFAULT '0',
  `permission_id` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_role_permission
-- ----------------------------
