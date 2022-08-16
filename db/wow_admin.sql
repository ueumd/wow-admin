/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : localhost:3306
 Source Schema         : wow_admin

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 16/08/2022 18:07:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for citys
-- ----------------------------
DROP TABLE IF EXISTS `citys`;
CREATE TABLE `citys`  (
  `id` int(11) NOT NULL COMMENT '城市ID',
  `cityName` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '城市名称',
  `parentId` int(11) NOT NULL COMMENT '父城市ID',
  `shortName` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '简介',
  `levelType` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '城市等级',
  `cityCode` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '城市代码',
  `zipCode` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '邮编',
  `mergerName` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '城市名称',
  `longitude` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '经度',
  `latitude` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '纬度',
  `pinyin` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '简拼',
  `isDel` tinyint(4) NOT NULL DEFAULT 0 COMMENT '1 已删除  0 未删除',
  `createTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建日期',
  `updateId` int(11) NOT NULL DEFAULT 0 COMMENT '修改人Id',
  `updateTime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '修改日期',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_parentId`(`parentId`) USING BTREE,
  INDEX `idx_cityName`(`cityName`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '城市信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of citys
-- ----------------------------
INSERT INTO `citys` VALUES (100000, '中国', 0, '中国', '0', '', '', '中国', '116.3683244', '39.915085', 'China', 0, '2019-11-27 11:14:53', 0, '2019-12-13 19:23:57');
INSERT INTO `citys` VALUES (110000, '北京', 100000, '北京', '1', '', '', '中国,北京', '116.405285', '39.904989', 'Beijing', 0, '2019-11-27 11:14:53', 0, '2019-12-13 19:23:57');
INSERT INTO `citys` VALUES (110100, '北京市', 110000, '北京', '2', '010', '100000', '中国,北京,北京市', '116.405285', '39.904989', 'Beijing', 0, '2019-11-27 11:14:53', 0, '2019-12-13 19:23:57');
INSERT INTO `citys` VALUES (110101, '东城区', 110100, '东城', '3', '010', '100010', '中国,北京,北京市,东城区', '116.41005', '39.93157', 'Dongcheng', 0, '2019-11-27 11:14:53', 0, '2019-12-13 19:23:57');

-- ----------------------------
-- Table structure for log_operation_infos
-- ----------------------------
DROP TABLE IF EXISTS `log_operation_infos`;
CREATE TABLE `log_operation_infos`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `userId` int(11) NOT NULL COMMENT '用户id',
  `phone` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `opType` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '操作类型\n   login\n   logout\n   updateInfo',
  `opContent` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `isDel` tinyint(4) NOT NULL DEFAULT 0 COMMENT '1已删除 0未删除',
  `createTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updateTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_phone`(`phone`) USING BTREE,
  INDEX `idx_userId`(`userId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 256 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '操作流水表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log_operation_infos
-- ----------------------------

-- ----------------------------
-- Table structure for register_info
-- ----------------------------
DROP TABLE IF EXISTS `register_info`;
CREATE TABLE `register_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `userId` int(11) NOT NULL,
  `callName` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `childAge` int(2) NOT NULL DEFAULT 0,
  `childName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `childSex` int(1) NOT NULL DEFAULT 0 COMMENT '1 男 2 女 3保密',
  `parentWho` tinyint(2) NOT NULL DEFAULT 0 COMMENT '	',
  `cityCode` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `cityName` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `budget` int(2) NOT NULL DEFAULT 0,
  `learningLesson` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `studyWishSkill` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `givePriorityToLesson` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `intentionStudyType` int(2) NOT NULL DEFAULT 0,
  `isDel` tinyint(4) NOT NULL DEFAULT 0 COMMENT '1已删除 0未删除',
  `createTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updateId` int(11) NOT NULL DEFAULT 0,
  `updateTime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`, `intentionStudyType`) USING BTREE,
  UNIQUE INDEX `idx_userId`(`userId`) USING BTREE,
  INDEX `idx_phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '注册信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of register_info
-- ----------------------------
INSERT INTO `register_info` VALUES (17, '1391245678', 21, '', 1, 'hsd', 1, 1, '310101', '黄浦区', 1, '美术 ,编程', '2,3', '1,2,4,3', 1, 0, '2019-12-24 15:02:59', 0, '2022-08-16 17:45:35');

-- ----------------------------
-- Table structure for wechat_user
-- ----------------------------
DROP TABLE IF EXISTS `wechat_user`;
CREATE TABLE `wechat_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `unionid` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `openid` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `phone` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '呢称',
  `gender` tinyint(1) NOT NULL DEFAULT 3 COMMENT '1男，2:女, 3:保密',
  `headImgUrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `sketch` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '简述',
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '当前位置',
  `longitude` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '经度',
  `latitude` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '纬度',
  `isDel` tinyint(4) NOT NULL DEFAULT 0 COMMENT '1已删除 0未删除',
  `createTime` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updateId` int(11) NOT NULL DEFAULT 0,
  `updateTime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_phone`(`phone`) USING BTREE,
  INDEX `idx_unionid`(`unionid`) USING BTREE,
  INDEX `idx_nickname`(`nickname`) USING BTREE,
  INDEX `idx_openid`(`openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10015 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '微信用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wechat_user
-- ----------------------------
INSERT INTO `wechat_user` VALUES (10012, 'oITiajhIe-C52LD0aTV2XPXiyITA', 'oC-fN4rpL3GBoZZNK06xWPl-5MfI', '13912345678', 'hsd', 1, '', '123456', '', '', 'JingShanghai', '', '', 0, '2020-03-03 14:08:43', 0, '2022-08-16 18:00:29');

SET FOREIGN_KEY_CHECKS = 1;
