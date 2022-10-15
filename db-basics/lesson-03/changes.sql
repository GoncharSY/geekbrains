-- Изменение имени таблицы.
ALTER TABLE 'average_grades' RENAME TO 'grades';

-- Изменение имени колонки.
ALTER TABLE 'streams' RENAME COLUMN 'start_date' TO 'started_at';

-- Добавление новой колонки.
ALTER TABLE 'streams' ADD COLUMN 'finished_at' TEXT;

-- Добавление преподавателей.
INSERT INTO 'teachers' ('name', 'surname', 'email') VALUES
    ('Николай', 'Савельев', 'saveliev.n@mail.ru'),
    ('Наталья', 'Петрова', 'petrova.n@yandex.ru'),
    ('Елена', 'Малышева', 'malisheva.e@google.com');

-- Добавление курсов.
INSERT INTO 'courses' ('name') VALUES
    ('Базы данных'),
    ('Основы Python'),
    ('Linux. Рабочая станция');

-- Добавление потоков.
INSERT INTO 'streams' ('course_id', 'number', 'started_at', 'finished_at', 'student_amount')
VALUES
    (3, 165, '18.08.2020', '18.09.2020', 34),
    (2, 178, '02.10.2020', '02.11.2020', 37),
    (1, 203, '12.11.2020', '12.12.2020', 35),
    (1, 210, '03.12.2020', '03.01.2021', 41);

-- Добавление оценок.
INSERT INTO 'grades' ('teacher_id', 'stream_id', 'performance') VALUES
    (3, 1, 4.7),
    (2, 2, 4.9),
    (1, 3, 4.8),
    (1, 4, 4.9);

-- Далее идет набор команд изменения БД для выполнения 4-го задания.
-- 1. Создание новой временной таблицы.
CREATE TABLE IF NOT EXISTS 'grades_tmp' (
    'teacher_id' INTEGER NOT NULL,
    'stream_id' REAL NOT NULL,
    'performance' REAL NOT NULL DEFAULT 0,
    PRIMARY KEY ('teacher_id', 'stream_id'),
    FOREIGN KEY ('teacher_id') REFERENCES 'teachers' ('id'),
    FOREIGN KEY ('stream_id') REFERENCES 'streams' ('id')   
);

-- 2. Перенос данных из старой таблицы во временную.
INSERT INTO 'grades_tmp' SELECT * FROM 'grades';

-- 3. Удаление старой таблицы.
DROP TABLE IF EXISTS 'grades';

-- 4. Переименование временной таблицы (изменение ее статуса на постояннную).
ALTER TABLE 'grades_tmp' RENAME TO 'grades';
