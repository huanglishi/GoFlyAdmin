/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : gofly_single

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 09/09/2023 13:37:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
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
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 1, 1, 3, 'gofly', '131ffac800502aee306d42a3c83ff6c4', '1693074293', '开发管理员', '管家人', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230310/b4ac2e2246073c50c9dc764d5b426720.png', '88422345', '18988274072', '550325@qq.com', '', 1668909071, 0, 0, 1666161776, 0, '对的', '昆明', '开发测试账号', 'GoFLy科技', '', 'chaoyang');
INSERT INTO `admin_user` VALUES (3, 1, 1, 4, 'test', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试账号2', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667142475, 0, '', '', '', '试试', '', '');
INSERT INTO `admin_user` VALUES (4, 1, 1, 1, '123ss', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '销售员de', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667144713, 0, '', '', '', '', '', '');
INSERT INTO `admin_user` VALUES (9, 1, 1, 1, '22334', '166d2832ebcc7672e59d13f37a79f59e', '3305628230121721621', '新增账号', '', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230309/162ba4b5924cc0fe399d7a2ffd1d1110.png', '', '', '', '', 0, 0, 0, 1678370986, 1678373636, '五华区霖雨路江东耀龙康城27幢二单元502', '昆明市', '', '云律科技（云南）有限公司', '', '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 733 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '附件管理' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 106 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户端附件' ROW_FORMAT = Dynamic;

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
INSERT INTO `business_attachment` VALUES (99, 99, 1, 1, 0, '微信截图_20230731223049.png', '微信截图_20230731223049', 0, 'resource/uploads/20230811/a67d2c864ad5b33dd628bcd5643c54a4.png', '', '', 182240, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/a67d2c864ad5b33dd628bcd5643c54a4.png', '', 'a22cb0313ce2570434030398b6f4535b', 0, 1691759331);
INSERT INTO `business_attachment` VALUES (100, 100, 1, 1, 0, '微信截图_20230731223220.png', '微信截图_20230731223220', 0, 'resource/uploads/20230811/672ab21428230ea9d664cfb74a4b2095.png', '', '', 205064, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/672ab21428230ea9d664cfb74a4b2095.png', '', 'b4130cce914ac7091007e219e057f792', 0, 1691759331);
INSERT INTO `business_attachment` VALUES (101, 101, 1, 1, 0, '微信截图_20230802001743.png', '微信截图_20230802001743', 0, 'resource/uploads/20230811/a8cd963dc01e7fa40256544cb7276c9e.png', '', '', 560460, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/a8cd963dc01e7fa40256544cb7276c9e.png', '', '54e6b4860ddb71e1336105e82915e1c8', 0, 1691759333);
INSERT INTO `business_attachment` VALUES (102, 102, 1, 1, 0, '微信截图_20230801235142.png', '微信截图_20230801235142', 0, 'resource/uploads/20230811/4bfae33e4192e0d313d64219f466e533.png', '', '', 545847, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/4bfae33e4192e0d313d64219f466e533.png', '', '4e77bd0e1b59bb7c038f0af6da222779', 0, 1691759333);
INSERT INTO `business_attachment` VALUES (103, 103, 1, 1, 0, '微信截图_20230802001834 - 副本.png', '微信截图_20230802001834 - 副本', 0, 'resource/uploads/20230811/0bac4e160efa826f4d2637b999c614de.png', '', '', 686638, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/0bac4e160efa826f4d2637b999c614de.png', '', '8090e3b8227dde0a31e6cb352b62570f', 0, 1691759333);
INSERT INTO `business_attachment` VALUES (104, 104, 1, 1, 0, '微信截图_20230802001432.png', '微信截图_20230802001432', 0, 'resource/uploads/20230811/44764a0ff44456b9d44c1c0a8f6d44d3.png', '', '', 784369, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/44764a0ff44456b9d44c1c0a8f6d44d3.png', '', 'e8e5b00efb806430a84b1cfe2e2bdc48', 0, 1691759333);
INSERT INTO `business_attachment` VALUES (105, 105, 1, 1, 0, '微信截图_20230801234815.png', '微信截图_20230801234815', 0, 'resource/uploads/20230811/f82c9410d28d5e7ea8efc992917bbb4e.png', '', '', 1271571, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/f82c9410d28d5e7ea8efc992917bbb4e.png', '', '6a6213d840281da94340c9cd93352dc2', 0, 1691759334);
INSERT INTO `business_attachment` VALUES (106, 106, 1, 1, 0, 'GoFLy发布文章封面.png', 'GoFLy发布文章封面', 0, 'resource/uploads/20230811/3bb630bd364c652d336e4776f6a7dd21.png', '', '', 40902, 'image/png', '', '/dataDB/project/go/gofly_single/resource/uploads/20230811/3bb630bd364c652d336e4776f6a7dd21.png', '', 'b98da546d168f3e1d91d32585aaf719e', 0, 1691759379);

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
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理后台部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_dept
-- ----------------------------
INSERT INTO `business_auth_dept` VALUES (1, 1, 1, '市场部门', 0, 1, 0, '营销', 1666972562);
INSERT INTO `business_auth_dept` VALUES (2, 1, 1, '第一组', 1, 2, 0, '', 1660493279);
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
INSERT INTO `business_auth_role` VALUES (19, 1, 1, 1, '新增权限', '8,6', '[8]', 0, 0, '', 19, 1678596680);
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
INSERT INTO `business_auth_role_access` VALUES (1, 1);
INSERT INTO `business_auth_role_access` VALUES (3, 5);
INSERT INTO `business_auth_role_access` VALUES (10, 5);
INSERT INTO `business_auth_role_access` VALUES (11, 1);
INSERT INTO `business_auth_role_access` VALUES (12, 1);
INSERT INTO `business_auth_role_access` VALUES (13, 1);

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
) ENGINE = InnoDB AUTO_INCREMENT = 145 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'C端-菜单' ROW_FORMAT = Dynamic;

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
INSERT INTO `business_auth_rule` VALUES (78, 1, '微信管理', '', 68, 0, 0, 'svgfont-weixinshezhi', '/wxsys', 'wxsys', 'LAYOUT', '/wxsys/wxuser', '', 0, 0, 0, 1, 0, 0, 0, 0, 1680788675);
INSERT INTO `business_auth_rule` VALUES (79, 1, '账号配置', '', 87, 1, 78, '', 'wxsetting', 'wxsetting', '/wxsys/wxsetting/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1680788805);
INSERT INTO `business_auth_rule` VALUES (84, 1, '公众号菜单', '', 84, 1, 78, '', 'wxmenu', 'wxmenu', '/wxsys/wxmenu/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1680790715);
INSERT INTO `business_auth_rule` VALUES (87, 1, '微信用户', '', 79, 1, 78, '', 'wxuser', 'wxuser', '/wxsys/wxuser/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1680791150);
INSERT INTO `business_auth_rule` VALUES (97, 1, '生成代码', '', 97, 1, 74, '', 'generatecode', 'generatecode', '/developer/generatecode/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1685366576);
INSERT INTO `business_auth_rule` VALUES (110, 1, '生成代码示例', '', 110, 0, 0, 'icon-english-fill', '/createcode', 'createcode', 'LAYOUT', '/createcode/test/index', '', 0, 0, 0, 1, 0, 0, 0, 0, 1685550835);
INSERT INTO `business_auth_rule` VALUES (120, 1, '测试api代码生成', '', 120, 1, 110, '', 'api', 'api', 'createcode/api/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686115348);
INSERT INTO `business_auth_rule` VALUES (121, 1, '数据中心', '', 79, 0, 0, 'icon-storage', '/datacenter', 'datacenter', 'LAYOUT', '/datacenter/dictionary', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686132257);
INSERT INTO `business_auth_rule` VALUES (123, 1, '字典数据', '', 123, 1, 121, '', 'data', 'data', '/datacenter/dictionary/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686132869);
INSERT INTO `business_auth_rule` VALUES (137, 1, '附件管理', '', 137, 1, 121, '', 'attachment', 'attachment', 'datacenter/attachment/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1686286572);
INSERT INTO `business_auth_rule` VALUES (141, 1, '测试代码生成', '', 141, 1, 110, '', 'code', 'code', 'createcode/code/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1687794349);
INSERT INTO `business_auth_rule` VALUES (143, 1, '配置管理', '', 143, 1, 121, '', 'configuration', 'configuration', '/datacenter/configuration/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1690646744);
INSERT INTO `business_auth_rule` VALUES (144, 1, '测试品牌', '', 144, 1, 110, '', 'band', 'band', 'createcode/band/index', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 1692896109);

-- ----------------------------
-- Table structure for business_createcode_api
-- ----------------------------
DROP TABLE IF EXISTS `business_createcode_api`;
CREATE TABLE `business_createcode_api`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名',
  `nickename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `image` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片',
  `file` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附件',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试api代码生成' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_createcode_api
-- ----------------------------
INSERT INTO `business_createcode_api` VALUES (2, 0, 1, 0, '张三', '第三', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230811/672ab21428230ea9d664cfb74a4b2095.png', '', 2, '测试', 1685722937, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-gksdbjfivbs0\"><br></p></div>');
INSERT INTO `business_createcode_api` VALUES (5, 0, 1, 0, '张8', '第三', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230811/0bac4e160efa826f4d2637b999c614de.png', '', 5, '测试', 1685726627, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-1gu952f48yww0\"><br></p></div>');
INSERT INTO `business_createcode_api` VALUES (6, 0, 1, 0, '李四', '李师傅', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230811/a67d2c864ad5b33dd628bcd5643c54a4.png', '', 6, '测2', 1685760124, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-1hw7ixeuxu5c0\"><br></p></div>');
INSERT INTO `business_createcode_api` VALUES (8, 0, 1, 0, '刚睡醒', '是撒', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230608/e3ce4dc222fff37ec130ea7daff3b474.png', '', 8, '测4', 1685808707, '');
INSERT INTO `business_createcode_api` VALUES (9, 0, 1, 0, 'www', 'ww', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230608/eaf1511fa669c7dd54af301d50c9478e.png', '', 9, 'ddd', 1686219848, '');
INSERT INTO `business_createcode_api` VALUES (14, 0, 1, 0, '张三', '第三', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230811/f82c9410d28d5e7ea8efc992917bbb4e.png', '', 14, '测试', 1690645019, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-b2jazhoe0kg0\"><br></p></div>');
INSERT INTO `business_createcode_api` VALUES (15, 0, 1, 0, 'dd', 'ffd', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230811/0bac4e160efa826f4d2637b999c614de.png', '', 15, 'dd', 1692894314, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-wshdjnfmbo00\"><br></p></div>');

-- ----------------------------
-- Table structure for business_createcode_band
-- ----------------------------
DROP TABLE IF EXISTS `business_createcode_band`;
CREATE TABLE `business_createcode_band`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `des` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态:0=启用,1=禁用',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试品牌' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_createcode_band
-- ----------------------------
INSERT INTO `business_createcode_band` VALUES (1, '测试', '测试备注', 0, 0, '');
INSERT INTO `business_createcode_band` VALUES (3, '大众', '德国品牌', 0, 1, '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-3lj9sbsgnle0\"><br></p></div>');

-- ----------------------------
-- Table structure for business_createcode_cate
-- ----------------------------
DROP TABLE IF EXISTS `business_createcode_cate`;
CREATE TABLE `business_createcode_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试代码生成分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_createcode_cate
-- ----------------------------
INSERT INTO `business_createcode_cate` VALUES (1, 0, 1, 0, '分类1', 1, '等等', 1685635749);
INSERT INTO `business_createcode_cate` VALUES (2, 0, 1, 0, '大类', 2, '测试', 1686227141);
INSERT INTO `business_createcode_cate` VALUES (3, 0, 1, 0, '单服', 3, '等等', 1686228675);

-- ----------------------------
-- Table structure for business_createcode_code
-- ----------------------------
DROP TABLE IF EXISTS `business_createcode_code`;
CREATE TABLE `business_createcode_code`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `cid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类',
  `age` int(11) NOT NULL COMMENT '年龄',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `image` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试代码生成' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_createcode_code
-- ----------------------------
INSERT INTO `business_createcode_code` VALUES (1, 0, 0, 0, 0, '测试数据', '', '');
INSERT INTO `business_createcode_code` VALUES (4, 0, 1, 3, 25, '张三', '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p1dr9j7ls-1ifj69ez4uqo0\">大中华</p></div>', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230811/f82c9410d28d5e7ea8efc992917bbb4e.png');

-- ----------------------------
-- Table structure for business_email
-- ----------------------------
DROP TABLE IF EXISTS `business_email`;
CREATE TABLE `business_email`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `sender_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '发送者邮箱',
  `auth_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱授权码',
  `mail_title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件标题',
  `mail_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件内容,可以是html',
  `service_host` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件服务器',
  `service_port` int(11) NOT NULL DEFAULT 0 COMMENT '邮件服务器端口',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '业务端邮箱' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_email
-- ----------------------------
INSERT INTO `business_email` VALUES (1, 0, 1, '504500934@qq.com', 'amidmyjnnxyvbgfb', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);

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
-- Table structure for business_user
-- ----------------------------
DROP TABLE IF EXISTS `business_user`;
CREATE TABLE `business_user`  (
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
-- Records of business_user
-- ----------------------------
INSERT INTO `business_user` VALUES (1, 1, 1, 1, 3, 'gofly', '131ffac800502aee306d42a3c83ff6c4', '1693074293', '开发管理员', '黄师傅', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230506/d4bb8324ee87699ed727bc1fd0479b34.jpg', '88422345', '18988347563', '550325@qq.com', '', 1668909071, 0, 0, 1666161776, 1678544449, '王府井', '昆明', '开发测试账号', 'GoFLy科技1', '', 'chaoyang', 2147483647);
INSERT INTO `business_user` VALUES (3, 1, 1, 3, 4, 'test', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试账号biz', '', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667142475, 1678550309, '', '', '', '试试', '', '', 2147483647);
INSERT INTO `business_user` VALUES (4, 1, 1, 4, 1, '123ss', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '销售员de', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1667144713, 0, '', '', '', '', '', '', 2147483647);
INSERT INTO `business_user` VALUES (9, 1, 1, 9, 1, '22334', '166d2832ebcc7672e59d13f37a79f59e', '3305628230121721621', '新增账号', '', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230309/162ba4b5924cc0fe399d7a2ffd1d1110.png', '', '', '', '', 0, 0, 0, 1678370986, 1678373636, '五华区霖雨路江东耀龙康城27幢二单元502', '昆明市', '', '云律科技（云南）有限公司', '', '', 2147483647);
INSERT INTO `business_user` VALUES (10, 1, 1, 10, 0, 'tssss', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试22', '', 'resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1678585983, 0, '', '', '', '', '', '', 2147483647);
INSERT INTO `business_user` VALUES (12, 1, 1, 1, 1, 'c1', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试1号', 'ww是', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1678976464, 1678976476, '五华区霖雨路江东耀龙康城27幢二单元502', '昆明市', '', '云律科技（云南）有限公司', '', '', 2147483647);
INSERT INTO `business_user` VALUES (13, 1, 1, 13, 0, 'jinpopo', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', 'jinpopo', 'jinpopo', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/staticfile/avatar.png', '', '', '', '', 0, 0, 0, 1690209684, 1690209722, '', '', '', '', '', '', 3787456512);

-- ----------------------------
-- Table structure for business_wxsys_officonfig
-- ----------------------------
DROP TABLE IF EXISTS `business_wxsys_officonfig`;
CREATE TABLE `business_wxsys_officonfig`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公众号名称',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
  `AppID` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '开发者ID(AppID)',
  `AppSecret` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '开发者密码(AppSecret)',
  `expires_access_token` int(11) NOT NULL DEFAULT 0 COMMENT '获取access_token时间',
  `expires_jsapi_ticket` int(11) NOT NULL DEFAULT 0 COMMENT '获取jsapi_ticket时间',
  `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'access_token',
  `jsapi_ticket` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'jsapi_ticket',
  `qrcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '二维码',
  `Token` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Token 长度为3-32字符',
  `EncodingAESKey` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '消息加密密钥由43位字符组成',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信公众号配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_wxsys_officonfig
-- ----------------------------
INSERT INTO `business_wxsys_officonfig` VALUES (4, 19, 0, '云律法务咨询', '云律科技', 'wx73fd73f2ba85dd5d', '3bd86536b5dfa0fc2058cce6e9e4b485', 1680945724, 1678878997, '67_D0hSc0ohwRWkRb68_DLRp1zCESdpRG1xAFhTQPVTObZGh16hjwZ7lZwOjPZzd6TsyVEEJoFbhKx6FRvoZDqQN3vHYRXCFsB6tFy_X45sxBqKG0PqC5zkZsnkaqUGNCeAEAQXM', 'O3SMpm8bG7kJnF36aXbe8-dOkbHtC1tZ-OT1t0-u3u5gk4d9wHinoVzgAtE8ftrBddRNAVPLdfHIV71PemXSZQ', '', '', '', 1672367636);
INSERT INTO `business_wxsys_officonfig` VALUES (6, 1, 1, '测试号管理', '微信号： gh_6d5eb37e43d8', 'wx3c20a30ae0ab44b5', 'e4ad51399c9e4db1d99aa2a1fe5e2af1', 1694236702, 0, '72_m9OOpSaKX5c3OAWn9c4CPyfXPwyy2tuRFti3Fyi2lLRzCCBlC3IzbmYy5WrMzLHyXzF6Ft_iWDZcaXYUZz2UR_od1iX_GdwXTnKwaWSEwo2rTWj2nFR4BQeRyJATYXaAIAMRQ', '', 'https://sg.goflys.cn/common/uploadfile/get_image?url=resource/uploads/20230507/7992812e7e9f2b140968ba3874de1d1a.jpg', '19ibcTXUf', 'WPhI4Izo8aLtcvO9EjYSfA7LolcEyqPCKiqWGM44xrS', 1688536087);

-- ----------------------------
-- Table structure for business_wxsys_user
-- ----------------------------
DROP TABLE IF EXISTS `business_wxsys_user`;
CREATE TABLE `business_wxsys_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `dept_id` int(11) NOT NULL DEFAULT 0 COMMENT '属于部门',
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `password` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码盐',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '姓名',
  `openid` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户的标识，对当前公众号唯一',
  `wxapp_openid` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '小程序的openid',
  `unionid` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。',
  `subscribe` tinyint(1) NOT NULL COMMENT '用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。',
  `subscribe_time` int(11) NOT NULL DEFAULT 0 COMMENT '用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(145) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `mobile` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '邮箱',
  `city` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '城市',
  `area` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户所在县区',
  `address` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户所在详细地址',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态 1=禁用',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信关注用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_wxsys_user
-- ----------------------------

-- ----------------------------
-- Table structure for business_wxsys_wxappconfig
-- ----------------------------
DROP TABLE IF EXISTS `business_wxsys_wxappconfig`;
CREATE TABLE `business_wxsys_wxappconfig`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公众号名称',
  `AppID` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '开发者ID(AppID)',
  `AppSecret` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '开发者密码(AppSecret)',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
  `expires_access_token` int(11) NOT NULL DEFAULT 0 COMMENT '获取access_token时间',
  `expires_jsapi_ticket` int(11) NOT NULL DEFAULT 0 COMMENT '获取jsapi_ticket时间',
  `qrcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '二维码',
  `Token` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Token 长度为3-32字符',
  `EncodingAESKey` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '消息加密密钥由43位字符组成',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信小程序配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_wxsys_wxappconfig
-- ----------------------------
INSERT INTO `business_wxsys_wxappconfig` VALUES (4, 19, 0, '云律法务咨询', 'wx73fd73f2ba85dd5d', '3bd86536b5dfa0fc2058cce6e9e4b485', '云律科技', 1680945724, 1678878997, '', '', '', 1672367636);
INSERT INTO `business_wxsys_wxappconfig` VALUES (6, 1, 1, '小程序测试号', 'wxe91102f7f7dead52', '6371e1f0ff2d9ed840ca8ff2809cd838', 'wxid_jdduiwcre0iv22的接口测试号', 0, 0, '', '', '', 1681195635);

-- ----------------------------
-- Table structure for business_wxsys_wxmenu
-- ----------------------------
DROP TABLE IF EXISTS `business_wxsys_wxmenu`;
CREATE TABLE `business_wxsys_wxmenu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountID` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `businessID` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公众号名称',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
  `menu` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单内容',
  `select` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否选择用1=是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微站微信菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_wxsys_wxmenu
-- ----------------------------
INSERT INTO `business_wxsys_wxmenu` VALUES (9, 1, 1, '测试号管理', 'sssdsss', '{\"button\":[{\"name\":\"菜单名称\",\"type\":\"view\",\"url\":\"\",\"sub_button\":{\"list\":[{\"name\":\"子菜单名称\",\"url\":\"\",\"type\":\"view\"}]}}]}', 0);
INSERT INTO `business_wxsys_wxmenu` VALUES (11, 1, 1, '测试号管理', '', '{\"button\":[{\"type\":\"view\",\"url\":\"\",\"sub_button\":{\"list\":[{\"type\":\"view\",\"name\":\"子菜单名称\",\"url\":\"\"}]},\"name\":\"菜单名称\"}]}', 0);
INSERT INTO `business_wxsys_wxmenu` VALUES (13, 1, 1, '公众号菜单1', '', '{\"button\":[{\"type\":\"view\",\"url\":\"\",\"sub_button\":{\"list\":[{\"type\":\"view\",\"name\":\"子菜单名称\",\"url\":\"\"}]},\"name\":\"菜单名称\"},{\"type\":\"view\",\"name\":\"菜单名称\",\"url\":\"\"},{\"type\":\"view\",\"name\":\"菜单名称\",\"url\":\"\"}]}', 0);

-- ----------------------------
-- Table structure for common_apitest
-- ----------------------------
DROP TABLE IF EXISTS `common_apitest`;
CREATE TABLE `common_apitest`  (
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
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口测试数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apitest
-- ----------------------------
INSERT INTO `common_apitest` VALUES (1, 'biz', 2, '登录接口-获取openid', '1请求参数：\ncode： wx.login()中的code\nappid：wx.getAccountInfoSync().miniProgram.appId中取\n2返回参数：\nuserinfo:用户信息\ntoken：用户token，统一在请求中封装返回', '/user/get_openid', '{\n\"code\":\"\",\n\"appid\":\"\"\n}', 'get', 'business_wxsys_user', 0, 0, 1683386592);
INSERT INTO `common_apitest` VALUES (5, 'biz', 3, '后端测试接口', '业务端测试端，请求接口', '/test/api/get_data', '', 'get', '', 0, 1, 1683380403);
INSERT INTO `common_apitest` VALUES (6, 'biz', 3, '测试获取列表数据接口', '', '/test/api/get_list', '', 'get', 'business_auth_rule', 0, 1, 1683380802);
INSERT INTO `common_apitest` VALUES (7, 'biz', 2, '获取小程序数据', '', '/test/wxapi/get_data', '', 'get', '', 0, 1, 1683386443);

-- ----------------------------
-- Table structure for common_apitest_group
-- ----------------------------
DROP TABLE IF EXISTS `common_apitest_group`;
CREATE TABLE `common_apitest_group`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'admin' COMMENT '分类接口属于那端，admin=管理，biz=B端，client=C端',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父级',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `type_id` int(11) NOT NULL DEFAULT 0 COMMENT '接口类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台端接口测试分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apitest_group
-- ----------------------------
INSERT INTO `common_apitest_group` VALUES (1, 'biz', 0, 'app端', 0, 3);
INSERT INTO `common_apitest_group` VALUES (2, 'biz', 0, '小程序', 0, 1);
INSERT INTO `common_apitest_group` VALUES (3, 'biz', 0, '后台管理', 0, 2);

-- ----------------------------
-- Table structure for common_apitest_type
-- ----------------------------
DROP TABLE IF EXISTS `common_apitest_type`;
CREATE TABLE `common_apitest_type`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型名称',
  `rooturl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求服务器地址',
  `verifyEncrypt` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密验证字符串',
  `isself` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是本端1=是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口类型' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apitest_type
-- ----------------------------
INSERT INTO `common_apitest_type` VALUES (1, '小程序', 'https://sg.goflys.cn/wxapp', 'gofly@888', 0);
INSERT INTO `common_apitest_type` VALUES (2, '本端', '', '', 1);
INSERT INTO `common_apitest_type` VALUES (3, '手机APP', 'https://sg.goflys.cn/mbapp', 'gofly@888', 0);

-- ----------------------------
-- Table structure for common_apitext
-- ----------------------------
DROP TABLE IF EXISTS `common_apitext`;
CREATE TABLE `common_apitext`  (
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
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口测试数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apitext
-- ----------------------------
INSERT INTO `common_apitext` VALUES (1, 'biz', 2, '登录接口-获取openid', '1请求参数：\ncode： wx.login()中的code\n2返回参数：\nuserinfo:用户信息\ntoken：用户token，统一在请求中封装返回', '/user/get_openid', '{\n\"code\": \"\"\n}', 'get', 'business_wxsys_user', 0, 0, 0, 0, '', '', 1690649631);
INSERT INTO `common_apitext` VALUES (5, 'biz', 3, '后端测试接口', '业务端测试端，请求接口', '/test/api/get_data', '', 'get', '', 0, 1, 0, 0, '', '', 1683380403);
INSERT INTO `common_apitext` VALUES (6, 'biz', 3, '测试获取列表数据接口', '', '/test/api/get_list', '', 'get', 'business_auth_rule', 0, 1, 0, 0, '', '', 1683380802);
INSERT INTO `common_apitext` VALUES (7, 'biz', 2, '获取小程序数据', '', '/test/wxapi/get_data', '', 'get', '', 0, 0, 0, 0, '', '', 1685722258);
INSERT INTO `common_apitext` VALUES (8, 'biz', 2, '测试代码生成接口', '', '/testcode/codes/get_list', '{\n\"cid\":1\n}', 'get', 'business_createcode_api', 0, 0, 1, 0, 'des,id,name,nickename,status,weigh', 'list', 1685726288);
INSERT INTO `common_apitext` VALUES (9, 'biz', 2, '测试代码生成接口', '', '/testcode/codes/save', '{\n\"name\":\"张三\",\n\"nickename\":\"第三\",\n\"des\":\"测试\"\n\n}', 'post', 'business_createcode_api', 0, 1, 1, 0, '', 'list', 1685726614);
INSERT INTO `common_apitext` VALUES (10, 'biz', 2, '删除', '', '/testcode/codes/del', '{\n\"ids\":[1]\n}', 'delete', '', 0, 0, 1, 0, '', 'list', 1685726578);
INSERT INTO `common_apitext` VALUES (11, 'biz', 2, '测试生成获取详情', '', '/testcode/codes/get_detail', '{\n\"id\":2\n}', 'get', 'business_createcode_api', 0, 0, 1, 0, 'accountID,businessID,id,des', 'detail', 1685726269);
INSERT INTO `common_apitext` VALUES (12, 'biz', 2, '通用_list', '', '/common/api/get_list', '', 'get', 'business_createcode_api', 0, 0, 2, 0, 'name,id,nickename,des,status,createtime', 'list', 1685760228);
INSERT INTO `common_apitext` VALUES (14, 'biz', 2, '通用接口_detail', '', '/common/api/get_detail', '{\n\"id\":5\n}', 'get', 'business_createcode_api', 0, 0, 2, 0, 'accountID,createtime,businessID,des,id,name,nickename,status,weigh', 'detail', 1685757846);
INSERT INTO `common_apitext` VALUES (15, 'biz', 2, '通用接口_添加数据save', '', '/common/api/save', '', 'post', 'business_createcode_api', 0, 1, 2, 0, '', 'list', 1690209824);
INSERT INTO `common_apitext` VALUES (17, 'biz', 1, '通用接口_删除del', '', '/common/api/del', '', 'delete', 'business_createcode_api', 0, 0, 2, 0, '', 'list', 1685882013);
INSERT INTO `common_apitext` VALUES (18, 'biz', 2, '生成api代码', '', '/createcode/api/get_list', '', 'get', 'business_createcode_api', 0, 0, 1, 0, 'accountID,businessID,createtime,des,id,name,weigh,status', 'list', 1685965785);

-- ----------------------------
-- Table structure for common_apitext_group
-- ----------------------------
DROP TABLE IF EXISTS `common_apitext_group`;
CREATE TABLE `common_apitext_group`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'admin' COMMENT '分类接口属于那端，admin=管理，biz=B端，client=C端',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父级',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `type_id` int(11) NOT NULL DEFAULT 0 COMMENT '接口类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台端接口测试分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apitext_group
-- ----------------------------
INSERT INTO `common_apitext_group` VALUES (1, 'biz', 0, 'app端', 0, 3);
INSERT INTO `common_apitext_group` VALUES (2, 'biz', 0, '小程序', 0, 1);
INSERT INTO `common_apitext_group` VALUES (3, 'biz', 0, '后台管理', 0, 2);

-- ----------------------------
-- Table structure for common_apitext_type
-- ----------------------------
DROP TABLE IF EXISTS `common_apitext_type`;
CREATE TABLE `common_apitext_type`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型名称',
  `rooturl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求服务器地址',
  `verifyEncrypt` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密验证字符串',
  `isself` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是本端1=是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口类型' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_apitext_type
-- ----------------------------
INSERT INTO `common_apitext_type` VALUES (1, '小程序', 'https://sg.goflys.cn/wxapp', 'gofly@888', 0);
INSERT INTO `common_apitext_type` VALUES (2, '本端', '', '', 1);
INSERT INTO `common_apitext_type` VALUES (3, '手机APP', 'https://sg.goflys.cn/mbapp', 'gofly@888', 0);

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
INSERT INTO `common_config` VALUES (2, 'common', 0, 'rooturl', 'https://sg.goflys.cn/common/uploadfile/get_image?url=', '图片路径', 0);

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据-测试数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_dictionary_data
-- ----------------------------
INSERT INTO `common_dictionary_data` VALUES (1, 'common', '管理层', 'mteam', '公司领导', 0, 1, 1686156976, 0);
INSERT INTO `common_dictionary_data` VALUES (2, 'common', '业务员', 'salesman', '', 0, 2, 1691760155, 0);

-- ----------------------------
-- Table structure for common_dictionary_integral
-- ----------------------------
DROP TABLE IF EXISTS `common_dictionary_integral`;
CREATE TABLE `common_dictionary_integral`  (
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
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '积分等级-测试数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_dictionary_integral
-- ----------------------------
INSERT INTO `common_dictionary_integral` VALUES (1, 'common', '普通', '12', '', 0, 1, 1686157762, 0);
INSERT INTO `common_dictionary_integral` VALUES (2, 'common', '高级', '100', '', 0, 2, 1686157775, 0);
INSERT INTO `common_dictionary_integral` VALUES (3, 'common', '特价', '500', '', 0, 3, 1686157786, 0);

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_dictionary_table
-- ----------------------------
INSERT INTO `common_dictionary_table` VALUES (1, 0, 'common', '积分等级', '积分等级分类', 'common_dictionary_integral', 0, 1, 1686152038);
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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '业务端邮箱' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_email
-- ----------------------------
INSERT INTO `common_email` VALUES (1, 'business', 1, '504500934@qq.com', 'amidmyjnnxy(youkey)', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);
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
  `is_install` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否安装0=未安装，1=已安装，2=已卸载',
  `tpl_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'list' COMMENT '模板类型list=仅一个数据，cate=数据加分类',
  `cate_tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类表名称',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '上传时间',
  `updatetime` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_generatecode
-- ----------------------------
INSERT INTO `common_generatecode` VALUES (1, 'admin_auth_dept', '管理后台部门', 'InnoDB', 2, 'utf8mb4_general_ci', 7, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (2, 'admin_auth_role', '权限分组', 'InnoDB', 8, 'utf8mb4_general_ci', 17, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (3, 'admin_auth_role_access', 'admin端菜单权限', 'InnoDB', 6, 'utf8mb4_general_ci', 0, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (4, 'admin_auth_rule', 'C端-菜单', 'InnoDB', 22, 'utf8mb4_general_ci', 81, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (5, 'admin_user', '用户端-用户信息', 'InnoDB', 4, 'utf8mb4_general_ci', 10, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (6, 'attachment', '附件管理', 'InnoDB', 9, 'utf8mb4_general_ci', 734, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (7, 'business_attachment', '客户端附件', 'InnoDB', 66, 'utf8mb4_general_ci', 107, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (8, 'business_auth_dept', '管理后台部门', 'InnoDB', 5, 'utf8mb4_general_ci', 7, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (9, 'business_auth_role', '权限分组', 'InnoDB', 10, 'utf8mb4_general_ci', 21, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (10, 'business_auth_role_access', '商务端菜单授权', 'InnoDB', 10, 'utf8mb4_general_ci', 0, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (11, 'business_auth_rule', 'C端-菜单', 'InnoDB', 25, 'utf8mb4_general_ci', 144, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (12, 'business_createcode_api', '测试api代码生成', 'InnoDB', 6, 'utf8mb4_general_ci', 15, 0, 110, '', 'api', 'api', 'createcode/api/index', 'business/createcode', 'api.go', 'id,accountID,businessID,status,name,nickename,image,file,weigh,createtime,des,content', 120, 1, 'list', 'business_createcode_cate', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (13, 'business_createcode_band', '测试品牌', 'InnoDB', 2, 'utf8mb4_general_ci', 4, 0, 110, '', 'band', 'band', 'createcode/band/index', 'business/createcode', 'band.go', 'id,name,des,status,businessID,content', 144, 1, 'list', 'business_createcode_cate', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (14, 'business_createcode_cate', '测试代码生成分类', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 0, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (15, 'business_createcode_code', '测试代码生成', 'InnoDB', 2, 'utf8mb4_general_ci', 5, 0, 110, '', 'code', 'code', 'createcode/code/index', 'business/createcode', 'code.go', 'accountID,businessID,cid,age,id,name,content,image', 141, 1, 'cate', 'business_createcode_cate', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (16, 'business_home_quickop', '首页快捷操作', 'InnoDB', 2, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (17, 'business_user', '用户端-用户信息', 'InnoDB', 7, 'utf8mb4_general_ci', 14, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (18, 'business_wxsys_officonfig', '微信公众号配置', 'InnoDB', 2, 'utf8mb4_general_ci', 7, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (19, 'business_wxsys_user', '微信关注用户', 'InnoDB', 0, 'utf8mb4_general_ci', 1, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (20, 'business_wxsys_wxappconfig', '微信小程序配置', 'InnoDB', 2, 'utf8mb4_general_ci', 7, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (21, 'business_wxsys_wxmenu', '微站微信菜单', 'InnoDB', 3, 'utf8mb4_general_ci', 14, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (22, 'common_apitext', '接口测试数据', 'InnoDB', 13, 'utf8mb4_general_ci', 19, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (23, 'common_apitext_group', '后台端接口测试分组', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (24, 'common_apitext_type', '接口类型', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (25, 'common_config', '系统配置参数', 'InnoDB', 0, 'utf8mb4_general_ci', 3, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (26, 'common_dictionary_data', '字典数据-测试数据', 'InnoDB', 2, 'utf8mb4_general_ci', 3, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (27, 'common_dictionary_integral', '积分等级-测试数据', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (28, 'common_dictionary_table', '字典表', 'InnoDB', 0, 'utf8mb4_general_ci', 10, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (29, 'common_email', '业务端邮箱', 'InnoDB', 0, 'utf8mb4_general_ci', 3, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (30, 'common_generatecode', '代码生成', 'InnoDB', 0, 'utf8mb4_general_ci', 346, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (31, 'common_logininfo', '登录页面内容', 'InnoDB', 0, 'utf8mb4_general_ci', 4, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (32, 'common_message', '系统通用消息', 'InnoDB', 0, 'utf8mb4_general_ci', 0, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (33, 'common_picture', '图片库', 'InnoDB', 0, 'utf8mb4_general_ci', 9, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (34, 'common_picture_cate', '分类名称', 'InnoDB', 0, 'utf8mb4_general_ci', 28, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (35, 'common_verify_code', '验证码存储', 'InnoDB', 0, 'utf8mb4_general_ci', 2, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (36, 'login_logs', '（平台及客户）后台登录日志', 'InnoDB', 337, 'utf8mb4_general_ci', 897, 1, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692894701, 1692897217);
INSERT INTO `common_generatecode` VALUES (37, 'business_email', '业务端邮箱', 'InnoDB', 0, 'utf8mb4_general_ci', 2, 0, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692896953, 1692897217);
INSERT INTO `common_generatecode` VALUES (38, 'common_apitest', '接口测试数据', 'InnoDB', 4, 'utf8mb4_general_ci', 8, 0, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692896953, 1692897217);
INSERT INTO `common_generatecode` VALUES (39, 'common_apitest_group', '后台端接口测试分组', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 0, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692896953, 1692897217);
INSERT INTO `common_generatecode` VALUES (40, 'common_apitest_type', '接口类型', 'InnoDB', 3, 'utf8mb4_general_ci', 4, 0, 0, '', '', '', '', '', '', '', 0, 0, 'list', '', 1692896953, 1692897217);

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
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '图片库' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分类名称' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '验证码存储' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_verify_code
-- ----------------------------
INSERT INTO `common_verify_code` VALUES (1, 'huang_li_shi@163.com', '380466', 1676913544);

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
) ENGINE = InnoDB AUTO_INCREMENT = 904 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '（平台及客户）后台登录日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login_logs
-- ----------------------------
INSERT INTO `login_logs` VALUES (560, 1, 1, 'in', '183.225.0.57', 1683383371);
INSERT INTO `login_logs` VALUES (561, 1, 1, 'in', '183.225.0.57', 1683383399);
INSERT INTO `login_logs` VALUES (562, 1, 1, 'in', '183.225.0.57', 1683383646);
INSERT INTO `login_logs` VALUES (563, 1, 1, 'in', '', 1683384019);
INSERT INTO `login_logs` VALUES (564, 1, 1, 'in', '', 1683384037);
INSERT INTO `login_logs` VALUES (565, 1, 1, 'in', '183.225.0.57', 1683389457);
INSERT INTO `login_logs` VALUES (566, 1, 1, 'in', '183.225.0.57', 1683391918);
INSERT INTO `login_logs` VALUES (567, 1, 1, 'in', '183.225.0.57', 1683393134);
INSERT INTO `login_logs` VALUES (568, 1, 1, 'in', '183.225.0.57', 1683428426);
INSERT INTO `login_logs` VALUES (569, 1, 1, 'in', '183.225.0.57', 1683431732);
INSERT INTO `login_logs` VALUES (570, 1, 1, 'in', '183.225.0.57', 1683462618);
INSERT INTO `login_logs` VALUES (571, 1, 1, 'in', '183.225.0.57', 1683478432);
INSERT INTO `login_logs` VALUES (572, 1, 1, 'in', '183.225.0.57', 1683479260);
INSERT INTO `login_logs` VALUES (573, 1, 1, 'in', '183.225.0.57', 1683712303);
INSERT INTO `login_logs` VALUES (574, 1, 1, 'in', '27.153.182.231', 1683854201);
INSERT INTO `login_logs` VALUES (575, 1, 1, 'in', '183.225.0.44', 1683861545);
INSERT INTO `login_logs` VALUES (576, 1, 1, 'in', '120.228.134.49', 1683867360);
INSERT INTO `login_logs` VALUES (577, 1, 1, 'in', '183.225.0.44', 1683900659);
INSERT INTO `login_logs` VALUES (578, 1, 1, 'in', '39.128.24.104', 1684943010);
INSERT INTO `login_logs` VALUES (579, 1, 1, 'in', '39.128.24.191', 1685169495);
INSERT INTO `login_logs` VALUES (580, 1, 1, 'in', '', 1685366364);
INSERT INTO `login_logs` VALUES (581, 1, 1, 'in', '', 1685370063);
INSERT INTO `login_logs` VALUES (582, 1, 1, 'in', '', 1685421859);
INSERT INTO `login_logs` VALUES (583, 1, 1, 'in', '', 1685453427);
INSERT INTO `login_logs` VALUES (584, 1, 1, 'in', '', 1685461284);
INSERT INTO `login_logs` VALUES (585, 1, 1, 'in', '', 1685461568);
INSERT INTO `login_logs` VALUES (586, 1, 1, 'in', '', 1685533278);
INSERT INTO `login_logs` VALUES (587, 1, 1, 'in', '', 1685547586);
INSERT INTO `login_logs` VALUES (588, 1, 1, 'in', '116.54.72.27', 1685592278);
INSERT INTO `login_logs` VALUES (589, 1, 1, 'in', '', 1685594423);
INSERT INTO `login_logs` VALUES (590, 1, 1, 'in', '', 1685628956);
INSERT INTO `login_logs` VALUES (591, 1, 1, 'in', '', 1685635364);
INSERT INTO `login_logs` VALUES (592, 1, 1, 'in', '39.128.24.147', 1685638104);
INSERT INTO `login_logs` VALUES (593, 1, 1, 'in', '116.54.72.27', 1685665344);
INSERT INTO `login_logs` VALUES (594, 1, 1, 'in', '', 1685665855);
INSERT INTO `login_logs` VALUES (595, 1, 1, 'in', '', 1685697256);
INSERT INTO `login_logs` VALUES (596, 1, 1, 'in', '', 1685713340);
INSERT INTO `login_logs` VALUES (597, 1, 1, 'in', '', 1685751879);
INSERT INTO `login_logs` VALUES (598, 1, 1, 'in', '116.249.222.229', 1685760939);
INSERT INTO `login_logs` VALUES (599, 1, 1, 'in', '117.28.107.0', 1685779866);
INSERT INTO `login_logs` VALUES (600, 1, 1, 'in', '111.163.91.193', 1685780375);
INSERT INTO `login_logs` VALUES (601, 1, 1, 'in', '112.232.224.221', 1685780545);
INSERT INTO `login_logs` VALUES (602, 1, 1, 'in', '112.224.75.103', 1685781487);
INSERT INTO `login_logs` VALUES (603, 1, 1, 'in', '123.147.251.84', 1685782272);
INSERT INTO `login_logs` VALUES (604, 1, 1, 'in', '58.39.165.192', 1685786769);
INSERT INTO `login_logs` VALUES (605, 1, 1, 'in', '39.128.24.147', 1685801513);
INSERT INTO `login_logs` VALUES (606, 1, 1, 'in', '123.139.44.90', 1685808569);
INSERT INTO `login_logs` VALUES (607, 1, 1, 'in', '183.225.0.181', 1685863883);
INSERT INTO `login_logs` VALUES (608, 1, 1, 'in', '124.127.147.126', 1685875104);
INSERT INTO `login_logs` VALUES (609, 1, 1, 'in', '125.70.176.170', 1685875348);
INSERT INTO `login_logs` VALUES (610, 1, 1, 'in', '125.70.176.170', 1685875385);
INSERT INTO `login_logs` VALUES (611, 1, 1, 'in', '39.163.103.186', 1685875638);
INSERT INTO `login_logs` VALUES (612, 1, 1, 'in', '106.34.163.79', 1685876598);
INSERT INTO `login_logs` VALUES (613, 1, 1, 'in', '124.116.36.13', 1685881919);
INSERT INTO `login_logs` VALUES (614, 1, 1, 'in', '113.118.225.103', 1685881927);
INSERT INTO `login_logs` VALUES (615, 1, 1, 'in', '113.118.225.103', 1685881982);
INSERT INTO `login_logs` VALUES (616, 1, 1, 'in', '42.95.230.81', 1685887348);
INSERT INTO `login_logs` VALUES (617, 1, 1, 'in', '114.254.1.145', 1685888455);
INSERT INTO `login_logs` VALUES (618, 1, 1, 'in', '36.113.194.92', 1685891705);
INSERT INTO `login_logs` VALUES (619, 1, 1, 'in', '183.225.0.181', 1685898018);
INSERT INTO `login_logs` VALUES (620, 1, 1, 'in', '36.60.202.75', 1685917748);
INSERT INTO `login_logs` VALUES (621, 1, 1, 'in', '36.60.202.75', 1685917766);
INSERT INTO `login_logs` VALUES (622, 1, 1, 'in', '113.57.43.181', 1685929108);
INSERT INTO `login_logs` VALUES (623, 1, 1, 'in', '222.185.146.105', 1685939321);
INSERT INTO `login_logs` VALUES (624, 1, 1, 'in', '113.89.232.171', 1685950288);
INSERT INTO `login_logs` VALUES (625, 1, 1, 'in', '183.240.204.116', 1685951601);
INSERT INTO `login_logs` VALUES (626, 1, 1, 'in', '', 1685965253);
INSERT INTO `login_logs` VALUES (627, 1, 1, 'in', '115.57.136.107', 1685966646);
INSERT INTO `login_logs` VALUES (628, 1, 1, 'in', '115.57.136.107', 1685966840);
INSERT INTO `login_logs` VALUES (629, 1, 1, 'in', '122.4.121.255', 1685975572);
INSERT INTO `login_logs` VALUES (630, 1, 1, 'in', '220.248.57.68', 1686011156);
INSERT INTO `login_logs` VALUES (631, 1, 1, 'in', '', 1686038571);
INSERT INTO `login_logs` VALUES (632, 1, 1, 'in', '125.120.231.236', 1686114957);
INSERT INTO `login_logs` VALUES (633, 1, 1, 'in', '125.120.231.236', 1686114987);
INSERT INTO `login_logs` VALUES (634, 1, 1, 'in', '183.14.133.124', 1686120913);
INSERT INTO `login_logs` VALUES (635, 1, 1, 'in', '', 1686127258);
INSERT INTO `login_logs` VALUES (636, 1, 1, 'in', '', 1686130034);
INSERT INTO `login_logs` VALUES (637, 1, 1, 'in', '', 1686132114);
INSERT INTO `login_logs` VALUES (638, 1, 1, 'in', '', 1686136534);
INSERT INTO `login_logs` VALUES (639, 1, 1, 'in', '', 1686136567);
INSERT INTO `login_logs` VALUES (640, 1, 1, 'in', '', 1686138311);
INSERT INTO `login_logs` VALUES (641, 1, 1, 'in', '', 1686138486);
INSERT INTO `login_logs` VALUES (642, 1, 1, 'in', '116.77.28.207', 1686151082);
INSERT INTO `login_logs` VALUES (643, 1, 1, 'in', '', 1686151967);
INSERT INTO `login_logs` VALUES (644, 1, 1, 'in', '', 1686152238);
INSERT INTO `login_logs` VALUES (645, 1, 1, 'in', '', 1686155225);
INSERT INTO `login_logs` VALUES (646, 1, 1, 'in', '183.225.0.34', 1686158124);
INSERT INTO `login_logs` VALUES (647, 1, 1, 'in', '183.225.0.34', 1686158193);
INSERT INTO `login_logs` VALUES (648, 1, 1, 'in', '183.225.0.34', 1686158253);
INSERT INTO `login_logs` VALUES (649, 1, 1, 'in', '183.225.0.34', 1686191668);
INSERT INTO `login_logs` VALUES (650, 1, 1, 'in', '149.129.70.207', 1686192306);
INSERT INTO `login_logs` VALUES (651, 1, 1, 'in', '', 1686192346);
INSERT INTO `login_logs` VALUES (652, 1, 1, 'in', '149.129.70.207', 1686192478);
INSERT INTO `login_logs` VALUES (653, 1, 1, 'in', '', 1686199374);
INSERT INTO `login_logs` VALUES (654, 1, 1, 'in', '', 1686200495);
INSERT INTO `login_logs` VALUES (655, 1, 1, 'in', '183.237.232.198', 1686203719);
INSERT INTO `login_logs` VALUES (656, 1, 1, 'in', '', 1686210161);
INSERT INTO `login_logs` VALUES (657, 1, 1, 'in', '', 1686217809);
INSERT INTO `login_logs` VALUES (658, 1, 1, 'in', '', 1686220716);
INSERT INTO `login_logs` VALUES (659, 1, 1, 'in', '', 1686229208);
INSERT INTO `login_logs` VALUES (660, 1, 1, 'in', '183.225.0.34', 1686234921);
INSERT INTO `login_logs` VALUES (661, 1, 1, 'in', '183.225.0.34', 1686235091);
INSERT INTO `login_logs` VALUES (662, 1, 1, 'in', '', 1686235743);
INSERT INTO `login_logs` VALUES (663, 1, 1, 'in', '', 1686237720);
INSERT INTO `login_logs` VALUES (664, 1, 1, 'in', '', 1686269055);
INSERT INTO `login_logs` VALUES (665, 1, 1, 'in', '', 1686270222);
INSERT INTO `login_logs` VALUES (666, 1, 1, 'in', '', 1686279599);
INSERT INTO `login_logs` VALUES (667, 1, 1, 'in', '183.225.0.34', 1686284704);
INSERT INTO `login_logs` VALUES (668, 1, 1, 'in', '115.236.90.234', 1686284729);
INSERT INTO `login_logs` VALUES (669, 1, 1, 'in', '115.236.90.234', 1686285158);
INSERT INTO `login_logs` VALUES (670, 1, 1, 'in', '115.236.90.234', 1686285178);
INSERT INTO `login_logs` VALUES (671, 1, 1, 'in', '', 1686285550);
INSERT INTO `login_logs` VALUES (672, 1, 1, 'in', '117.183.5.190', 1686286743);
INSERT INTO `login_logs` VALUES (673, 1, 1, 'in', '183.225.0.34', 1686296900);
INSERT INTO `login_logs` VALUES (674, 1, 1, 'in', '27.38.202.157', 1686318507);
INSERT INTO `login_logs` VALUES (675, 1, 1, 'in', '', 1686320079);
INSERT INTO `login_logs` VALUES (676, 1, 1, 'in', '183.225.0.34', 1686322903);
INSERT INTO `login_logs` VALUES (677, 1, 1, 'in', '183.225.0.34', 1686329165);
INSERT INTO `login_logs` VALUES (678, 1, 1, 'in', '183.225.0.34', 1686329791);
INSERT INTO `login_logs` VALUES (679, 1, 1, 'in', '106.61.235.132', 1686359452);
INSERT INTO `login_logs` VALUES (680, 1, 1, 'in', '183.225.0.34', 1686375943);
INSERT INTO `login_logs` VALUES (681, 1, 1, 'in', '183.225.0.34', 1686378894);
INSERT INTO `login_logs` VALUES (682, 1, 1, 'in', '27.16.231.24', 1686383470);
INSERT INTO `login_logs` VALUES (683, 1, 1, 'in', '', 1686388629);
INSERT INTO `login_logs` VALUES (684, 1, 1, 'in', '', 1686403623);
INSERT INTO `login_logs` VALUES (685, 1, 1, 'in', '183.225.0.34', 1686410698);
INSERT INTO `login_logs` VALUES (686, 1, 1, 'in', '183.225.0.34', 1686414374);
INSERT INTO `login_logs` VALUES (687, 1, 1, 'in', '220.163.70.136', 1686551892);
INSERT INTO `login_logs` VALUES (688, 1, 1, 'in', '113.118.174.185', 1686558973);
INSERT INTO `login_logs` VALUES (689, 1, 1, 'in', '223.160.200.97', 1686562051);
INSERT INTO `login_logs` VALUES (690, 1, 1, 'in', '220.163.70.136', 1686562071);
INSERT INTO `login_logs` VALUES (691, 1, 1, 'in', '39.130.117.90', 1686562101);
INSERT INTO `login_logs` VALUES (692, 1, 1, 'in', '117.147.35.29', 1686578780);
INSERT INTO `login_logs` VALUES (693, 1, 1, 'in', '183.225.0.117', 1686583954);
INSERT INTO `login_logs` VALUES (694, 1, 1, 'in', '14.20.91.29', 1686641228);
INSERT INTO `login_logs` VALUES (695, 1, 1, 'in', '61.134.63.210', 1686641228);
INSERT INTO `login_logs` VALUES (696, 1, 1, 'in', '220.163.91.184', 1686649826);
INSERT INTO `login_logs` VALUES (697, 1, 1, 'in', '220.163.91.184', 1686709874);
INSERT INTO `login_logs` VALUES (698, 1, 1, 'in', '220.163.91.184', 1686710239);
INSERT INTO `login_logs` VALUES (699, 1, 1, 'in', '220.163.91.184', 1686731611);
INSERT INTO `login_logs` VALUES (700, 1, 1, 'in', '117.22.140.239', 1686743366);
INSERT INTO `login_logs` VALUES (701, 1, 1, 'in', '115.236.90.234', 1686798210);
INSERT INTO `login_logs` VALUES (702, 1, 1, 'in', '110.185.192.95', 1686800668);
INSERT INTO `login_logs` VALUES (703, 1, 1, 'in', '115.236.90.234', 1686810386);
INSERT INTO `login_logs` VALUES (704, 1, 1, 'in', '113.105.103.89', 1686822443);
INSERT INTO `login_logs` VALUES (705, 1, 1, 'in', '183.208.6.93', 1686964476);
INSERT INTO `login_logs` VALUES (706, 1, 1, 'in', '223.246.219.120', 1686982569);
INSERT INTO `login_logs` VALUES (707, 1, 1, 'in', '111.85.215.134', 1687043610);
INSERT INTO `login_logs` VALUES (708, 1, 1, 'in', '121.225.190.232', 1687080861);
INSERT INTO `login_logs` VALUES (709, 1, 1, 'in', '121.225.190.232', 1687081066);
INSERT INTO `login_logs` VALUES (710, 1, 1, 'in', '112.20.84.119', 1687107136);
INSERT INTO `login_logs` VALUES (711, 1, 1, 'in', '117.147.34.14', 1687398679);
INSERT INTO `login_logs` VALUES (712, 1, 1, 'in', '117.147.34.14', 1687398727);
INSERT INTO `login_logs` VALUES (713, 1, 1, 'in', '183.197.158.128', 1687459030);
INSERT INTO `login_logs` VALUES (714, 1, 1, 'in', '60.255.73.54', 1687460254);
INSERT INTO `login_logs` VALUES (715, 1, 1, 'in', '39.149.191.202', 1687475475);
INSERT INTO `login_logs` VALUES (716, 1, 1, 'in', '112.195.25.13', 1687479568);
INSERT INTO `login_logs` VALUES (717, 1, 1, 'in', '223.104.68.239', 1687480454);
INSERT INTO `login_logs` VALUES (718, 1, 1, 'in', '39.149.191.202', 1687481439);
INSERT INTO `login_logs` VALUES (719, 1, 1, 'in', '115.193.230.255', 1687485912);
INSERT INTO `login_logs` VALUES (720, 1, 1, 'in', '112.25.223.101', 1687501074);
INSERT INTO `login_logs` VALUES (721, 1, 1, 'in', '111.183.143.85', 1687503373);
INSERT INTO `login_logs` VALUES (722, 1, 1, 'in', '221.4.204.245', 1687507669);
INSERT INTO `login_logs` VALUES (723, 1, 1, 'in', '180.138.19.124', 1687509160);
INSERT INTO `login_logs` VALUES (724, 1, 1, 'in', '39.128.24.111', 1687512944);
INSERT INTO `login_logs` VALUES (725, 1, 1, 'in', '222.137.181.195', 1687523257);
INSERT INTO `login_logs` VALUES (726, 1, 1, 'in', '120.41.157.124', 1687527393);
INSERT INTO `login_logs` VALUES (727, 1, 1, 'in', '39.130.70.71', 1687530461);
INSERT INTO `login_logs` VALUES (728, 1, 1, 'in', '39.130.70.71', 1687530482);
INSERT INTO `login_logs` VALUES (729, 1, 1, 'in', '39.128.24.111', 1687599192);
INSERT INTO `login_logs` VALUES (730, 1, 1, 'in', '182.244.221.142', 1687661328);
INSERT INTO `login_logs` VALUES (731, 1, 1, 'in', '120.41.157.124', 1687674668);
INSERT INTO `login_logs` VALUES (732, 1, 1, 'in', '120.41.157.124', 1687674671);
INSERT INTO `login_logs` VALUES (733, 1, 1, 'in', '202.45.129.182', 1687679246);
INSERT INTO `login_logs` VALUES (734, 1, 1, 'in', '115.236.90.234', 1687740546);
INSERT INTO `login_logs` VALUES (735, 1, 1, 'in', '61.166.195.42', 1687743559);
INSERT INTO `login_logs` VALUES (736, 1, 1, 'in', '39.128.24.111', 1687793972);
INSERT INTO `login_logs` VALUES (737, 1, 1, 'in', '', 1687794139);
INSERT INTO `login_logs` VALUES (738, 1, 1, 'in', '60.191.9.124', 1687831001);
INSERT INTO `login_logs` VALUES (739, 1, 1, 'in', '114.35.24.110', 1687897525);
INSERT INTO `login_logs` VALUES (740, 1, 1, 'in', '183.48.244.214', 1687940436);
INSERT INTO `login_logs` VALUES (741, 1, 1, 'in', '183.48.244.214', 1687956465);
INSERT INTO `login_logs` VALUES (742, 1, 1, 'in', '106.61.196.136', 1687999125);
INSERT INTO `login_logs` VALUES (743, 1, 1, 'in', '183.158.108.178', 1688004002);
INSERT INTO `login_logs` VALUES (744, 1, 1, 'in', '183.158.108.178', 1688005038);
INSERT INTO `login_logs` VALUES (745, 1, 1, 'in', '221.4.204.245', 1688005517);
INSERT INTO `login_logs` VALUES (746, 1, 1, 'in', '221.4.204.245', 1688007850);
INSERT INTO `login_logs` VALUES (747, 1, 1, 'in', '183.158.108.178', 1688013097);
INSERT INTO `login_logs` VALUES (748, 1, 1, 'in', '222.221.175.89', 1688027303);
INSERT INTO `login_logs` VALUES (749, 1, 1, 'in', '222.221.175.89', 1688027322);
INSERT INTO `login_logs` VALUES (750, 1, 1, 'in', '106.61.227.20', 1688084288);
INSERT INTO `login_logs` VALUES (751, 1, 1, 'in', '1.203.80.242', 1688110280);
INSERT INTO `login_logs` VALUES (752, 1, 1, 'in', '1.203.80.242', 1688110593);
INSERT INTO `login_logs` VALUES (753, 1, 1, 'in', '115.56.103.166', 1688189776);
INSERT INTO `login_logs` VALUES (754, 1, 1, 'in', '183.202.138.146', 1688190644);
INSERT INTO `login_logs` VALUES (755, 1, 1, 'in', '183.202.138.146', 1688190716);
INSERT INTO `login_logs` VALUES (756, 1, 1, 'in', '120.69.112.191', 1688281972);
INSERT INTO `login_logs` VALUES (757, 1, 1, 'in', '223.104.68.171', 1688293889);
INSERT INTO `login_logs` VALUES (758, 1, 1, 'in', '1.193.97.16', 1688373612);
INSERT INTO `login_logs` VALUES (759, 1, 1, 'in', '1.193.97.16', 1688374545);
INSERT INTO `login_logs` VALUES (760, 1, 1, 'in', '61.174.157.34', 1688403754);
INSERT INTO `login_logs` VALUES (761, 1, 1, 'in', '61.174.157.34', 1688403919);
INSERT INTO `login_logs` VALUES (762, 1, 1, 'in', '36.142.138.174', 1688429962);
INSERT INTO `login_logs` VALUES (763, 1, 1, 'in', '183.158.119.54', 1688435423);
INSERT INTO `login_logs` VALUES (764, 1, 1, 'in', '183.158.119.54', 1688435432);
INSERT INTO `login_logs` VALUES (765, 1, 1, 'in', '123.232.96.222', 1688441872);
INSERT INTO `login_logs` VALUES (766, 1, 1, 'in', '183.14.29.32', 1688450887);
INSERT INTO `login_logs` VALUES (767, 1, 1, 'in', '222.221.187.15', 1688456123);
INSERT INTO `login_logs` VALUES (768, 1, 1, 'in', '222.71.226.50', 1688459716);
INSERT INTO `login_logs` VALUES (769, 1, 1, 'in', '112.10.134.18', 1688476281);
INSERT INTO `login_logs` VALUES (770, 1, 1, 'in', '223.12.75.200', 1688523765);
INSERT INTO `login_logs` VALUES (771, 1, 1, 'in', '113.110.229.99', 1688524074);
INSERT INTO `login_logs` VALUES (772, 1, 1, 'in', '182.150.24.124', 1688536019);
INSERT INTO `login_logs` VALUES (773, 1, 1, 'in', '182.150.24.124', 1688536500);
INSERT INTO `login_logs` VALUES (774, 1, 1, 'in', '182.150.24.124', 1688536508);
INSERT INTO `login_logs` VALUES (775, 1, 1, 'in', '182.150.24.124', 1688543867);
INSERT INTO `login_logs` VALUES (776, 1, 1, 'in', '171.107.26.36', 1688575313);
INSERT INTO `login_logs` VALUES (777, 1, 1, 'in', '14.104.143.200', 1688614689);
INSERT INTO `login_logs` VALUES (778, 1, 1, 'in', '14.104.143.200', 1688615012);
INSERT INTO `login_logs` VALUES (779, 1, 1, 'in', '183.158.119.54', 1688623789);
INSERT INTO `login_logs` VALUES (780, 1, 1, 'in', '183.158.119.54', 1688624949);
INSERT INTO `login_logs` VALUES (781, 1, 1, 'in', '14.104.143.200', 1688629532);
INSERT INTO `login_logs` VALUES (782, 1, 1, 'in', '175.8.163.208', 1688638408);
INSERT INTO `login_logs` VALUES (783, 1, 1, 'in', '117.147.35.109', 1688656457);
INSERT INTO `login_logs` VALUES (784, 1, 1, 'in', '61.170.161.207', 1688711128);
INSERT INTO `login_logs` VALUES (785, 1, 1, 'in', '117.136.33.187', 1688731900);
INSERT INTO `login_logs` VALUES (786, 1, 1, 'in', '111.49.59.248', 1688764298);
INSERT INTO `login_logs` VALUES (787, 1, 1, 'in', '111.49.59.248', 1688764377);
INSERT INTO `login_logs` VALUES (788, 1, 1, 'in', '110.7.176.86', 1688783656);
INSERT INTO `login_logs` VALUES (789, 1, 1, 'in', '110.88.104.233', 1688800752);
INSERT INTO `login_logs` VALUES (790, 1, 1, 'in', '1.85.202.18', 1688841481);
INSERT INTO `login_logs` VALUES (791, 1, 1, 'in', '1.85.202.18', 1688841524);
INSERT INTO `login_logs` VALUES (792, 1, 1, 'in', '223.86.234.135', 1688872091);
INSERT INTO `login_logs` VALUES (793, 1, 1, 'in', '1.85.202.18', 1688888146);
INSERT INTO `login_logs` VALUES (794, 1, 1, 'in', '1.85.202.18', 1688888159);
INSERT INTO `login_logs` VALUES (795, 1, 1, 'in', '183.225.0.101', 1688889828);
INSERT INTO `login_logs` VALUES (796, 1, 1, 'in', '115.150.244.124', 1688895154);
INSERT INTO `login_logs` VALUES (797, 1, 1, 'in', '115.150.244.124', 1688895222);
INSERT INTO `login_logs` VALUES (798, 1, 1, 'in', '1.85.202.18', 1688912279);
INSERT INTO `login_logs` VALUES (799, 1, 1, 'in', '1.85.202.18', 1688912380);
INSERT INTO `login_logs` VALUES (800, 1, 1, 'in', '223.104.244.160', 1688925372);
INSERT INTO `login_logs` VALUES (801, 1, 1, 'in', '61.173.26.90', 1688964803);
INSERT INTO `login_logs` VALUES (802, 1, 1, 'in', '110.87.75.186', 1688977571);
INSERT INTO `login_logs` VALUES (803, 1, 1, 'in', '60.186.173.189', 1689002171);
INSERT INTO `login_logs` VALUES (804, 1, 1, 'in', '116.21.231.164', 1689044716);
INSERT INTO `login_logs` VALUES (805, 1, 1, 'in', '116.21.231.164', 1689044791);
INSERT INTO `login_logs` VALUES (806, 1, 1, 'in', '119.163.195.46', 1689066108);
INSERT INTO `login_logs` VALUES (807, 1, 1, 'in', '113.78.173.162', 1689144766);
INSERT INTO `login_logs` VALUES (808, 1, 1, 'in', '222.173.111.58', 1689149436);
INSERT INTO `login_logs` VALUES (809, 1, 1, 'in', '58.34.165.178', 1689149567);
INSERT INTO `login_logs` VALUES (810, 1, 1, 'in', '122.14.46.131', 1689315596);
INSERT INTO `login_logs` VALUES (811, 1, 1, 'in', '182.150.22.78', 1689320917);
INSERT INTO `login_logs` VALUES (812, 1, 1, 'in', '182.150.22.78', 1689321463);
INSERT INTO `login_logs` VALUES (813, 1, 1, 'in', '61.166.195.146', 1689325562);
INSERT INTO `login_logs` VALUES (814, 1, 1, 'in', '110.7.176.231', 1689557027);
INSERT INTO `login_logs` VALUES (815, 1, 1, 'in', '110.86.29.134', 1689575096);
INSERT INTO `login_logs` VALUES (816, 1, 1, 'in', '110.86.29.134', 1689575209);
INSERT INTO `login_logs` VALUES (817, 1, 1, 'in', '14.104.139.152', 1689647002);
INSERT INTO `login_logs` VALUES (818, 1, 1, 'in', '39.152.131.47', 1689662488);
INSERT INTO `login_logs` VALUES (819, 1, 1, 'in', '59.173.19.133', 1689666613);
INSERT INTO `login_logs` VALUES (820, 1, 1, 'in', '117.63.239.52', 1689730021);
INSERT INTO `login_logs` VALUES (821, 1, 1, 'in', '117.63.239.52', 1689730219);
INSERT INTO `login_logs` VALUES (822, 1, 1, 'in', '110.86.29.134', 1689821562);
INSERT INTO `login_logs` VALUES (823, 1, 1, 'in', '110.86.29.134', 1689821590);
INSERT INTO `login_logs` VALUES (824, 1, 1, 'in', '222.221.187.129', 1689824926);
INSERT INTO `login_logs` VALUES (825, 1, 1, 'in', '106.61.95.31', 1689898699);
INSERT INTO `login_logs` VALUES (826, 1, 1, 'in', '120.192.185.122', 1689921047);
INSERT INTO `login_logs` VALUES (827, 1, 1, 'in', '222.137.220.220', 1689928044);
INSERT INTO `login_logs` VALUES (828, 1, 1, 'in', '121.227.80.53', 1689929232);
INSERT INTO `login_logs` VALUES (829, 1, 1, 'in', '121.227.80.53', 1689929984);
INSERT INTO `login_logs` VALUES (830, 1, 1, 'in', '121.227.80.53', 1689931330);
INSERT INTO `login_logs` VALUES (831, 1, 1, 'in', '112.0.138.230', 1689942989);
INSERT INTO `login_logs` VALUES (832, 1, 1, 'in', '117.82.37.135', 1689946332);
INSERT INTO `login_logs` VALUES (833, 1, 1, 'in', '124.114.98.206', 1689997721);
INSERT INTO `login_logs` VALUES (834, 1, 1, 'in', '119.123.42.89', 1690123857);
INSERT INTO `login_logs` VALUES (835, 1, 1, 'in', '119.123.42.89', 1690123927);
INSERT INTO `login_logs` VALUES (836, 1, 1, 'in', '106.61.122.238', 1690158307);
INSERT INTO `login_logs` VALUES (837, 1, 1, 'in', '124.238.61.145', 1690170095);
INSERT INTO `login_logs` VALUES (838, 1, 1, 'in', '58.33.97.119', 1690185856);
INSERT INTO `login_logs` VALUES (839, 1, 1, 'in', '121.227.80.53', 1690193831);
INSERT INTO `login_logs` VALUES (840, 1, 1, 'in', '121.227.80.53', 1690193866);
INSERT INTO `login_logs` VALUES (841, 1, 1, 'in', '223.102.248.98', 1690209638);
INSERT INTO `login_logs` VALUES (842, 1, 1, 'in', '223.102.248.98', 1690209664);
INSERT INTO `login_logs` VALUES (843, 1, 13, 'in', '223.102.248.98', 1690209706);
INSERT INTO `login_logs` VALUES (844, 1, 1, 'in', '106.61.122.238', 1690243970);
INSERT INTO `login_logs` VALUES (845, 1, 1, 'in', '222.221.175.66', 1690250601);
INSERT INTO `login_logs` VALUES (846, 1, 1, 'in', '222.221.175.66', 1690270573);
INSERT INTO `login_logs` VALUES (847, 1, 1, 'in', '112.0.138.230', 1690293239);
INSERT INTO `login_logs` VALUES (848, 1, 1, 'in', '222.221.175.66', 1690364628);
INSERT INTO `login_logs` VALUES (849, 1, 1, 'in', '222.221.175.66', 1690364698);
INSERT INTO `login_logs` VALUES (850, 1, 1, 'in', '114.88.161.112', 1690371745);
INSERT INTO `login_logs` VALUES (851, 1, 1, 'in', '114.88.161.112', 1690371753);
INSERT INTO `login_logs` VALUES (852, 1, 1, 'in', '114.88.161.112', 1690373394);
INSERT INTO `login_logs` VALUES (853, 1, 1, 'in', '222.221.187.72', 1690438006);
INSERT INTO `login_logs` VALUES (854, 1, 1, 'in', '111.3.191.226', 1690448145);
INSERT INTO `login_logs` VALUES (855, 1, 1, 'in', '115.60.172.39', 1690452309);
INSERT INTO `login_logs` VALUES (856, 1, 1, 'in', '115.60.172.39', 1690452368);
INSERT INTO `login_logs` VALUES (857, 1, 1, 'in', '49.75.44.16', 1690508041);
INSERT INTO `login_logs` VALUES (858, 1, 1, 'in', '39.164.170.164', 1690525060);
INSERT INTO `login_logs` VALUES (859, 1, 1, 'in', '39.164.170.164', 1690525083);
INSERT INTO `login_logs` VALUES (860, 1, 1, 'in', '222.221.187.72', 1690532416);
INSERT INTO `login_logs` VALUES (861, 1, 1, 'in', '218.207.196.93', 1690532732);
INSERT INTO `login_logs` VALUES (862, 1, 1, 'in', '111.18.139.7', 1690555788);
INSERT INTO `login_logs` VALUES (863, 1, 1, 'in', '124.119.122.237', 1690576882);
INSERT INTO `login_logs` VALUES (864, 1, 1, 'in', '113.228.184.139', 1690609618);
INSERT INTO `login_logs` VALUES (865, 1, 1, 'in', '39.164.170.164', 1690614395);
INSERT INTO `login_logs` VALUES (866, 1, 1, 'in', '39.164.170.164', 1690614609);
INSERT INTO `login_logs` VALUES (867, 1, 1, 'in', '39.128.22.89', 1690615065);
INSERT INTO `login_logs` VALUES (868, 1, 1, 'in', '116.52.196.121', 1690619651);
INSERT INTO `login_logs` VALUES (869, 1, 1, 'in', '106.61.170.103', 1690625995);
INSERT INTO `login_logs` VALUES (870, 1, 1, 'in', '', 1690643451);
INSERT INTO `login_logs` VALUES (871, 1, 1, 'in', '', 1690644822);
INSERT INTO `login_logs` VALUES (872, 1, 1, 'in', '', 1690644986);
INSERT INTO `login_logs` VALUES (873, 1, 1, 'in', '', 1690646600);
INSERT INTO `login_logs` VALUES (874, 1, 1, 'in', '', 1690647375);
INSERT INTO `login_logs` VALUES (875, 1, 1, 'in', '', 1690649810);
INSERT INTO `login_logs` VALUES (876, 1, 1, 'in', '', 1690651298);
INSERT INTO `login_logs` VALUES (877, 1, 1, 'in', '', 1690652618);
INSERT INTO `login_logs` VALUES (878, 1, 1, 'in', '39.128.22.89', 1690654842);
INSERT INTO `login_logs` VALUES (879, 1, 1, 'in', '39.128.22.89', 1690654844);
INSERT INTO `login_logs` VALUES (880, 1, 1, 'in', '39.128.22.89', 1690654846);
INSERT INTO `login_logs` VALUES (881, 1, 1, 'in', '39.128.22.89', 1690654849);
INSERT INTO `login_logs` VALUES (882, 1, 1, 'in', '39.128.22.89', 1690654854);
INSERT INTO `login_logs` VALUES (883, 1, 1, 'in', '39.128.22.89', 1690654855);
INSERT INTO `login_logs` VALUES (884, 1, 1, 'in', '39.128.22.89', 1690654862);
INSERT INTO `login_logs` VALUES (885, 1, 1, 'in', '39.128.22.89', 1690655282);
INSERT INTO `login_logs` VALUES (886, 1, 1, 'in', '223.104.19.93', 1690686406);
INSERT INTO `login_logs` VALUES (887, 1, 1, 'in', '', 1690690457);
INSERT INTO `login_logs` VALUES (888, 1, 1, 'in', '', 1690690617);
INSERT INTO `login_logs` VALUES (889, 1, 1, 'in', '', 1690691111);
INSERT INTO `login_logs` VALUES (890, 1, 1, 'in', '', 1690691335);
INSERT INTO `login_logs` VALUES (891, 1, 1, 'in', '', 1690691351);
INSERT INTO `login_logs` VALUES (892, 1, 1, 'in', '112.42.28.37', 1690691577);
INSERT INTO `login_logs` VALUES (893, 1, 1, 'in', '', 1690691791);
INSERT INTO `login_logs` VALUES (894, 1, 1, 'in', '', 1690694819);
INSERT INTO `login_logs` VALUES (895, 1, 1, 'in', '39.128.22.63', 1691759555);
INSERT INTO `login_logs` VALUES (896, 1, 1, 'in', '106.84.161.169', 1691759668);
INSERT INTO `login_logs` VALUES (897, 1, 1, 'in', '', 1692894262);
INSERT INTO `login_logs` VALUES (898, 1, 1, 'in', '', 1693020706);
INSERT INTO `login_logs` VALUES (899, 1, 1, 'in', '', 1693039484);
INSERT INTO `login_logs` VALUES (900, 1, 1, 'in', '', 1693068953);
INSERT INTO `login_logs` VALUES (901, 1, 1, 'in', '', 1693106638);
INSERT INTO `login_logs` VALUES (902, 1, 1, 'in', '', 1693106816);
INSERT INTO `login_logs` VALUES (903, 1, 1, 'in', '', 1694235940);
INSERT INTO `login_logs` VALUES (904, 1, 1, 'in', '', 1694236616);
INSERT INTO `login_logs` VALUES (905, 1, 1, 'in', '', 1694237243);

SET FOREIGN_KEY_CHECKS = 1;
