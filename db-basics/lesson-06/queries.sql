-- ================================================================================================
-- Задание №1: ====================================================================================
-- ================================================================================================

SELECT
    "number" AS 'stream_number',
    "name" AS 'course_name',
    "started_at" AS 'start_date'
FROM "streams" LEFT JOIN "courses" ON "course_id" = "courses"."id";

-- ================================================================================================
-- Задание №2: ====================================================================================
-- ================================================================================================

SELECT
    "name" AS "course_name",
    SUM("student_amount") AS 'student_total'
FROM "streams" LEFT JOIN "courses" ON "course_id" = "courses"."id"
GROUP BY "course_id";

-- ================================================================================================
-- Задание №3: ====================================================================================
-- ================================================================================================

-- Учитель без потоков:
INSERT INTO "teachers" ("name", "surname", "email") VALUES ("Иван", "Иванов", "ivanov.i@mail.ru");

-- Решение:
SELECT
    "id" AS 'teacher_id',
    "surname" AS 'teacher_surname',
    "name" AS 'teacher_name',
    AVG("performance") AS 'avg_grade'
FROM "teachers" LEFT JOIN "grades" ON "teacher_id" = "teachers"."id"
GROUP BY "teacher_id";

-- ================================================================================================
-- Задание №4: ====================================================================================
-- ================================================================================================

-- Будущий поток поток (id: 5).
INSERT INTO "streams" ("number", "course_id", "student_amount", "started_at", "finished_at")
VALUES (211, 3, 36, '2022-11-20', '2022-12-20');

-- Назначение будущего потока преподавателю.
INSERT INTO "grades" ("teacher_id", "stream_id")
VALUES (3, 5);

-- Прошедшие курсы преподавателей с худшими потоками.
SELECT
    "g"."teacher_id",
    "c"."name" AS 'course_name',
    MIN("g"."performance") AS 'grade'
FROM "grades" AS 'g'
LEFT JOIN "streams" AS 's' ON "g"."stream_id" = "s"."id"
LEFT JOIN "courses" AS 'c' ON "s"."course_id" = "c"."id"
WHERE "s"."finished_at" < CURRENT_DATE
GROUP BY "g"."teacher_id";

-- Прошедшие курсы преподавателей с лучшими потоками.
SELECT
    "g"."teacher_id",
    "c"."name" AS 'course_name',
    MAX("g"."performance") AS 'grade'
FROM "grades" AS 'g'
LEFT JOIN "streams" AS 's' ON "g"."stream_id" = "s"."id"
LEFT JOIN "courses" AS 'c' ON "s"."course_id" = "c"."id"
WHERE "s"."finished_at" < CURRENT_DATE
GROUP BY "g"."teacher_id";

-- Грядующие потоки у преподавателей.
SELECT
    "teacher_id",
    MIN("started_at") AS 'date'
FROM "grades" LEFT JOIN "streams" ON "stream_id" = "streams"."id"
WHERE "started_at" > CURRENT_DATE
GROUP BY "teacher_id";

-- Итоговое решение.
SELECT
    "t"."name" AS 'teacher_name',
    "t"."surname" AS 'teacher_surname',
    "min"."grade" AS 'min_grade',
    "min"."course_name" AS 'min_course',
    "max"."grade" AS 'max_grade',
    "max"."course_name" AS 'max_course',
    "nxt"."date" AS 'next_stream'
FROM "teachers" AS 't'
LEFT JOIN (
    SELECT
        "g"."teacher_id",
        "c"."name" AS 'course_name',
        MIN("g"."performance") AS 'grade'
    FROM "grades" AS 'g'
    LEFT JOIN "streams" AS 's' ON "g"."stream_id" = "s"."id"
    LEFT JOIN "courses" AS 'c' ON "s"."course_id" = "c"."id"
    WHERE "s"."finished_at" < CURRENT_DATE
    GROUP BY "g"."teacher_id"
) AS 'min' ON "t"."id" = "min"."teacher_id"
LEFT JOIN (
    SELECT
        "g"."teacher_id",
        "c"."name" AS 'course_name',
        MAX("g"."performance") AS 'grade'
    FROM "grades" AS 'g'
    LEFT JOIN "streams" AS 's' ON "g"."stream_id" = "s"."id"
    LEFT JOIN "courses" AS 'c' ON "s"."course_id" = "c"."id"
    WHERE "s"."finished_at" < CURRENT_DATE
    GROUP BY "g"."teacher_id"
) AS 'max' ON "t"."id" = "max"."teacher_id"
LEFT JOIN (
    SELECT
        "teacher_id",
        MIN("started_at") AS 'date'
    FROM "grades" LEFT JOIN "streams" ON "stream_id" = "streams"."id"
    WHERE "started_at" > CURRENT_DATE
    GROUP BY "teacher_id"
) AS 'nxt' ON "t"."id" = "nxt"."teacher_id";

-- ================================================================================================
-- ВНЕ ЗАДАНИЙ: ===================================================================================
-- ================================================================================================

-- Потоки, когда, курс, преподаватель, успеваемость.
-- Этот запрос для ручной проверки 4-го задания.
SELECT
    "s"."number" AS "stream",
    "s"."started_at" AS "start",
    "s"."finished_at" AS "finish",
    "c"."name" AS "course",
    "t"."surname" || ' ' || "t"."name" AS 'teacher',
    "g"."performance" AS "grade"
FROM "streams" AS 's'
LEFT JOIN "courses" AS 'c' ON "s"."course_id" = "c"."id"
LEFT JOIN "grades" AS 'g' ON "s"."id" = "g"."stream_id"
LEFT JOIN "teachers" AS 't' ON "g"."teacher_id" = "t"."id";
