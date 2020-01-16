/* 
    3. (по желанию) Напишите хранимую функцию для вычисления произвольного числа 
    Фибоначчи. Числами Фибоначчи называется последовательность в которой число равно 
    сумме двух предыдущих чисел. Вызов функции FIBONACCI(10) должен возвращать число 55.
*/

/*=====================================================================================*/
/*=====================================================================================*/
/*=====================================================================================*/

-- Создадим функцию:
DELIMITER $$
DROP FUNCTION IF EXISTS `fibonacci` $$
CREATE FUNCTION `fibonacci`(arg INT) RETURNS INT DETERMINISTIC
BEGIN
    DECLARE itr, s0, s1, result INT;

    SET s0 = 0;
    SET s1 = 1;
    SET itr = 2;

    -- Обобые случаи.
    IF arg = 0 THEN
        RETURN s0;
    ELSEIF arg = 1 THEN
        RETURN s1;
    END IF;

    -- Нормальное вычисление.
    1st: WHILE itr <= arg DO
        SET result = s0 + s1;
        SET s0 = s1;
        SET s1 = result;
        SET itr = itr + 1;
    END WHILE 1st;

    RETURN result;
END $$
DELIMITER ;

-- Создадим таблицу.
DROP TABLE IF EXISTS `fibonacci_numbers`;
CREATE TEMPORARY TABLE `fibonacci_numbers` (`num` INT);

-- Наполним таблицу.
INSERT INTO `fibonacci_numbers` VALUES (0), (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);

-- Отобразим результат:
SELECT *, fibonacci(num) FROM `fibonacci_numbers`;

-- В исходное состояние:
DROP TABLE IF EXISTS `fibonacci_numbers`;
DROP FUNCTION IF EXISTS `fibonacci`;