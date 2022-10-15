-- Решение:
CREATE TRIGGER 'new_stream' BEFORE INSERT ON streams BEGIN
    SELECT CASE
	    WHEN NEW.started_at <= CURRENT_DATE
        THEN RAISE (ABORT,'Start date must be in the future')
        WHEN NEW.number <= (SELECT MAX(number) FROM streams)
        THEN RAISE (ABORT,'Number must be greater than existing one')
    END;
END;

-- Проверка: Прошедшая дата:
INSERT INTO streams (id, number, course_id, student_amount, started_at, finished_at)
VALUES (6, 212, 3, 33, '2022-10-06', '2022-11-06');

-- Проверка: Нормер потока меньше существующих:
INSERT INTO streams (id, number, course_id, student_amount, started_at, finished_at)
VALUES (6, 209, 3, 33, '2022-10-08', '2022-11-08');

-- Проверка: Нормальное исполнение:
INSERT INTO streams (id, number, course_id, student_amount, started_at, finished_at)
VALUES (6, 212, 3, 33, '2022-10-08', '2022-11-08');

-- Удаление ненужной записи:
DELETE FROM streams WHERE id = 6;
