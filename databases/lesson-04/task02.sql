-- Пример исходных данных, для проверки.
SELECT
	u.id,
	u.name,
	u.birthday,
    @tmp:=CONCAT(
		SUBSTRING(CURDATE(), 1, 4),
		SUBSTRING(u.birthday, 5, 6)
	) AS 'this_year_date',
	DAYOFWEEK(@tmp) AS 'day of week',
    d.name AS 'name of day'
FROM users AS u
JOIN days_of_week as d ON d.id = DAYOFWEEK(CONCAT(
		SUBSTRING(CURDATE(), 1, 4),
		SUBSTRING(u.birthday, 5, 6)
	));

-- Решение задачи.
SELECT dow.name, IF(ubds.total IS NULL, 0, ubds.total) AS 'birthdays'
FROM days_of_week AS dow
LEFT JOIN(
	SELECT
		DAYOFWEEK(CONCAT(
			SUBSTRING(CURDATE(), 1, 4),
			SUBSTRING(birthday, 5, 6)
		)) AS 'day_of_week',
		COUNT(id) AS 'total'
	FROM users
	GROUP BY day_of_week
) AS ubds ON ubds.day_of_week = dow.id;