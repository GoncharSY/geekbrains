/*
(по желанию) Пусть имеется таблица с календарным полем created_at. В ней 
размещены разряженые календарные записи за август 2018 года '2018-08-01', 
'2016-08-04', '2018-08-16' и '2018-08-17'. 
Составьте запрос, который выводит полный список дат за август, выставляя в 
соседнем поле значение 1, если дата присутствует в исходном таблице и 0, если 
она отсутствует.
*/

-- Создадим таблицу с отдельными днями месяца:
DROP TABLE IF EXISTS `tmp_some_dates`;
CREATE TEMPORARY TABLE `tmp_some_dates` (created_at DATE);

-- Добави даты:
INSERT INTO `tmp_some_dates` VALUES
    ('2018-08-01'),
    ('2018-08-04'),
    ('2018-08-16'),
    ('2018-08-17');

-- Создадим таблицу со всеми днями месяца:
DROP TABLE IF EXISTS `tmp_all_dates`;
CREATE TEMPORARY TABLE `tmp_all_dates` (created_at DATE);

-- Добавим даты:
DELIMITER $$ 
DROP PROCEDURE IF EXISTS tmp_add_all_dates $$
CREATE PROCEDURE tmp_add_all_dates () BEGIN
    SET @mm = '08';
    SET @date = '2018-08-01';
    
    WHILE MONTH(@date) = @mm DO
        INSERT INTO tmp_all_dates VALUES (@date);

        SET @date = @date + INTERVAL 1 DAY;
    END WHILE;
END $$
DELIMITER ;

CALL tmp_add_all_dates();

-- Результат:
SELECT ta.*, (ts.created_at IS NOT NULL) AS 'Exists'
FROM `tmp_all_dates` AS ta 
LEFT JOIN `tmp_some_dates` AS ts ON ta.created_at = ts.created_at
ORDER BY ta.created_at;
