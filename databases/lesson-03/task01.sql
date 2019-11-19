-- Пересоздадим базу.
DROP DATABASE IF EXISTS `gsy_lesson03_task01`;
CREATE DATABASE `gsy_lesson03_task01` 
    CHARACTER SET `utf8` 
    COLLATE `utf8_unicode_ci`;

-- Установим ее по умолчанию.
USE `gsy_lesson03_task01`;

-- Создадим таблицу.
CREATE TABLE `users` (
  `id`          SERIAL,
  `name`        VARCHAR(255),
  `created_at`  DATETIME,
  `updated_at`  DATETIME,
  PRIMARY KEY (`id`)
);

-- Заполним таблицу.
INSERT INTO `users` (`name`) VALUES
    ('Иванов Иван'),
    ('Борисов Борис'),
    ('Антонов Антон'),
    ('Петров Пётр');

-- Отобразим результат.
SELECT * FROM `users`;

-- Дозаполним поля 'updated_at' и 'created_at'.
UPDATE `users` SET
    updated_at = NOW(),
    created_at = NOW();

-- Отобразим результат.
SELECT * FROM `users`;