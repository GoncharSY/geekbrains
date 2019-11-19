-- Установим БД по умолчанию.
USE `db_example2`;

-- Изменим кодировку в БД.
ALTER DATABASE `db_example2`
DEFAULT CHARACTER SET `utf8`
DEFAULT COLLATE `utf8_unicode_ci`;

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