SET
  @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS,
  UNIQUE_CHECKS = 0;
SET
  @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS,
  FOREIGN_KEY_CHECKS = 0;
SET
  @OLD_SQL_MODE = @@SQL_MODE,
  SQL_MODE = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET GLOBAL max_allowed_packet=100000000;

CREATE SCHEMA IF NOT EXISTS `swipe_shukatu` DEFAULT CHARACTER SET utf8mb4;

SET
  CHARSET utf8mb4;

-- -----------------------------------------------------
-- user
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`user` (
  `uuid` varchar(255) NOT NULL COMMENT 'firebase uuid',
  `registered_method_id` tinyint(3) NULL COMMENT '登録方法',
  `gender` tinyint(3) NULL COMMENT '性別(1->M, 2->F, 3->X)',
  `graduate_year` tinyint(3) NULL COMMENT '卒業年次',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '退会日時',
  PRIMARY KEY (`uuid`)
) ENGINE = InnoDB COMMENT = 'ユーザ';

-- -----------------------------------------------------
-- recruit
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`recruit` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID', 
  `job_type_id` smallint(5) unsigned NULL COMMENT '職種ID',
  `educational_background_id` smallint(5) unsigned NULL COMMENT '必要な最終学歴ID',
  `average_salary` int(11) unsigned NULL COMMENT '平均年収',
  `starting_salary` int(11) unsigned NULL COMMENT '初任給', 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `is_invalid_at` DATETIME NULL COMMENT '無効日時', 
  PRIMARY KEY (`id`), 
  INDEX `idx_ltd_id` (`ltd_id` ASC)
) ENGINE = InnoDB COMMENT = '求人';

-- -----------------------------------------------------
-- ltd_image
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`ltd_image` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `image_path` varchar(255) NOT NULL COMMENT 'プロフィール画像保存先',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順', 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`), 
  INDEX `idx_id` (`ltd_id` ASC)
) ENGINE = InnoDB COMMENT = '企業プロフィール画像';

-- -----------------------------------------------------
-- ltd_benefit
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`ltd_benefit` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ltd_id` bigint(20) unsigned NOT NULL COMMENT '企業ID',
  `benefit_id` smallint(5) unsigned NOT NULL COMMENT '福利厚生ID',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順', 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
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
  `location_id` smallint(5) unsigned NOT NULL COMMENT '所在地ID',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順', 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '退会日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '企業の所在地';

-- -----------------------------------------------------
-- like
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`like` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `recruit_id` bigint(20) unsigned NOT NULL COMMENT '求人ID', 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`user_id`, `recruit_id`)
) ENGINE = InnoDB COMMENT = 'いいね';

-- -----------------------------------------------------
-- nope
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`nope` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `recruit_id` bigint(20) unsigned NOT NULL COMMENT '求人ID', 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`user_id`, `recruit_id`)
) ENGINE = InnoDB COMMENT = 'いまいち';

-- -----------------------------------------------------
-- superlike
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`superlike` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `recruit_id` bigint(20) unsigned NOT NULL COMMENT '求人ID', 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`user_id`, `recruit_id`)
) ENGINE = InnoDB COMMENT = '最高';

-- -----------------------------------------------------
-- message
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
  `recruit_id` bigint(20) unsigned NOT NULL COMMENT '求人ID', 
  `type` tinyint(3) unsigned NOT NULL COMMENT 'メッセージタイプ(1->text, 2->remind, 3->image)',
  `content` text NULL COMMENT 'メッセージ内容',
  `image_path` varchar(255) NOT NULL COMMENT '画像保存先',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`),
  INDEX `idx_recruit_id_user_id` (`recruit_id`, `user_id`)
) ENGINE = InnoDB COMMENT = 'メッセージ';

-- -----------------------------------------------------
-- remind_message
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`remind_message` (
  `message_id` bigint(20) NOT NULL COMMENT 'メッセージID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `execute_at` DATETIME NULL COMMENT '実行日時',
  PRIMARY KEY (`message_id`)
) ENGINE = InnoDB COMMENT = 'リマインドメッセージ';

-- -----------------------------------------------------
-- ### MASTER DATA ###
-- -----------------------------------------------------

-- -----------------------------------------------------
-- ltd
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_ltd` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '企業名',
  `profile` text NULL COMMENT '企業紹介',
  `employee_number` int(11) unsigned NULL COMMENT '従業員数',
  `average_age` tinyint(3) unsigned NULL COMMENT '平均年齢',
  `industry_id` smallint(5) NULL COMMENT '業種ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '退会日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '企業';

-- -----------------------------------------------------
-- m_ltd_official_message
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_ltd_official_message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `recruit_id` bigint(20) unsigned NOT NULL COMMENT '求人ID',
  `type` tinyint(3) unsigned NOT NULL COMMENT 'メッセージタイプ(1->Text, 2->Image)',
  `content` text NOT NULL COMMENT 'メッセージ内容',
  `image_path` varchar(255) NOT NULL COMMENT '画像保存先',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `is_invalid_at` DATETIME NULL COMMENT '無効日時',
  PRIMARY KEY (`id`),
  INDEX `idx_recruit_id` (`recruit_id` ASC)
) ENGINE = InnoDB COMMENT = '企業の自動送信メッセージ';

-- -----------------------------------------------------
-- m_registered_method
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_registered_method` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '登録方法',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = 'ログイン方法';

-- -----------------------------------------------------
-- m_job_type
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_job_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '職種',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '職種';

-- -----------------------------------------------------
-- m_educational_background
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_educational_background` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '学歴',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '学歴';

-- -----------------------------------------------------
-- m_benefit
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_benefit` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '福利厚生',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '福利厚生';

-- -----------------------------------------------------
-- m_location
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_location` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '所在地',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '所在地';

-- -----------------------------------------------------
-- m_industry
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `swipe_shukatu`.`m_industry` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '業種',
  `sort_id` smallint(5) unsigned NOT NULL COMMENT '表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '業種';

SET
  SQL_MODE = @OLD_SQL_MODE;
SET
  FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS;
SET
  UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS;
