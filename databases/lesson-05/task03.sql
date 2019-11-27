/*
    (по желанию) Есть таблица рейсов flights (id, from, to) и таблица 
    городов cities (label, name). Поля from, to и label содержат английские 
    названия городов, поле name — русское. Выведите список рейсов (flights) 
    с русскими названиями городов.
*/

/* ========================================================================= */
/* БАЗА ==================================================================== */
/* ========================================================================= */

-- Снесем базу, если такая была.
DROP DATABASE IF EXISTS `gsy_lesson05_task03`;

-- Создадим базу данных заново.
CREATE DATABASE `gsy_lesson05_task03` 
    CHARACTER SET `utf8` 
    COLLATE `utf8_unicode_ci`;

-- Установим ее по умолчанию.
USE `gsy_lesson05_task03`;

/* ========================================================================= */
/* ТАБЛИЦЫ ================================================================= */
/* ========================================================================= */

-- Создадим в БД таблицу городов.
DROP TABLE IF EXISTS `cities`;
CREATE TABLE `cities` (
  `label`   VARCHAR(3),
  `name`    VARCHAR(25),
  PRIMARY KEY (`label`)
) COMMENT = 'Города';

-- Наполним таблицу данными.
INSERT INTO `cities` VALUES
    ('msk', 'Москва'),
    ('irk', 'Иркутск'),
    ('nov', 'Новгород'),
    ('kaz', 'Казань'),
    ('oms', 'Омск');

-- Создадим в БД таблицу авиарейсов.
DROP TABLE IF EXISTS `flights`;
CREATE TABLE `flights` (
    `id`    SERIAL,
    `from`  VARCHAR(3),
    `to`    VARCHAR(3),
    PRIMARY KEY (`id`),
    CONSTRAINT FOREIGN KEY `from_city` (`from`)
        REFERENCES `cities` (`label`)
        ON DELETE RESTRICT
        ON UPDATE RESTRICT,
    CONSTRAINT FOREIGN KEY `to_city` (`to`)
        REFERENCES `cities` (`label`)
        ON DELETE RESTRICT
        ON UPDATE RESTRICT
) COMMENT = 'Авиарейсы';

-- Наполним таблицу данными.
INSERT INTO `flights` VALUES
    (DEFAULT, 'msk', 'oms'),
    (DEFAULT, 'nov', 'kaz'),
    (DEFAULT, 'irk', 'msk'),
    (DEFAULT, 'oms', 'irk'),
    (DEFAULT, 'msk', 'kaz');

/* ========================================================================= */
/* ОТЧЕТЫ ================================================================== */
/* ========================================================================= */

-- Решение:
SELECT 
	flt.id,
	cf.name AS 'From city',
	ct.name AS 'To city'
FROM flights AS flt
LEFT JOIN cities AS cf ON cf.label = flt.`from`
LEFT JOIN cities AS ct ON ct.label = flt.`to`;