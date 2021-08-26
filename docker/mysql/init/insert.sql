use `swipe_shukatu`;

SET NAMES utf8mb4;

INSERT INTO `m_location` (`id`, `name`, `sort_id`, `created_at`, `updated_at`) VALUES
(1, '北海道', 1, NOW(), NOW()),
(2, '青森県', 2, NOW(), NOW()),
(3, '岩手県', 3, NOW(), NOW()),
(4, '宮城県', 4, NOW(), NOW()),
(5, '秋田県', 5, NOW(), NOW()),
(6, '山形県', 6, NOW(), NOW()),
(7, '福島県', 7, NOW(), NOW()),
(8, '茨城県', 8, NOW(), NOW()),
(9, '栃木県', 9, NOW(), NOW()),
(10, '群馬県', 10, NOW(), NOW()),
(11, '埼玉県', 11, NOW(), NOW()),
(12, '千葉県', 12, NOW(), NOW()),
(13, '東京都', 13, NOW(), NOW()),
(14, '神奈川県', 14, NOW(), NOW()),
(15, '新潟県', 15, NOW(), NOW()),
(16, '富山県', 16, NOW(), NOW()),
(17, '石川県', 17, NOW(), NOW()),
(18, '福井県', 18, NOW(), NOW()),
(19, '山梨県', 19, NOW(), NOW()),
(20, '長野県', 20, NOW(), NOW()),
(21, '岐阜県', 21, NOW(), NOW()),
(22, '静岡県', 22, NOW(), NOW()),
(23, '愛知県', 23, NOW(), NOW()),
(24, '三重県', 24, NOW(), NOW()),
(25, '滋賀県', 25, NOW(), NOW()),
(26, '京都府', 26, NOW(), NOW()),
(27, '大阪府', 27, NOW(), NOW()),
(28, '兵庫県', 28, NOW(), NOW()),
(29, '奈良県', 29, NOW(), NOW()),
(30, '和歌山県', 30, NOW(), NOW()),
(31, '鳥取県', 31, NOW(), NOW()),
(32, '島根県', 32, NOW(), NOW()),
(33, '岡山県', 33, NOW(), NOW()),
(34, '広島県', 34, NOW(), NOW()),
(35, '山口県', 35, NOW(), NOW()),
(36, '徳島県', 36, NOW(), NOW()),
(37, '香川県', 37, NOW(), NOW()),
(38, '愛媛県', 38, NOW(), NOW()),
(39, '高知県', 39, NOW(), NOW()),
(40, '福岡県', 40, NOW(), NOW()),
(41, '佐賀県', 41, NOW(), NOW()),
(42, '長崎県', 42, NOW(), NOW()),
(43, '熊本県', 43, NOW(), NOW()),
(44, '大分県', 44, NOW(), NOW()),
(45, '宮崎県', 45, NOW(), NOW()),
(46, '鹿児島県', 46, NOW(), NOW()),
(47, '沖縄県', 47, NOW(), NOW()),
(48, '海外', 48, NOW(), NOW());
INSERT INTO `m_educational_background` (`id`, `name`, `sort_id`, `created_at`, `updated_at`) VALUES
(1, '中卒', 1, NOW(), NOW()),
(2, '高卒', 2, NOW(), NOW()),
(3, '専門卒', 3, NOW(), NOW()),
(4, '大卒', 4, NOW(), NOW()),
(5, '院卒', 5, NOW(), NOW());
INSERT INTO `m_registered_method` (`id`, `name`, `sort_id`, `created_at`, `updated_at`) VALUES
(1, 'Google', 1, NOW(), NOW());
INSERT INTO `m_industry` (`id`, `sort_id`, `name`, `created_at`, `updated_at`) VALUES
(1, 1, '流通・小売', NOW(), NOW()),
(2, 2, '不動産', NOW(), NOW()),
(3, 3, '専門店', NOW(), NOW()),
(4, 4, '食品', NOW(), NOW()),
(5, 5, '商社', NOW(), NOW()),
(6, 6, '住宅', NOW(), NOW()),
(7, 7, '広告・出版・マスコミ', NOW(), NOW()),
(8, 8, '金融', NOW(), NOW()),
(9, 9, '官公庁・公社・団体', NOW(), NOW()),
(10, 10, '化粧品', NOW(), NOW()),
(11, 11, 'メーカー', NOW(), NOW()),
(12, 12, 'ソフトウェア・通信', NOW(), NOW()),
(13, 13, 'サービス・インフラ', NOW(), NOW());
INSERT INTO `m_job_type` (`id`, `sort_id`, `name`, `created_at`, `updated_at`) VALUES
(1, 1, '総合職', NOW(), NOW()),
(2, 2, '事務職', NOW(), NOW()),
(3, 3, '技術職', NOW(), NOW());
