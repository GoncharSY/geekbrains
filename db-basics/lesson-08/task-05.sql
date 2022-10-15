-- Решение:

SELECT
    t.name AS 'teacher_name',
    t.surname AS 'teacher_surname',
    min.grade1 AS 'min_grade',
    min.course_name AS 'min_course',
    max.grade2 AS 'max_grade',
    max.course_name AS 'max_course',
    nxt.date AS 'next_stream'
FROM teachers AS 't'
LEFT JOIN (
    SELECT DISTINCT
        g.teacher_id,
        c.name AS 'course_name',
        MIN(g.performance) OVER teachers_w AS 'grade1'
    FROM grades AS 'g'
    LEFT JOIN streams AS 's' ON g.stream_id = s.id
    LEFT JOIN courses AS 'c' ON s.course_id = c.id
    WHERE s.finished_at < CURRENT_DATE
    WINDOW 'teachers_w' AS (PARTITION BY g.teacher_id)
) AS 'min' ON t.id = min.teacher_id
LEFT JOIN (
    SELECT DISTINCT
        g.teacher_id,
        c.name AS 'course_name',
        MAX(g.performance) OVER teachers_w AS 'grade2'
    FROM grades AS 'g'
    LEFT JOIN streams AS 's' ON g.stream_id = s.id
    LEFT JOIN courses AS 'c' ON s.course_id = c.id
    WHERE s.finished_at < CURRENT_DATE
    WINDOW 'teachers_w' AS (PARTITION BY g.teacher_id)
) AS 'max' ON t.id = max.teacher_id
LEFT JOIN (
    SELECT DISTINCT
        teacher_id,
        MIN(started_at) OVER teachers_w AS 'date'
    FROM grades LEFT JOIN streams ON stream_id = streams.id
    WHERE started_at > CURRENT_DATE
    WINDOW 'teachers_w' AS (PARTITION BY teacher_id)
) AS 'nxt' ON t.id = nxt.teacher_id;

-- Решение из 6-го урока:

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
