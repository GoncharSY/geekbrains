-- Снесем базу, если такая была.
DROP DATABASE IF EXISTS `db_example`;

-- Создадим базу данных заново.
CREATE DATABASE `db_example` 
    CHARACTER SET `utf8` 
    COLLATE `utf8_unicode_ci`;

-- Установим ее по умолчанию.
USE `db_example`;

-- Создадим в БД таблицу.
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id`      SERIAL,
  `name`    VARCHAR(255),
  PRIMARY KEY (`id`)
) COMMENT = 'Пользователи';

-- Наполним таблицу данными.
INSERT INTO `users` VALUES
    (DEFAULT, 'Иванов Иван'),
    (DEFAULT, 'Петров Пётр'),
    (DEFAULT, 'Алексеев Алексей');