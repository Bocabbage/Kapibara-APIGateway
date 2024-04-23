USE `auth_test_db0`;

-- create tables

CREATE TABLE IF NOT EXISTS `users`(
    `account_id` INT AUTO_INCREMENT,
    `account` VARCHAR(50) UNIQUE NOT NULL,
    `user_name` VARCHAR(50) NOT NULL,
    `pwd_hash` VARCHAR(100) NOT NULL,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` SMALLINT DEFAULT 0 COMMENT '0: not-in-use, 1: not-active, 2: in-use, 3: black-list',
    `role_bitmap` BIGINT DEFAULT 1,
    PRIMARY KEY (`account_id`),
    UNIQUE KEY `account_key` (`account`),
    UNIQUE KEY `user_name_key` (`user_name`)
) DEFAULT CHARSET=utf8;

-- CREATE TABLE IF NOT EXISTS `roles`(
--     `roleBitmap` INT UNIQUE,
--     `roleName` VARCHAR(50)
-- );

-- insert test-data

-- INSERT INTO users ( account, username, pwdHash, status, roleBitmap)
-- VALUES ("KapibaraTest", "UuuuserName", "pppppppwdHash!", 2, 1);

-- INSERT INTO users ( account, username, pwdHash, status, roleBitmap)
-- VALUES ("KapibaraTest2", "UuuuserName2", "cb61b53e13b064123310afac8464dc07234675cc0d1c9c0980fa69acea31e927", 2, 1);

-- INSERT INTO roles (roleBitmap, roleName)
-- VALUES (1, "admin");