/*
(по желанию) Есть таблица (accounts), включающая в себя три столбца: 
id, name, password, которые содержат первичный ключ, имя пользователя 
и его пароль. 

Создайте представление username таблицы accounts, предоставляющее 
доступ к столбцам id и name. 

Создайте пользователя user_read, который бы не имел доступа к таблице 
accounts, однако мог извлекать записи из представления username.
*/

-- Создадим базу данных заново.
CREATE DATABASE `db_example` CHARACTER SET `utf8` COLLATE `utf8_unicode_ci`;

-- Установим ее по умолчанию.
USE `db_example`;

-- Создадим в БД таблицу.
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id`      SERIAL,
  `name`    VARCHAR(255),
  `pass`    VARCHAR(255),
  PRIMARY KEY (`id`)
) COMMENT = 'Учетные записи пользователей';

-- Наполним таблицу данными.
INSERT INTO `accounts` VALUES
    (DEFAULT, 'User-01', 'Pass-Of-User-01'),
    (DEFAULT, 'User-02', 'Pass-Of-User-02'),
    (DEFAULT, 'User-03', 'Pass-Of-User-03');

-- Создадим вертикальное представление таблицы.
CREATE OR REPLACE VIEW `username` AS SELECT `id`, `name` FROM `accounts`;

-- Создадим пользователя.
DROP USER IF EXISTS 'user_read';
CREATE USER IF NOT EXISTS 'user_read';

-- Устанавливаю права.
GRANT SELECT ON `db_example`.`username` TO 'user_read';

-- Удаляю пользователя.
-- DROP USER IF EXISTS 'user_read';
