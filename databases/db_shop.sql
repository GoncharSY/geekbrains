/*
    Это учебная база данных магазина.
    Здесь я отрабатываю:
    - построение различных таблиц;
    - использование разных типов полей;
    - настройки колючевых полей;
    - настройки связей таблиц;
    - выборки, сортировки, фильтрации;
*/

/* ============================================================================================= */
/* СОЗДАНИЕ БАЗЫ ДАННЫХ ======================================================================== */
/* ============================================================================================= */

-- Снесем базу, если такая была.
DROP DATABASE IF EXISTS `gsy_shop`;

-- Создадим базу данных заново.
CREATE DATABASE `gsy_shop` 
    CHARACTER SET `utf8` 
    COLLATE `utf8_unicode_ci`;

-- Установим ее по умолчанию.
USE `gsy_shop`;


/* ============================================================================================= */
/* СОЗДАНИЕ И ЗАПОЛНЕНИЕ ТАБЛИЦ ================================================================ */
/* ============================================================================================= */

DROP TABLE IF EXISTS `catalogs`;
CREATE TABLE `catalogs` (
    `id`      SERIAL,                               -- bigint(20) unsigned NOT NULL AUTO_INCREMENT
    `name`    VARCHAR(255) DEFAULT 'NO NAME' COMMENT 'Наименование раздела',
    PRIMARY KEY (`id`),                             -- Первичный ключ
    UNIQUE `unique_name` (`name` (10))              -- Индексируем первые 10 символов 
) COMMENT = 'Разделы интернет-магазина';

INSERT IGNORE INTO `catalogs` (`id`, `name`) VALUES 
    (DEFAULT, 'Процессоры'),
    (DEFAULT, 'Материнские платы'),
    (DEFAULT, 'Видеокарты'),
    (DEFAULT, 'Жесткие диски (HDD)'),
    (DEFAULT, 'Твердотельные накопители памяти (SSD)');

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id`          SERIAL,
    `name`        VARCHAR(255) NOT NULL   COMMENT 'Имя пользователя',
    `birthday`    DATE                    COMMENT 'День рождения',
    `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) COMMENT = 'Пользователи';

INSERT INTO `users` (`id`, `name`, `birthday`) VALUES
    (0, 'Администратор', '1991-01-01'),   -- 0 = DEFAULT when SERIAL/AUTO_INCREMENT 
    (0, 'Петров Пётр', '1995-01-01'), 
    (0, 'Антонов Антон', '1997-01-01'), 
    (0, 'Сергеев Сергей', '2000-01-01'), 
    (0, 'Борисов Борис', '2001-01-01');

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
    `id`          SERIAL,
    `name`        VARCHAR(255) DEFAULT 'NO NAME'  COMMENT 'Название товара',
    `description` TEXT                            COMMENT 'Описание товара',
    `price`       DECIMAL(11,2)                   COMMENT 'Цена товара',
    `catalog_id`  INT UNSIGNED                    COMMENT 'Раздел',
    `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY index_of_catalog_id(`catalog_id`)
) COMMENT = 'Товары';

INSERT INTO `products` (
    `id`, 
    `name`, 
    `description`,
    `price`,
    `catalog_id`,
    `created_at`,
    `updated_at`
) VALUES
    (DEFAULT, 'Модель-001', 'Описание для модели товара 001', 100.00, 1, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-002', 'Описание для модели товара 002', 130.00, 1, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-003', 'Описание для модели товара 003', 170.00, 1, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-004', 'Описание для модели товара 004', 200.00, 1, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-005', 'Описание для модели товара 005', 140.00, 1, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-001', 'Описание для модели товара 001', 100.00, 2, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-002', 'Описание для модели товара 002', 130.00, 2, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-003', 'Описание для модели товара 003', 170.00, 2, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-004', 'Описание для модели товара 004', 200.00, 2, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-005', 'Описание для модели товара 005', 140.00, 2, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-001', 'Описание для модели товара 001', 100.00, 3, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-002', 'Описание для модели товара 002', 130.00, 3, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-003', 'Описание для модели товара 003', 170.00, 3, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-004', 'Описание для модели товара 004', 200.00, 3, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-005', 'Описание для модели товара 005', 140.00, 3, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-001', 'Описание для модели товара 001', 100.00, 4, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-002', 'Описание для модели товара 002', 130.00, 4, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-003', 'Описание для модели товара 003', 170.00, 4, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-004', 'Описание для модели товара 004', 200.00, 4, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-005', 'Описание для модели товара 005', 140.00, 4, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-001', 'Описание для модели товара 001', 100.00, 5, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-002', 'Описание для модели товара 002', 130.00, 5, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-003', 'Описание для модели товара 003', 170.00, 5, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-004', 'Описание для модели товара 004', 200.00, 5, DEFAULT, DEFAULT),
    (DEFAULT, 'Модель-005', 'Описание для модели товара 005', 140.00, 5, DEFAULT, DEFAULT);

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
    `id`          SERIAL,
    `user_id`     INT UNSIGNED COMMENT 'Покупатель',
    `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY index_of_user_id(`user_id`)
) COMMENT = 'Заказы';

