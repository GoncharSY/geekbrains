-- -- Пересоздадим базу.
DROP DATABASE IF EXISTS `gsy_lesson03_task03`;
CREATE DATABASE `gsy_lesson03_task03` 
    CHARACTER SET `utf8` 
    COLLATE `utf8_unicode_ci`;

-- Установим ее по умолчанию.
USE `gsy_lesson03_task03`;

-- Создадим таблицу.
DROP TABLE IF EXISTS `storehouses_products`;
CREATE TABLE `storehouses_products` (
  `id`              SERIAL,
  `storehouse_id`   INT UNSIGNED NOT NULL           COMMENT 'Склад',
  `product_id`      INT UNSIGNED NOT NULL           COMMENT 'Товар',
  `total`           INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Количество единиц товара',
  `created_at`      DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_of_storehouse_id` (`storehouse_id`),
  KEY `index_of_product_id` (`product_id`)
) COMMENT = 'Товары на складах';

-- Наполним таблицу данными.
INSERT INTO `storehouses_products` (
    `storehouse_id`, 
    `product_id`, 
    `total`
) VALUES
    (1, 1, 1),
    (1, 2, 7),
    (1, 3, 0),
    (1, 4, 12),
    (1, 5, 0),
    (1, 6, 11);

-- Отобразим содержимое заданном порядке.
SELECT *
FROM `storehouses_products`
ORDER BY `total` = 0, `total`;