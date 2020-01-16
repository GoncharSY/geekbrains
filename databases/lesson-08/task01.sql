/* 
    1. Создайте хранимую функцию hello(), которая будет возвращать приветствие, 
    в зависимости от текущего времени суток. С 6:00 до 12:00 функция должна 
    возвращать фразу "Доброе утро", с 12:00 до 18:00 функция должна возвращать 
    фразу "Добрый день", с 18:00 до 00:00 — "Добрый вечер", с 00:00 до 6:00 — 
    "Доброй ночи".
*/

/*=====================================================================================*/
/*=====================================================================================*/
/*=====================================================================================*/

-- Выберем базу.
USE db_example;

-- Удалим фукнкцию, если уже была.
DROP FUNCTION IF EXISTS `hello`;
DROP FUNCTION IF EXISTS `hello_on_time`;

-- Создадим новую функцию с параметром времени.
DELIMITER $$
CREATE FUNCTION `hello_on_time`(`arg_time` TIME) RETURNS VARCHAR(255) DETERMINISTIC
BEGIN
    DECLARE `phrase` VARCHAR(255) DEFAULT 'Доброй ночи!';
    DECLARE `now_time` TIME DEFAULT TIME(NOW());
    DECLARE `now_hour` TINYINT;

    SET `now_time` = COALESCE(`arg_time`, `now_time`);
    SET `now_hour` = HOUR(`now_time`);

    IF `now_hour` > 17 THEN
        SET `phrase` = 'Добрый вечер!';
    ELSEIF `now_hour` > 11 THEN
        SET `phrase` = 'Добрый день!';
    ELSEIF  `now_hour` > 5 THEN
        SET `phrase` = 'Доброе утро!';
    END IF;

    RETURN `phrase`;
END $$

-- Создадим функцию, с сигнатурой, как в задании.
CREATE FUNCTION `hello`() RETURNS VARCHAR(255) NOT DETERMINISTIC
BEGIN
    RETURN `hello_on_time`(NULL);
END $$
DELIMITER ;

-- Выполним примерочное тестирование:
-- SELECT hello_on_time('00:00:00') AS 'Приветсвие';
-- SELECT hello_on_time('05:59:59') AS 'Приветсвие';
-- SELECT hello_on_time('06:00:00') AS 'Приветсвие';
-- SELECT hello_on_time('11:59:59') AS 'Приветсвие';
-- SELECT hello_on_time('12:00:00') AS 'Приветсвие';
-- SELECT hello_on_time('17:59:59') AS 'Приветсвие';
-- SELECT hello_on_time('18:00:00') AS 'Приветсвие';
-- SELECT hello_on_time('23:59:59') AS 'Приветсвие';

-- Вызовем функцию.
-- SELECT hello() AS 'Приветсвие';

-- Удалим фукнкции.
-- DROP FUNCTION IF EXISTS `hello`;
-- DROP FUNCTION IF EXISTS `hello_on_time`;
