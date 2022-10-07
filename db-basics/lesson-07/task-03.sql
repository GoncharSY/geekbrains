-- Решение:
CREATE TRIGGER 'check_grade' BEFORE UPDATE ON grades BEGIN
    SELECT CASE
	    WHEN NEW.performance NOT BETWEEN 0 AND 5
        THEN RAISE (ABORT,'Performance can only be between 0 and 5')
    END;
END;

-- Проверка ниже 0:
UPDATE grades SET performance = -2
WHERE teacher_id = 2 AND stream_id = 2;

-- Проверка выше 5:
UPDATE grades SET performance = 5.1
WHERE teacher_id = 2 AND stream_id = 2;

-- Проверка нормального выполнения:
UPDATE grades SET performance = 0
WHERE teacher_id = 2 AND stream_id = 2;

UPDATE grades SET performance = 5
WHERE teacher_id = 2 AND stream_id = 2;

UPDATE grades SET performance = 4.9
WHERE teacher_id = 2 AND stream_id = 2;
