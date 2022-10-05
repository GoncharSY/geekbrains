-- ================================================================================================
-- Задание №1: Потоки и ученики ===================================================================
-- ================================================================================================

SELECT
    "number" AS '№',
    (SELECT "name" FROM "courses" WHERE "id" = "course_id") AS 'course',
    "student_amount" AS 'students'
FROM "streams" WHERE "student_amount" >= 40;

-- ================================================================================================
-- Задание №2: Два потока низкой успеваемостью ====================================================
-- ================================================================================================

SELECT
    (SELECT "number" FROM "streams" WHERE "id" = "stream_id") AS 'stream',
    (SELECT "name" FROM "courses" WHERE "id" = (
        SELECT "course_id" FROM "streams" WHERE "id" = "stream_id")) AS 'course',
    (SELECT "surname" || ' ' || "name" FROM "teachers" WHERE "id" = "teacher_id") AS 'teacher',
    "performance" AS 'grade'
FROM "grades" ORDER BY "performance" ASC LIMIT 2;

-- ================================================================================================
-- Задание №3: Потоки Савельева Н. ================================================================
-- ================================================================================================

SELECT "teacher_id", AVG("performance") AS 'avg_grade' FROM "grades" 
WHERE "teacher_id" = (
    SELECT "id" FROM "teachers" WHERE "name" == "Николай" AND "surname" = "Савельев");

-- ================================================================================================
-- Задание №4: Потоки Петровой Н. =================================================================
-- ================================================================================================

SELECT 
    "stream_id",
    (SELECT "surname" FROM "teachers" WHERE "id" = "teacher_id") AS 'teacher_surname',
    (SELECT "name" FROM "teachers" WHERE "id" = "teacher_id") AS 'teacher_name'
FROM "grades" 
WHERE "teacher_id" = (
    SELECT "id" FROM "teachers" WHERE "name" == "Наталья" AND "surname" = "Петрова")
UNION
SELECT 
    "stream_id",
    (SELECT "surname" FROM "teachers" WHERE "id" = "teacher_id") AS 'teacher_surname',
    (SELECT "name" FROM "teachers" WHERE "id" = "teacher_id") AS 'teacher_name'
FROM "grades" 
WHERE "performance" < 4.8;

-- ================================================================================================
-- Задание №5: Разность успеваемости ==============================================================
-- ================================================================================================

-- Вариант 1:
SELECT (
    SELECT AVG("performance") AS 'avg_grade'
    FROM "grades" GROUP BY "teacher_id" 
    ORDER BY "avg_grade" DESC LIMIT 1
) - (
    SELECT AVG("performance") AS 'avg_grade'
    FROM "grades" GROUP BY "teacher_id" 
    ORDER BY "avg_grade" LIMIT 1
);

-- Вариант 2:
SELECT (
    SELECT MAX("value") FROM (SELECT AVG("performance") AS 'value' FROM "grades" GROUP BY "teacher_id")
) - (
    SELECT MIN("value") FROM (SELECT AVG("performance") AS 'value' FROM "grades" GROUP BY "teacher_id")
);

-- Наглядное представление:
SELECT 'max_avg' AS 'value_name', MAX("value") AS 'value' 
FROM (SELECT AVG("performance") AS 'value' FROM "grades" GROUP BY "teacher_id")
UNION
SELECT 'min_avg' AS 'value_name', MIN("value") AS 'value'
FROM (SELECT AVG("performance") AS 'value' FROM "grades" GROUP BY "teacher_id")
UNION
SELECT 'difference' AS 'name_value', (
    SELECT MAX("value") FROM (SELECT AVG("performance") AS 'value' FROM "grades" GROUP BY "teacher_id")
) - (
    SELECT MIN("value") FROM (SELECT AVG("performance") AS 'value' FROM "grades" GROUP BY "teacher_id")
) AS 'value' ORDER BY "value" DESC;
