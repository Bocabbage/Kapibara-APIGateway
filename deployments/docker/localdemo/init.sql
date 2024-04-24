USE `auth_test_db0`;

-- create tables

CREATE TABLE IF NOT EXISTS `user`(
    `id` UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `account` VARCHAR(64) UNIQUE NOT NULL,
    `user_name` VARCHAR(64) UNIQUE NOT NULL,
    `pwd_hash` VARCHAR(100) NOT NULL,
    `create_time` datetime NOT NULL, -- DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL, -- DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime DEFAULT NULL,
    `status` SMALLINT DEFAULT 0 COMMENT '0: not-in-use, 1: not-active, 2: in-use, 3: black-list',
    `role_bitmap` BIGINT DEFAULT 1,
) DEFAULT CHARSET=utf8;