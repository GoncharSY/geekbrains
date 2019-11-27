/* 
    Составьте список пользователей users, которые осуществили хотя бы 
    один заказ (orders) в интернет-магазине. 
*/

-- Заказы пользователей:
SELECT 
	usr.id,
	usr.name,
	COUNT(ord.id) AS 'count of orders'
FROM orders AS ord 
RIGHT JOIN users AS usr ON usr.id = ord.user_id
GROUP BY usr.id, usr.name;

-- Решение 1:
SELECT usr.*
FROM users AS usr
LEFT JOIN orders AS ord ON usr.id = ord.user_id
GROUP BY usr.id
HAVING COUNT(ord.id) != 0;

-- Решение 2:
SELECT * 
FROM users 
WHERE EXISTS (
    SELECT 1 
    FROM orders 
    WHERE user_id = users.id
);