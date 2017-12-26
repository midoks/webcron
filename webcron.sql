/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50624
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50624
File Encoding         : 65001

Date: 2017-11-27 15:04:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `app_debugs`
-- ----------------------------
DROP TABLE IF EXISTS `wc_debugs`;
CREATE TABLE `app_debugs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL,
  `msg` text NOT NULL,
  `add_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of app_debugs
-- ----------------------------

-- ----------------------------
-- Table structure for `sys_func`
-- ----------------------------
DROP TABLE IF EXISTS `sys_func`;
CREATE TABLE `sys_func` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '名称',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父ID 0顶级',
  `controller` varchar(100) NOT NULL,
  `action` varchar(100) DEFAULT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型：0controller 1action',
  `is_menu` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是菜单：0不是1是',
  `icon` varchar(100) DEFAULT NULL,
  `desc` varchar(200) DEFAULT NULL COMMENT '介绍',
  `sort` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：0无效1有效-1软删除',
  `update_time` int(10),
  `create_time` int(10),
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COMMENT='权限列表';

-- ----------------------------
-- Records of sys_func
-- ----------------------------
INSERT INTO `sys_func` VALUES ('1', '系统设置', '0', 'sys', 'index', '0', '1', 'fa fa-cog', '系统相关参数设置', '0', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('2', '管理员管理', '1', 'sysuser', 'index', '1', '1', 'fa fa-users', '添加、删除、编辑系统管理员的权限。', '0', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('3', '系统功能添加', '1', 'sysfunc', 'add', '1', '0', 'glyphicon glyphicon-th', '系统功能添加', '6', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('5', '功能管理', '1', 'sysfunc', 'index', '1', '1', '', '功能列表', '7', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('6', '系统功能删除', '1', 'sysfunc', 'del', '1', '0', '', '系统功能删除', '8', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('7', '添加管理员', '1', 'sysuser', 'add', '1', '0', 'glyphicon glyphicon-user', '添加管理员', '1', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('9', '管理员删除', '1', 'sysuser', 'del', '1', '0', '', '管理员删除', '2', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('11', '重置管理员密码', '1', 'sysuser', 'repwd', '1', '0', '', '重置管理员密码', '3', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('12', '锁定管理员', '1', 'sysuser', 'lock', '1', '0', '', '锁定管理员', '4', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('13', '系统功能锁定', '1', 'sysfunc', 'lock', '1', '0', '', '系统功能锁定', '9', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('14', '角色管理', '1', 'sysrole', 'index', '1', '1', 'fa fa-users', '系统功能锁定', '10', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('15', '添加角色', '1', 'sysrole', 'add', '1', '0', 'fa fa-users', '添加角色', '11', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('17', '删除角色', '1', 'sysrole', 'del', '1', '0', 'fa fa-users', '删除角色', '12', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('18', '锁定角色', '1', 'sysrole', 'lock', '1', '0', 'fa fa-users', '锁定角色', '13', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('19', '功能设置菜单', '1', 'sysfunc', 'setmenu', '1', '0', '', '功能设置菜单', '9', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('22', '加密相关', '0', 'mima', 'init', '0', '1', 'fa fa-bell', '', '1', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('23', 'authcode', '22', 'encrypt', 'index', '1', '1', '', 'discuz加密解密方案', '0', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('30', '功能升降序', '1', 'sysfunc', 'sort', '1', '0', '', '功能升降序', '5', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('31', '后台模版事例', '0', 'tpl', 'index', '0', '1', 'glyphicon glyphicon-star-empty', '后台模版事例', '2', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('32', '异步上传', '31', 'tpl', 'index', '1', '1', '', '', '0', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('33', '监控', '0', 'monitor', 'index', '0', '1', 'fa fa-bullhorn', '', '3', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('34', '列表', '33', 'monitor', 'index', '1', '1', '', '列表', '99', '1', '1489429439', '1489429439');
INSERT INTO `sys_func` VALUES ('37', '日志管理', '1', 'syslog', 'index', '1', '1', '', '日志管理', '99', '1', '1489429439', '1489429439');

-- ----------------------------
-- Table structure for `sys_logs`
-- ----------------------------
DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE `sys_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `msg` text NOT NULL,
  `add_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_logs
-- ----------------------------
INSERT INTO `sys_logs` VALUES ('1', '1', '1', '1', '1510207207');
INSERT INTO `sys_logs` VALUES ('2', '3', '1', '登录成功', '1510208513');
INSERT INTO `sys_logs` VALUES ('3', '3', '1', '登录成功', '1511332681');
INSERT INTO `sys_logs` VALUES ('4', '3', '1', '登录成功', '1511759612');

-- ----------------------------
-- Table structure for `sys_role`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `desc` varchar(200) DEFAULT NULL COMMENT '角色介绍',
  `list` text NOT NULL COMMENT '权限列表JSON',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态1有效0无效',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='管理员权限表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '系统管理员', '系统总管理员', '2,7,9,11,12,30,3,5,6,13,19,14,15,17,18,37,23,32,34', '1', '2017-03-15 15:39:20', '2017-11-09 13:49:30');
INSERT INTO `sys_role` VALUES ('3', '编辑', '普通编辑人员', '2,14,15,17,18,23,32', '1', '2017-11-03 17:35:17', '2017-11-08 18:51:20');

-- ----------------------------
-- Table structure for `sys_user`
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(150) NOT NULL COMMENT '登录名',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `nick` varchar(50) DEFAULT NULL COMMENT '昵称',
  `sex` tinyint(4) DEFAULT '0' COMMENT '1男0女',
  `mail` varchar(150) DEFAULT NULL COMMENT '邮箱',
  `tel` varchar(11) DEFAULT NULL COMMENT '手机号',
  `roleid` int(11) DEFAULT '0' COMMENT '所属角色',
  `status` tinyint(4) DEFAULT '1' COMMENT '状体1有效0无效',
  `create_time` int(10),
  `update_time` int(10),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='管理员';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('3', 'admin', '21232f297a57a5a743894a0e4a801fc3', '管理员', '1', 'admin@localhost', '13000000000', '1', '1', '1489429439', '1489429439');
INSERT INTO `sys_user` VALUES ('4', 'guest', '084e0343a0486ff05530df6c705c8bb4', 'guest', '1', '13800138000@qq.com', '13800138000', '3', '1', '1489429439', '1489429439');
