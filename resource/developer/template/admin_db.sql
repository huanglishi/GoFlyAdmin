/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : gofly_enterprise

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 24/01/2024 15:10:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_account
-- ----------------------------
DROP TABLE IF EXISTS `admin_account`;
CREATE TABLE `admin_account`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户端-用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_account
-- ----------------------------
INSERT INTO `admin_account` VALUES (1, 1, 1, 3, 'gofly', '47e3cee18368271b2dbe9e5a22caef88', '1697472561', '开发管理员', '管家人', 'resource/staticfile/avatar.png', '88422345', '18988274072', '550325@qq.com', '', 1668909071, 0, 0, 1666161776, 0, '对的', '昆明', '开发测试账号', 'GoFLy科技', '', 'chaoyang');
INSERT INTO `admin_account` VALUES (3, 1, 1, 4, 'test', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试账号2', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667142475, 0, '', '', '', '试试', '', '');
INSERT INTO `admin_account` VALUES (4, 1, 1, 1, '123ss', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '销售员de', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667144713, 0, '', '', '', '', '', '');
INSERT INTO `admin_account` VALUES (9, 1, 1, 1, '22334', '166d2832ebcc7672e59d13f37a79f59e', '3305628230121721621', '新增账号', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1678370986, 1678373636, '五华区霖雨路江东耀龙康城27幢二单元502', '昆明市', '', '云律科技（云南）有限公司', '', '');

-- ----------------------------
-- Table structure for admin_auth_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_dept`;
CREATE TABLE `admin_auth_dept`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '上级部门',
  `weigh` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理后台部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_dept
-- ----------------------------
INSERT INTO `admin_auth_dept` VALUES (1, 1, '市场部门', 0, 1, 0, '营销', 1666972562);
INSERT INTO `admin_auth_dept` VALUES (2, 1, '第一组', 1, 2, 1, '', 1660493279);
INSERT INTO `admin_auth_dept` VALUES (3, 1, '研发部门', 1, 3, 0, '', 1660493302);
INSERT INTO `admin_auth_dept` VALUES (4, 1, '领导部门', 0, 4, 1, '', 1660493325);
INSERT INTO `admin_auth_dept` VALUES (6, 1, '人事组', 1, 6, 0, '', 1667827895);

-- ----------------------------
-- Table structure for admin_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_role`;
CREATE TABLE `admin_auth_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
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
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_role
-- ----------------------------
INSERT INTO `admin_auth_role` VALUES (1, 1, 0, '超级管理组', '*', '*', 0, 0, '账号的总管理员', 1, 1666336830);
INSERT INTO `admin_auth_role` VALUES (5, 1, 1, '销售员', '8,9,10,4,7,11,13,14,32,33,34', '[8,9,10,4,7,11,13,14,32,33,34]', 0, 0, '产品销售组', 2, 1667101829);
INSERT INTO `admin_auth_role` VALUES (6, 1, 1, '管理员', '7,11,13,32,8,64,61,12,63,6', '[7,11,13,32,8,64,61,12,63]', 0, 0, '', 3, 1678293133);
INSERT INTO `admin_auth_role` VALUES (7, 1, 6, '编辑组', '7,34,33,11,12,4', '[7,34,33,11,12]', 0, 0, '', 4, 1660725985);
INSERT INTO `admin_auth_role` VALUES (8, 1, 6, '兼职组', '11,12,34,7,33', '[11,12,34,7,33]', 0, 0, 'ceshi', 5, 1667105411);
INSERT INTO `admin_auth_role` VALUES (11, 1, 0, '管理组', '8,9,10', '[8,9,10]', 0, 0, '', 11, 1660904496);
INSERT INTO `admin_auth_role` VALUES (13, 1, 0, '市场部门', '8', '[8]', 0, 0, '', 13, 1667117642);
INSERT INTO `admin_auth_role` VALUES (16, 1, 0, '财务室', '8,48,49,59,69,6', '[8,48,49,59,69]', 0, 0, '修改', 16, 1678292260);

-- ----------------------------
-- Table structure for admin_auth_role_access
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_role_access`;
CREATE TABLE `admin_auth_role_access`  (
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `role_id` int(11) NOT NULL DEFAULT 0 COMMENT '授权id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'admin端菜单权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_role_access
-- ----------------------------
INSERT INTO `admin_auth_role_access` VALUES (1, 1);
INSERT INTO `admin_auth_role_access` VALUES (3, 5);
INSERT INTO `admin_auth_role_access` VALUES (4, 1);
INSERT INTO `admin_auth_role_access` VALUES (5, 6);
INSERT INTO `admin_auth_role_access` VALUES (9, 6);
INSERT INTO `admin_auth_role_access` VALUES (9, 5);

-- ----------------------------
-- Table structure for admin_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_rule`;
CREATE TABLE `admin_auth_rule`  (
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
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'C端-菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_rule
-- ----------------------------
INSERT INTO `admin_auth_rule` VALUES (8, 1, '概况', '', 1, 1, 0, 'icon-dashboard', '/home', 'home', '/dashboard/workplace/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1666167477);
INSERT INTO `admin_auth_rule` VALUES (11, 1, '角色管理', '', 12, 1, 61, '', 'role', 'role', '/system/role/index', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 1666336763);
INSERT INTO `admin_auth_rule` VALUES (12, 1, '菜单管理', '', 11, 1, 61, '', 'rule', 'rule', '/system/rule/index', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 1657817329);
INSERT INTO `admin_auth_rule` VALUES (13, 1, '部门管理', '', 13, 1, 61, '', 'dept', 'dept', '/system/dept/index', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 1660818242);
INSERT INTO `admin_auth_rule` VALUES (32, 1, '详情', '', 32, 1, 61, '', 'account_detail/:id', 'AccountDetail', '/system/account/AccountDetail', '', '', 0, 0, 1, 0, 1, 1, 0, 1, 1660635610);
INSERT INTO `admin_auth_rule` VALUES (48, 1, '业务端管理', '', 2, 0, 0, 'icon-book', '/business', 'business', 'LAYOUT', '/business/bizuser', '', 0, 0, 0, 0, 0, 0, 0, 0, 1666167571);
INSERT INTO `admin_auth_rule` VALUES (49, 1, '账号管理', '', 49, 1, 48, '', 'bizuser', 'bizuser', '/business/bizuser/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1666060273);
INSERT INTO `admin_auth_rule` VALUES (59, 1, '业务角色', '', 59, 1, 48, '', 'bizrole', 'bizrole', '/business/bizrole/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1666165594);
INSERT INTO `admin_auth_rule` VALUES (61, 1, '系统设置', '', 8, 0, 0, 'icon-english-fill', '/system', 'system', 'LAYOUT', '/system/account', '', 0, 0, 0, 0, 0, 0, 0, 0, 1667145000);
INSERT INTO `admin_auth_rule` VALUES (63, 1, '账户管理', '', 63, 1, 61, '', 'account', 'account', '/system/account/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1666167258);
INSERT INTO `admin_auth_rule` VALUES (64, 1, '添加账号', '', 64, 2, 7, '', '', '', '', '', 'add', 0, 0, 0, 0, 0, 0, 0, 0, 1667142600);
INSERT INTO `admin_auth_rule` VALUES (68, 1, '个人中心', '', 79, 0, 0, 'icon-user', '/user', 'user', 'LAYOUT', '/user/info', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678027122);
INSERT INTO `admin_auth_rule` VALUES (69, 1, '账号信息', '', 69, 1, 68, '', 'info', 'info', '/user/info/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678027200);
INSERT INTO `admin_auth_rule` VALUES (70, 1, '用户设置', '', 70, 1, 68, '', 'setting', 'setting', '/user/setting/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678027237);
INSERT INTO `admin_auth_rule` VALUES (73, 1, '菜单管理', '', 73, 1, 48, '', 'bizrule', 'bizrule', '/business/bizrule/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1678504239);
INSERT INTO `admin_auth_rule` VALUES (74, 1, '数据中心', '', 74, 0, 0, 'icon-storage', '/datacenter', 'datacenter', 'LAYOUT', '/datacenter/logininfo', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686136637);
INSERT INTO `admin_auth_rule` VALUES (75, 1, '登录页轮播', '', 76, 1, 74, '', 'logininfo', 'logininfo', 'datacenter/logininfo/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686136825);
INSERT INTO `admin_auth_rule` VALUES (76, 1, '字典数据', '', 75, 1, 74, '', 'dictionary', 'dictionary', 'datacenter/dictionary/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686229379);
INSERT INTO `admin_auth_rule` VALUES (77, 1, '系统附件', '', 77, 1, 78, '', 'attachment', 'attachment', '/matter/attachment/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686235517);
INSERT INTO `admin_auth_rule` VALUES (78, 1, '素材管理', '', 78, 0, 0, 'icon-folder', '/matter', 'matter', 'LAYOUT', '/matter/attachment', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686240839);
INSERT INTO `admin_auth_rule` VALUES (79, 1, '公共图片库', '', 79, 1, 78, '', 'picture', 'picture', '/matter/picture/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686241299);
INSERT INTO `admin_auth_rule` VALUES (80, 1, '配置管理', '', 80, 1, 74, '', 'configuration', 'configuration', '/datacenter/configuration/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1690691471);

SET FOREIGN_KEY_CHECKS = 1;
