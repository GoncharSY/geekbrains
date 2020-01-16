/* 
    2. В таблице products есть два текстовых поля: name с названием товара и 
    description с его описанием. Допустимо присутствие обоих полей или одно из них. 
    Ситуация, когда оба поля принимают неопределенное значение NULL неприемлема. 
    Используя триггеры, добейтесь того, чтобы одно из этих полей или оба поля были 
    заполнены. При попытке присвоить полям NULL-значение необходимо отменить операцию.
*/

/*=====================================================================================*/
/*=====================================================================================*/
/*=====================================================================================*/

-- Выберем базу:
USE `gsy_shop`;

-- Проверим содержимое таблицы:
-- SELECT `id`, `name`, `description` 
--     FROM `products` 
--     ORDER BY `id` DESC 
--     LIMIT 4;


-- Удалим триггер, если такой был:
DROP TRIGGER IF EXISTS `product_fill_control`;

DELIMITER $$

-- Создадим тирггер снова:
CREATE TRIGGER `product_fill_control` BEFORE INSERT ON `products`
FOR EACH ROW BEGIN
    DECLARE filled_value VARCHAR(255);

    SET filled_value = COALESCE(NEW.name, NEW.description);

    IF filled_value IS NULL THEN
        SIGNAL SQLSTATE '45000' 
        SET MESSAGE_TEXT = 'Операция отменена: Некорректный ввод';
    END IF;
END$$

DELIMITER ;


-- Выполним проверочное тестирование:
-- INSERT INTO `products` (`id`, `name`, `description`) VALUES
--     (1001, 'Процессор', 'Описание для процессора');
-- INSERT INTO `products` (`id`, `name`, `description`) VALUES
--     (1002, 'Процессор без описания', NULL);
-- INSERT INTO `products` (`id`, `name`, `description`) VALUES
--     (1003, NULL, 'Описание для процессора без имени');
-- INSERT INTO `products` (`id`, `name`, `description`) VALUES
--     (1004, NULL, NULL);


-- Проверим содержимое таблицы:
-- SELECT `id`, `name`, `description` 
--     FROM `products` 
--     ORDER BY `id` DESC 
--     LIMIT 5;


-- Вернем таблицу в исходное состояние:
-- DELETE FROM `products` WHERE `id` > 1000;


-- Проверим содержимое таблицы:
-- SELECT `id`, `name`, `description` 
--     FROM `products` 
--     ORDER BY `id` DESC 
--     LIMIT 5;
