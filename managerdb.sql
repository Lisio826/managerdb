/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 80012
Source Host           : localhost:3306
Source Database       : managerdb

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2019-09-12 16:04:09
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for my_db_user
-- ----------------------------
DROP TABLE IF EXISTS `my_db_user`;
CREATE TABLE `my_db_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_pwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_status` int(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of my_db_user
-- ----------------------------
INSERT INTO `my_db_user` VALUES ('1', 'lf', '9d6b464ef48b40c12064c8b249cbf4ec', '1');

-- ----------------------------
-- Table structure for t_manage_permission
-- ----------------------------
DROP TABLE IF EXISTS `t_manage_permission`;
CREATE TABLE `t_manage_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `permission_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `permission_code` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `permission_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `permission_status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  `operate_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `updata_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of t_manage_permission
-- ----------------------------

-- ----------------------------
-- Table structure for t_manage_resource
-- ----------------------------
DROP TABLE IF EXISTS `t_manage_resource`;
CREATE TABLE `t_manage_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `resource_code` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `resource_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `resource_status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  `operate_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `updata_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of t_manage_resource
-- ----------------------------

-- ----------------------------
-- Table structure for t_manage_role
-- ----------------------------
DROP TABLE IF EXISTS `t_manage_role`;
CREATE TABLE `t_manage_role` (
  `id` int(11) NOT NULL,
  `role_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `role_code` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `role_status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  `operate_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `updata_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of t_manage_role
-- ----------------------------

-- ----------------------------
-- Table structure for t_manage_role_resource_scope
-- ----------------------------
DROP TABLE IF EXISTS `t_manage_role_resource_scope`;
CREATE TABLE `t_manage_role_resource_scope` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) DEFAULT NULL,
  `resouce_id` int(11) DEFAULT NULL,
  `permission_id` int(11) DEFAULT NULL,
  `status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `operate_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  `id_delete` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of t_manage_role_resource_scope
-- ----------------------------

-- ----------------------------
-- Table structure for t_manage_user
-- ----------------------------
DROP TABLE IF EXISTS `t_manage_user`;
CREATE TABLE `t_manage_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `real_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `sur_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `user_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `identity` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `mobile` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `role_id` int(11) NOT NULL,
  `avatar` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  `operate_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of t_manage_user
-- ----------------------------

-- ----------------------------
-- Table structure for t_manage_user_resouce_scope
-- ----------------------------
DROP TABLE IF EXISTS `t_manage_user_resouce_scope`;
CREATE TABLE `t_manage_user_resouce_scope` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `resouce_id` int(11) DEFAULT NULL,
  `permission_id` int(11) DEFAULT NULL,
  `status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `operate_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  `is_delete` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of t_manage_user_resouce_scope
-- ----------------------------
