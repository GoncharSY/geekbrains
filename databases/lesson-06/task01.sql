/*
В базе данных shop и sample присутвуют одни и те же таблицы учебной базы 
данных. Переместите запись id = 1 из таблицы shop.users в таблицу sample.users. 
Используйте транзакции.
*/

SET @user_id = 1;


-- Очистим таблицу базы 'sample':
TRUNCATE gsy_sample.users;


-- Отобразим исходное состояние таблиц:
SELECT 'shop' AS 'DB', u1.id, u1.name 
    FROM gsy_shop.users AS u1
    WHERE id = @user_id
UNION ALL
SELECT 'sample' AS 'DB', u2.id, u2.name 
    FROM gsy_sample.users AS u2
    WHERE id = @user_id;


-- Решение:
START TRANSACTION;

INSERT INTO gsy_sample.users
SELECT * FROM gsy_shop.users 
WHERE id = @user_id;

COMMIT;


-- Результат:
SELECT 'shop' AS 'DB', u1.id, u1.name 
    FROM gsy_shop.users AS u1
    WHERE id = @user_id
UNION ALL
SELECT 'sample' AS 'DB', u2.id, u2.name 
    FROM gsy_sample.users AS u2
    WHERE id = @user_id;
