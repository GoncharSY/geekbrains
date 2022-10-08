-- Решение:
SELECT DISTINCT
    courses.name AS 'course_name',
    SUM(streams.student_amount) OVER w_courses AS 'student_total'
FROM streams LEFT JOIN courses ON streams.course_id = courses.id
WINDOW 'w_courses' AS (PARTITION BY courses.id);

-- Решение из 6-го урока:
SELECT
    name AS course_name,
    SUM(student_amount) AS 'student_total'
FROM streams LEFT JOIN courses ON course_id = courses.id
GROUP BY course_id;
