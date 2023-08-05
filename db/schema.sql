-- DBスキーマ初期化用スクリプト
-- Cloud SQL では以下の手順でローカルから実行できる
-- 1. まだの場合、mysql-clientをインストール: sudo apt-get install mysql-client
-- 2. mysql-client 経由でDBと接続: gcloud sql connect go-server-poc --user=root
-- 3. DB.sql を実行: source ./db/schema.sql

DROP DATABASE IF EXISTS `go-server-poc`;
CREATE DATABASE `go-server-poc`;

USE `go-server-poc`;

CREATE TABLE IF NOT EXISTS `user` (
  `user_id` CHAR(36) NOT NULL,
  `user_name` VARCHAR(255) UNIQUE NOT NULL,
  `country_id` SMALLINT,
  `banned` BOOLEAN NOT NULL DEFAULT 0,
  `resigned` BOOLEAN NOT NULL DEFAULT 0,
  `signin_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
);

CREATE TABLE IF NOT EXISTS `health` (
		`health_id` INT NOT NULL AUTO_INCREMENT,
		`message` VARCHAR(255) NOT NULL,
		`updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (`health_id`)
);
