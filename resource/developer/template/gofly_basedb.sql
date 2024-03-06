/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : gofly_enterprise_test

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 19/02/2024 00:18:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for attachment
-- ----------------------------
DROP TABLE IF EXISTS `attachment`;
CREATE TABLE `attachment`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '商户账号',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '上传用户',
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '分类',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '访问路径',
  `imagewidth` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '宽度',
  `imageheight` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '高度',
  `imagetype` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '图片类型',
  `imageframes` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片帧数',
  `filesize` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `mimetype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'mime类型',
  `extparam` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '透传数据',
  `storage` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'local' COMMENT '存储位置',
  `sha1` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件 sha1编码',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件名称',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '附件名称',
  `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频封面',
  `updatetime` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `uploadtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 745 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '附件管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of attachment
-- ----------------------------
INSERT INTO `attachment` VALUES (722, 1, 0, 1, 0, 'resource/uploads/20230309/162ba4b5924cc0fe399d7a2ffd1d1110.png', '', '', '', 0, 21902, 'image/png', '', 'E:\\Project\\go\\src\\GoFlyAdmin\\tmpresource\\uploads\\20230309\\162ba4b5924cc0fe399d7a2ffd1d1110.png', '820751820adeadea0c353765e7c7137b', '接种后定时回访缩略图', '接种后定时回访缩略图.png', '', 1678373466, 1678373466);
INSERT INTO `attachment` VALUES (723, 1, 0, 1, 0, 'resource/uploads/20230310/b4ac2e2246073c50c9dc764d5b426720.png', '', '', '', 0, 21902, 'image/png', '', 'E:\\Project\\go\\src\\GoFlyAdmin\\tmpresource\\uploads\\20230310\\b4ac2e2246073c50c9dc764d5b426720.png', 'd811adccbd7d0b424f6de118c09e7ed3', '接种后定时回访缩略图', '接种后定时回访缩略图.png', '', 1678455956, 1678455956);
INSERT INTO `attachment` VALUES (725, 1, 0, 1, 0, 'resource/uploads/20230506/d4bb8324ee87699ed727bc1fd0479b34.jpg', '', '', '', 0, 2621283, 'image/jpeg', '', '/dataDB/project/go/gofly_singleresource\\uploads\\20230506\\d4bb8324ee87699ed727bc1fd0479b34.jpg', 'b089570f68f71c6f4175b1e91cac6014', 'Default', 'Default.jpg', '', 1683383609, 1683383609);
INSERT INTO `attachment` VALUES (726, 1, 0, 1, 0, 'resource/uploads/20230507/7992812e7e9f2b140968ba3874de1d1a.jpg', '', '', '', 0, 23454, 'image/jpeg', '', '/dataDB/project/go/gofly_singleresource\\uploads\\20230507\\7992812e7e9f2b140968ba3874de1d1a.jpg', 'c378ea13ca636384da41926588b95fdb', '0', '0.jpg', '', 1683389595, 1683389595);
INSERT INTO `attachment` VALUES (727, 1, 0, 1, 0, 'resource/uploads/20230607/f1fbf7039464d632d9b5fcecb1e41fab.png', '', '', '', 0, 337597, 'image/png', '', '/dataDB/project/go/gofly_singleresource\\uploads\\20230607\\f1fbf7039464d632d9b5fcecb1e41fab.png', 'ca7ce059dbafb26a728f8d3c66ef5cb6', 'loginbanner1', 'loginbanner1.png', '', 1686135515, 1686135515);
INSERT INTO `attachment` VALUES (728, 1, 0, 1, 0, 'resource/uploads/20230607/4825b3bc4721d2e6266b9696f47b23c5.png', '', '', '', 0, 277224, 'image/png', '', '/dataDB/project/go/gofly_singleresource\\uploads\\20230607\\4825b3bc4721d2e6266b9696f47b23c5.png', '5120d2bcc567803f0435e4782e857457', 'loginbanner2', 'loginbanner2.png', '', 1686135612, 1686135612);
INSERT INTO `attachment` VALUES (729, 1, 0, 1, 0, 'resource/uploads/20230607/33926ec2fcbc2da95e9cae158e00019e.png', '', '', '', 0, 136610, 'image/png', '', '/dataDB/project/go/gofly_singleresource\\uploads\\20230607\\33926ec2fcbc2da95e9cae158e00019e.png', '8b13300e268f2f4ddf100eaf8c2876b9', 'loginbanner3', 'loginbanner3.png', '', 1686135659, 1686135659);
INSERT INTO `attachment` VALUES (730, 1, 0, 1, 0, 'resource/uploads/20230608/eaf1511fa669c7dd54af301d50c9478e.png', '', '', '', 0, 277224, 'image/png', '', '/dataDB/project/go/gofly_singleresource\\uploads\\20230608\\eaf1511fa669c7dd54af301d50c9478e.png', '53bc5d1a0ac48e75121295b5c1e004ce', 'loginbanner2', 'loginbanner2.png', '', 1686219829, 1686219829);
INSERT INTO `attachment` VALUES (733, 1, 0, 1, 0, 'resource/uploads/20230609/82b4e47320cd007879ff180ca63fe2b2.png', '', '', '', 0, 351874, 'image/png', '', '/dataDB/project/go/gofly_singleresource\\uploads\\20230609\\82b4e47320cd007879ff180ca63fe2b2.png', '8536ed44d16b2e7c1f354ced43f50b7b', 'menu', 'menu.png', '', 1686273353, 1686273353);
INSERT INTO `attachment` VALUES (734, 1, 0, 1, 0, 'resource/uploads/20231202/95ca6221f1ec5ba2f9739c7f4cb736c6.zip', '', '', '', 0, 129849, 'application/x-zip-compressed', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20231202\\95ca6221f1ec5ba2f9739c7f4cb7', '4aa32d3b5af3ea81d9ea362fad210b3b', 'wxsys', 'wxsys.zip', '', 1701522368, 1701522368);
INSERT INTO `attachment` VALUES (735, 1, 0, 1, 0, 'resource/uploads/20231202/363e90e3ae28ab59569347cb610a30f6.zip', '', '', '', 0, 124345, 'application/x-zip-compressed', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20231202\\363e90e3ae28ab59569347cb610a', '98e7ce8be96f1530d1de1786a891c8fc', 'wxplus', 'wxplus.zip', '', 1701529388, 1701529388);
INSERT INTO `attachment` VALUES (736, 1, 0, 1, 0, 'resource/uploads/20231203/b014ea30ed4e749fe61cd87f4d66ca5c.zip', '', '', '', 0, 129849, 'application/x-zip-compressed', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20231203\\b014ea30ed4e749fe61cd87f4d66', 'c3a1da93c79a72dd17a17b8a3faab206', 'wxsys', 'wxsys.zip', '', 1701533312, 1701533312);
INSERT INTO `attachment` VALUES (737, 1, 0, 1, 0, 'resource/uploads/20231203/f803f6e9cbab2eb6bb64ad1a18e37bab.zip', '', '', '', 0, 129849, 'application/x-zip-compressed', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20231203\\f803f6e9cbab2eb6bb64ad1a18e3', 'c3a1da93c79a72dd17a17b8a3faab206', 'wxsys', 'wxsys.zip', '', 1701533753, 1701533753);
INSERT INTO `attachment` VALUES (738, 1, 0, 1, 0, 'resource/uploads/20231203/99fa7bda307258942a6e84c41cf88eff.zip', '', '', '', 0, 129849, 'application/x-zip-compressed', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20231203\\99fa7bda307258942a6e84c41cf8', 'c3a1da93c79a72dd17a17b8a3faab206', 'wxsys', 'wxsys.zip', '', 1701534327, 1701534327);
INSERT INTO `attachment` VALUES (739, 1, 0, 1, 1, 'resource/uploads/20240120/269281af0350cb928a6ce8c9c59b7335.png', '', '', '', 0, 36400, 'image/png', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20240120\\269281af0350cb928a6ce8c9c59b', '8886f29db3f9396e566920d4bfa27adf', 'electron', 'electron.png', '', 1705736990, 1705736990);
INSERT INTO `attachment` VALUES (740, 1, 0, 1, 1, 'resource/uploads/20240120/c9b0b82ab76bafbf2c237c683a9ca6d3.png', '', '', '', 0, 55614, 'image/png', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20240120\\c9b0b82ab76bafbf2c237c683a9c', '3df0d0f5c7f0753d75cea532be054472', '103011293', '103011293.png', '', 1705737010, 1705737010);
INSERT INTO `attachment` VALUES (741, 1, 0, 1, 1, 'resource/uploads/20240120/bd1f5a91d41a69efa81c1d9e06447824.png', '', '', '', 0, 337597, 'image/png', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20240120\\bd1f5a91d41a69efa81c1d9e0644', '5da9e68d5ad04f4ea5f38ebccca908d9', 'loginbanner1', 'loginbanner1.png', '', 1705755853, 1705755853);
INSERT INTO `attachment` VALUES (742, 1, 0, 1, 1, 'resource/uploads/20240120/407161e2dbdc4cc4848bdc1a1478fde7.jpg', '', '', '', 0, 52295, 'image/jpeg', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20240120\\407161e2dbdc4cc4848bdc1a1478', '5d19ce226813f6755921e67ac606d875', '微信图片_20230816001705', '微信图片_20230816001705.jpg', '', 1705756877, 1705756877);
INSERT INTO `attachment` VALUES (743, 1, 0, 1, 1, 'resource/uploads/20240120/3f5ef54b8f57c7de9a702b7302567a87.png', '', '', '', 0, 3953, 'image/png', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20240120\\3f5ef54b8f57c7de9a702b730256', 'd363a8e018ba91331c43e2b2468a6717', 'getimage', 'getimage.png', '', 1705758575, 1705758575);
INSERT INTO `attachment` VALUES (744, 1, 0, 1, 1, 'resource/uploads/20240120/70180de5ee294a6a78ee88a2be52a796.png', '', '', '', 0, 686638, 'image/png', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmpresource\\uploads\\20240120\\70180de5ee294a6a78ee88a2be52', '31eeb6f38f852d46c94417b55c19c0a4', '微信截图_20230802001834 - 副本', '微信截图_20230802001834 - 副本.png', '', 1705758720, 1705758720);

-- ----------------------------
-- Table structure for business_account
-- ----------------------------
DROP TABLE IF EXISTS `business_account`;
CREATE TABLE `business_account`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `dept_id` int(11) NOT NULL COMMENT '部门id',
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `password` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码盐',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '姓名',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(145) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像',
  `tel` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备用电话用户自己填写',
  `mobile` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '邮箱',
  `lastLoginIp` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '最后登录IP',
  `lastLoginTime` int(11) NOT NULL DEFAULT 0 COMMENT '最后登录时间',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态0=正常，1=禁用',
  `validtime` int(11) NOT NULL DEFAULT 0 COMMENT '账号有效时间0=无限',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updatetime` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地址',
  `city` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '城市',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `company` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公司名称',
  `province` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '省份',
  `area` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地区',
  `fileSize` int(10) UNSIGNED NOT NULL DEFAULT 3787456512 COMMENT '附件存储空间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户端-用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_account
