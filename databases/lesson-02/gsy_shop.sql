/*
    Это учебная база данных магазина.
    Здесь я отрабатываю:
    - построение различных таблиц;
    - использование разных типов полей;
    - настройки колючевых полей;
    - настройки связей таблиц;
    - выборки, сортировки, фильтрации;
*/


-- Это поможет добавлять значения в текстовые поля.
-- Иначе mysql будет сообщать об ошибке:
-- ERROR 1366 (HY000): Incorrect string value: ...
SET NAMES `utf8` COLLATE `utf8_unicode_ci`;


DROP TABLE IF EXISTS `catalogs`;
CREATE TABLE `catalogs` (
  `id`      SERIAL,
  `name`    VARCHAR(255) DEFAULT 'NO NAME' COMMENT 'Наименование раздела',
  PRIMARY KEY (`id`),
  UNIQUE `unique_name` (`name` (10)) -- Индексируем первые 10 символов 
) COMMENT = 'Разделы интернет-магазина';

INSERT IGNORE INTO `catalogs` (`id`, `name`) VALUES 
    (DEFAULT, 'Процессоры'),
    (DEFAULT, 'Мат.платы'),
    (DEFAULT, 'Видеокарты'),
    (DEFAULT, 'Видеокарты');
SELECT 'TABLE catalogs CREATED' AS 'SUCCESS';



DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id`          SERIAL,
  `name`        VARCHAR(255) NOT NULL   COMMENT 'Имя пользователя',
  `birthday`    DATE                    COMMENT 'День рождения',
  `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) COMMENT = 'Пользователи';

INSERT INTO `users` (`id`, `name`, `birthday`) VALUES (0, 'Администратор', '2000-01-01');
SELECT 'TABLE users CREATED' AS 'SUCCESS';



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

-- Создание индекса в уже существующей таблице:
-- CREATE INDEX `index_of_catalog_id` ON `products` (`ctalog_id`);
-- CREATE INDEX `index_of_catalog_id` USING BTREE ON `products` (`ctalog_id`);
-- CREATE INDEX `index_of_catalog_id` USING HASH ON `products` (`ctalog_id`);
-- DROP INSERT `index_of_catalog_id` ON `products`;
SELECT 'TABLE products CREATED' AS 'SUCCESS';



DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id`          SERIAL,
  `user_id`     INT UNSIGNED COMMENT 'Покупатель',
  `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY index_of_user_id(`user_id`)
) COMMENT = 'Заказы';

SELECT 'TABLE orders CREATED' AS 'SUCCESS';



DROP TABLE IF EXISTS `orders_products`;
CREATE TABLE `orders_products` (
  `id`          SERIAL,
  `order_id`    INT UNSIGNED NOT NULL           COMMENT 'Заказ',
  `product_id`  INT UNSIGNED NOT NULL           COMMENT 'Товар',
  `total`       INT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'Количество единиц товара',
  `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) COMMENT = 'Товары в заказах';

SELECT 'TABLE orders_products CREATED' AS 'SUCCESS';



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



DROP TABLE IF EXISTS `storehouses`;
CREATE TABLE `storehouses` (
  `id`          SERIAL,
  `name`        VARCHAR(255) DEFAULT 'NO NAME' COMMENT 'Наименование склада',
  `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) COMMENT = 'Склады';

SELECT 'TABLE storehouses CREATED' AS 'SUCCESS';



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
