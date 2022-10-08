-- Решение:
SELECT DISTINCT
    id AS 'teacher_id',
    surname AS 'teacher_surname',
    name AS 'teacher_name',
    AVG(performance) OVER w_teachers AS 'avg_grade'
FROM teachers
LEFT JOIN (
    SELECT teacher_id, performance
    FROM grades LEFT JOIN streams
    ON streams.id = stream_id
    WHERE finished_at < CURRENT_DATE
) ON teacher_id = teachers.id
WINDOW 'w_teachers' AS (PARTITION BY teacher_id)
ORDER BY teachers.id;

-- Решение из 6-го урока:
SELECT
    "id" AS 'teacher_id',
    "surname" AS 'teacher_surname',
    "name" AS 'teacher_name',
    AVG("performance") AS 'avg_grade'
FROM "teachers" LEFT JOIN "grades" ON "teacher_id" = "teachers"."id"
GROUP BY "teacher_id";