-- ----------------------------
INSERT INTO `business_account` VALUES (1, 1, 1, 1, 3, 'gofly', '47e3cee18368271b2dbe9e5a22caef88', '1697472561', '开发管理员', '黄师傅', 'resource/staticfile/avatar.png', '88422345', '18988347563', '550325@qq.com', '', 1668909071, 0, 0, 1666161776, 1701187998, '王府井', '昆明', '开发测试账号', 'GoFLy科技1', '', 'chaoyang', 2147483647);
INSERT INTO `business_account` VALUES (3, 1, 1, 3, 4, 'test', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试账号biz', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667142475, 1678550309, '', '', '', '试试', '', '', 2147483647);
INSERT INTO `business_account` VALUES (4, 1, 1, 4, 1, '123ss', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '销售员de', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667144713, 0, '', '', '', '', '', '', 2147483647);
INSERT INTO `business_account` VALUES (9, 1, 1, 9, 1, '22334', '166d2832ebcc7672e59d13f37a79f59e', '3305628230121721621', '新增账号', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1678370986, 1678373636, '五华区霖雨路江东耀龙康城27幢二单元502', '昆明市', '', '云律科技（云南）有限公司', '', '', 2147483647);

-- ----------------------------
-- Table structure for business_attachment
-- ----------------------------
DROP TABLE IF EXISTS `business_attachment`;
CREATE TABLE `business_attachment`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '附件',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '附件原来名称',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件名称',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '文件类型0=图片，1=文件夹,2=视频，3=音频',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '访问路径',
  `imagewidth` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '宽度',
  `imageheight` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '高度',
  `filesize` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `mimetype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'mime类型',
  `extparam` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '透传数据',
  `storage` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'local' COMMENT '存储位置',
  `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频封面',
  `sha1` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件 sha1编码',
  `is_common` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否公共1=是',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 108 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户端附件' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_attachment
-- ----------------------------
INSERT INTO `business_attachment` VALUES (1, 1, 0, 0, 0, '', '默认文件', 1, '', '', '', 0, '', '', 'local', '', '', 1, 1686302484);
INSERT INTO `business_attachment` VALUES (3, 3, 0, 0, 0, '', '新建文件夹', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686302628);
INSERT INTO `business_attachment` VALUES (4, 4, 0, 0, 0, '', '新建文件夹', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686303346);
INSERT INTO `business_attachment` VALUES (7, 7, 0, 1, 5, '', '新建文件夹', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686320109);
INSERT INTO `business_attachment` VALUES (9, 9, 0, 1, 6, '', '新建文件夹', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686320125);
INSERT INTO `business_attachment` VALUES (10, 10, 0, 1, 0, '', '新建文件夹', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686321444);
INSERT INTO `business_attachment` VALUES (11, 11, 0, 1, 0, '', '新建文件夹4', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686321777);
INSERT INTO `business_attachment` VALUES (24, 24, 0, 1, 0, '', '新建文件夹7', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325948);
INSERT INTO `business_attachment` VALUES (25, 25, 0, 1, 0, '', '新建文件夹8', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325949);
INSERT INTO `business_attachment` VALUES (26, 26, 0, 1, 0, '', '新建文件夹9', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325949);
INSERT INTO `business_attachment` VALUES (27, 27, 0, 1, 0, '', '新建文件夹10', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325949);
INSERT INTO `business_attachment` VALUES (28, 28, 0, 1, 0, '', '新建文件夹11', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325949);
INSERT INTO `business_attachment` VALUES (29, 29, 0, 1, 0, '', '新建文件夹12', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325950);
INSERT INTO `business_attachment` VALUES (30, 30, 0, 1, 0, '', '新建文件夹13', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325950);
INSERT INTO `business_attachment` VALUES (31, 31, 0, 1, 0, '', '新建文件夹13', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325950);
INSERT INTO `business_attachment` VALUES (32, 32, 0, 1, 0, '', '新建文件夹15', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325950);
INSERT INTO `business_attachment` VALUES (33, 33, 0, 1, 0, '', '新建文件夹16', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325950);
INSERT INTO `business_attachment` VALUES (34, 34, 0, 1, 0, '', '新建文件夹16', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325951);
INSERT INTO `business_attachment` VALUES (35, 35, 0, 1, 0, '', '新建文件夹18', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325951);
INSERT INTO `business_attachment` VALUES (36, 36, 0, 1, 0, '', '新建文件夹19', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325951);
INSERT INTO `business_attachment` VALUES (37, 37, 0, 1, 0, '', '新建文件夹20', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325951);
INSERT INTO `business_attachment` VALUES (38, 38, 0, 1, 0, '', '里面有文件', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686325951);
INSERT INTO `business_attachment` VALUES (45, 45, 0, 1, 38, '', '新建文件夹1', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326291);
INSERT INTO `business_attachment` VALUES (49, 49, 0, 1, 38, '', '新建文件夹2', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326342);
INSERT INTO `business_attachment` VALUES (50, 50, 0, 1, 38, '', '新建文件夹3', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326342);
INSERT INTO `business_attachment` VALUES (51, 51, 0, 1, 38, '', '新建文件夹3', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326342);
INSERT INTO `business_attachment` VALUES (52, 52, 0, 1, 38, '', '新建文件夹5', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326343);
INSERT INTO `business_attachment` VALUES (53, 53, 0, 1, 38, '', '新建文件夹5', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326342);
INSERT INTO `business_attachment` VALUES (54, 54, 0, 1, 38, '', '新建文件夹7', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326343);
INSERT INTO `business_attachment` VALUES (55, 55, 0, 1, 38, '', '新建文件夹8', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326343);
INSERT INTO `business_attachment` VALUES (56, 56, 0, 1, 38, '', '新建文件夹9', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326343);
INSERT INTO `business_attachment` VALUES (57, 57, 0, 1, 38, '', '新建文件夹10', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326343);
INSERT INTO `business_attachment` VALUES (58, 58, 0, 1, 38, '', '新建文件夹11', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326344);
INSERT INTO `business_attachment` VALUES (59, 59, 0, 1, 38, '', '新建文件夹11', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326344);
INSERT INTO `business_attachment` VALUES (60, 60, 0, 1, 38, '', '新建文件夹13', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326344);
INSERT INTO `business_attachment` VALUES (61, 61, 0, 1, 38, '', '新建文件夹14', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326344);
INSERT INTO `business_attachment` VALUES (62, 62, 0, 1, 38, '', '新建文件夹15', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326344);
INSERT INTO `business_attachment` VALUES (63, 63, 0, 1, 38, '', '新建文件夹16', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326345);
INSERT INTO `business_attachment` VALUES (64, 64, 0, 1, 38, '', '新建文件夹17', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326345);
INSERT INTO `business_attachment` VALUES (65, 65, 0, 1, 38, '', '新建文件夹18', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326345);
INSERT INTO `business_attachment` VALUES (66, 66, 0, 1, 38, '', '新建文件夹19', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326345);
INSERT INTO `business_attachment` VALUES (67, 67, 0, 1, 38, '', '新建文件夹20', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326345);
INSERT INTO `business_attachment` VALUES (68, 68, 0, 1, 38, '', '新建文件夹21', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326346);
INSERT INTO `business_attachment` VALUES (69, 69, 0, 1, 38, '', '新建文件夹22', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326346);
INSERT INTO `business_attachment` VALUES (70, 70, 0, 1, 38, '', '新建文件夹23', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326346);
INSERT INTO `business_attachment` VALUES (71, 71, 0, 1, 38, '', '新建文件夹24', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326346);
INSERT INTO `business_attachment` VALUES (72, 72, 0, 1, 38, '', '新建文件夹25', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326346);
INSERT INTO `business_attachment` VALUES (73, 73, 0, 1, 38, '', '新建文件夹26', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326347);
INSERT INTO `business_attachment` VALUES (74, 74, 0, 1, 38, '', '新建文件夹27', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326347);
INSERT INTO `business_attachment` VALUES (75, 75, 0, 1, 38, '', '新建文件夹28', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326347);
INSERT INTO `business_attachment` VALUES (76, 76, 0, 1, 38, '', '新建文件夹29', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326347);
INSERT INTO `business_attachment` VALUES (77, 77, 0, 1, 38, '', '新建文件夹30', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326347);
INSERT INTO `business_attachment` VALUES (78, 78, 0, 1, 38, '', '新建文件夹31', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326348);
INSERT INTO `business_attachment` VALUES (85, 85, 0, 1, 0, '', '新建文件夹21', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1686326628);
INSERT INTO `business_attachment` VALUES (90, 90, 0, 1, 0, '', '新建文件夹20', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1688543891);
INSERT INTO `business_attachment` VALUES (91, 91, 0, 1, 0, '', '新建文件夹21', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1688543892);
INSERT INTO `business_attachment` VALUES (94, 94, 0, 1, 0, '', '111', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1688543894);
INSERT INTO `business_attachment` VALUES (97, 97, 0, 1, 1, '', '新建文件夹1', 1, '', '', '', 0, '', '', 'local', '', '', 0, 1690647894);
INSERT INTO `business_attachment` VALUES (107, 107, 1, 1, 0, 'electron.png', 'electron', 0, 'resource/uploads/20231221/d397e65e2d50e0dcc2ced152fd8c224e.png', '', '', 36400, 'image/png', '', 'D:\\Project\\develop\\go\\src\\gofly_enterprise\\tmp/resource/uploads/20231221/d397e65e2d50e0dcc2ced152fd8', '', '62d3dbda952ff1bb762f92988fcbb293', 0, 1703132158);

-- ----------------------------
-- Table structure for business_auth_dept
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_dept`;
CREATE TABLE `business_auth_dept`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '上级部门',
  `weigh` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理后台部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_dept
-- ----------------------------
INSERT INTO `business_auth_dept` VALUES (1, 1, 1, '市场部门', 0, 1, 0, '营销', 1666972562);
INSERT INTO `business_auth_dept` VALUES (2, 1, 1, '第一组', 1, 2, 0, '', 1701187668);
INSERT INTO `business_auth_dept` VALUES (3, 1, 1, '研发部门', 1, 3, 0, '', 1660493302);
INSERT INTO `business_auth_dept` VALUES (4, 2, 2, '领导部门', 0, 4, 0, '', 1660493325);
INSERT INTO `business_auth_dept` VALUES (6, 2, 2, '人事组', 4, 6, 0, '', 1667827895);

-- ----------------------------
-- Table structure for business_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_role`;
CREATE TABLE `business_auth_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户id',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父级',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `rules` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '规则ID 所拥有的权限包扣父级',
  `menu` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '选择的id，用于编辑赋值',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `data_access` tinyint(1) NOT NULL DEFAULT 0 COMMENT '数据权限0=自己1=自己及子权限，2=全部',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `weigh` int(11) NOT NULL COMMENT '排序',
  `createtime` int(11) NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_role
-- ----------------------------
INSERT INTO `business_auth_role` VALUES (1, 0, 1, 0, '超级管理组', '*', '*', 0, 0, '账号的总管理员', 1, 1678549942);
INSERT INTO `business_auth_role` VALUES (5, 0, 1, 1, '销售员2', '8,11,13,49,59,6', '[8,11,13,49,59]', 0, 0, '产品销售组', 2, 1678550158);
INSERT INTO `business_auth_role` VALUES (6, 0, 1, 1, '管理员', '7,11,13,32,8,64,61,12,63,6', '[7,11,13,32,8,64,61,12,63]', 0, 0, '', 3, 1678549964);
INSERT INTO `business_auth_role` VALUES (7, 0, 1, 6, '编辑组', '7,34,33,11,12,6', '[7,34,33,11,12]', 0, 0, '', 4, 1678549960);
INSERT INTO `business_auth_role` VALUES (8, 0, 1, 6, '兼职组', '11,12,34,7,33', '[11,12,34,7,33]', 0, 0, 'ceshi', 5, 1667105411);
INSERT INTO `business_auth_role` VALUES (11, 0, 1, 0, '管理组', '8,9,10,6', '[8,9,10]', 0, 0, '', 11, 1678549957);
INSERT INTO `business_auth_role` VALUES (13, 0, 1, 0, '市场部门', '8,6', '[8]', 0, 0, '', 13, 1678549952);
INSERT INTO `business_auth_role` VALUES (16, 0, 1, 0, '财务室', '8,48,49,59,69,6', '[8,48,49,59,69]', 0, 0, '修改', 16, 1678549955);
INSERT INTO `business_auth_role` VALUES (19, 1, 1, 1, '新增权限2', '70,68', '[70]', 0, 0, '', 19, 1701191528);
INSERT INTO `business_auth_role` VALUES (20, 1, 1, 1, '123', '97,74,75', '[97,74,75]', 0, 0, '', 20, 1687527490);

-- ----------------------------
-- Table structure for business_auth_role_access
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_role_access`;
CREATE TABLE `business_auth_role_access`  (
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `role_id` int(11) NOT NULL DEFAULT 0 COMMENT '授权id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商务端菜单授权' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_role_access
-- ----------------------------
INSERT INTO `business_auth_role_access` VALUES (4, 1);
INSERT INTO `business_auth_role_access` VALUES (5, 6);
INSERT INTO `business_auth_role_access` VALUES (9, 6);
INSERT INTO `business_auth_role_access` VALUES (9, 5);
INSERT INTO `business_auth_role_access` VALUES (3, 5);
INSERT INTO `business_auth_role_access` VALUES (10, 5);
INSERT INTO `business_auth_role_access` VALUES (11, 1);
INSERT INTO `business_auth_role_access` VALUES (12, 1);
INSERT INTO `business_auth_role_access` VALUES (13, 1);
INSERT INTO `business_auth_role_access` VALUES (1, 1);

-- ----------------------------
-- Table structure for business_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_rule`;
CREATE TABLE `business_auth_rule`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `title` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `locale` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '中英文标题key',
  `orderNo` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型 0=目录，1=菜单，2=按钮',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '上一级',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图标',
  `routePath` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路由地址',
  `routeName` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路由名称',
  `component` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '组件路径',
  `redirect` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '重定向地址',
  `permission` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限标识',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态 0=启用1=禁用',
  `isExt` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否外链 0=否1=是',
  `keepalive` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否缓存 0=否1=是',
  `requiresAuth` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否需要登录鉴权 0=否1=是',
  `hideInMenu` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否在左侧菜单中隐藏该项 0=否1=是',
  `hideChildrenInMenu` tinyint(1) NOT NULL DEFAULT 0 COMMENT '强制在左侧菜单中显示单项 0=否1=是',
  `activeMenu` tinyint(1) NOT NULL DEFAULT 0 COMMENT '高亮设置的菜单项 0=否1=是',
  `noAffix` tinyint(1) NOT NULL DEFAULT 0 COMMENT '如果设置为true，标签将不会添加到tab-bar中 0=否1=是',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 437 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'C端-菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_rule
-- ----------------------------
INSERT INTO `business_auth_rule` VALUES (8, 1, '概况', '', 1, 1, 0, 'icon-dashboard', '/home', 'home', '/dashboard/workplace/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1666167477);
INSERT INTO `business_auth_rule` VALUES (11, 1, '角色管理', '', 11, 1, 61, '', 'role', 'role', '/system/role/index', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 1666336763);
INSERT INTO `business_auth_rule` VALUES (12, 1, '菜单管理(dev)', '', 12, 1, 61, '', 'rule', 'rule', '/system/rule/index', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 1657817329);
INSERT INTO `business_auth_rule` VALUES (13, 1, '部门管理', '', 13, 1, 61, '', 'dept', 'dept', '/system/dept/index', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 1660818242);
INSERT INTO `business_auth_rule` VALUES (61, 1, '系统设置', '', 78, 0, 0, 'icon-settings', '/system', 'system', 'LAYOUT', '/system/account', '', 0, 0, 0, 0, 0, 0, 0, 0, 1667145000);
INSERT INTO `business_auth_rule` VALUES (63, 1, '账户管理', '', 32, 1, 61, '', 'account', 'account', '/system/account/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1666167258);
INSERT INTO `business_auth_rule` VALUES (64, 1, '添加账号', '', 64, 2, 7, '', '', '', '', '', 'add', 0, 0, 0, 0, 0, 0, 0, 0, 1667142600);
INSERT INTO `business_auth_rule` VALUES (68, 1, '个人中心', '', 77, 0, 0, 'icon-user', '/user', 'user', 'LAYOUT', '/user/info', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678027122);
INSERT INTO `business_auth_rule` VALUES (69, 1, '账号信息', '', 70, 1, 68, '', 'info', 'info', '/user/info/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678027200);
INSERT INTO `business_auth_rule` VALUES (70, 1, '用户设置', '', 69, 1, 68, '', 'setting', 'setting', '/user/setting/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678027237);
INSERT INTO `business_auth_rule` VALUES (74, 1, '开发者', '', 80, 0, 0, 'icon-code', '/developer', 'developer', 'LAYOUT', '/developer/apidoc', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678550475);
INSERT INTO `business_auth_rule` VALUES (75, 1, '接口文档', '', 75, 1, 74, '', 'devapi', 'devapi', '/developer/devapi/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678550506);
INSERT INTO `business_auth_rule` VALUES (97, 1, '生成代码', '', 97, 1, 74, '', 'generatecode', 'generatecode', '/developer/generatecode/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1685366576);
INSERT INTO `business_auth_rule` VALUES (121, 1, '数据中心', '', 79, 0, 0, 'icon-storage', '/datacenter', 'datacenter', 'LAYOUT', '/datacenter/dictionary', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686132257);
INSERT INTO `business_auth_rule` VALUES (123, 1, '字典数据', '', 123, 1, 121, '', 'data', 'data', '/datacenter/dictionary/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686132869);
INSERT INTO `business_auth_rule` VALUES (137, 1, '附件管理', '', 137, 1, 121, '', 'attachment', 'attachment', 'datacenter/attachment/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686286572);
INSERT INTO `business_auth_rule` VALUES (143, 1, '配置管理', '', 143, 1, 121, '', 'configuration', 'configuration', '/datacenter/configuration/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1690646744);
INSERT INTO `business_auth_rule` VALUES (374, 1, '代码生成器', '', 374, 1, 74, '', 'codemaker', 'codemaker', '/developer/generatecode/codemaker.vue', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 1704534684);
INSERT INTO `business_auth_rule` VALUES (383, 1, '生成代码测试', '', 383, 0, 0, 'icon-archive', '/makecode', 'makecode', 'LAYOUT', '/makecode/test', '', 0, 0, 0, 1, 0, 0, 0, 0, 1704639220);
INSERT INTO `business_auth_rule` VALUES (435, 1, '测试代码产品', '', 435, 1, 383, 'icon-sun-fill', 'product', 'product', 'makecode/product/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1708272596);
INSERT INTO `business_auth_rule` VALUES (436, 1, '测试代码产品分类', '', 436, 1, 383, '', 'cate', 'cate', 'makecode/cate/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1708272637);

-- ----------------------------
-- Table structure for business_home_quickop
-- ----------------------------
DROP TABLE IF EXISTS `business_home_quickop`;
CREATE TABLE `business_home_quickop`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加人',
  `is_common` tinyint(1) NOT NULL DEFAULT 0 COMMENT '公共1=是',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型1=外部',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '快捷名称',
  `path_url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '跳转路径',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '权重',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '首页快捷操作' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_home_quickop
-- ----------------------------
INSERT INTO `business_home_quickop` VALUES (1, 1, 1, 0, 0, '文档接口', 'devapi', 'icon-common', 1);
INSERT INTO `business_home_quickop` VALUES (2, 1, 1, 0, 0, '生成代码', 'generatecode', 'icon-mobile', 2);

-- ----------------------------
-- Table structure for common_apidoc
-- ----------------------------
DROP TABLE IF EXISTS `common_apidoc`;
CREATE TABLE `common_apidoc`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'admin' COMMENT '分类接口属于那端，admin=管理，biz=B端，client=C端',
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '分组',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '接口名称',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '说明',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '接口路径',
  `param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '参数',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '方法get,post,delete',
  `tablename` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '数据表名',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `istoken` tinyint(1) NOT NULL DEFAULT 1 COMMENT '需要token 1=需要',
  `apicode_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '接口代码类型0=手动写，1=自动生成，2=通用接口',
  `is_install` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否安装0=未生成1=已生成=2已卸载',
  `fields` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '查询字段',
  `getdata_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '获取数据类型list=多条，detail=单条',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口测试数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apidoc
-- ----------------------------
INSERT INTO `common_apidoc` VALUES (1, 'biz', 2, '登录接口-获取openid', '1请求参数：\ncode： wx.login()中的code\n2返回参数：\nuserinfo:用户信息\ntoken：用户token，统一在请求中封装返回', '/wxapp/user/get_openid', '{\n\"code\": \"\"\n}', 'get', 'business_wxsys_user', 0, 0, 0, 0, '', '', 1697535606);
INSERT INTO `common_apidoc` VALUES (5, 'biz', 3, '后端测试接口', '业务端测试端，请求接口', '/business/test/api/get_data', '', 'get', '', 0, 1, 0, 0, '', '', 1697643557);
INSERT INTO `common_apidoc` VALUES (6, 'biz', 3, '测试获取列表数据接口', '', '/business/test/api/get_list', '', 'get', 'business_auth_rule', 0, 1, 0, 0, '', '', 1697535645);
INSERT INTO `common_apidoc` VALUES (7, 'biz', 2, '获取小程序数据', '', '/wxapp/test/wxapi/get_data', '', 'get', '', 0, 0, 0, 0, '', '', 1697535559);

-- ----------------------------
-- Table structure for common_apidoc_group
-- ----------------------------
DROP TABLE IF EXISTS `common_apidoc_group`;
CREATE TABLE `common_apidoc_group`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'admin' COMMENT '分类接口属于那端，admin=管理，biz=B端，client=C端',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父级0=一级',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `type_id` int(11) NOT NULL DEFAULT 0 COMMENT '接口类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '后台端接口测试分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apidoc_group
-- ----------------------------
INSERT INTO `common_apidoc_group` VALUES (1, 'biz', 0, 'app端', 0, 3);
INSERT INTO `common_apidoc_group` VALUES (2, 'biz', 0, '小程序', 0, 1);
INSERT INTO `common_apidoc_group` VALUES (3, 'biz', 0, '后台管理', 0, 2);
INSERT INTO `common_apidoc_group` VALUES (4, 'biz', 2, '小程序-疫苗计划', 0, 1);

-- ----------------------------
-- Table structure for common_apidoc_type
-- ----------------------------
DROP TABLE IF EXISTS `common_apidoc_type`;
CREATE TABLE `common_apidoc_type`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型名称',
  `rooturl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求服务器地址',
  `verifyEncrypt` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密验证字符串',
  `isself` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是本端1=是',
  `user_tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '测试授权用户数据表名',
  `user_id` int(10) NOT NULL DEFAULT 0 COMMENT '测试用户id',
  `login_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录地址',
  `model_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '模块目录',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口类型' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apidoc_type
-- ----------------------------
INSERT INTO `common_apidoc_type` VALUES (1, '小程序', 'https://yg.goflys.cn', 'gofly@888', 0, 'business_wxsys_user', 6, '/wxapp/user/get_apitoken', 'wxapp');
INSERT INTO `common_apidoc_type` VALUES (2, '本端', '', '', 1, '', 0, '', '');
INSERT INTO `common_apidoc_type` VALUES (3, '手机APP', 'https://yg.goflys.cn', 'gofly@888', 0, '', 0, '', '');

-- ----------------------------
-- Table structure for common_config
-- ----------------------------
DROP TABLE IF EXISTS `common_config`;
CREATE TABLE `common_config`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `data_from` enum('common','business') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT '数据来源common=公共，business=商业端',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `keyname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置名称',
  `keyvalue` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置值',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统配置参数' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_config
-- ----------------------------
INSERT INTO `common_config` VALUES (2, 'common', 0, 'rooturl', 'http://localhost:8108/common/uploadfile/get_image?url=', '图片路径', 0);

-- ----------------------------
-- Table structure for common_dictionary_data
-- ----------------------------
DROP TABLE IF EXISTS `common_dictionary_data`;
CREATE TABLE `common_dictionary_data`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `data_from` enum('common','business') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT '数据来源common=公共，business=商业端',
  `keyname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典名称',
  `keyvalue` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典项值',
  `des` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典描述',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updatetime` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据-测试数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_dictionary_data
-- ----------------------------
INSERT INTO `common_dictionary_data` VALUES (1, 'common', '管理层', 'mteam', '公司领导', 0, 1, 1686156976, 0);
INSERT INTO `common_dictionary_data` VALUES (2, 'common', '业务员', 'salesman', '', 0, 2, 1691760155, 0);

-- ----------------------------
-- Table structure for common_dictionary_table
-- ----------------------------
DROP TABLE IF EXISTS `common_dictionary_table`;
CREATE TABLE `common_dictionary_table`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `data_from` enum('common','business') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT '数据来源common=公共，business=商业端',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典名称',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数据表名称',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_dictionary_table
-- ----------------------------
INSERT INTO `common_dictionary_table` VALUES (2, 0, 'business', '用户分组', '用户分组', 'common_dictionary_data', 0, 2, 1686152478);

-- ----------------------------
-- Table structure for common_email
-- ----------------------------
DROP TABLE IF EXISTS `common_email`;
CREATE TABLE `common_email`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `data_from` enum('common','business') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT '数据来源common=公共，business=商业端',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `sender_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '发送者邮箱',
  `auth_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱授权码',
  `mail_title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件标题',
  `mail_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件内容,可以是html',
  `service_host` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件服务器',
  `service_port` int(11) NOT NULL DEFAULT 0 COMMENT '邮件服务器端口',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '业务端邮箱' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_email
-- ----------------------------
INSERT INTO `common_email` VALUES (1, 'business', 1, '504500934@qq.com', 'amidmyjnnxy(youwkey)', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);
INSERT INTO `common_email` VALUES (2, 'common', 0, '504500934@qq.com', 'amidmyjnnxy(youkey)', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);

-- ----------------------------
-- Table structure for common_generatecode
-- ----------------------------
DROP TABLE IF EXISTS `common_generatecode`;
CREATE TABLE `common_generatecode`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表名称',
  `comment` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表备注',
  `engine` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '引擎',
  `table_rows` int(11) NOT NULL DEFAULT 0 COMMENT '记录数',
  `collation` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '编码',
  `auto_increment` int(11) NOT NULL DEFAULT 1 COMMENT '自增索引',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '菜单上级',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图标',
  `routePath` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由地址',
  `routeName` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路由名称',
  `component` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '组件路径',
  `api_path` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '后端业务接口',
  `api_filename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '后端文件名',
  `fields` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '查询字段',
  `rule_id` int(11) NOT NULL DEFAULT 0 COMMENT '生成菜单id',
  `rule_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `is_install` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否安装0=未安装，1=已安装，2=已卸载',
  `tpl_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'list' COMMENT '模板类型list=仅一个数据，cate=数据加分类',
  `cate_tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类表名称',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  `updatetime` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 61 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_generatecode
-- ----------------------------
INSERT INTO `common_generatecode` VALUES (1, 'admin_auth_dept', '管理后台部门', 'InnoDB', 5, 'utf8mb4_general_ci', 6, 1, 0, '', '', '', '', '', '', '', 0, '管理后台部门', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (2, 'admin_auth_role', '权限分组', 'InnoDB', 8, 'utf8mb4_general_ci', 16, 1, 0, '', '', '', '', '', '', '', 0, '权限分组', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (3, 'admin_auth_role_access', 'admin端菜单权限', 'InnoDB', 6, 'utf8mb4_general_ci', 0, 1, 0, '', '', '', '', '', '', '', 0, 'admin端菜单权限', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (4, 'admin_auth_rule', 'C端-菜单', 'InnoDB', 22, 'utf8mb4_general_ci', 80, 1, 0, '', '', '', '', '', '', '', 0, 'C端-菜单', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (6, 'attachment', '附件管理', 'InnoDB', 19, 'utf8mb4_general_ci', 744, 1, 0, '', '', '', '', '', '', '', 0, '附件管理', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (7, 'business_attachment', '客户端附件', 'InnoDB', 60, 'utf8mb4_general_ci', 108, 1, 0, '', '', '', '', '', '', '', 0, '客户端附件', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (8, 'business_auth_dept', '管理后台部门', 'InnoDB', 5, 'utf8mb4_general_ci', 7, 1, 0, '', '', '', '', '', '', '', 0, '管理后台部门', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (9, 'business_auth_role', '权限分组', 'InnoDB', 10, 'utf8mb4_general_ci', 21, 1, 0, '', '', '', '', '', '', '', 0, '权限分组', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (10, 'business_auth_role_access', '商务端菜单授权', 'InnoDB', 10, 'utf8mb4_general_ci', 0, 1, 0, '', '', '', '', '', '', '', 0, '商务端菜单授权', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (11, 'business_auth_rule', 'C端-菜单', 'InnoDB', 20, 'utf8mb4_general_ci', 435, 1, 0, '', '', '', '', '', '', '', 0, 'C端-菜单', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (16, 'business_home_quickop', '首页快捷操作', 'InnoDB', 2, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, '首页快捷操作', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (18, 'business_wxsys_officonfig', '微信公众号配置', 'InnoDB', 2, 'utf8mb4_general_ci', 6, 1, 0, '', '', '', '', '', '', '', 0, '微信公众号配置', 0, 'list', '', 1692894701, 1706080118);
INSERT INTO `common_generatecode` VALUES (19, 'business_wxsys_user', '微信关注用户', 'InnoDB', 0, 'utf8mb4_general_ci', 0, 1, 0, '', '', '', '', '', '', '', 0, '微信关注用户', 0, 'list', '', 1692894701, 1706080118);
INSERT INTO `common_generatecode` VALUES (20, 'business_wxsys_wxappconfig', '微信小程序配置', 'InnoDB', 2, 'utf8mb4_general_ci', 6, 1, 0, '', '', '', '', '', '', '', 0, '微信小程序配置', 0, 'list', '', 1692894701, 1706080118);
INSERT INTO `common_generatecode` VALUES (21, 'business_wxsys_wxmenu', '微站微信菜单', 'InnoDB', 2, 'utf8mb4_general_ci', 13, 1, 0, '', '', '', '', '', '', '', 0, '微站微信菜单', 0, 'list', '', 1692894701, 1706080118);
INSERT INTO `common_generatecode` VALUES (25, 'common_config', '系统配置参数', 'InnoDB', 0, 'utf8mb4_general_ci', 3, 1, 0, '', '', '', '', '', '', '', 0, '系统配置参数', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (26, 'common_dictionary_data', '字典数据-测试数据', 'InnoDB', 2, 'utf8mb4_general_ci', 2, 1, 0, '', '', '', '', '', '', '', 0, '字典数据-测试数据', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (27, 'common_dictionary_integral', '积分等级-测试数据', 'InnoDB', 3, 'utf8mb4_general_ci', 3, 1, 0, '', '', '', '', '', '', '', 0, '', 0, 'list', '', 1692894701, 1700667455);
INSERT INTO `common_generatecode` VALUES (28, 'common_dictionary_table', '字典表', 'InnoDB', 2, 'utf8mb4_general_ci', 2, 1, 0, '', '', '', '', '', '', '', 0, '字典表', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (29, 'common_email', '业务端邮箱', 'InnoDB', 2, 'utf8mb4_general_ci', 2, 1, 0, '', '', '', '', '', '', '', 0, '业务端邮箱', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (30, 'common_generatecode', '代码生成', 'InnoDB', 43, 'utf8mb4_general_ci', 61, 1, 0, '', '', '', '', '', '', '', 0, '代码生成', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (31, 'common_logininfo', '登录页面内容', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, '登录页面内容', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (32, 'common_message', '系统通用消息', 'InnoDB', 0, 'utf8mb4_general_ci', 0, 1, 0, '', '', '', '', '', '', '', 0, '系统通用消息', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (33, 'common_picture', '图片库', 'InnoDB', 4, 'utf8mb4_general_ci', 8, 1, 0, '', '', '', '', '', '', '', 0, '图片库', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (34, 'common_picture_cate', '分类名称', 'InnoDB', 27, 'utf8mb4_general_ci', 27, 1, 0, '', '', '', '', '', '', '', 0, '分类名称', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (35, 'common_verify_code', '验证码存储', 'InnoDB', 0, 'utf8mb4_general_ci', 1, 1, 0, '', '', '', '', '', '', '', 0, '验证码存储', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (36, 'login_logs', '（平台及客户）后台登录日志', 'InnoDB', 346, 'utf8mb4_general_ci', 976, 1, 0, '', '', '', '', '', '', '', 0, '（平台及客户）后台登录日志', 0, 'list', '', 1692894701, 1708273043);
INSERT INTO `common_generatecode` VALUES (41, 'business_email', '业务端邮箱', 'InnoDB', 0, 'utf8mb4_general_ci', 1, 1, 0, '', '', '', '', '', '', '', 0, '', 0, 'list', '', 1699718494, 1700667455);
INSERT INTO `common_generatecode` VALUES (42, 'admin_account', '用户端-用户信息', 'InnoDB', 4, 'utf8mb4_general_ci', 9, 1, 0, '', '', '', '', '', '', '', 0, '用户端-用户信息', 0, 'list', '', 1699972481, 1708273043);
INSERT INTO `common_generatecode` VALUES (43, 'business_account', '用户端-用户信息', 'InnoDB', 7, 'utf8mb4_general_ci', 14, 1, 0, '', '', '', '', '', '', '', 0, '用户端-用户信息', 0, 'list', '', 1699972481, 1708273043);
INSERT INTO `common_generatecode` VALUES (44, 'common_apidoc', '接口测试数据', 'InnoDB', 4, 'utf8mb4_general_ci', 19, 1, 0, '', '', '', '', '', '', '', 0, '接口测试数据', 0, 'list', '', 1699972481, 1708273043);
INSERT INTO `common_generatecode` VALUES (45, 'common_apidoc_group', '后台端接口测试分组', 'InnoDB', 4, 'utf8mb4_0900_ai_ci', 5, 1, 0, '', '', '', '', '', '', '', 0, '后台端接口测试分组', 0, 'list', '', 1699972481, 1708273043);
INSERT INTO `common_generatecode` VALUES (46, 'common_apidoc_type', '接口类型', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, '接口类型', 0, 'list', '', 1699972481, 1708273043);
INSERT INTO `common_generatecode` VALUES (51, 'createcode_product', '测试代码产品', 'InnoDB', 2, 'utf8mb4_general_ci', 4, 0, 383, 'icon-sun-fill', 'product', 'product', 'makecode/product/index', 'business/makecode', 'product.go', 'id,title,price,num,createtime', 435, '测试代码产品', 1, 'contentlist', '', 1700667455, 1708273043);
INSERT INTO `common_generatecode` VALUES (52, 'common_generatecode_field', '生成代码字段管理', 'InnoDB', 9, 'utf8mb4_general_ci', 45, 1, 0, '', '', '', '', '', '', '', 0, '生成代码字段管理', 0, 'list', '', 1704615488, 1708273043);
INSERT INTO `common_generatecode` VALUES (53, 'createcode_product_cate', '测试代码产品分类', 'InnoDB', 0, 'utf8mb4_general_ci', 0, 0, 383, '', 'cate', 'cate', 'makecode/cate/index', 'business/makecode', 'cate.go', '', 436, '测试代码产品分类', 1, 'list', 'createcode_product_cate', 1704707236, 1708273043);
INSERT INTO `common_generatecode` VALUES (54, 'business_website_article_cate', '网站管理-文章分类', 'InnoDB', 15, 'utf8_general_ci', 21, 1, 0, '', '', '', '', '', '', '', 0, '网站管理-文章分类', 0, 'list', '', 1705508094, 1706080118);
INSERT INTO `common_generatecode` VALUES (55, 'business_website_article_content', '网站管理-文章内容', 'InnoDB', 11, 'utf8_general_ci', 24, 1, 0, '', '', '', '', '', '', '', 0, '网站管理-文章内容', 0, 'list', '', 1705508094, 1706080118);
INSERT INTO `common_generatecode` VALUES (56, 'business_website_leavemessage', '网站管理-留言', 'InnoDB', 4, 'utf8_general_ci', 23, 1, 0, '', '', '', '', '', '', '', 0, '网站管理-留言', 0, 'list', '', 1705508094, 1706080118);
INSERT INTO `common_generatecode` VALUES (57, 'business_website_link', '网站-友情链接', 'InnoDB', 2, 'utf8_general_ci', 3, 1, 0, '', '', '', '', '', '', '', 0, '网站-友情链接', 0, 'list', '', 1705508094, 1706080118);
INSERT INTO `common_generatecode` VALUES (58, 'business_website_module', '网站模块', 'InnoDB', 7, 'utf8_general_ci', 9, 1, 0, '', '', '', '', '', '', '', 0, '网站模块', 0, 'list', '', 1705508094, 1706080118);
INSERT INTO `common_generatecode` VALUES (59, 'business_website_site', '网站管理-站点', 'InnoDB', 0, 'utf8_general_ci', 1, 1, 0, '', '', '', '', '', '', '', 0, '网站管理-站点', 0, 'list', '', 1705508094, 1706080118);
INSERT INTO `common_generatecode` VALUES (60, 'business_website_visit_record', '网站访问记录', 'InnoDB', 302, 'utf8mb4_general_ci', 1780, 1, 0, '', '', '', '', '', '', '', 0, '', 0, 'list', '', 1705508094, 1706080118);

-- ----------------------------
-- Table structure for common_generatecode_field
-- ----------------------------
DROP TABLE IF EXISTS `common_generatecode_field`;
CREATE TABLE `common_generatecode_field`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `generatecode_id` int(10) NOT NULL COMMENT '关联列表',
  `islist` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是列表1=是',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段名称',
  `field` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段',
  `isorder` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否参与排序',
  `align` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'left' COMMENT '对齐方向',
  `width` int(10) NOT NULL DEFAULT 0 COMMENT '宽度',
  `isform` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为表单字段',
  `required` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为必填项',
  `formtype` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表单类型',
  `datatable` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联数据表',
  `datatablename` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联显示字段',
  `issearch` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否查询',
  `searchway` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '=' COMMENT '查询方式',
  `searchtype` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '查询文本类型',
  `field_weigh` int(10) NOT NULL COMMENT '表单排序',
  `list_weigh` int(10) NOT NULL COMMENT '列表排序',
  `search_weigh` int(10) NOT NULL DEFAULT 0 COMMENT '搜索排序',
  `def_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '默认选项json',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 49 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '生成代码字段管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_generatecode_field
-- ----------------------------
INSERT INTO `common_generatecode_field` VALUES (37, 51, 1, 'ID', 'id', 1, 'left', 0, 0, 0, 'number', '', '', 0, '=', 'text', 1, 1, 1, '[]');
INSERT INTO `common_generatecode_field` VALUES (38, 51, 1, '标题1', 'title', 1, 'center', 200, 1, 1, 'text', '', '', 0, '=', 'text', 2, 2, 2, '[]');
INSERT INTO `common_generatecode_field` VALUES (39, 51, 1, '库存', 'num', 1, 'left', 0, 1, 1, 'number', '', '', 1, '=', 'text', 4, 5, 4, '[]');
INSERT INTO `common_generatecode_field` VALUES (40, 51, 1, '价格', 'price', 0, 'left', 0, 1, 0, 'number', 'business_auth_dept', 'pid', 0, '=', 'text', 3, 4, 3, '[]');
INSERT INTO `common_generatecode_field` VALUES (41, 51, 0, '内容', 'content', 1, 'left', 220, 1, 1, 'editor', 'business_auth_role', 'data_access', 0, '=', 'text', 5, 3, 5, '[]');
INSERT INTO `common_generatecode_field` VALUES (42, 51, 1, '上传时间', 'createtime', 0, 'left', 0, 0, 0, 'number', '', '', 0, '=', 'text', 6, 6, 6, '[]');
INSERT INTO `common_generatecode_field` VALUES (43, 50, 0, 'ID', 'id', 1, 'left', 0, 0, 0, 'number', '', '', 0, '=', 'text', 0, 0, 0, '[]');
INSERT INTO `common_generatecode_field` VALUES (44, 50, 0, '名称', 'name', 0, 'left', 0, 0, 0, 'text', '', '', 0, '=', 'text', 0, 0, 0, '[]');
INSERT INTO `common_generatecode_field` VALUES (45, 50, 0, '上传时间', 'createtime', 0, 'left', 0, 0, 0, 'number', '', '', 0, '=', 'text', 0, 0, 0, '[]');
INSERT INTO `common_generatecode_field` VALUES (46, 53, 0, '上传时间', 'createtime', 0, 'left', 0, 0, 0, 'number', '', '', 0, '=', 'text', 0, 0, 0, '[]');
INSERT INTO `common_generatecode_field` VALUES (47, 53, 0, 'ID', 'id', 1, 'left', 0, 0, 0, 'number', '', '', 0, '=', 'text', 0, 0, 0, '[]');
INSERT INTO `common_generatecode_field` VALUES (48, 53, 0, '名称', 'name', 0, 'left', 0, 0, 0, 'text', '', '', 0, '=', 'text', 0, 0, 0, '[]');

-- ----------------------------
-- Table structure for common_logininfo
-- ----------------------------
DROP TABLE IF EXISTS `common_logininfo`;
CREATE TABLE `common_logininfo`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` enum('admin','business','common') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT 'admin=管理端，business=商业端 common=公共',
  `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `image` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登录页面内容' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_logininfo
-- ----------------------------
INSERT INTO `common_logininfo` VALUES (1, 'common', '开箱即用Go应用系统', '为您开发搭建基础，您只专注业务开发，并能快速开始，我们做到同等功能少其他语言框架的的一半代码，这是做好用的Go语言web开发框架。', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230607/f1fbf7039464d632d9b5fcecb1e41fab.png', 0, 1, 1686135358);
INSERT INTO `common_logininfo` VALUES (2, 'common', 'GoFly简单易用的服务端', '我们为您做了自动路由，统一配置文件，日志文件，简单易用ORM，接口验证，防止高频攻击，Token验证，简介开发架构等等！', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230607/4825b3bc4721d2e6266b9696f47b23c5.png', 0, 2, 1686135616);
INSERT INTO `common_logininfo` VALUES (3, 'common', '轻量集成快速开发Admin/中台', '我们采用vue3、vite、ts搭建，采用最优美的UI，集成了自动加载路由，权限管理，账号管理，多语言包自动加载等等，用上我们框架保你快人一倍！', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230607/33926ec2fcbc2da95e9cae158e00019e.png', 0, 3, 1686135660);

-- ----------------------------
-- Table structure for common_message
-- ----------------------------
DROP TABLE IF EXISTS `common_message`;
CREATE TABLE `common_message`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `usertype` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户类型0=全部，1=系统-2=boss端，3=C端',
  `accountID` int(11) NOT NULL COMMENT '账号id',
  `adduid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `touid` int(11) NOT NULL DEFAULT 0 COMMENT '接收用户',
  `type` tinyint(1) NOT NULL DEFAULT 2 COMMENT '类型1=通知，2=消息，3=代办',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息标题',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '跳转路由',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `isread` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否已读1=已读',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '发送时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统通用消息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_message
-- ----------------------------

-- ----------------------------
-- Table structure for common_picture
-- ----------------------------
DROP TABLE IF EXISTS `common_picture`;
CREATE TABLE `common_picture`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加账号',
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '分类id',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '附件原来名称',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件名称',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型0=素材图1=插图,2=视频，3=音频',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '访问路径',
  `imagewidth` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '宽度',
  `imageheight` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '高度',
  `filesize` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `mimetype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'mime类型',
  `storage` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'local' COMMENT '存储位置',
  `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频封面',
  `sha1` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件 sha1编码',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '图片库' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_picture
-- ----------------------------
INSERT INTO `common_picture` VALUES (5, 1, 20, 5, 'GoFLy发布文章封面.png', 'GoFLy发布文章封面', 0, 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230609/00658402ef4e5ba229f3935eca6701d8.png', '', '', 40902, 'image/png', '/dataDB/project/go/gofly_singleresource\\uploads\\20230609\\00658402ef4e5ba229f3935eca6701d8.png', '', 'b98da546d168f3e1d91d32585aaf719e', 1686285651, 0);
INSERT INTO `common_picture` VALUES (6, 1, 24, 6, '信息.png', '信息', 1, 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230609/46e5cc40453791e1db8c0e25a1c8ff9c.png', '', '', 65892, 'image/png', '/dataDB/project/go/gofly_singleresource\\uploads\\20230609\\46e5cc40453791e1db8c0e25a1c8ff9c.png', '', 'd58b80c230362875af642143b6bd3a70', 1686296929, 0);
INSERT INTO `common_picture` VALUES (7, 1, 25, 7, '宣传.png', '宣传', 1, 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230609/d43a77c266fd59f23b438a7204e80173.png', '', '', 42539, 'image/png', '/dataDB/project/go/gofly_singleresource\\uploads\\20230609\\d43a77c266fd59f23b438a7204e80173.png', '', 'a226b08471c634ebd11b4d32ac138176', 1686296942, 0);
INSERT INTO `common_picture` VALUES (8, 1, 19, 8, 'sw1.jpg', 'sw1', 0, 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230609/c895e724853152e06b5915f046348808.jpg', '', '', 25384, 'image/jpeg', '/dataDB/project/go/gofly_singleresource\\uploads\\20230609\\c895e724853152e06b5915f046348808.jpg', '', '8a81b3c0d0f346d7a36a4573e7196408', 1686296961, 0);

-- ----------------------------
-- Table structure for common_picture_cate
-- ----------------------------
DROP TABLE IF EXISTS `common_picture_cate`;
CREATE TABLE `common_picture_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加账号',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型0=素材图1=插图,2=两种共有',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分类名称' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_picture_cate
-- ----------------------------
INSERT INTO `common_picture_cate` VALUES (1, 1, 1, 0, '商务', 0, '', 1667977623);
INSERT INTO `common_picture_cate` VALUES (2, 1, 2, 2, '科技', 0, '', 1667977649);
INSERT INTO `common_picture_cate` VALUES (3, 1, 3, 0, '教育', 0, '', 1667978004);
INSERT INTO `common_picture_cate` VALUES (4, 1, 4, 0, '风景', 0, '', 1667978020);
INSERT INTO `common_picture_cate` VALUES (5, 1, 5, 0, '建筑', 0, '', 1667978039);
INSERT INTO `common_picture_cate` VALUES (6, 1, 6, 2, '人物', 0, '', 1667978055);
INSERT INTO `common_picture_cate` VALUES (7, 1, 7, 0, '金融', 0, '', 1667978090);
INSERT INTO `common_picture_cate` VALUES (8, 1, 8, 0, '城市', 0, '', 1667978107);
INSERT INTO `common_picture_cate` VALUES (9, 1, 9, 0, '运动', 0, '', 1667978124);
INSERT INTO `common_picture_cate` VALUES (10, 1, 10, 2, '美食', 0, '', 1667978141);
INSERT INTO `common_picture_cate` VALUES (11, 1, 11, 0, '交通', 0, '', 1667978158);
INSERT INTO `common_picture_cate` VALUES (12, 1, 12, 0, '植物', 0, '', 1667978173);
INSERT INTO `common_picture_cate` VALUES (13, 1, 13, 2, '动物', 0, '', 1667978191);
INSERT INTO `common_picture_cate` VALUES (14, 1, 14, 0, '生活', 0, '', 1667978207);
INSERT INTO `common_picture_cate` VALUES (15, 1, 15, 0, '创意', 0, '', 1667978238);
INSERT INTO `common_picture_cate` VALUES (16, 1, 16, 0, '艺术', 0, '', 1667978255);
INSERT INTO `common_picture_cate` VALUES (17, 1, 17, 0, '场景', 0, '', 1667978277);
INSERT INTO `common_picture_cate` VALUES (18, 1, 18, 0, '生产', 0, '', 1667978296);
INSERT INTO `common_picture_cate` VALUES (19, 1, 19, 0, '军事', 0, '', 1667978316);
INSERT INTO `common_picture_cate` VALUES (20, 1, 20, 0, '背景', 0, '', 1667978321);
INSERT INTO `common_picture_cate` VALUES (21, 1, 21, 1, '产品', 0, '', 1667978359);
INSERT INTO `common_picture_cate` VALUES (22, 1, 22, 1, '浮漂', 0, '', 1667978385);
INSERT INTO `common_picture_cate` VALUES (23, 1, 23, 1, '水墨', 0, '', 1667978403);
INSERT INTO `common_picture_cate` VALUES (24, 1, 24, 1, '特效', 0, '', 1667978414);
INSERT INTO `common_picture_cate` VALUES (25, 1, 25, 1, '动物', 0, '', 1667978423);
INSERT INTO `common_picture_cate` VALUES (26, 1, 26, 1, '自然', 0, '', 1667978438);
INSERT INTO `common_picture_cate` VALUES (27, 1, 27, 1, '文字', 0, '', 1667978447);

-- ----------------------------
-- Table structure for common_verify_code
-- ----------------------------
DROP TABLE IF EXISTS `common_verify_code`;
CREATE TABLE `common_verify_code`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `keyname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '存储key',
  `code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '验证码',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '验证码存储' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_verify_code
-- ----------------------------
INSERT INTO `common_verify_code` VALUES (1, 'huang_li_shi@163.com', '380466', 1676913544);

-- ----------------------------
-- Table structure for createcode_product
-- ----------------------------
DROP TABLE IF EXISTS `createcode_product`;
CREATE TABLE `createcode_product`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `num` int(10) NOT NULL DEFAULT 0 COMMENT '库存',
  `price` decimal(10, 2) NOT NULL COMMENT '价格',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试代码产品' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of createcode_product
-- ----------------------------
INSERT INTO `createcode_product` VALUES (1, '测试', 1, 3.00, '你爹', 0);
INSERT INTO `createcode_product` VALUES (2, '对对对22', 0, 0.00, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-4cpevtz8u2w0\"><img src=\"http://localhost:8108/common/uploadfile/get_image?url=resource/uploads/20240120/1705761650.png\" style=\"visibility: visible; width: 352px; height: 352px;\" data-type=\"inline\"></p><p data-id=\"p838747a-PMhdXQU9\"><br></p></div>', 1705652178);
INSERT INTO `createcode_product` VALUES (4, '图片', 12, 0.00, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p838747a-uWGGoSeu\"><div data-id=\"d0124757-qLk9v7rb\" data-card-editable=\"false\" data-syntax=\"plain\" auto-wrap=\"false\"><div class=\"\" style=\"border: 1px solid rgb(232, 232, 232); padding: 8px; background: rgb(249, 249, 249);\"><div style=\"font-family: monospace;font-size: 13px; line-height: 21px; color: #595959; direction: ltr; height: auto; overflow: hidden;background: transparent;\"><pre style=\"color: rgb(89, 89, 89); margin: 0px; padding: 0px; background: none 0% 0% / auto repeat scroll padding-box border-box rgba(0, 0, 0, 0);\"></pre></div></div></div></p><p style=\"text-align:center;\"><img src=\"http://localhost:8108/common/uploadfile/get_image?url=resource/uploads/20240120/407161e2dbdc4cc4848bdc1a1478fde7.jpg\" style=\"width: 272px; visibility: visible; height: 272px;\" data-type=\"block\"></p><p data-id=\"p838747a-BAwmTSAD\">ddddd</p><p data-id=\"p838747a-jpEEYoLa\">ddff</p><p style=\"text-align:center;\"><img src=\"http://localhost:8108/common/uploadfile/get_image?url=resource/uploads/20231221/d397e65e2d50e0dcc2ced152fd8c224e.png\" style=\"visibility: visible; width: 400px; height: 400px;\" data-type=\"block\"></p><p style=\"text-align:center;\"><img src=\"http://localhost:8108/common/uploadfile/get_image?url=resource/uploads/20240120/70180de5ee294a6a78ee88a2be52a796.png\" style=\"visibility: visible; width: 267px; height: 351px;\" data-type=\"block\"></p></div>', 1705756900);

-- ----------------------------
-- Table structure for createcode_product_cate
-- ----------------------------
DROP TABLE IF EXISTS `createcode_product_cate`;
CREATE TABLE `createcode_product_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试代码产品分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of createcode_product_cate
-- ----------------------------

-- ----------------------------
-- Table structure for login_logs
-- ----------------------------
DROP TABLE IF EXISTS `login_logs`;
CREATE TABLE `login_logs`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型1=平台。2=b端，3=C端',
  `uid` int(11) NOT NULL COMMENT '用户id',
  `out_in` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录或退出 out in',
  `loginIP` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录IP',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 976 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '（平台及客户）后台登录日志' ROW_FORMAT = Dynamic;


SET FOREIGN_KEY_CHECKS = 1;
