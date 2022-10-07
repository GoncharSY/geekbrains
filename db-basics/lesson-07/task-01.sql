-- Последние потоки курсов:
SELECT
    number,
    course_id,
    MAX(started_at) AS 'start_date'
FROM streams
GROUP BY course_id;

-- Средняя успеваемость курса:
-- Учитываются только завершенные потоки.
SELECT
    course_id,
    AVG(performance) AS 'avg_grade'
FROM grades
LEFT JOIN streams ON id = stream_id
WHERE finished_at < CURRENT_DATE
GROUP BY course_id;

-- Решение:
CREATE VIEW 'course_statuses' AS SELECT
    crs.name AS 'course_name',
    lst.number AS 'last_stream',
    lst.start_date,
    grd.avg_grade
FROM courses AS 'crs'
LEFT JOIN (
    SELECT
        number,
        course_id,
        MAX(started_at) AS 'start_date'
    FROM streams
    GROUP BY course_id
) AS 'lst' ON crs.id = lst.course_id
LEFT JOIN (
    SELECT
        course_id,
        AVG(performance) AS 'avg_grade'
    FROM grades
    LEFT JOIN streams ON id = stream_id
    WHERE finished_at < CURRENT_DATE
    GROUP BY course_id
) AS 'grd' ON crs.id = grd.course_id;

-- Запрос:
SELECT * FROM course_statuses;