INSERT INTO `catalogs` (`id`, `user_id`, `created_at`, `updated_at`) VALUES 
    (DEFAULT, 1, DEFAULT, DEFAULT),
    (DEFAULT, 2, DEFAULT, DEFAULT),
    (DEFAULT, 3, DEFAULT, DEFAULT),
    (DEFAULT, 4, DEFAULT, DEFAULT),
    (DEFAULT, 5, DEFAULT, DEFAULT);

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */



DROP TABLE IF EXISTS `orders_products`;
CREATE TABLE `orders_products` (
    `id`          SERIAL,
    `order_id`    INT UNSIGNED NOT NULL           COMMENT 'Заказ',
    `product_id`  INT UNSIGNED NOT NULL           COMMENT 'Товар',
    `total`       INT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'Количество единиц товара',
    `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY index_of_order_product (`order_id`, `product_id`) -- Индексируем по двум полям
) COMMENT = 'Товары в заказах';

INSERT INTO `orders_products` (
    `order_id`,
    `product_id`,
    `total`,
    `created_at`,
    `updated_at`,
) VALUES 
    (DEFAULT, 1, 1, 1, DEFAULT, DEFAULT),
    (DEFAULT, 2, 2, 1, DEFAULT, DEFAULT),
    (DEFAULT, 3, 3, 1, DEFAULT, DEFAULT),
    (DEFAULT, 4, 4, 1, DEFAULT, DEFAULT),
    (DEFAULT, 5, 5, 1, DEFAULT, DEFAULT);

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */



DROP TABLE IF EXISTS `discounts`;
CREATE TABLE `discounts` (
  `id`          SERIAL,
  `user_id`     INT UNSIGNED    COMMENT 'Покупатель',
  `product_id`  INT UNSIGNED    COMMENT 'Товар',
  `discount`    FLOAT UNSIGNED  COMMENT 'Доля скидки',
  `started_at`  DATETIME        COMMENT 'Начало действия',
  `finished_at` DATETIME        COMMENT 'Окончание действия',
  `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_of_user_id` (`user_id`),
  KEY `index_of_product_id` (`product_id`)
) COMMENT = 'Скидки';

SELECT 'TABLE discounts CREATED' AS 'SUCCESS';

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */



DROP TABLE IF EXISTS `storehouses`;
CREATE TABLE `storehouses` (
  `id`          SERIAL,
  `name`        VARCHAR(255) DEFAULT 'NO NAME' COMMENT 'Наименование склада',
  `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) COMMENT = 'Склады';

SELECT 'TABLE storehouses CREATED' AS 'SUCCESS';

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */



DROP TABLE IF EXISTS `storehouses_products`;
CREATE TABLE `storehouses_products` (
  `id`              SERIAL,
  `storehouse_id`  INT UNSIGNED NOT NULL           COMMENT 'Склад',
  `product_id`      INT UNSIGNED NOT NULL           COMMENT 'Товар',
  `total`           INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Количество единиц товара',
  `created_at`      DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_of_storehouse_id` (`storehouse_id`),
  KEY `index_of_product_id` (`product_id`)
) COMMENT = 'Товары на складах';

SELECT 'TABLE storehouses_products CREATED' AS 'SUCCESS';

/* ============================================================================================= */
/* ============================================================================================= */
/* ============================================================================================= */
