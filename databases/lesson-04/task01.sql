-- Список пользователей, пример исходных данных.
SELECT 
	id,
	name,
	birthday,
	now() as 'now',
	TIMESTAMPDIFF(YEAR, birthday, NOW()) AS 'age'
FROM users;

-- Решение задачи.
SELECT
	COUNT(`id`) AS 'Count of users',
	SUM(TIMESTAMPDIFF(YEAR, birthday, NOW())) AS 'Sum of ages',
	AVG(TIMESTAMPDIFF(YEAR, birthday, NOW())) AS 'Average age'
FROM users;