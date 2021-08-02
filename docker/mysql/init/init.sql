SET
  @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS,
  UNIQUE_CHECKS = 0;
SET
  @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS,
  FOREIGN_KEY_CHECKS = 0;
SET
  @OLD_SQL_MODE = @@SQL_MODE,
  SQL_MODE = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `swipe_shukatu` DEFAULT CHARACTER SET utf8mb4;

USE `swipe_shukatu`;
SET
  CHARSET utf8mb4;

-- -----------------------------------------------------
-- user
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL COMMENT 'firebase uuid',
  `rand_id` varchar(255) NOT NULL COMMENT 'セキュリティのためのハッシュ化されたID',
  `registered_method_id` tinyint(3) NULL COMMENT '登録方法',
  `gender` enum('M', 'F', 'X') NULL COMMENT '性別',
  `graduate_year` DATETIME NULL COMMENT '卒業年次',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '退会日時',
  PRIMARY KEY (`id`),
  INDEX `idx_uuid` (`uuid` ASC),
  INDEX `idx_rand_id` (`rand_id` ASC)
) ENGINE = InnoDB COMMENT = 'ユーザ';

-- -----------------------------------------------------
-- ltd
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`ltd` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '企業名',
  profile` text(500) NULL COMMENT '企業紹介',
  `employee_number` int(11) unsigned NULL COMMENT '従業員数',
  `job_type_id` smallint(5) unsigned NULL COMMENT '職種ID',
  `educational_background_id` smallint(5) unsigned NULL COMMENT '必要な最終学歴ID',
  `average_salary` int(11) unsigned NULL COMMENT '平均年収',
  `average_age` tinyint(3) unsigned NULL COMMENT '平均年齢',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '退会日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '企業';

-- -----------------------------------------------------
-- ltd_image
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`ltd_image` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `image_path` varchar(255) NOT NULL COMMENT 'プロフィール画像保存先',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '企業プロフィール画像';

-- -----------------------------------------------------
-- ltd_benefit
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`ltd_benefit` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `benefit_id` smallint(11) unsigned NOT NULL COMMENT '福利厚生ID',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '退会日時',
  PRIMARY KEY (`id`),
  INDEX `idx_ltd_id` (`ltd_id` ASC)
) ENGINE = InnoDB COMMENT = '企業の福利厚生';

-- -----------------------------------------------------
-- ltd_location
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`ltd_location` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `location_id` smallint(11) unsigned NOT NULL COMMENT '所在地ID',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '退会日時',
  PRIMARY KEY (`id`),
  INDEX `idx_ltd_id` (`ltd_id` ASC)
) ENGINE = InnoDB COMMENT = '企業の所在地';

-- -----------------------------------------------------
-- like
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`like` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  PRIMARY KEY (`user_id`, `ltd_id`
) ENGINE = InnoDB COMMENT = 'いいね';

-- -----------------------------------------------------
-- nope
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`nope` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  PRIMARY KEY (`user_id`, `ltd_id`
) ENGINE = InnoDB COMMENT = 'いまいち';

-- -----------------------------------------------------
-- superlike
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`superlike` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  PRIMARY KEY (`user_id`, `ltd_id`)
) ENGINE = InnoDB COMMENT = '最高';

-- -----------------------------------------------------
-- message
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `type` tinyint(3) unsigned NOT NULL COMMENT 'メッセージタイプ(1->text, 2->remind, 3->image)',
  `content` text(500) NULL COMMENT 'メッセージ内容',
  `image_path` varchar(255) NOT NULL COMMENT '画像保存先',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id` ASC),
  INDEX `idx_ltd_id` (`ltd_id` ASC)
) ENGINE = InnoDB COMMENT = 'メッセージ';

-- -----------------------------------------------------
-- remind_message
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`remind_message` (
  `message_id` bigint(20) NOT NULL COMMENT 'メッセージID',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `execute_at` DATETIME NULL COMMENT '実行日時',
  PRIMARY KEY (`message_id`)
) ENGINE = InnoDB COMMENT = 'リマインドメッセージ';

-- -----------------------------------------------------
-- ### MASTER DATA ###
-- -----------------------------------------------------

-- -----------------------------------------------------
-- m_ltd_official_message
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_ltd_official_message` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `type` tinyint(3) unsigned NOT NULL COMMENT 'メッセージタイプ(1->Text, 2->Image)',
  `content` text(500) NOT NULL COMMENT 'メッセージ内容',
  `image_path` varchar(255) NOT NULL COMMENT '画像保存先',
  `created_at` DATETIME NULL COMMENT '登録日時',
  `updated_at` DATETIME NULL COMMENT '更新日時',
  `is_invalid_at` DATETIME NULL COMMENT '無効日時',
  PRIMARY KEY (`user_id`, `ltd_id`)
) ENGINE = InnoDB COMMENT = '企業の自動送信メッセージ';

-- -----------------------------------------------------
-- m_registered_method
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_registered_method` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '登録方法',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = 'ログイン方法';

-- -----------------------------------------------------
-- m_job_type
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_job_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '職種',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '職種';

-- -----------------------------------------------------
-- m_educational_background
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_educational_background` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '学歴',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '学歴';

-- -----------------------------------------------------
-- m_benefit
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_benefit` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '福利厚生',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '福利厚生';

-- -----------------------------------------------------
-- m_location
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_location` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '所在地',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '所在地';

SET
  SQL_MODE = @OLD_SQL_MODE;
SET
  FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS;
SET
  UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS;
