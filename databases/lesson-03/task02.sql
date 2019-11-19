-- -- Пересоздадим базу.
-- DROP DATABASE IF EXISTS `gsy_lesson03_task02`;
-- CREATE DATABASE `gsy_lesson03_task02` 
--     CHARACTER SET `utf8` 
--     COLLATE `utf8_unicode_ci`;

-- Установим ее по умолчанию.
USE `gsy_lesson03_task02`;

-- Создадим таблицу.
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id`          SERIAL,
  `name`        VARCHAR(255),
  `created_at`  VARCHAR(20),
  `updated_at`  VARCHAR(20),
  PRIMARY KEY (`id`)
);
-- Заполним таблицу.
INSERT INTO `users` (`name`, `created_at`, `updated_at`) VALUES
    ('Иванов Иван', '1.10.2017 8:10', '2.10.2017 8:10'),
    ('Борисов Борис', '9.10.2017 9:10', '11.10.2017 9:10'),
    ('Антонов Антон', '10.10.2017 10:10', '19.10.2017 10:10'),
    ('Петров Пётр', '20.10.2017 11:10', '22.10.2017 11:10');

-- Отобразим исходное описание таблицы.
DESCRIBE `users`;

-- Отобразим исходное содержимое таблицы.
SELECT * FROM `users`;



-- Заполним поля, но в правильном формате.
UPDATE `users` AS u1 
INNER JOIN (
    SELECT 
        `id` AS 'id', 
        @tmp_date := SUBSTRING_INDEX(`created_at`, ' ', 1)              AS 'c_date',
        @tmp_ddmm := SUBSTRING_INDEX(@tmp_date, '.', 2)                 AS 'c_ddmm',
        @tmp_yyyy := SUBSTRING_INDEX(@tmp_date, '.', -1)                AS 'c_yyyy',
        @tmp_mm := LPAD(SUBSTRING_INDEX(@tmp_ddmm, '.', -1), 2, 0)      AS 'c_mm',
        @tmp_dd := LPAD(SUBSTRING_INDEX(@tmp_ddmm, '.', 1), 2, 0)       AS 'c_dd',
        @tmp_yyyymmdd := CONCAT(@tmp_yyyy, '-', @tmp_mm, '-', @tmp_dd)  AS 'c_yyyymmdd',
        --
        @tmp_time := SUBSTRING_INDEX(`created_at`, ' ', -1)             AS 'c_time',
        @tmp_hours := LPAD(SUBSTRING_INDEX(@tmp_time, ':', 1), 2, 0)    AS 'c_hours',
        @tmp_mins := LPAD(SUBSTRING_INDEX(@tmp_time, ':', -1), 2, 0)    AS 'c_minutes',
        @tmp_hhmmss := CONCAT(@tmp_hours, ':', @tmp_mins, ':00')        AS 'c_hhmmss',
        CONCAT(@tmp_yyyymmdd, ' ', @tmp_hhmmss)                         AS 'str_created_at',
        --
        @tmp_date:=SUBSTRING_INDEX(`updated_at`, ' ', 1)                AS 'u_date',
        @tmp_ddmm := SUBSTRING_INDEX(@tmp_date, '.', 2)                 AS 'u_ddmm',
        @tmp_yyyy := SUBSTRING_INDEX(@tmp_date, '.', -1)                AS 'u_yyyy',
        @tmp_mm := LPAD(SUBSTRING_INDEX(@tmp_ddmm, '.', -1), 2, 0)      AS 'u_mm',
        @tmp_dd := LPAD(SUBSTRING_INDEX(@tmp_ddmm, '.', 1), 2, 0)       AS 'u_dd',
        @tmp_yyyymmdd := CONCAT(@tmp_yyyy, '-', @tmp_mm, '-', @tmp_dd)  AS 'u_yyyymmdd',
        --
        @tmp_time := SUBSTRING_INDEX(`updated_at`, ' ', -1)             AS 'u_time',
        @tmp_hours := LPAD(SUBSTRING_INDEX(@tmp_time, ':', 1), 2, 0)    AS 'u_hours',
        @tmp_mins := LPAD(SUBSTRING_INDEX(@tmp_time, ':', -1), 2, 0)    AS 'u_minutes',
        @tmp_hhmmss := CONCAT(@tmp_hours, ':', @tmp_mins, ':00')        AS 'u_hhmmss',
        CONCAT(@tmp_yyyymmdd, ' ', @tmp_hhmmss)                         AS 'str_updated_at'
    FROM `users`
) AS u2 ON u1.id = u2.id
SET 
    u1.created_at = u2.str_created_at,
    u1.updated_at = u2.str_updated_at;

-- Изменим типы колонок.
ALTER TABLE `users` 
    MODIFY `created_at` DATETIME,
    MODIFY `updated_at` DATETIME;

-- Отобразим конечное описание таблицы.
DESCRIBE `users`;

-- Отобразим конечное содержимое таблицы.
SELECT * FROM `users`;