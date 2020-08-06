/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50723
Source Host           : localhost:3306
Source Database       : world

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2020-08-06 17:21:53
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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_admin
-- ----------------------------
INSERT INTO `w_admin` VALUES ('1', '0', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '13839958207', '127.0.0.1', '1593324826', '1562057575');
INSERT INTO `w_admin` VALUES ('13', '0', '测试', 'e10adc3949ba59abbe56e057f20f883e', '11111111111', '', '0', '1596621542');

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
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_menu
-- ----------------------------
INSERT INTO `w_menu` VALUES ('2', '学校设置', 'fa-cog', 'config', '0', '1', '1', '1562304625', '1563170886');
INSERT INTO `w_menu` VALUES ('3', '用户管理', 'fa-users', 'user', '0', '3', '1', '1562304648', '1565053001');
INSERT INTO `w_menu` VALUES ('4', '系统设置', 'fa-cogs', 'sys', '0', '101', '0', '1562304671', '1568252505');
INSERT INTO `w_menu` VALUES ('9', '基本设置', '', 'config/index', '2', '1', '1', '1562315240', '1563170509');
INSERT INTO `w_menu` VALUES ('10', '学生列表', '', 'user/student', '3', '1', '1', '1562315268', '1562315268');
INSERT INTO `w_menu` VALUES ('11', '菜单管理', '', 'sys/menu', '4', '100', '1', '1562315306', '1562315306');
INSERT INTO `w_menu` VALUES ('14', '管理员列表', '', 'sys/admin', '4', '1', '1', '1562565735', '1562565735');
INSERT INTO `w_menu` VALUES ('15', '角色管理', '', 'sys/role', '4', '100', '1', '1562567947', '1562567947');
INSERT INTO `w_menu` VALUES ('16', '考勤管理', 'fa-bars', 'attence', '0', '5', '1', '1563846134', '1565576370');
INSERT INTO `w_menu` VALUES ('17', '识别记录', '', 'attence/index', '16', '100', '1', '1563846539', '1567910326');
INSERT INTO `w_menu` VALUES ('18', '考勤记录', '', 'attence/attenceCount', '16', '100', '1', '1563846576', '1565576570');
INSERT INTO `w_menu` VALUES ('19', '考勤设置', '', 'config/attence', '2', '100', '1', '1564984458', '1564984458');
INSERT INTO `w_menu` VALUES ('20', '班级管理', 'fa-building', 'depart', '0', '2', '1', '1564987633', '1564987633');
INSERT INTO `w_menu` VALUES ('21', '班级列表', '', 'depart/index', '20', '100', '1', '1564987736', '1564987736');
INSERT INTO `w_menu` VALUES ('22', '年级列表', '', 'depart/grade', '20', '100', '1', '1564987811', '1564987811');
INSERT INTO `w_menu` VALUES ('23', '家长列表', '', 'user/parent', '3', '3', '1', '1565057495', '1565057495');
INSERT INTO `w_menu` VALUES ('24', '人脸照片', '', 'user/face', '3', '2', '1', '1565075093', '1565075093');
INSERT INTO `w_menu` VALUES ('25', '设备管理', 'fa-video-camera', 'device', '0', '4', '1', '1565080559', '1567819161');
INSERT INTO `w_menu` VALUES ('26', '人脸识别', '', 'device/index', '25', '100', '1', '1565080649', '1567819178');
INSERT INTO `w_menu` VALUES ('27', '添加', '', 'device/add', '26', '100', '1', '1565082377', '1565082377');
INSERT INTO `w_menu` VALUES ('28', '添加', '', 'config/attenceAdd', '19', '100', '1', '1567774031', '1567774031');
INSERT INTO `w_menu` VALUES ('29', '班级考勤', '', 'attence/depart', '16', '1', '1', '1568021138', '1568021138');

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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_role
-- ----------------------------
INSERT INTO `w_role` VALUES ('1', '超级管理员', '1562573108');
INSERT INTO `w_role` VALUES ('9', '网站管理员', '1562585552');

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
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of w_role_menu
-- ----------------------------
INSERT INTO `w_role_menu` VALUES ('88', '1', '2');
INSERT INTO `w_role_menu` VALUES ('89', '1', '13');
INSERT INTO `w_role_menu` VALUES ('90', '1', '9');
INSERT INTO `w_role_menu` VALUES ('91', '1', '16');
INSERT INTO `w_role_menu` VALUES ('92', '1', '3');
INSERT INTO `w_role_menu` VALUES ('93', '1', '10');
INSERT INTO `w_role_menu` VALUES ('94', '1', '4');
INSERT INTO `w_role_menu` VALUES ('95', '1', '14');
INSERT INTO `w_role_menu` VALUES ('96', '1', '11');
INSERT INTO `w_role_menu` VALUES ('97', '1', '15');
INSERT INTO `w_role_menu` VALUES ('98', '9', '2');
INSERT INTO `w_role_menu` VALUES ('99', '9', '13');
INSERT INTO `w_role_menu` VALUES ('100', '9', '9');
INSERT INTO `w_role_menu` VALUES ('101', '9', '16');
INSERT INTO `w_role_menu` VALUES ('102', '9', '3');
INSERT INTO `w_role_menu` VALUES ('103', '9', '10');

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
