SET NAMES `utf8` COLLATE `utf8_unicode_ci`;

ALTER DATABASE `gsy_media_files`
DEFAULT CHARACTER SET utf8
DEFAULT COLLATE utf8_unicode_ci;



DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id`          SERIAL,
  `name`        VARCHAR(255) NOT NULL COMMENT 'Имя пользователя',
  PRIMARY KEY (`id`),
  UNIQUE `index_of_name` (`name` (30))
) COMMENT = 'Пользователи';

INSERT INTO `users` VALUES
    (DEFAULT, 'Иванов Иван'),
    (DEFAULT, 'Петров Пётр'),
    (DEFAULT, 'Алексеев Алексей');



DROP TABLE IF EXISTS `file_types`;
CREATE TABLE `file_types` (
  `id`          SERIAL,
  `name`        VARCHAR(255) NOT NULL COMMENT 'Имя типа',
  PRIMARY KEY (`id`),
  UNIQUE `index_of_name` (`name`)
) COMMENT = 'Типы файлов';

INSERT INTO `file_types` VALUES
    (DEFAULT, 'Фото'),
    (DEFAULT, 'Аудио'),
    (DEFAULT, 'Видео');



DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id`          SERIAL,
  `name`        VARCHAR(255) NOT NULL       COMMENT 'Имя файла',
  `description` TEXT                        COMMENT 'Описание файла',
  `path`        VARCHAR(255) NOT NULL       COMMENT 'Путь к файлу',
  `type_id`     BIGINT UNSIGNED NOT NULL    COMMENT 'Код типа файла',
  `user_id`     BIGINT UNSIGNED NOT NULL    COMMENT 'Код владельца файла',
  PRIMARY KEY (`id`),
  KEY `index_of_type` (`type_id`),
  KEY `index_of_user` (`user_id`)
) COMMENT = 'Медиафайлы';


INSERT INTO `files` (`id`, `name`, `path`, `type_id`, `user_id`, `description`) VALUES
    (DEFAULT, 'Фото-001', '/files/photo/001.jpg', 1, 1, 'Описание для фотографии'),
    (DEFAULT, 'Аудио-001', '/files/audio/001.mp3', 2, 2, 'Описание для аудиозаписи'),
    (DEFAULT, 'Видео-001', '/files/video/001.mp4', 3, 3, 'Описание для видеозаписи');
