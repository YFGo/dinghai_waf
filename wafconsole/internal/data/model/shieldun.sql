/*
 Navicat Premium Data Transfer

 Source Server         : 张云飞
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 152.136.50.60:3306
 Source Schema         : shieldun

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 06/10/2024 12:46:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_sites
-- ----------------------------
DROP TABLE IF EXISTS `app_sites`;
CREATE TABLE `app_sites`  (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '应用id',
                              `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'app名称',
                              `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'WEB程序URL',
                              `match` tinyint default 1 COMMENT '匹配方式（1: full match, 2: prefix match, 3: suffix）',
                              `parent_id` bigint NOT NULL COMMENT '父级ID',
                              `create_at` datetime NOT NULL,
                              `update_at` datetime NOT NULL,
                              `delete_at` datetime NULL DEFAULT NULL COMMENT '软删',
                              PRIMARY KEY (`id`) USING BTREE,
                              UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '应用级防护对象' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for server_sites
-- ----------------------------
DROP TABLE IF EXISTS `server_sites`;
CREATE TABLE `server_sites`  (
                                 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '服务器id',
                                 `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '站点名称',
                                 `ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'ip',
                                 `port` int NOT NULL COMMENT '监听端口',
                                 `tls` tinyint NOT NULL COMMENT '是否使用加密传输（0: disable, 1: enable）',
                                 `cert` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'base64格式的证书',
                                 `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'base64格式的私钥',
                                 `upstreams` json  COMMENT '上游服务器列表',
                                 `parent_id` bigint COMMENT '父级ID（服务组ID）',
                                 `create_at` datetime NOT NULL,
                                 `update_at` datetime NOT NULL,
                                 `delete_at` datetime NULL DEFAULT NULL COMMENT '软删',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服务器级防护对象' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for site_groups
-- ----------------------------
DROP TABLE IF EXISTS `site_groups`;
CREATE TABLE `site_groups`  (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '服务器组id',
                                `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '站点组名称',
                                `mod` tinyint NOT NULL COMMENT '防护模式(0: offline, 1: audit, 2: defence)',
                                `hosts` json NOT NULL COMMENT '防护的host列表',
                                `create_at` datetime NOT NULL,
                                `update_at` datetime NOT NULL,
                                `delete_at` datetime NULL DEFAULT NULL COMMENT '软删',
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '防护站点组，也成服务器组' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for buildin_rules
-- ----------------------------
DROP TABLE IF EXISTS `buildin_rules`;
CREATE TABLE `buildin_rules`  (
                                  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '内核规则表',
                                  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '规则名称',
                                  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '规则功能描述',
                                  `pattern` json NOT NULL COMMENT 'JSON结构描述的规则表达式',
                                  `risk_level` tinyint NOT NULL COMMENT '风险等级',
                                  `group_id` bigint NOT NULL COMMENT '所属规则组',
                                  `status` tinyint NOT NULL COMMENT '状态（0: 不启用， 1:启用）',
                                  `update_at` datetime NOT NULL,
                                  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删',
                                  `create_at` datetime NOT NULL,
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统内置规则表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for custom_rules
-- ----------------------------
DROP TABLE IF EXISTS `custom_rules`;
CREATE TABLE `custom_rules`  (
                                 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户自定义规则集',
                                 `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '规则名称',
                                 `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '规则功能描述',
                                 `pattern` json NOT NULL COMMENT 'JSON结构描述的规则表达式',
                                 `risk_level` tinyint NOT NULL COMMENT '风险等级',
                                 `group_id` bigint NOT NULL COMMENT '所属规则组',
                                 `create_at` datetime NOT NULL,
                                 `status` tinyint NOT NULL COMMENT '状态（0: 不启用， 1:启用）',
                                 `delete_at` datetime NULL DEFAULT NULL COMMENT '软删',
                                 `update_at` datetime NOT NULL,
                                 PRIMARY KEY (`id`) USING BTREE,
                                 UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '自定义规则表' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for rule_groups
-- ----------------------------
DROP TABLE IF EXISTS `rule_groups`;
CREATE TABLE `rule_groups`  (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户自定义规则集',
                                `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '规则组名称',
                                `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '规则组描述',
                                `is_buildin` tinyint NOT NULL DEFAULT 0 COMMENT '是否内置（0:用户定义组，1: 系统内置组）',
                                `create_at` datetime NOT NULL,
                                `update_at` datetime NOT NULL,
                                `delete_at` datetime NULL DEFAULT NULL COMMENT '软删',
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '规则组，也成特征组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for deny_allow_list
-- ----------------------------
DROP TABLE IF EXISTS `deny_allow_list`;
CREATE TABLE `deny_allow_list`  (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
                                    `pattern` json NOT NULL COMMENT '规则配置',
                                    `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'defence' COMMENT '规则动作 defence： 阻止',
                                    `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态（1: 启用， 2:不启用）',
                                    `list_type` tinyint NOT NULL COMMENT '名单类型（1: 黑名单， 2:白名单）',
                                    `create_at` datetime NOT NULL,
                                    `update_at` datetime NOT NULL,
                                    PRIMARY KEY (`id`) USING BTREE,
                                    UNIQUE INDEX `idx_deny_allow_list_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for list_apply
-- ----------------------------
DROP TABLE IF EXISTS `list_apply`;
CREATE TABLE `list_apply`  (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `list_id` bigint NOT NULL COMMENT '名单ID',
                               `key_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '应用/服务器/服务器组ID',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 49 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;







-- ----------------------------
-- Table structure for site_strategies
-- ----------------------------
DROP TABLE IF EXISTS `site_strategies`;
CREATE TABLE `site_strategies`  (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `site_id` bigint NOT NULL COMMENT '站点（应用/服务器/服务器组）ID',
                                    `strategy_id` bigint NOT NULL COMMENT '绑定的策略ID',
                                    `site_level` tinyint NOT NULL COMMENT '站点层级（1: 服务器组， 2:服务器，3: 应用）',
                                    `create_at` datetime NOT NULL,
                                    `update_at` datetime NOT NULL,
                                    PRIMARY KEY (`id`) USING BTREE,
                                    UNIQUE INDEX `site_strategies_index_0`(`site_id` ASC, `strategy_id` ASC, `site_level` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '各级站点应用的策略' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for strategies
-- ----------------------------
DROP TABLE IF EXISTS `strategies`;
CREATE TABLE `strategies`  (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '策略名称',
                               `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '策略描述',
                               `kind` tinyint NOT NULL COMMENT '策略类别（1: 网络， 2:服务器， 3: 应用）',
                               `default` tinyint NOT NULL DEFAULT 0 COMMENT '是否默认策略',
                               `is_buildin` tinyint NOT NULL DEFAULT 0 COMMENT '是否内置（0:用户定义策略，1: 系统内置策略）',
                               `create_at` datetime NOT NULL,
                               `update_at` datetime NULL DEFAULT NULL,
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '防御策略表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for strategy_configs
-- ----------------------------
DROP TABLE IF EXISTS `strategy_configs`;
CREATE TABLE `strategy_configs`  (
                                     `id` bigint NOT NULL AUTO_INCREMENT,
                                     `strategy_id` bigint NOT NULL COMMENT '策略ID',
                                     `rule_group_id` bigint NOT NULL COMMENT '规则组ID',
                                     `status` tinyint NOT NULL COMMENT '状态（0: 不启用， 1:启用）',
                                     `alert_level` tinyint NOT NULL COMMENT '告警级别（0: 无，1: 低， 2：中， 3: 高）',
                                     `action` tinyint NOT NULL COMMENT '命中策略时动作（0: 记录， 1： 阻止）',
                                     `next_action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '命中后的后续动作',
                                     `create_at` datetime NOT NULL,
                                     `update_at` datetime NULL DEFAULT NULL,
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '防御策略表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
