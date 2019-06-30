-- ---
-- Globals
-- ---

-- SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- SET FOREIGN_KEY_CHECKS=0;

-- ---
-- Table 'user'
-- a
-- ---

DROP TABLE IF EXISTS `user`;
		
CREATE TABLE `user` (
  `ulid` VARCHAR(32) NOT NULL,
  `firebase_uid` VARCHAR(36) NULL DEFAULT NULL,
  `stripe_id` CHAR(32) NULL DEFAULT NULL,
  `cart_ulid` VARCHAR(32) NULL DEFAULT NULL,
  `address_ulid` VARCHAR(32) NULL DEFAULT NULL,
  `name` VARCHAR(32) NULL DEFAULT NULL,
  `email` VARCHAR(128) NULL DEFAULT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`),
     UNIQUE KEY (`firebase_uid`),
     UNIQUE KEY (`stripe_id`)
) COMMENT 'a';

-- ---
-- Table 'address'
-- 
-- ---

DROP TABLE IF EXISTS `address`;
		
CREATE TABLE `address` (
  `ulid` VARCHAR(32) NOT NULL,
  `postcode` CHAR(128) NULL DEFAULT NULL,
  `address1` VARCHAR(256) NULL DEFAULT NULL,
  `address2` VARCHAR(256) NULL DEFAULT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'user_address'
-- 
-- ---

DROP TABLE IF EXISTS `user_address`;
		
CREATE TABLE `user_address` (
  `ulid` VARCHAR(32) NOT NULL,
  `user_ulid` VARCHAR(32) NOT NULL,
  `address_ulid` VARCHAR(32) NOT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'order'
-- 
-- ---

DROP TABLE IF EXISTS `order`;
		
CREATE TABLE `order` (
  `ulid` VARCHAR(32) NOT NULL,
  `user_ulid` VARCHAR(32) NOT NULL,
  `address_ulid` VARCHAR(32) NOT NULL,
  `is_canceled` INTEGER(1) NOT NULL DEFAULT 0,
  `status` INTEGER NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'cart'
-- 
-- ---

DROP TABLE IF EXISTS `cart`;
		
CREATE TABLE `cart` (
  `ulid` VARCHAR(32) NOT NULL,
  `user_ulid` VARCHAR(32) NOT NULL,
  `bought_at` DATETIME NULL DEFAULT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'cart_item'
-- 
-- ---

DROP TABLE IF EXISTS `cart_item`;
		
CREATE TABLE `cart_item` (
  `ulid` VARCHAR(32) NOT NULL,
  `cart_ulid` VARCHAR(32) NOT NULL,
  `item_ulid` VARCHAR(32) NOT NULL,
  `amount` INTEGER NOT NULL DEFAULT 0,
  `is_deleted` INTEGER(1) NOT NULL DEFAULT 0,
  `is_keep` INTEGER(1) NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'item'
-- 
-- ---

DROP TABLE IF EXISTS `item`;
		
CREATE TABLE `item` (
  `ulid` VARCHAR(32) NOT NULL,
  `name` VARCHAR(128) NULL DEFAULT NULL,
  `context` VARCHAR(2048) NULL DEFAULT NULL,
  `amount` INTEGER NOT NULL DEFAULT 0,
  `priority` INTEGER NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'price'
-- 
-- ---

DROP TABLE IF EXISTS `price`;
		
CREATE TABLE `price` (
  `ulid` VARCHAR(32) NOT NULL,
  `item_ulid` VARCHAR(32) NOT NULL,
  `price` INTEGER NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'order_item'
-- 
-- ---

DROP TABLE IF EXISTS `order_item`;
		
CREATE TABLE `order_item` (
  `ulid` VARCHAR(32) NOT NULL,
  `order_ulid` VARCHAR(32) NOT NULL,
  `item_ulid` VARCHAR(32) NOT NULL,
  `amount` INTEGER NOT NULL DEFAULT 0,
  `price` INTEGER NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Table 'item_image'
-- 
-- ---

DROP TABLE IF EXISTS `item_image`;
		
CREATE TABLE `item_image` (
  `ulid` CHAR(32) NOT NULL,
  `item_ulid` VARCHAR(32) NOT NULL,
  `path` VARCHAR(1024) NULL DEFAULT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`ulid`)
);

-- ---
-- Foreign Keys 
-- ---

ALTER TABLE `user` ADD FOREIGN KEY (cart_ulid) REFERENCES `cart` (`ulid`);
ALTER TABLE `user` ADD FOREIGN KEY (address_ulid) REFERENCES `address` (`ulid`);
ALTER TABLE `user_address` ADD FOREIGN KEY (user_ulid) REFERENCES `user` (`ulid`);
ALTER TABLE `user_address` ADD FOREIGN KEY (address_ulid) REFERENCES `address` (`ulid`);
ALTER TABLE `order` ADD FOREIGN KEY (user_ulid) REFERENCES `user` (`ulid`);
ALTER TABLE `order` ADD FOREIGN KEY (address_ulid) REFERENCES `address` (`ulid`);
ALTER TABLE `cart` ADD FOREIGN KEY (user_ulid) REFERENCES `user` (`ulid`);
ALTER TABLE `cart_item` ADD FOREIGN KEY (cart_ulid) REFERENCES `cart` (`ulid`);
ALTER TABLE `cart_item` ADD FOREIGN KEY (item_ulid) REFERENCES `item` (`ulid`);
ALTER TABLE `price` ADD FOREIGN KEY (item_ulid) REFERENCES `item` (`ulid`);
ALTER TABLE `order_item` ADD FOREIGN KEY (order_ulid) REFERENCES `order` (`ulid`);
ALTER TABLE `order_item` ADD FOREIGN KEY (item_ulid) REFERENCES `item` (`ulid`);
ALTER TABLE `item_image` ADD FOREIGN KEY (item_ulid) REFERENCES `item` (`ulid`);

-- ---
-- Table Properties
-- ---

-- ALTER TABLE `user` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `address` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `user_address` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `order` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `cart` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `cart_item` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `item` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `price` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `order_item` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `item_image` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ---
-- Test Data
-- ---
INSERT INTO `user` (`ulid`,`firebase_uid`,`cart_ulid`,`address_ulid`,`name`,`email`,`created_at`,`updated_at`) VALUES
('01DDSX4GRR2PSRKCDVTRFA96WH','01DDS-X4GRR2P-SRKCDV-TRFA96WH',NULL,NULL,'テストユーザ','test@gmail.com','2019-06-11 13:58:22','2019-06-20 13:58:22');
INSERT INTO `address` (`ulid`,`postcode`,`address1`,`address2`,`created_at`,`updated_at`) VALUES
('01DDTFT8B2EPJ2MKJK1BDWGQGA','1650021','Virginia, U.S.','Arlington, ','2019-06-11 13:58:22','2019-06-15 13:58:22');
INSERT INTO `user_address` (`ulid`,`user_ulid`,`address_ulid`,`created_at`,`updated_at`) VALUES
('01DDTG1NM077AP5PEZJ6KJJR8X','01DDSX4GRR2PSRKCDVTRFA96WH','01DDTFT8B2EPJ2MKJK1BDWGQGA','2019-06-11 13:58:22','2019-06-13 13:58:22');
-- INSERT INTO `order` (`ulid`,`user_ulid`,`address_ulid`,`is_canceled`,`status`,`created_at`,`updated_at`) VALUES
-- ('','','','','','','');
INSERT INTO `cart` (`ulid`,`user_ulid`,`bought_at`,`created_at`,`updated_at`) VALUES
('01DDTKER1V853DJ7YMS4S59KD7','01DDSX4GRR2PSRKCDVTRFA96WH','2019-06-20 13:58:47','2019-06-11 13:58:22','2019-06-13 13:58:22');
-- INSERT INTO `cart_item` (`ulid`,`cart_ulid`,`item_ulid`,`amount`,`is_deleted`,`is_keep`,`created_at`,`updated_at`) VALUES
-- ('','','','','','','','');
INSERT INTO `item` (`ulid`,`name`,`context`,`amount`,`priority`,`created_at`,`updated_at`) VALUES
('01DDNE37NVJT1DVBT3ZB3C1Q0F','みかん','くだもの','23','5','2019-06-18 13:58:20','2019-06-18 13:58:22');
INSERT INTO `item` (`ulid`,`name`,`context`,`amount`,`priority`,`created_at`,`updated_at`) VALUES
('01DDNE42FBVZPPJE0DMW4PE2TK','りんご','くだもの2','25','1','2019-06-18 13:58:47','2019-06-18 13:58:47');
-- INSERT INTO `price` (`ulid`,`item_ulid`,`price`,`created_at`,`updated_at`) VALUES
-- ('','','','','');
-- INSERT INTO `order_item` (`ulid`,`order_ulid`,`item_ulid`,`amount`,`price`,`created_at`,`updated_at`) VALUES
-- ('','','','','','','');
-- INSERT INTO `item_image` (`ulid`,`item_ulid`,`path`,`created_at`,`updated_at`) VALUES
-- ('','','','','');